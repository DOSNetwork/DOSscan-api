// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_config "github.com/DOSNetwork/DOSscan-api/config"
	_repository "github.com/DOSNetwork/DOSscan-api/repository"
	_cache "github.com/DOSNetwork/DOSscan-api/repository/cache"
	_gorm "github.com/DOSNetwork/DOSscan-api/repository/gorm"
	_onchain "github.com/DOSNetwork/DOSscan-api/repository/onchain"
	_handler "github.com/DOSNetwork/DOSscan-api/server/handler"
	"github.com/DOSNetwork/DOSscan-api/server/middleware"
	_service "github.com/DOSNetwork/DOSscan-api/service"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

func main() {
	configPath := ""
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	}
	if configPath == "" {
		configPath = "./config.json"
	}
	config, err := _config.LoadConfig(configPath)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	//1)Init repositorys abd service
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	cacheRepo := _cache.NewCacheRepo(c)

	var db *gorm.DB
	postgres_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB_USER, config.DB_PASSWORD, config.DB_IP, config.DB_PORT, config.DB_NAME)
	if db, err = gorm.Open("postgres", postgres_url); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbRepo := _gorm.NewGormRepo(db)

	var onchainRepo _repository.Onchain
	for _, geth := range config.ChainNodePool {
		fmt.Println("Dial to ", geth)
		var client *ethclient.Client
		client, err = ethclient.Dial(geth)
		if err != nil {
			fmt.Printf("Dial %v\n", err)
			client.Close()
			continue
		}
		onchainRepo, err = _onchain.NewGethRepo(client, config.BRIDGE_ADDR)
		if err != nil {
			fmt.Printf("NewGethRepo %v\n", err)
			client.Close()
			continue
		}
		break
	}
	if onchainRepo != nil {
		search := _service.NewSearch(onchainRepo, dbRepo)

		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.ForwardedByClientIP = true
		r.Use(middleware.CORS())

		// Setup route group for the API
		searchHandler := _handler.NesSearchHandler(search, cacheRepo)
		api := r.Group("/api")
		v1 := api.Group("/explorer")
		v1.GET("/search", searchHandler.Search)
		v1.GET("/eventNames", searchHandler.SupportedEvents)
		bootStrapHandler := _handler.NesBootStrapHandler(config.BootStrapIPs)
		api.GET("/bootStrap", bootStrapHandler.BootStrap)
		server := &http.Server{
			Addr:           ":8080",
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()
		go func() {
			for _ = range ticker.C {
				searchHandler.UpdateLatestEvent()
			}
		}()
		go server.ListenAndServe()
		gracefulExitWeb(server)
	} else {
		fmt.Println("No onchainRepo")
	}
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("------exited--------", time.Since(now))
}
