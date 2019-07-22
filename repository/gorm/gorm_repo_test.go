package gorm

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	//"database/sql"
	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/repository"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

func initDB(user, password, dbName string) *gorm.DB {
	postgres_url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbName)
	var db *gorm.DB
	db, err := gorm.Open("postgres", postgres_url)
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.Transaction{}, &models.LogRegisteredNewPendingNode{},
		&models.LogGrouping{}, &models.LogPublicKeyAccepted{}, &models.LogGroupDissolve{},
		&models.Group{}, &models.Node{})
	return db
}

func mockGrouping(r repository.DB, t *testing.T, sender string, nonce uint64, hash string, blknum uint64, groupid string, nodes []string) {
	tx := models.Transaction{
		Hash:        hash,
		GasPrice:    2000000000,
		Value:       0,
		GasLimit:    6000000,
		Nonce:       nonce,
		Sender:      sender,
		To:          "0xproxyadd00000000000000000000000000000000",
		BlockNumber: blknum,
		Method:      "triggetCallbcak",
	}
	log := models.LogGrouping{Event: models.Event{
		Method:          "triggetCallbcak",
		EventLog:        "LogRegisteredNewPendingNode",
		TransactionHash: hash,
		TxIndex:         8,
		BlockNumber:     blknum,
		BlockHash:       "0xblockhashgrouping00000000000000000000000000000000000000000000000",
		LogIndex:        11,
		Removed:         false,
	},
		GroupId: groupid,
		NodeId:  pq.StringArray(nodes),
	}

	eventc := make(chan []interface{})
	var input []interface{}
	input = append(input, tx)
	input = append(input, log)

	err, errc := r.SaveModel(context.Background(), models.TypeGrouping, eventc)
	if err != nil {
		t.Errorf("TestSave SaveModel Error : %s", err.Error())
	}
	eventc <- input
	close(eventc)
	err = <-errc
	if err != nil {
		fmt.Println(err)
	}
}

func mockAccepted(r repository.DB, t *testing.T, sender string, nonce uint64, hash string, blknum uint64, groupid string) {
	tx := models.Transaction{
		Hash:        hash,
		GasPrice:    2000000000,
		Value:       0,
		GasLimit:    6000000,
		Nonce:       nonce,
		Sender:      sender,
		To:          "0xproxyadd00000000000000000000000000000000",
		BlockNumber: blknum,
		Method:      "methdoFormockAccepted",
	}
	log := models.LogPublicKeyAccepted{Event: models.Event{
		Method:          "methdoFormockAccepted",
		EventLog:        "LogPublicKeyAccepted",
		TransactionHash: hash,
		TxIndex:         8,
		BlockNumber:     blknum,
		BlockHash:       "0xblockhashgrouping00000000000000000000000000000000000000000000000",
		LogIndex:        11,
		Removed:         false,
	},
		GroupId:          groupid,
		AcceptedBlkNum:   blknum,
		PubKey:           pq.StringArray([]string{"PubKey1", "PubKey2"}),
		NumWorkingGroups: 1,
	}

	eventc := make(chan []interface{})
	var input []interface{}
	input = append(input, tx)
	input = append(input, log)

	err, errc := r.SaveModel(context.Background(), models.TypePublicKeyAccepted, eventc)
	if err != nil {
		t.Errorf("TestSave SaveModel Error : %s", err.Error())
	}
	eventc <- input
	close(eventc)
	err = <-errc
	if err != nil {
		fmt.Println(err)
	}
}

