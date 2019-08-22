package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_models "github.com/DOSNetwork/DOSscan-api/models"
	_gorm "github.com/DOSNetwork/DOSscan-api/repository/gorm"
	_onchain "github.com/DOSNetwork/DOSscan-api/repository/onchain"
	_service "github.com/DOSNetwork/DOSscan-api/service"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	errors "golang.org/x/xerrors"
)

// Config is the configuration for creating a DOS client instance.
type Config struct {
	DB_IP         string
	DB_PORT       string
	DB_USER       string
	DB_PASSWORD   string
	DB_NAME       string
	ChainNodePool []string
}

// LoadConfig loads configuration file from path.
func LoadConfig(path string) (c *Config, err error) {
	var jsonFile *os.File
	var byteValue []byte

	fmt.Println("Path ", path)
	// Open our jsonFile
	jsonFile, err = os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		errors.Errorf(": %w", err)
		return
	}
	fmt.Println("Successfully Opened json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		errors.Errorf(": %w", err)
		return
	}
	c = &Config{}
	err = json.Unmarshal(byteValue, c)
	if err != nil {
		errors.Errorf(": %w", err)
		return
	}
	return
}

func main() {
	configPath := ""
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	}
	if configPath == "" {
		configPath = "./config.json"
	}
	config, err := LoadConfig(configPath)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	//1)Init repositorys
	var db *gorm.DB
	postgres_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB_USER, config.DB_PASSWORD, config.DB_IP, config.DB_PORT, config.DB_NAME)
	if db, err = gorm.Open("postgres", postgres_url); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbRepo := _gorm.NewGormRepo(db)

	//Graceful shutdown of application
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		oscall := <-sigs
		log.Printf("system call:%+v", oscall)
		cancel()
		os.Exit(0)
	}()
	//4)Start periodic task
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
RECONN:
	var transformService *_service.Transformer
	for _, geth := range config.ChainNodePool {
		fmt.Println("Dial to ", geth)
		var client *ethclient.Client
		client, err = ethclient.Dial(geth)
		if err != nil {
			fmt.Printf("Dial %v\n", err)
			client.Close()
			continue
		}
		onchainRepo, err := _onchain.NewGethRepo(client)
		if err != nil {
			fmt.Printf("NewGethRepo %v\n", err)
			client.Close()
			continue
		}
		//2)Create a service
		modelsType := []int{_models.TypePublicKeyAccepted, _models.TypeGroupDissolve,
			_models.TypeUpdateRandom, _models.TypeUrl, _models.TypeRequestUserRandom,
			_models.TypeGuardianReward, _models.TypeCallbackTriggeredFor, _models.TypeError}
		transformService = _service.NewTransformer(onchainRepo, dbRepo, 4468400, modelsType)
		if transformService != nil {
			break
		}
	}
	if transformService == nil {
		fmt.Println("Failed to new transformService")
		os.Exit(0)
	}

	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker  ")
			err := transformService.FetchHistoricalLogs(context.Background())
			if err != nil {
				fmt.Printf("FetchHistoricalLogs %v\n", err)
				goto RECONN
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
	//TODO Add subbscriber task to get real time events
}
