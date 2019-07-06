package dosproxy

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	//	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/DOSNetwork/DOSscan-api/models"
	//	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/jinzhu/gorm"
)

var methodMap map[string]string
var proxyAbi abi.ABI

func init() {
	jsonFile, err := os.Open("/Users/chenhaonien/go/src/github.com/DOSNetwork/explorer-Api/subscriber/abi/DOSProxy.abi")
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

func getTx(txHash common.Hash, client *ethclient.Client, db *gorm.DB) *models.Transaction {
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Println("TransactionByHash err", err)
		return nil
	}
	methodName := "ExternalCall"
	if method, err := proxyAbi.MethodById(tx.Data()[:4]); err == nil {
		methodName = method.Name
	}
	mTx := models.Transaction{
		Hash:     fmt.Sprintf("%x", txHash.Big()),
		GasPrice: tx.GasPrice().Uint64(),
		Value:    tx.Value().Uint64(),
		GasLimit: tx.Gas(),
		Nonce:    tx.Nonce(),
		To:       fmt.Sprintf("%x", tx.To().Big()),
		Data:     tx.Data(),
		Method:   methodName,
	}
	if err := db.Where("Hash = ?", mTx.Hash).First(&mTx).Error; gorm.IsRecordNotFoundError(err) {
		db.Create(&mTx)
		fmt.Println("Saved Tx Log: ", mTx.Hash)
	} else {
		fmt.Println("duplicate Tx Log: ", mTx.Hash)
	}
	return &mTx

}
func FromBlockNumber(ctx context.Context, event string, db *gorm.DB) (chan uint64, chan error) {
	out := make(chan uint64)
	errc := make(chan error)
	go func() {
		var lastBlkNum uint64

		latestRecord := fmt.Sprintf("SELECT block_number FROM %s ORDER BY block_number DESC LIMIT 1;", event)
		rows, err := db.Raw(latestRecord).Rows() // (*sql.Rows, error)
		if err != nil {
			fmt.Println(event, " : lastblock err", err)
			lastBlkNum = 0
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&lastBlkNum)
		}
		if lastBlkNum < 4468402 {
			lastBlkNum = 4468402
		}
		fmt.Println(event, " : lastblock ", lastBlkNum)
		select {
		case <-ctx.Done():
		case out <- lastBlkNum:
		}
	}()
	return out, errc
}

