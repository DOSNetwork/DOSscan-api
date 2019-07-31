package main

import (
	"context"
	"fmt"
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
	//1)Init repositorys
	var db *gorm.DB
	var err error
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

	//2)Create a service
	modelsType := []int{_models.TypeNewPendingNode,
		_models.TypeGrouping, _models.TypePublicKeySuggested, _models.TypePublicKeyAccepted, _models.TypeGroupDissolve,
		_models.TypeUpdateRandom, _models.TypeUrl, _models.TypeRequestUserRandom, _models.TypeValidationResult,
		_models.TypeGuardianReward, _models.TypeCallbackTriggeredFor, _models.TypeError}
	transformService := _service.NewTransformer(onchainRepo, dbRepo, 4468400, modelsType)

	//3)Graceful shutdown of application
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
	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker  ")
			err, errc := transformService.FetchHistoricalLogs(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			for err := range errc {
				fmt.Println(err)
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}

	//TODO Add subbscriber task to get real time events

}
