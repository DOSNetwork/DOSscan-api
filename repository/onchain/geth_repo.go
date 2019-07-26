package onchain

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/repository"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type gethRepo struct {
	client   *ethclient.Client
	proxy    *models.DosproxySession
	bridge   *models.DosbridgeSession
	proxyAbi abi.ABI
}

const (
	bridgeAddress = "0xf0CEFfc4209e38EA3Cd1926DDc2bC641cbFFd1cF"
)

func NewGethRepo(client *ethclient.Client) repository.Onchain {
	ctx := context.Background()
	d, err := models.NewDosbridge(common.HexToAddress(bridgeAddress), client)
	if err != nil {
		fmt.Println("NewDosbridge err ", err)
		return nil
	}

	bridge := &models.DosbridgeSession{Contract: d, CallOpts: bind.CallOpts{Context: ctx}}
	proxyAddr, err := bridge.GetProxyAddress()
	if err != nil {
		return nil
	}

	p, err := models.NewDosproxy(proxyAddr, client)
	if err != nil {
		return nil
	}
	proxy := &models.DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: ctx}}

	jsonFile, err := os.Open("./abi/DOSProxy.abi")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	abiJsonByte, _ := ioutil.ReadAll(jsonFile)
	proxyAbi, err := abi.JSON(strings.NewReader(string(abiJsonByte)))
	if err != nil {
		fmt.Println(err)
	}

	return &gethRepo{
		client:   client,
		proxy:    proxy,
		bridge:   bridge,
		proxyAbi: proxyAbi,
	}
}

