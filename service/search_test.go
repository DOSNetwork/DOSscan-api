package service

import (
	"context"
	"fmt"
	"testing"

	_models "github.com/DOSNetwork/DOSscan-api/models"
	_gorm "github.com/DOSNetwork/DOSscan-api/repository/gorm"
	_onchain "github.com/DOSNetwork/DOSscan-api/repository/onchain"
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

//TODO: Change to mock test data

func testInit(t *testing.T) (client *ethclient.Client, db *gorm.DB) {
	var err error
	postgres_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		DB_USER, DB_PASSWORD, DB_IP, DB_PORT, DB_NAME)
	if db, err = gorm.Open("postgres", postgres_url); err != nil {
		t.Errorf(err.Error())
	}

	client, err = ethclient.Dial(ETH_URL)
	if err != nil {
		t.Errorf(err.Error())
	}
	return client, db
}

func TestSearchEventByName(t *testing.T) {
	client, db := testInit(t)
	defer db.Close()
	defer client.Close()

	search := NewSearch(_onchain.NewGethRepo(client), _gorm.NewGormRepo(db))
	_, results, resultType, _ := search.Search(context.Background(), "LogGrouping", 200, 0)
	if resultType != _models.TypeGrouping {
		t.Errorf("TestSearchEventByName : Expect type %d Actual type %d ", _models.TypeGrouping, resultType)
	}
	if len(results) > 200 {
		t.Errorf("TestSearchEventByName : Expect len > 200 Actual len %d ", len(results))
	}
}

func TestSearchNode(t *testing.T) {
	client, db := testInit(t)
	defer db.Close()
	defer client.Close()

	search := NewSearch(_onchain.NewGethRepo(client), _gorm.NewGormRepo(db))
	_, results, resultType, _ := search.Search(context.Background(), "0x6F3BA7F3E5DB7B056b2558E4D6b9063c1ABD09aB", 200, 0)
	if resultType != _models.TypeNode {
		t.Errorf("TestSubscribeTable : Expect type %d Actual type %d ", _models.TypeNode, resultType)
	}
	if len(results) != 1 {
		t.Errorf("TestSubscribeTable : Expect len %d Actual len %d ", 1, len(results))
	}
}

func TestSearchGroup(t *testing.T) {
	client, db := testInit(t)
	defer db.Close()
	defer client.Close()

	search := NewSearch(_onchain.NewGethRepo(client), _gorm.NewGormRepo(db))
	_, results, resultType, _ := search.Search(context.Background(), "0x8ffd4588d68622ba96ac3da40e5bb1176e9876b1c2240da293b6070e98884597", 200, 0)
	if resultType != _models.TypeGroup {
		t.Errorf("TestSubscribeTable : Expect type %d Actual type %d ", _models.TypeGroup, resultType)
	}
	if len(results) != 1 {
		t.Errorf("TestSubscribeTable : Expect len %d Actual len %d ", 1, len(results))
	}
}

func TestSearchRequestURL(t *testing.T) {
	client, db := testInit(t)
	defer db.Close()
	defer client.Close()

	search := NewSearch(_onchain.NewGethRepo(client), _gorm.NewGormRepo(db))
	_, results, resultType, _ := search.Search(context.Background(), "0x0e0e9563734e782e51c1fe10c8b67a1dbad3c2d9a106a3fdf6c509dde779a556", 200, 0)
	if resultType != _models.TypeUrlRequest {
		t.Errorf("TestSubscribeTable : Expect type %d Actual type %d ", _models.TypeUrlRequest, resultType)
	}
	if len(results) != 1 {
		t.Errorf("TestSubscribeTable : Expect len %d Actual len %d ", 1, len(results))
	}
}

func TestSearchRequestRandom(t *testing.T) {
	client, db := testInit(t)
	defer db.Close()
	defer client.Close()

	search := NewSearch(_onchain.NewGethRepo(client), _gorm.NewGormRepo(db))
	_, results, resultType, _ := search.Search(context.Background(), "0xb0d4922cbb20291088e7a70a86ad754b3d2eb532cf7e40c1de6153d0ef250314", 200, 0)
	if resultType != _models.TypeRandomRequest {
		t.Errorf("TestSubscribeTable : Expect type %d Actual type %d ", _models.TypeRandomRequest, resultType)
	}
	if len(results) != 1 {
		t.Errorf("TestSubscribeTable : Expect len %d Actual len %d ", 1, len(results))
	}
}
