package dosproxy

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

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
	jsonFile, err := os.Open("./abi/DOSProxy.abi")
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogUrl{
						EventLog:          "LogUrl",
						Method:            tx.Method,
						Topics:            topics,
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         log.Raw.BlockHash.Hex(),
						TransactionHash:   log.Raw.TxHash.Hex(),
						TxIndex:           log.Raw.TxIndex,
						LogIndex:          log.Raw.Index,
						Removed:           log.Raw.Removed,
						QueryId:           hexutil.Encode(log.QueryId.Bytes()),
						Timeout:           hexutil.Encode(log.Timeout.Bytes()),
						DataSource:        log.DataSource,
						Selector:          log.Selector,
						Randomness:        hexutil.Encode(log.Randomness.Bytes()),
						DispatchedGroupId: hexutil.Encode(log.DispatchedGroupId.Bytes()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						tx.LogUrl = append(tx.LogUrl, mLog)
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogRequestUserRandom{
						EventLog:             "LogRequestUserRandom",
						Method:               tx.Method,
						Topics:               topics,
						BlockNumber:          log.Raw.BlockNumber,
						BlockHash:            log.Raw.BlockHash.Hex(),
						TransactionHash:      log.Raw.TxHash.Hex(),
						TxIndex:              log.Raw.TxIndex,
						LogIndex:             log.Raw.Index,
						Removed:              log.Raw.Removed,
						RequestId:            hexutil.Encode(log.RequestId.Bytes()),
						LastSystemRandomness: hexutil.Encode(log.LastSystemRandomness.Bytes()),
						UserSeed:             hexutil.Encode(log.UserSeed.Bytes()),
						DispatchedGroupId:    hexutil.Encode(log.DispatchedGroupId.Bytes()),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogNonSupportedType{
						EventLog:        "LogNonSupportedType",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						InvalidSelector: log.InvalidSelector,
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogNonContractCall{
						EventLog:        "LogNonContractCall",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						CallAddr:        log.From.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogCallbackTriggeredFor{
						EventLog:        "LogCallbackTriggeredFor",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						CallbackAddr:    log.CallbackAddr.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogRequestFromNonExistentUC{
						EventLog:        "LogRequestFromNonExistentUC",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogUpdateRandom{
						EventLog:          "LogUpdateRandom",
						Method:            tx.Method,
						Topics:            topics,
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         log.Raw.BlockHash.Hex(),
						TransactionHash:   log.Raw.TxHash.Hex(),
						TxIndex:           log.Raw.TxIndex,
						LogIndex:          log.Raw.Index,
						Removed:           log.Raw.Removed,
						LastRandomness:    hexutil.Encode(log.LastRandomness.Bytes()),
						DispatchedGroupId: hexutil.Encode(log.DispatchedGroupId.Bytes()),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogValidationResult{
						EventLog:        "LogValidationResult",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						TrafficType:     log.TrafficType,
						TrafficId:       hexutil.Encode(log.TrafficId.Bytes()),
						Signature:       []string{hexutil.Encode(log.Signature[0].Bytes()), hexutil.Encode(log.Signature[1].Bytes())},
						PubKey:          []string{hexutil.Encode(log.PubKey[0].Bytes()), hexutil.Encode(log.PubKey[1].Bytes()), hexutil.Encode(log.PubKey[2].Bytes()), hexutil.Encode(log.PubKey[3].Bytes())},
						Pass:            log.Pass,
					}

					if log.TrafficType == 2 {
						mLog.Message = string(log.Message)
					} else {
						mLog.Message = hexutil.Encode(log.Message)
						return
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogInsufficientPendingNode{
						EventLog:        "LogInsufficientPendingNode",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogInsufficientWorkingGroup{
						EventLog:         "LogInsufficientWorkingGroup",
						Method:           tx.Method,
						Topics:           topics,
						BlockNumber:      log.Raw.BlockNumber,
						BlockHash:        log.Raw.BlockHash.Hex(),
						TransactionHash:  log.Raw.TxHash.Hex(),
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
						EventLog:        "LogGrouping",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         hexutil.Encode(log.GroupId.Bytes()),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogPublicKeyAccepted{
						EventLog:         "LogPublicKeyAccepted",
						Method:           tx.Method,
						Topics:           topics,
						BlockNumber:      log.Raw.BlockNumber,
						BlockHash:        log.Raw.BlockHash.Hex(),
						TransactionHash:  log.Raw.TxHash.Hex(),
						TxIndex:          log.Raw.TxIndex,
						LogIndex:         log.Raw.Index,
						Removed:          log.Raw.Removed,
						GroupId:          hexutil.Encode(log.GroupId.Bytes()),
						PubKey:           []string{hexutil.Encode(log.PubKey[0].Bytes()), hexutil.Encode(log.PubKey[1].Bytes()), hexutil.Encode(log.PubKey[2].Bytes()), hexutil.Encode(log.PubKey[3].Bytes())},
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogPublicKeySuggested{
						EventLog:        "LogPublicKeySuggested",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         hexutil.Encode(log.GroupId.Bytes()),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogGroupDissolve{
						EventLog:        "LogGroupDissolve",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         hexutil.Encode(log.GroupId.Bytes()),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogRegisteredNewPendingNode{
						EventLog:        "LogRegisteredNewPendingNode",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						Node:            log.Node.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogGroupingInitiated{
						EventLog:          "LogGroupingInitiated",
						Method:            tx.Method,
						Topics:            topics,
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         log.Raw.BlockHash.Hex(),
						TransactionHash:   log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogNoPendingGroup{
						EventLog:        "LogNoPendingGroup",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         hexutil.Encode(log.GroupId.Bytes()),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogPendingGroupRemoved{
						EventLog:        "LogPendingGroupRemoved",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						GroupId:         hexutil.Encode(log.GroupId.Bytes()),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.LogError{
						EventLog:        "LogError",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateGroupToPick{
						EventLog:        "UpdateGroupToPick",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateGroupSize{
						EventLog:        "UpdateGroupSize",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateGroupingThreshold{
						EventLog:        "UpdateGroupingThreshold",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateGroupMaturityPeriod{
						EventLog:        "UpdateGroupMaturityPeriod",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateBootstrapCommitDuration{
						EventLog:        "UpdateBootstrapCommitDuration",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdateBootstrapRevealDuration{
						EventLog:        "UpdateBootstrapRevealDuration",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdatebootstrapStartThreshold{
						EventLog:        "UpdatebootstrapStartThreshold",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.UpdatePendingGroupMaxLife{
						EventLog:        "UpdatePendingGroupMaxLife",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
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
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, client, db)
					if tx == nil {
						continue
					}
					mLog := models.GuardianReward{
						EventLog:        "GuardianReward",
						Method:          tx.Method,
						Topics:          topics,
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       log.Raw.BlockHash.Hex(),
						TransactionHash: log.Raw.TxHash.Hex(),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						BlkNum:          log.BlkNum.Uint64(),
						Guardian:        log.Guardian.Hex(),
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