var fetchTable = []func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error){
	models.TypeNewPendingNode: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
		out := make(chan []interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogRegisteredNewPendingNode")
			if toBlock <= fromBlock {
				return
			}
			for fromBlock <= toBlock {
				nextBlk := toBlock
				if nextBlk-fromBlock > blockLimit {
					nextBlk = fromBlock + blockLimit
				}
				fmt.Println("LogRegisteredNewPendingNode ", fromBlock, " - ", nextBlk)
				//get the historic data from proxy that start from lastBlkNum to latest
				logs, err := filter.FilterLogRegisteredNewPendingNode(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogRegisteredNewPendingNode err ", err)
				}
				for logs.Next() {
					var result []interface{}
					log := logs.Event
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
					if tx == nil {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					mLog := &models.LogRegisteredNewPendingNode{
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
					result = append(result, tx)
					result = append(result, mLog)
					select {
					case <-ctx.Done():
					case out <- result:
					}
				}
				fromBlock = nextBlk + 1

			}
		}()
		return out, errc
	},
	models.TypeGrouping: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
		out := make(chan []interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogGrouping")
			if toBlock <= fromBlock {
				return
			}
			for fromBlock <= toBlock {
				nextBlk := toBlock
				if nextBlk-fromBlock > blockLimit {
					nextBlk = fromBlock + blockLimit
				}
				fmt.Println("LogGrouping ", fromBlock, " - ", nextBlk)
				//2) get the historic data from proxy that start from lastBlkNum to latest
				logs, err := filter.FilterLogGrouping(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogGrouping err ", err)
				}
				for logs.Next() {
					var result []interface{}
					log := logs.Event
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
					if tx == nil {
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
					mLog := &models.LogGrouping{
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
					result = append(result, tx)
					result = append(result, mLog)
					select {
					case <-ctx.Done():
					case out <- result:
					}
				}
				fromBlock = nextBlk + 1
			}
			fmt.Println("LogGrouping Done fetch")

		}()
		return out, errc
	},
	models.TypePublicKeyAccepted: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
		out := make(chan []interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogPublicKeyAccepted")
			if toBlock <= fromBlock {
				return
			}
			for fromBlock <= toBlock {
				nextBlk := toBlock
				if nextBlk-fromBlock > blockLimit {
					nextBlk = fromBlock + blockLimit
				}
				fmt.Println("LogPublicKeyAccepted ", fromBlock, " - ", nextBlk)
				//get the historic data from proxy that start from lastBlkNum to latest
				logs, err := filter.FilterLogPublicKeyAccepted(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogPublicKeyAccepted err ", err)
				}
				for logs.Next() {
					var result []interface{}
					log := logs.Event
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
					if tx == nil {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					mLog := &models.LogPublicKeyAccepted{
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
					result = append(result, tx)
					result = append(result, mLog)
					select {
					case <-ctx.Done():
					case out <- result:
					}
				}
				fromBlock = nextBlk + 1
			}
		}()
		return out, errc
	},
	models.TypeGroupDissolve: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
		out := make(chan []interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogGroupDissolve")
			if toBlock <= fromBlock {
				return
			}
			for fromBlock <= toBlock {
				nextBlk := toBlock
				if nextBlk-fromBlock > blockLimit {
					nextBlk = fromBlock + blockLimit
				}
				fmt.Println("LogGroupDissolve ", fromBlock, " - ", nextBlk)
				//get the historic data from proxy that start from lastBlkNum to latest
				logs, err := filter.FilterLogGroupDissolve(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogGroupDissolve err ", err)
				}
				for logs.Next() {
					var result []interface{}
					log := logs.Event
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
					if tx == nil {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					mLog := &models.LogGroupDissolve{
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
					result = append(result, tx)
					result = append(result, mLog)
					select {
					case <-ctx.Done():
					case out <- result:
					}
				}
				fromBlock = nextBlk + 1

			}

		}()
		return out, errc
	},
	models.TypeRequestUserRandom: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
		out := make(chan []interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogRequestUserRandom")
			if toBlock <= fromBlock {
				return
			}
			for fromBlock <= toBlock {
				nextBlk := toBlock
				if nextBlk-fromBlock > blockLimit {
					nextBlk = fromBlock + blockLimit
				}
				fmt.Println("LogRequestUserRandom ", fromBlock, " - ", nextBlk)
				//get the historic data from proxy that start from lastBlkNum to latest
				logs, err := filter.FilterLogRequestUserRandom(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogRequestUserRandom err ", err)
				}
				for logs.Next() {
					var result []interface{}
					log := logs.Event
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
					if tx == nil {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					mLog := &models.LogRequestUserRandom{
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
					result = append(result, tx)
					result = append(result, mLog)
					select {
					case <-ctx.Done():
					case out <- result:
					}
				}
				fromBlock = nextBlk + 1

			}
		}()
		return out, errc
	},
	models.TypeUrl: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
		out := make(chan []interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			if toBlock <= fromBlock {
				return
			}
			for fromBlock <= toBlock {
				nextBlk := toBlock
				if nextBlk-fromBlock > blockLimit {
					nextBlk = fromBlock + blockLimit
				}
				fmt.Println("LogUrl ", fromBlock, " - ", nextBlk)
				//1) get the historic data from proxy that start from lastBlkNum to latest

				logs, err := filter.FilterLogUrl(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogUrl err ", err)
				}
				for logs.Next() {
					var result []interface{}
					log := logs.Event
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
					if tx == nil {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					mLog := &models.LogUrl{
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
					result = append(result, tx)
					result = append(result, mLog)
					select {
					case <-ctx.Done():
					case out <- result:
					}
				}
				fromBlock = nextBlk + 1
			}
			fmt.Println("LogUrl Done")

		}()
		return out, errc
	},
	models.TypeValidationResult: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
		out := make(chan []interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			if toBlock <= fromBlock {
				return
			}
			for fromBlock <= toBlock {
				nextBlk := toBlock
				if nextBlk-fromBlock > blockLimit {
					nextBlk = fromBlock + blockLimit
				}
				fmt.Println("LogValidationResult ", fromBlock, " - ", nextBlk)
				logs, err := filter.FilterLogValidationResult(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogValidationResult err ", err)
				}
				for logs.Next() {
					var result []interface{}
					log := logs.Event
					tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
					if tx == nil {
						continue
					}
					var topics []string
					for i := range log.Raw.Topics {
						topics = append(topics, log.Raw.Topics[i].Hex())
					}
					mLog := &models.LogValidationResult{
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
					result = append(result, tx)
					result = append(result, mLog)
					select {
					case <-ctx.Done():
					case out <- result:
					}
				}
				fromBlock = nextBlk + 1
			}
		}()
		return out, errc
	},
}

var subscriptionTable = []func(ctx context.Context, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error){
	models.TypeNewPendingNode: func(ctx context.Context, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *models.DosproxyLogRegisteredNewPendingNode)
		sub, err := filter.WatchLogRegisteredNewPendingNode(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for log := range ch {
				var result []interface{}
				//txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client
				tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
				if tx == nil {
					continue
				}
				result = append(result, tx)
				result = append(result, log)
				select {
				case <-ctx.Done():
					return
				case out <- result:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeGrouping: func(ctx context.Context, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *models.DosproxyLogGrouping)
		sub, err := filter.WatchLogGrouping(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for log := range ch {
				var result []interface{}
				//txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client
				tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
				if tx == nil {
					continue
				}
				fmt.Println("FetchLogs ", log.Raw.TxHash)
				result = append(result, tx)
				result = append(result, log)
				select {
				case <-ctx.Done():
					return
				case out <- result:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypePublicKeyAccepted: func(ctx context.Context, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *models.DosproxyLogPublicKeyAccepted)
		sub, err := filter.WatchLogPublicKeyAccepted(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for log := range ch {
				var result []interface{}
				//txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client
				tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
				if tx == nil {
					continue
				}
				result = append(result, tx)
				result = append(result, log)
				select {
				case <-ctx.Done():
					return
				case out <- result:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeGroupDissolve: func(ctx context.Context, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *models.DosproxyLogGroupDissolve)
		sub, err := filter.WatchLogGroupDissolve(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for log := range ch {
				var result []interface{}
				//txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client
				tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
				if tx == nil {
					continue
				}
				result = append(result, tx)
				result = append(result, log)
				select {
				case <-ctx.Done():
					return
				case out <- result:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeRequestUserRandom: func(ctx context.Context, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *models.DosproxyLogRequestUserRandom)
		sub, err := filter.WatchLogRequestUserRandom(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for log := range ch {
				var result []interface{}
				//txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client
				tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
				if tx == nil {
					continue
				}
				result = append(result, tx)
				result = append(result, log)
				select {
				case <-ctx.Done():
					return
				case out <- result:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeUrl: func(ctx context.Context, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *models.DosproxyLogUrl)
		sub, err := filter.WatchLogUrl(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for log := range ch {
				var result []interface{}
				//txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client
				tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
				if tx == nil {
					continue
				}
				result = append(result, tx)
				result = append(result, log)
				select {
				case <-ctx.Done():
					return
				case out <- result:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeValidationResult: func(ctx context.Context, filter *models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *models.DosproxyLogValidationResult)
		sub, err := filter.WatchLogValidationResult(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for log := range ch {
				var result []interface{}
				//txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client
				tx := getTx(log.Raw.TxHash, log.Raw.BlockNumber, log.Raw.BlockHash, log.Raw.Index, proxyAbi, client)
				if tx == nil {
					continue
				}
				result = append(result, tx)
				result = append(result, log)
				select {
				case <-ctx.Done():
					return
				case out <- result:
				}
			}
		}()
		return err, out, sub.Err()
	},
}

func (g *gethRepo) CurrentBlockNum(ctx context.Context) (blknum uint64, err error) {
	var header *types.Header
	header, err = g.client.HeaderByNumber(context.Background(), nil)
	if err == nil {
		blknum = header.Number.Uint64()
	}
	return
}

func (g *gethRepo) Balance(ctx context.Context, hexAddr string) (string, error) {
	if !common.IsHexAddress(hexAddr) {
		return "", errors.New("Not a valid hex address")
	}
	addr := common.HexToAddress(hexAddr)
	wei, err := g.client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return "", err
	}

	balance := new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))
	return balance.String(), nil
}

func caseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}

//ctx context.Context, fromBlockc chan uint64, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer
//FetchLogs(ctx context.Context, logType int, eventc chan []interface{}) (err error, errc chan error)
func (g *gethRepo) FetchLogs(ctx context.Context, logType int, fromBlock, toBlock uint64, blockLimit uint64) (err error, eventc chan []interface{}, errc chan error) {
	if logType >= len(fetchTable) {
		return errors.New("Not support model type"), nil, nil
	}
	out, errc := fetchTable[logType](ctx, fromBlock, toBlock, blockLimit, &g.proxy.Contract.DosproxyFilterer, g.proxyAbi, g.client)
	return nil, out, errc
}

func (g *gethRepo) SubscribeLogs(ctx context.Context, logType int) (err error, eventc chan []interface{}, errc <-chan error) {
	if logType >= len(subscriptionTable) {
		return errors.New("Not support model type"), nil, nil
	}
	return subscriptionTable[logType](ctx, &g.proxy.Contract.DosproxyFilterer, g.proxyAbi, g.client)
}

func getTx(txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client) *models.Transaction {
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
	return &mTx
}
