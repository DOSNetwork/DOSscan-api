package gorm

import (
	"context"
	"fmt"
	"testing"

	//"database/sql"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/jinzhu/gorm"
	//"github.com/lib/pq"
)

func initDB(user, password, dbName string) *gorm.DB {
	postgres_url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbName)
	var db *gorm.DB
	db, err := gorm.Open("postgres", postgres_url)
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.Transaction{}, &models.LogRegisteredNewPendingNode{}, &models.LogGrouping{})
	return db
}

func TestSave(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	r := NewGethRepo(db)
	tx := models.Transaction{
		Hash:        "0xa3ac8994c0c5a97e7206f073c4ce050405e6b03a2d041c464cce9e9c6d228b87",
		GasPrice:    2000000000,
		Value:       0,
		GasLimit:    6000000,
		Nonce:       18927,
		Sender:      "0x3E268ECB08CF59B5c2aDBf98651ccD8041C60f67",
		To:          "0x7fD667a87E2ef724f19315124755558cAA18836E",
		BlockNumber: 4468429,
		Method:      "registerNewNode",
	}
	log := models.LogRegisteredNewPendingNode{Event: models.Event{
		Method:          "registerNewNode",
		EventLog:        "LogRegisteredNewPendingNode",
		TransactionHash: "0xa3ac8994c0c5a97e7206f073c4ce050405e6b03a2d041c464cce9e9c6d228b87",
		TxIndex:         8,
		BlockNumber:     4468429,
		BlockHash:       "0x878a44fde8051b5deb917ca8e13a2cecf90ce5d5f4d56fda45f43c5bd540bde3",
		LogIndex:        11,
		Removed:         false,
	},
		Node: "0x3E268ECB08CF59B5c2aDBf98651ccD8041C60f67",
	}

	eventc := make(chan []interface{})
	var input []interface{}
	input = append(input, tx)
	input = append(input, log)
	errc := r.SaveModel(context.Background(), models.TypeNewPendingNode, eventc)
	eventc <- input
	close(eventc)
	err := <-errc
	if err != nil {
		fmt.Println(err)
	}

	var logFromDB models.LogRegisteredNewPendingNode
	db.Model(&tx).First(&tx).Related(&logFromDB, "LogRegisteredNewPendingNodes")
	if tx.Hash != logFromDB.TransactionHash || logFromDB.TransactionHash == "" {
		t.Errorf("TestSave Error : Expect %s Acctual %s", tx.Hash, logFromDB.TransactionHash)
	}

	var txFromDB models.Transaction
	db.Model(&log).First(&log).Related(&txFromDB)
	if log.TransactionHash != txFromDB.Hash || txFromDB.Hash == "" {
		t.Errorf("TestSave Error : Expect %s Acctual %s", log.TransactionHash, txFromDB.Hash)
	}

	db.Unscoped().Delete(&log)
	db.Unscoped().Delete(&tx)
}
func TestGetEvent(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	r := NewGethRepo(db)

	log := models.LogRegisteredNewPendingNode{}
	r.GetEventsByModel(context.Background(), log)
}