func mockDissolve(r repository.DB, t *testing.T, sender string, nonce uint64, hash string, blknum uint64, groupid string) {
	tx := models.Transaction{
		Hash:        hash,
		GasPrice:    2000000000,
		Value:       0,
		GasLimit:    6000000,
		Nonce:       nonce,
		Sender:      sender,
		To:          "0xproxyadd00000000000000000000000000000000",
		BlockNumber: blknum,
		Method:      "methdoForLogGroupDissolve",
	}
	log := models.LogGroupDissolve{Event: models.Event{
		Method:          "methdoForLogGroupDissolve",
		EventLog:        "LogGroupDissolve",
		TransactionHash: hash,
		TxIndex:         8,
		BlockNumber:     blknum,
		BlockHash:       "0xblockhashgrouping00000000000000000000000000000000000000000000000",
		LogIndex:        10,
		Removed:         false,
	},
		GroupId:         groupid,
		DissolvedBlkNum: blknum,
	}

	eventc := make(chan []interface{})
	var input []interface{}
	input = append(input, tx)
	input = append(input, log)

	err, errc := r.SaveModel(context.Background(), models.TypeGroupDissolve, eventc)
	if err != nil {
		t.Errorf("TestSave SaveModel Error : %s", err.Error())
	}
	eventc <- input
	close(eventc)
	err = <-errc
	if err != nil {
		fmt.Println(err)
	}
}

func mockNewPendingNode(r repository.DB, t *testing.T) {
	tx := models.Transaction{
		Hash:        "0xtxhashLogRegisteredNewPendingNode0000000000000000000000000000001",
		GasPrice:    2000000000,
		Value:       0,
		GasLimit:    6000000,
		Nonce:       18927,
		Sender:      "0xnode000000000000000000000000000000000001",
		To:          "0xproxyaddrproxyaddrproxyaddrproxyaddr0000",
		BlockNumber: 4468429,
		Method:      "registerNewNode",
	}
	log := models.LogRegisteredNewPendingNode{Event: models.Event{
		Method:          "registerNewNode",
		EventLog:        "LogRegisteredNewPendingNode",
		TransactionHash: "0xtxhashLogRegisteredNewPendingNode0000000000000000000000000000001",
		TxIndex:         8,
		BlockNumber:     4468429,
		BlockHash:       "0xblockhashgrouping00000000000000000000000000000000000000000000000",
		LogIndex:        11,
		Removed:         false,
	},
		Node: "0xnode000000000000000000000000000000000001",
	}

	eventc := make(chan []interface{})
	var input []interface{}
	input = append(input, tx)
	input = append(input, log)

	err, errc := r.SaveModel(context.Background(), models.TypeNewPendingNode, eventc)
	if err != nil {
		t.Errorf("TestSave SaveModel Error : %s", err.Error())
	}
	eventc <- input
	close(eventc)
	err = <-errc
	if err != nil {
		fmt.Println(err)
	}
}

func TestSave(t *testing.T) {
	db := initDB("postgres", "postgres", "test")
	r := NewGethRepo(db)
	mockNewPendingNode(r, t)

	var txFromDB models.Transaction
	var pendingNodeFromDB models.LogRegisteredNewPendingNode

	db.Model(&pendingNodeFromDB).First(&pendingNodeFromDB).Related(&txFromDB)
	if pendingNodeFromDB.TransactionHash != txFromDB.Hash ||
		pendingNodeFromDB.TransactionHash == "" || txFromDB.Hash == "" {
		t.Errorf("TestSave Error : Expect %s Acctual %s", pendingNodeFromDB.TransactionHash, txFromDB.Hash)
	}
	db.Unscoped().Delete(&pendingNodeFromDB)
	db.Unscoped().Delete(&txFromDB)
}

func TestEventsByType(t *testing.T) {
	db := initDB("postgres", "postgres", "test")
	r := NewGethRepo(db)
	mockNewPendingNode(r, t)
	results, err := r.EventsByModelType(context.Background(), models.TypeNewPendingNode, -1, -1)
	if err != nil {
		t.Errorf("TestEventsByType Error : %s", err.Error())
	}
	fmt.Println(len(results))
}

