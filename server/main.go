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

	_cache "github.com/DOSNetwork/DOSscan-api/repository/cache"
	_gorm "github.com/DOSNetwork/DOSscan-api/repository/gorm"
	_onchain "github.com/DOSNetwork/DOSscan-api/repository/onchain"
	_handler "github.com/DOSNetwork/DOSscan-api/server/handler"
	"github.com/DOSNetwork/DOSscan-api/server/middleware"
	_service "github.com/DOSNetwork/DOSscan-api/service"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

const (
	DB_IP       = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
	ETH_URL     = "wss://rinkeby.infura.io/ws/v3/3a3e5d776961418e93a8b33fef2f6642"
)

func main() {
	//TODO : Add configuration

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
		DB_USER, DB_PASSWORD, DB_IP, DB_PORT, DB_NAME)
	if db, err = gorm.Open("postgres", postgres_url); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbRepo := _gorm.NewGormRepo(db)

	var client *ethclient.Client
	client, err = ethclient.Dial(ETH_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	onchainRepo := _onchain.NewGethRepo(client)
	search := _service.NewSearch(onchainRepo, dbRepo)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.Use(middleware.CORS())

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./view", true)))
	r.Use(static.Serve("/explorer", static.LocalFile("./view", true)))
	r.Use(static.Serve("/myaccount", static.LocalFile("./view", true)))
	r.Use(static.Serve("/nodelist", static.LocalFile("./view", true)))

	// Setup route group for the API
	searchHandler := _handler.NesSearchHandler(search, cacheRepo)
	api := r.Group("/api")
	v1 := api.Group("/explorer")
	v1.GET("/search", searchHandler.Search)
	v1.GET("/eventNames", searchHandler.SupportedEvents)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go server.ListenAndServe()
	gracefulExitWeb(server)
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
