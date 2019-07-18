package dosproxy

import (
	"context"
	"fmt"
	"io/ioutil"

	//	"log"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

	//	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/DOSNetwork/DOSscan-api/models"
	//	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

var proxyAbi abi.ABI

const (
	LogRegisteredNewPendingNode int = iota
	LogGrouping
	LogPublicKeySuggested
	LogPublicKeyAccepted
	LogGroupDissolve

	LogUpdateRandom
	LogUrl
	LogRequestUserRandom
	LogValidationResult
	LogCallbackTriggeredFor
	GuardianReward

	LogError
)

var ProxyEvent = []interface{}{
	LogRegisteredNewPendingNode: &models.LogRegisteredNewPendingNode{},
	LogGrouping:                 &models.LogGrouping{},
	LogPublicKeySuggested:       &models.LogPublicKeySuggested{},
	LogPublicKeyAccepted:        &models.LogPublicKeyAccepted{},
	LogGroupDissolve:            &models.LogGroupDissolve{},

	LogUpdateRandom:         &models.LogUpdateRandom{},
	LogUrl:                  &models.LogUrl{},
	LogRequestUserRandom:    &models.LogRequestUserRandom{},
	LogValidationResult:     &models.LogValidationResult{},
	LogCallbackTriggeredFor: &models.LogCallbackTriggeredFor{},
	//LogError:                &models.LogError{},
	GuardianReward: &models.GuardianReward{},
}

func init() {
	jsonFile, err := os.Open("../../abi/DOSProxy.abi")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	abiJsonByte, _ := ioutil.ReadAll(jsonFile)
	proxyAbi, err = abi.JSON(strings.NewReader(string(abiJsonByte)))
	if err != nil {
		fmt.Println(err)
	}
}

func getBalance(client *ethclient.Client, addr common.Address) (balance *big.Float) {
	wei, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))

	return balance
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func getTx(txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, client *ethclient.Client, db *gorm.DB) *models.Transaction {
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Println("TransactionByHash err", err)
		return nil
	}
	sender, err := client.TransactionSender(context.Background(), tx, blockhash, index)
	if err != nil {
		fmt.Println("GetTransactionSender err", err)
		return nil
	}
	var methodName string
	if method, err := proxyAbi.MethodById(tx.Data()[:4]); err == nil {
		methodName = method.Name
	} else {
		methodName = fmt.Sprintf("ExternalCall 0x%x", tx.Data()[:4])
	}
	mTx := models.Transaction{
		Hash:        txHash.Hex(),
		GasPrice:    tx.GasPrice().Uint64(),
		Value:       tx.Value().Uint64(),
		GasLimit:    tx.Gas(),
		Nonce:       tx.Nonce(),
		Sender:      sender.Hex(),
		To:          tx.To().Hex(),
		BlockNumber: blockNum,
		Data:        tx.Data(),
		Method:      methodName,
	}
	if err := db.FirstOrCreate(&mTx, mTx).Error; err != nil {
		fmt.Println("Create tx err ", err.Error())
	}
	if mTx.Hash == "0x4f573abf3e004505166cd618f60545b7b1cca73b426f5644d43a8650e1e61a44" {
		fmt.Println("tx ", mTx.Hash)
		fmt.Println("BlockNumber ", mTx.BlockNumber)
	}
	/*
		if err := db.Where("hash = ?", mTx.Hash).First(&mTx).Error; gorm.IsRecordNotFoundError(err) {
			if err := db.Create(&mTx).Error; err != nil {
				// error handling...
				fmt.Println("Create tx err ", err.Error)
				db.First(&mTx)
			}
		}*/

	return &mTx
}

