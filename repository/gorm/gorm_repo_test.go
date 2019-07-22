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
		&models.Group{}, &models.Node{}, &models.LogRequestUserRandom{}, &models.LogUrl{},
		&models.LogValidationResult{}, &models.UrlRequest{}, &models.UserRandomRequest{})
	return db
}

func mockLogUrl(r repository.DB, t *testing.T, sender string, nonce uint64, hash string, blknum uint64, requestId, groupId string) {
	tx := models.Transaction{
		Hash:        hash,
		GasPrice:    2000000000,
		Value:       0,
		GasLimit:    6000000,
		Nonce:       nonce,
		Sender:      sender,
		To:          "0xproxyadd00000000000000000000000000000000",
		BlockNumber: blknum,
		Method:      "External",
	}
	log := models.LogUrl{Event: models.Event{
		Method:          "External",
		EventLog:        "LogUrl",
		TransactionHash: hash,
		TxIndex:         8,
		BlockNumber:     blknum,
	},
		RequestId:         requestId,
		DispatchedGroupId: groupId,
		DataSource:        "Fake Source",
	}

	eventc := make(chan []interface{})
	var input []interface{}
	input = append(input, tx)
	input = append(input, log)

	err, errc := r.SaveModel(context.Background(), models.TypeUrl, eventc)
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

func mockLogUserRandom(r repository.DB, t *testing.T, sender string, nonce uint64, hash string, blknum uint64, requestId, groupId string) {
	tx := models.Transaction{
		Hash:        hash,
		GasPrice:    2000000000,
		Value:       0,
		GasLimit:    6000000,
		Nonce:       nonce,
		Sender:      sender,
		To:          "0xproxyadd00000000000000000000000000000000",
		BlockNumber: blknum,
		Method:      "External",
	}
	log := models.LogRequestUserRandom{Event: models.Event{
		Method:          "External",
		EventLog:        "LogRequestUserRandom",
		TransactionHash: hash,
		TxIndex:         8,
		BlockNumber:     blknum,
	},
		RequestId:         requestId,
		DispatchedGroupId: groupId,
	}

	eventc := make(chan []interface{})
	var input []interface{}
	input = append(input, tx)
	input = append(input, log)

	err, errc := r.SaveModel(context.Background(), models.TypeRequestUserRandom, eventc)
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

func mockValidationResult(r repository.DB, t *testing.T, sender string, nonce uint64, hash string, blknum uint64, requestId string) {
	tx := models.Transaction{
		Hash:        hash,
		GasPrice:    2000000000,
		Value:       0,
		GasLimit:    6000000,
		Nonce:       nonce,
		Sender:      sender,
		To:          "0xproxyadd00000000000000000000000000000000",
		BlockNumber: blknum,
		Method:      "External",
	}
	log := models.LogValidationResult{Event: models.Event{
		Method:          "External",
		EventLog:        "ValidationResult",
		TransactionHash: hash,
		TxIndex:         8,
		BlockNumber:     blknum,
	},
		RequestId: requestId,
		Message:   "Fake Mesage",
		Pass:      true,
	}

	eventc := make(chan []interface{})
	var input []interface{}
	input = append(input, tx)
	input = append(input, log)

	err, errc := r.SaveModel(context.Background(), models.TypeValidationResult, eventc)
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

func TestBuildRequest(t *testing.T) {
	//sender string,nonce uint64, hash string ,blknum uint64,groupid string) {
	groupIdTemplate := "0xgroup000000000000000000000id00000000000000000000000000000000000"
	groupHashTemplate := "0xgrouping0000000000000000000000000000000000000000000000000000000"
	acceptedHashTemplate := "0xaccepted0000000000000000000000000000000000000000000000000000000"
	dissolvHashTemplate := "0xdissolv00000000000000000000000000000000000000000000000000000000"
	urlIdTemplate := "0xrequestUrl000000000000000000000id000000000000000000000000000000"
	urlHashTemplate := "0xUrlHash000000000000000000000id000000000000000000000000000000000"
	randomIdTemplate := "0xrequestRandom000000000000000000000id000000000000000000000000000"
	randomHashTemplate := "0xRandomHash000000000000000000000id000000000000000000000000000000"
	validationHashTemplate := "0xRulalidationHash000000000000000000000id000000000000000000000000"
	randomValidationHashTemplate := "0xRandomvalidationHash000000000000000000000id00000000000000000000"

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

	nonce = 29999
	blknum = 4468500
	for i := 1; i <= 3; i++ {
		groupId := groupIdTemplate + strconv.Itoa(i)
		urlId := urlIdTemplate + strconv.Itoa(i)
		urlHash := urlHashTemplate + strconv.Itoa(i)
		validationHash := validationHashTemplate + strconv.Itoa(i)
		sender := senderTemplate + strconv.Itoa(i)

		mockLogUrl(r, t, sender+strconv.Itoa(i), nonce, urlHash, blknum, urlId, groupId)
		mockValidationResult(r, t, sender+strconv.Itoa(i), nonce+1, validationHash, blknum+5, urlId)
		blknum++
		buildUrlRequest(db, urlId)
	}

	nonce = 30000
	blknum = 4468510
	for i := 1; i <= 3; i++ {
		groupId := groupIdTemplate + strconv.Itoa(i)
		randomId := randomIdTemplate + strconv.Itoa(i)
		randomHash := randomHashTemplate + strconv.Itoa(i)
		validationHash := randomValidationHashTemplate + strconv.Itoa(i)
		sender := senderTemplate + strconv.Itoa(i)

		mockLogUserRandom(r, t, sender+strconv.Itoa(i), nonce, randomHash, blknum, randomId, groupId)
		mockValidationResult(r, t, sender+strconv.Itoa(i), nonce+1, validationHash, blknum+5, randomId)
		blknum++
		buildRandomRequest(db, randomId)
	}

	for i := 1; i <= 3; i++ {
		groupId := groupIdTemplate + strconv.Itoa(i)
		var urls []models.UrlRequest
		var randoms []models.UserRandomRequest
		var group models.Group
		db.Where(&models.Group{GroupId: groupId}).First(&group)
		db.Model(&group).Related(&urls, "UrlRequests")
		db.Model(&group).Related(&randoms, "UserRandomRequests")
		fmt.Println(len(urls))
		fmt.Println(len(randoms))
	}
}
