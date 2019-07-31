package service

import (
	"context"
	"fmt"
	"testing"

	_models "github.com/DOSNetwork/DOSscan-api/models"
	//_repository "github.com/DOSNetwork/DOSscan-api/repository"
	_gorm "github.com/DOSNetwork/DOSscan-api/repository/gorm"
	_onchain "github.com/DOSNetwork/DOSscan-api/repository/onchain"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
)

func initDB(user, password, dbName string) *gorm.DB {
	postgres_url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbName)
	var db *gorm.DB
	db, err := gorm.Open("postgres", postgres_url)
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&_models.Transaction{}, &_models.LogRegisteredNewPendingNode{},
		&_models.LogGrouping{}, &_models.LogPublicKeyAccepted{}, &_models.LogGroupDissolve{},
		&_models.Group{}, &_models.Node{}, &_models.LogRequestUserRandom{}, &_models.LogUrl{},
		&_models.LogValidationResult{}, &_models.UrlRequest{}, &_models.UserRandomRequest{})
	return db
}

func initClient(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(url, ":Dial err ", err)
		return nil
	}
	return client
}

func TestFetchHistoricalLogs(t *testing.T) {

	db := initDB("postgres", "postgres", "test")
	repositoryGorm := _gorm.NewGormRepo(db)
	client := initClient("wss://rinkeby.infura.io/ws/v3/3a3e5d776961418e93a8b33fef2f6642")
	repositoryOnchain := _onchain.NewGethRepo(client)
	modelsType := []int{_models.TypeNewPendingNode,
		_models.TypeGrouping, _models.TypePublicKeySuggested, _models.TypePublicKeyAccepted, _models.TypeGroupDissolve,
		_models.TypeUpdateRandom, _models.TypeUrl, _models.TypeRequestUserRandom, _models.TypeValidationResult,
		_models.TypeGuardianReward, _models.TypeCallbackTriggeredFor, _models.TypeError}
	initBlkNum := uint64(4468400)
	transformer := NewTransformer(repositoryOnchain, repositoryGorm, initBlkNum, modelsType)

	err, errc := transformer.FetchHistoricalLogs(context.Background())
	if err != nil {
		t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
	}
	for err := range errc {
		t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
	}

	//Verify the results
	if logTotal, err := repositoryGorm.CountModel(context.Background(), _models.TypeNewPendingNode); err != nil {
		t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
	} else {
		if total, err := repositoryGorm.CountModel(context.Background(), _models.TypeNode); err != nil {
			t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
		} else {
			if logTotal != total {
				t.Errorf("TestFetchHistoricalLogs ,Expected %d Actual %d ", logTotal, total)
			}
		}
	}

	if logTotal, err := repositoryGorm.CountModel(context.Background(), _models.TypeGrouping); err != nil {
		t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
	} else {
		if total, err := repositoryGorm.CountModel(context.Background(), _models.TypeGroup); err != nil {
			t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
		} else {
			if logTotal != total {
				t.Errorf("TestFetchHistoricalLogs ,Expected %d Actual %d ", logTotal, total)
			}
		}
	}

	if logTotal, err := repositoryGorm.CountModel(context.Background(), _models.TypeRequestUserRandom); err != nil {
		t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
	} else {
		if total, err := repositoryGorm.CountModel(context.Background(), _models.TypeRandomRequest); err != nil {
			t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
		} else {
			if logTotal != total {
				t.Errorf("TestFetchHistoricalLogs ,Expected %d Actual %d ", logTotal, total)
			}
		}
	}
	if logTotal, err := repositoryGorm.CountModel(context.Background(), _models.TypeUrl); err != nil {
		t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
	} else {
		if total, err := repositoryGorm.CountModel(context.Background(), _models.TypeUrlRequest); err != nil {
			t.Errorf("TestFetchHistoricalLogs ,Error %v", err)
		} else {
			if logTotal != total {
				t.Errorf("TestFetchHistoricalLogs ,Expected %d Actual %d ", logTotal, total)
			}
		}
	}
}