func buildGroup(db *gorm.DB, grouId string) {
	var results []models.Group
	tempDb := db.Table("log_groupings").Select("log_groupings.group_id, log_public_key_accepteds.accepted_blk_num,log_group_dissolves.dissolved_blk_num, log_groupings.node_id, log_public_key_accepteds.pub_key")
	tempDb = tempDb.Joins("left join log_public_key_accepteds on log_public_key_accepteds.group_id = log_groupings.group_id")
	tempDb = tempDb.Joins("left join log_group_dissolves on log_group_dissolves.group_id = log_groupings.group_id")
	if grouId == "" {
		tempDb.Find(&results)
	} else {
		tempDb.Where("log_groupings.group_id = ? ", grouId).Find(&results)
	}
	fmt.Println(len(results))
	for _, group := range results {
		var existGroup models.Group
		if err := db.Where("group_id = ?", group.GroupId).First(&existGroup).Error; gorm.IsRecordNotFoundError(err) {
			db.Create(&group)
		} else {
			db.Model(&existGroup).Omit("group_id").Updates(&group)
			fmt.Println("Update group ", existGroup.GroupId, group.DissolvedBlkNum, group.AcceptedBlkNum)
		}
	}
}

func buildNode(db *gorm.DB, addr string) {
	var node models.Node
	db.Where(models.Node{Addr: addr}).FirstOrCreate(&node)
	//Find Group has node addr in node_id
	sel := "SELECT group_id FROM groups WHERE $1 <@ node_id"
	rows, err := db.DB().Query(sel, pq.Array([]string{addr}))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var group models.Group
		rows.Scan(&group.GroupId)
		if err := db.Where("group_id = ?", group.GroupId).First(&group).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println("Can't find group ", group.GroupId, " ", node.Addr)
		} else {
			res := db.Model(&node).Association("Groups").Append(&group)
			if res.Error != nil {
				fmt.Println("res ", res.Error)
			}
			fmt.Println("len ", db.Model(&node).Association("Groups").Count())
			fmt.Println("len ", node.Groups[0].GroupId)
		}
	}
}
func buildUrlRequest(db *gorm.DB, requestId string) {
	var results []models.UrlRequest
	tempDb := db.Table("log_urls").Select("log_urls.request_id, log_urls.dispatched_group_id,transactions.sender, transactions.block_number,transactions.hash,log_validation_results.message,log_validation_results.signature,log_validation_results.pub_key,log_validation_results.pass,log_urls.timeout,log_urls.data_source,log_urls.selector,log_urls.randomness")
	tempDb = tempDb.Joins("inner join log_validation_results on log_validation_results.request_id = log_urls.request_id")
	tempDb = tempDb.Joins("inner join transactions on log_validation_results.transaction_id = transactions.id")
	if requestId == "" {
		tempDb.Find(&results)
	} else {
		tempDb.Where("log_urls.request_id = ? ", requestId).Find(&results)
	}

	for _, request := range results {
		db.Where(request).FirstOrCreate(&request)
		var group models.Group
		db.Where(&models.Group{GroupId: request.DispatchedGroupId}).First(&group)
		res := db.Model(&group).Association("UrlRequests").Append(&request)
		if res.Error != nil {
			fmt.Println("res ", res.Error)
		}
		fmt.Println(group.GroupId, "-", " len ", db.Model(&group).Association("UrlRequests").Count())
	}
}
func buildRandomRequest(db *gorm.DB, requestId string) {
	var results []models.UserRandomRequest
	tempDb := db.Table("log_request_user_randoms").Select("log_request_user_randoms.request_id, log_request_user_randoms.dispatched_group_id,transactions.sender, transactions.block_number,transactions.hash,log_validation_results.message,log_validation_results.signature,log_validation_results.pub_key,log_validation_results.pass")
	tempDb = tempDb.Joins("inner join log_validation_results on log_validation_results.request_id = log_request_user_randoms.request_id")
	tempDb = tempDb.Joins("inner join transactions on log_validation_results.transaction_id = transactions.id")
	if requestId == "" {
		tempDb.Find(&results)
	} else {
		tempDb.Where("log_urls.request_id = ? ", requestId).Find(&results)
	}

	for _, request := range results {
		db.Where(request).FirstOrCreate(&request)
		var group models.Group
		db.Where(&models.Group{GroupId: request.DispatchedGroupId}).First(&group)
		res := db.Model(&group).Association("UrlRequests").Append(&request)
		if res.Error != nil {
			fmt.Println("res ", res.Error)
		}
		fmt.Println(group.GroupId, "-", " len ", db.Model(&group).Association("UrlRequests").Count())
	}
}
func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func BuildRelation(db *gorm.DB) {
	buildGroup(db, "")
	//Build node
	var addrs []string
	db.Model(&models.LogRegisteredNewPendingNode{}).Pluck("node", &addrs)
	fmt.Println(len(addrs))
	addrs = removeDuplicates(addrs)
	for i := 0; i < len(addrs); i++ {
		buildNode(db, addrs[i])
	}
	buildUrlRequest(db, "")
	buildRandomRequest(db, "")
}
func FromBlockNumber(ctx context.Context, event interface{}, db *gorm.DB) (chan uint64, chan error) {
	out := make(chan uint64)
	errc := make(chan error)
	go func() {
		defer close(out)
		defer close(errc)
		var lastBlkNum uint64

		var blkNums []uint64
		db.Limit(1).Order("block_number desc").Find(event).Pluck("block_number", &blkNums)
		if len(blkNums) == 0 {
			fmt.Printf("%T No recored \n", event)
			lastBlkNum = 4468402
		} else {
			lastBlkNum = blkNums[0]
		}

		fmt.Printf("%T lastblock = %d \n", event, lastBlkNum)
		select {
		case <-ctx.Done():
		case out <- lastBlkNum:
		}
	}()
	return out, errc
}