func TestBuildGroup(t *testing.T) {
	//sender string,nonce uint64, hash string ,blknum uint64,groupid string) {
	groupIdTemplate := "0xgroup000000000000000000000id00000000000000000000000000000000000"
	groupHashTemplate := "0xgrouping0000000000000000000000000000000000000000000000000000000"
	acceptedHashTemplate := "0xaccepted0000000000000000000000000000000000000000000000000000000"
	dissolvHashTemplate := "0xdissolv00000000000000000000000000000000000000000000000000000000"
	senderTemplate := "0xnode00000000000000000000000000000000000"
	nodes1 := []string{"0xnode000000000000000000000000000000000001", "0xnode000000000000000000000000000000000002"}
	var nonce, blknum uint64
	nonce = 18927
	blknum = 4468430

	db := initDB("postgres", "postgres", "test")
	r := NewGethRepo(db)

	for i := 1; i <= 3; i++ {
		groupId := groupIdTemplate + strconv.Itoa(i)
		groupHash := groupHashTemplate + strconv.Itoa(i)
		acceptedHash := acceptedHashTemplate + strconv.Itoa(i)
		dissolvHash := dissolvHashTemplate + strconv.Itoa(i)
		sender := senderTemplate + strconv.Itoa(i)
		mockGrouping(r, t, sender+strconv.Itoa(i), nonce, groupHash, blknum, groupId, nodes1)
		mockAccepted(r, t, sender+strconv.Itoa(i), nonce+1, acceptedHash, blknum+5, groupId)
		mockDissolve(r, t, sender+strconv.Itoa(i), nonce+2, dissolvHash, blknum+10, groupId)
		blknum++
	}
	buildGroup(db, "")
	blknum = 4468430
	for i := 1; i <= 3; i++ {
		groupId := groupIdTemplate + strconv.Itoa(i)
		group, err := r.GroupByID(context.Background(), groupId)
		if err != nil {
			t.Errorf("TestBuildGroup Error : %s", err.Error())
		}
		if group.AcceptedBlkNum != blknum+5 {
			t.Errorf("TestBuildGroup Error : Expected %d Actual %d ", blknum+5, group.AcceptedBlkNum)
		}
		if group.DissolvedBlkNum != blknum+10 {
			t.Errorf("TestBuildGroup Error : Expected %d Actual %d ", blknum+10, group.DissolvedBlkNum)
		}
		blknum++
	}
}

func TestBuildNode(t *testing.T) {
	//sender string,nonce uint64, hash string ,blknum uint64,groupid string) {
	groupIdTemplate := "0xgroup000000000000000000000id00000000000000000000000000000000000"
	groupHashTemplate := "0xgrouping0000000000000000000000000000000000000000000000000000000"
	acceptedHashTemplate := "0xaccepted0000000000000000000000000000000000000000000000000000000"
	dissolvHashTemplate := "0xdissolv00000000000000000000000000000000000000000000000000000000"
	senderTemplate := "0xnode00000000000000000000000000000000000"
	nodes1 := []string{"0xnode000000000000000000000000000000000001", "0xnode000000000000000000000000000000000002"}
	var nonce, blknum uint64
	nonce = 18927
	blknum = 4468430

	db := initDB("postgres", "postgres", "test")
	r := NewGethRepo(db)
	buildNode(db, "0xnode000000000000000000000000000000000001")
	buildNode(db, "0xnode000000000000000000000000000000000002")

	for i := 1; i <= 3; i++ {
		groupId := groupIdTemplate + strconv.Itoa(i)
		groupHash := groupHashTemplate + strconv.Itoa(i)
		acceptedHash := acceptedHashTemplate + strconv.Itoa(i)
		dissolvHash := dissolvHashTemplate + strconv.Itoa(i)
		sender := senderTemplate + strconv.Itoa(i)
		mockGrouping(r, t, sender+strconv.Itoa(i), nonce, groupHash, blknum, groupId, nodes1)
		mockAccepted(r, t, sender+strconv.Itoa(i), nonce+1, acceptedHash, blknum+5, groupId)
		mockDissolve(r, t, sender+strconv.Itoa(i), nonce+2, dissolvHash, blknum+10, groupId)
		blknum++
	}
	buildGroup(db, "")

	node, err := r.NodeByAddr(context.Background(), "0xnode000000000000000000000000000000000002")
	if err != nil {
		t.Errorf("TestBuildNode Error : %s", err.Error())
	}
	fmt.Println(node.Addr)

	var groups []models.Group
	db.Model(&node).Related(&groups, "Groups")
	if len(groups) != 3 {
		t.Errorf("TestBuildGroup Error : Expected %d Actual %d ", 3, len(groups))
	}
}