var ModelsTable = []func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error{
	0: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					fmt.Println("DosproxyLogUrl got event ")

					log, ok := event.(*DosproxyLogUrl)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogURL{
						Topics:            topics,
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash:   fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:           log.Raw.TxIndex,
						LogIndex:          log.Raw.Index,
						Removed:           log.Raw.Removed,
						QueryId:           fmt.Sprintf("%x", log.QueryId),
						Timeout:           fmt.Sprintf("%x", log.Timeout),
						DataSource:        fmt.Sprintf("%x", log.DataSource),
						Selector:          fmt.Sprintf("%x", log.Selector),
						Randomness:        fmt.Sprintf("%x", log.Randomness),
						DispatchedGroupId: fmt.Sprintf("%x", log.DispatchedGroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogURL = append(tx.LogURL, mLog)
						db.Save(&tx)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	1: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogRequestUserRandom{
						Topics:               topics,
						BlockNumber:          log.Raw.BlockNumber,
						BlockHash:            fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:              log.Raw.TxIndex,
						LogIndex:             log.Raw.Index,
						Removed:              log.Raw.Removed,
						RequestId:            fmt.Sprintf("%x", log.RequestId),
						LastSystemRandomness: fmt.Sprintf("%x", log.LastSystemRandomness),
						UserSeed:             fmt.Sprintf("%x", log.UserSeed),
						DispatchedGroupId:    fmt.Sprintf("%x", log.DispatchedGroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogRequestUserRandom = append(tx.LogRequestUserRandom, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	2: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogNonSupportedType)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogNonSupportedType{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						InvalidSelector: fmt.Sprintf("%x", log.InvalidSelector),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogNonSupportedType = append(tx.LogNonSupportedType, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	3: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogNonContractCall)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogNonContractCall{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						From:            fmt.Sprintf("%x", log.From.Big()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogNonContractCall = append(tx.LogNonContractCall, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	4: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogCallbackTriggeredFor{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						CallbackAddr:    fmt.Sprintf("%x", log.CallbackAddr.Big()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogCallbackTriggeredFor = append(tx.LogCallbackTriggeredFor, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	5: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogRequestFromNonExistentUC)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogRequestFromNonExistentUC{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogRequestFromNonExistentUC = append(tx.LogRequestFromNonExistentUC, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	6: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogUpdateRandom{
						Topics:            topics,
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash:   fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:           log.Raw.TxIndex,
						LogIndex:          log.Raw.Index,
						Removed:           log.Raw.Removed,
						LastRandomness:    fmt.Sprintf("%x", log.LastRandomness),
						DispatchedGroupId: fmt.Sprintf("%x", log.DispatchedGroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogUpdateRandom = append(tx.LogUpdateRandom, mLog)
						db.Save(&tx)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	7: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogValidationResult{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						TrafficType:     log.TrafficType,
						TrafficId:       fmt.Sprintf("%x", log.TrafficId),
						Message:         log.Message,
						Signature:       []string{fmt.Sprintf("%x", log.Signature[0]), fmt.Sprintf("%x", log.Signature[1])},
						PubKey:          []string{fmt.Sprintf("%x", log.PubKey[0]), fmt.Sprintf("%x", log.PubKey[1]), fmt.Sprintf("%x", log.PubKey[2]), fmt.Sprintf("%x", log.PubKey[3])},
						Pass:            log.Pass,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogValidationResult = append(tx.LogValidationResult, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	8: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogInsufficientPendingNode)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogInsufficientPendingNode{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						NumPendingNodes: log.NumPendingNodes.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogInsufficientPendingNode = append(tx.LogInsufficientPendingNode, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	9: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					fmt.Println("DosproxyLogInsufficientWorkingGroup")
					log, ok := event.(*DosproxyLogInsufficientWorkingGroup)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogInsufficientWorkingGroup{
						Topics:           topics,
						BlockNumber:      log.Raw.BlockNumber,
						BlockHash:        fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash:  fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:          log.Raw.TxIndex,
						LogIndex:         log.Raw.Index,
						Removed:          log.Raw.Removed,
						NumWorkingGroups: log.NumWorkingGroups.Uint64(),
						NumPendingGroups: log.NumPendingGroups.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogInsufficientWorkingGroup = append(tx.LogInsufficientWorkingGroup, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	10: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					var nodeIdstr []string
					for _, n := range log.NodeId {
						nodeIdstr = append(nodeIdstr, fmt.Sprintf("%x", n.Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogGrouping{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         fmt.Sprintf("%x", log.GroupId),
						NodeId:          nodeIdstr,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogGrouping = append(tx.LogGrouping, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	11: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogPublicKeyAccepted{
						Topics:           topics,
						BlockNumber:      log.Raw.BlockNumber,
						BlockHash:        fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash:  fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:          log.Raw.TxIndex,
						LogIndex:         log.Raw.Index,
						Removed:          log.Raw.Removed,
						GroupId:          fmt.Sprintf("%x", log.GroupId),
						PubKey:           []string{fmt.Sprintf("%x", log.PubKey[0]), fmt.Sprintf("%x", log.PubKey[1]), fmt.Sprintf("%x", log.PubKey[2]), fmt.Sprintf("%x", log.PubKey[3])},
						NumWorkingGroups: log.NumWorkingGroups.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogPublicKeyAccepted = append(tx.LogPublicKeyAccepted, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	12: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogPublicKeySuggested{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         fmt.Sprintf("%x", log.GroupId),
						PubKeyCount:     log.PubKeyCount.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogPublicKeySuggested = append(tx.LogPublicKeySuggested, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	13: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogGroupDissolve{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         fmt.Sprintf("%x", log.GroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogGroupDissolve = append(tx.LogGroupDissolve, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	14: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogRegisteredNewPendingNode{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						Node:            fmt.Sprintf("%x", log.Node),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogRegisteredNewPendingNode = append(tx.LogRegisteredNewPendingNode, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	15: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogGroupingInitiated)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogGroupingInitiated{
						Topics:            topics,
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash:   fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:           log.Raw.TxIndex,
						LogIndex:          log.Raw.Index,
						Removed:           log.Raw.Removed,
						PendingNodePool:   log.PendingNodePool.Uint64(),
						Groupsize:         log.Groupsize.Uint64(),
						Groupingthreshold: log.Groupingthreshold.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogGroupingInitiated = append(tx.LogGroupingInitiated, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	16: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogNoPendingGroup)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogNoPendingGroup{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         fmt.Sprintf("%x", log.GroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogNoPendingGroup = append(tx.LogNoPendingGroup, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	17: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogPendingGroupRemoved)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogPendingGroupRemoved{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         fmt.Sprintf("%x", log.GroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogPendingGroupRemoved = append(tx.LogPendingGroupRemoved, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	18: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogError{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						Err:             log.Err,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogError = append(tx.LogError, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	19: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateGroupToPick)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateGroupToPick{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						OldNum:          log.OldNum.Uint64(),
						NewNum:          log.NewNum.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.UpdateGroupToPick = append(tx.UpdateGroupToPick, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	20: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateGroupSize)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateGroupSize{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						OldSize:         log.OldSize.Uint64(),
						NewSize:         log.NewSize.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.UpdateGroupSize = append(tx.UpdateGroupSize, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	21: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateGroupingThreshold)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateGroupingThreshold{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						OldThreshold:    log.OldThreshold.Uint64(),
						NewThreshold:    log.NewThreshold.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.UpdateGroupingThreshold = append(tx.UpdateGroupingThreshold, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	22: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateGroupMaturityPeriod)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateGroupMaturityPeriod{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						OldPeriod:       log.OldPeriod.Uint64(),
						NewPeriod:       log.NewPeriod.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.UpdateGroupMaturityPeriod = append(tx.UpdateGroupMaturityPeriod, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	23: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateBootstrapCommitDuration)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateBootstrapCommitDuration{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						OldDuration:     log.OldDuration.Uint64(),
						NewDuration:     log.NewDuration.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.UpdateBootstrapCommitDuration = append(tx.UpdateBootstrapCommitDuration, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	24: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateBootstrapRevealDuration)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateBootstrapRevealDuration{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						OldDuration:     log.OldDuration.Uint64(),
						NewDuration:     log.NewDuration.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.UpdateBootstrapRevealDuration = append(tx.UpdateBootstrapRevealDuration, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	25: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdatebootstrapStartThreshold)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdatebootstrapStartThreshold{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						OldThreshold:    log.OldThreshold.Uint64(),
						NewThreshold:    log.NewThreshold.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.UpdatebootstrapStartThreshold = append(tx.UpdatebootstrapStartThreshold, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	26: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdatePendingGroupMaxLife)
					if !ok {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdatePendingGroupMaxLife{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						OldLifeBlocks:   log.OldLifeBlocks.Uint64(),
						NewLifeBlocks:   log.NewLifeBlocks.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.UpdatePendingGroupMaxLife = append(tx.UpdatePendingGroupMaxLife, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	27: func(ctx context.Context, db *gorm.DB, eventc chan interface{}, client *ethclient.Client) chan error {
		errc := make(chan error)
		go func() {
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
						topics = append(topics, fmt.Sprintf("%x", log.Raw.Topics[i].Big()))
					}
					tx := getTx(log.Raw.TxHash, client, db)
					if tx == nil {
						continue
					}
					mLog := models.GuardianReward{
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TransactionHash: fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						BlkNum:          log.BlkNum.Uint64(),
						Guardian:        fmt.Sprintf("%x", log.Guardian.Big()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.GuardianReward = append(tx.GuardianReward, mLog)
						db.Save(&tx)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
}