func ModelsForDashboard(ctx context.Context, db *gorm.DB, eventc chan interface{}, proxyS *DosproxySession, client *ethclient.Client) chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		for {
			select {
			case <-ctx.Done():
				return
			case event, ok := <-eventc:
				if !ok {
					return
				}
				switch t := event.(type) {

				case *models.LogUrl:
				case *models.LogRequestUserRandom:
				case *models.LogUpdateRandom:
				case *models.LogRegisteredNewPendingNode:

				case *models.LogGrouping:
					gInfo := models.Group{GroupId: t.GroupId}
					if err := db.Where("group_id = ?", t.GroupId).First(&gInfo).Error; gorm.IsRecordNotFoundError(err) {
						gInfo.NodeId = t.NodeId
						db.Save(&gInfo)
						fmt.Println("Save GroupInfo", t.GroupId)
					}
				case *models.LogPublicKeyAccepted:

				case *models.LogGroupDissolve:

				default:
					_ = t
					continue
				}
			}
		}
	}()
	return errc
}

var ModelsTable = []func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error{
	LogRegisteredNewPendingNode: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogRegisteredNewPendingNode)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogRegisteredNewPendingNode{
						Event: models.Event{
							EventLog:        "LogRegisteredNewPendingNode",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						Node: log.Node.Hex(),
					}

					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogRegisteredNewPendingNodes").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}

					}
				}
			}
		}()
		return errc
	},
	LogGrouping: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogGrouping)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					var nodeIdstr []string
					for _, n := range log.NodeId {
						nodeIdstr = append(nodeIdstr, n.Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogGrouping{
						Event: models.Event{
							EventLog:        "LogGrouping",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						GroupId: hexutil.Encode(log.GroupId.Bytes()),
						NodeId:  nodeIdstr,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogGroupings").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	LogPublicKeySuggested: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogPublicKeySuggested)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogPublicKeySuggested{
						Event: models.Event{
							EventLog:        "LogPublicKeySuggested",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						GroupId:     hexutil.Encode(log.GroupId.Bytes()),
						PubKeyCount: log.PubKeyCount.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogPublicKeySuggesteds").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}

					}
				}
			}
		}()
		return errc
	},
	LogPublicKeyAccepted: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogPublicKeyAccepted)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogPublicKeyAccepted{
						Event: models.Event{
							EventLog:        "LogPublicKeyAccepted",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						GroupId:          hexutil.Encode(log.GroupId.Bytes()),
						AcceptedBlkNum:   log.Raw.BlockNumber,
						PubKey:           []string{hexutil.Encode(log.PubKey[0].Bytes()), hexutil.Encode(log.PubKey[1].Bytes()), hexutil.Encode(log.PubKey[2].Bytes()), hexutil.Encode(log.PubKey[3].Bytes())},
						NumWorkingGroups: log.NumWorkingGroups.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogPublicKeyAccepteds").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	LogGroupDissolve: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogGroupDissolve)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogGroupDissolve{
						Event: models.Event{
							EventLog:        "LogGroupDissolve",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						GroupId:         hexutil.Encode(log.GroupId.Bytes()),
						DissolvedBlkNum: log.Raw.BlockNumber,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogGroupDissolves").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}

					}
				}
			}
		}()
		return errc
	},
	LogUpdateRandom: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogUpdateRandom)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogUpdateRandom{
						Event: models.Event{
							EventLog:        "LogUpdateRandom",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						LastRandomness:    hexutil.Encode(log.LastRandomness.Bytes()),
						DispatchedGroupId: hexutil.Encode(log.DispatchedGroupId.Bytes()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogUpdateRandoms").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	LogUrl: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}

					log, ok := event.(*DosproxyLogUrl)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogUrl{
						Event: models.Event{
							EventLog:        "LogUrl",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						RequestId:         hexutil.Encode(log.QueryId.Bytes()),
						Timeout:           log.Timeout.String(),
						DataSource:        log.DataSource,
						Selector:          log.Selector,
						Randomness:        hexutil.Encode(log.Randomness.Bytes()),
						DispatchedGroupId: hexutil.Encode(log.DispatchedGroupId.Bytes()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogUrls").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	LogRequestUserRandom: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogRequestUserRandom)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogRequestUserRandom{
						Event: models.Event{
							EventLog:        "LogRequestUserRandom",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						RequestId:            hexutil.Encode(log.RequestId.Bytes()),
						LastSystemRandomness: hexutil.Encode(log.LastSystemRandomness.Bytes()),
						UserSeed:             hexutil.Encode(log.UserSeed.Bytes()),
						DispatchedGroupId:    hexutil.Encode(log.DispatchedGroupId.Bytes()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogRequestUserRandoms").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	LogValidationResult: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogValidationResult)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogValidationResult{
						Event: models.Event{
							EventLog:        "LogValidationResult",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						RequestType: log.TrafficType,
						RequestId:   hexutil.Encode(log.TrafficId.Bytes()),
						Signature:   []string{hexutil.Encode(log.Signature[0].Bytes()), hexutil.Encode(log.Signature[1].Bytes())},
						PubKey:      []string{hexutil.Encode(log.PubKey[0].Bytes()), hexutil.Encode(log.PubKey[1].Bytes()), hexutil.Encode(log.PubKey[2].Bytes()), hexutil.Encode(log.PubKey[3].Bytes())},
						Pass:        log.Pass,
					}

					if log.TrafficType == 2 {
						mLog.Message = string(log.Message)
					} else {
						mLog.Message = hexutil.Encode(log.Message)
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogValidationResults").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	LogCallbackTriggeredFor: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogCallbackTriggeredFor)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogCallbackTriggeredFor{
						Event: models.Event{
							EventLog:        "LogCallbackTriggeredFor",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						CallbackAddr: log.CallbackAddr.Hex(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogCallbackTriggeredFors").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	LogError: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {

		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogError)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogError{
						Event: models.Event{
							EventLog:        "LogError",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						Err: log.Err,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("LogErrors").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	GuardianReward: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyGuardianReward)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.GuardianReward{
						Event: models.Event{
							EventLog:        "GuardianReward",
							Method:          tx.Method,
							Topics:          topics,
							BlockNumber:     log.Raw.BlockNumber,
							BlockHash:       log.Raw.BlockHash.Hex(),
							TransactionHash: log.Raw.TxHash.Hex(),
							TxIndex:         log.Raw.TxIndex,
							LogIndex:        log.Raw.Index,
							Removed:         log.Raw.Removed,
						},
						BlkNum:   log.BlkNum.Uint64(),
						Guardian: log.Guardian.Hex(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						res := db.Model(&tx).Association("GuardianRewards").Append(&mLog)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
}
