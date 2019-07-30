package onchain

import (
	"context"
	"fmt"

	_models "github.com/DOSNetwork/DOSscan-api/models"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

var fetchTable = []func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error){
	_models.TypeNewPendingNode: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
					mLog := &_models.LogRegisteredNewPendingNode{
						Event: _models.Event{
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
	_models.TypeGrouping: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
					mLog := &_models.LogGrouping{
						Event: _models.Event{
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
	_models.TypePublicKeySuggested: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
		out := make(chan []interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogPublicKeySuggested")
			if toBlock <= fromBlock {
				return
			}
			for fromBlock <= toBlock {
				nextBlk := toBlock
				if nextBlk-fromBlock > blockLimit {
					nextBlk = fromBlock + blockLimit
				}
				fmt.Println("LogPublicKeySuggested ", fromBlock, " - ", nextBlk)
				//get the historic data from proxy that start from lastBlkNum to latest
				logs, err := filter.FilterLogPublicKeySuggested(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogPublicKeySuggested err ", err)
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
					mLog := &_models.LogPublicKeySuggested{
						Event: _models.Event{
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
	_models.TypePublicKeyAccepted: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
					mLog := &_models.LogPublicKeyAccepted{
						Event: _models.Event{
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
	_models.TypeGroupDissolve: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
					mLog := &_models.LogGroupDissolve{
						Event: _models.Event{
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
	_models.TypeRequestUserRandom: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
					mLog := &_models.LogRequestUserRandom{
						Event: _models.Event{
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
	_models.TypeUpdateRandom: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
				fmt.Println("LogUpdateRandom ", fromBlock, " - ", nextBlk)
				//1) get the historic data from proxy that start from lastBlkNum to latest
				logs, err := filter.FilterLogUpdateRandom(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
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
					mLog := &_models.LogUpdateRandom{
						Event: _models.Event{
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
	_models.TypeUrl: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
					mLog := &_models.LogUrl{
						Event: _models.Event{
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
	_models.TypeValidationResult: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
					mLog := &_models.LogValidationResult{
						Event: _models.Event{
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
						Message:     log.Message,
						RequestType: log.TrafficType,
						RequestId:   hexutil.Encode(log.TrafficId.Bytes()),
						Signature:   []string{hexutil.Encode(log.Signature[0].Bytes()), hexutil.Encode(log.Signature[1].Bytes())},
						PubKey:      []string{hexutil.Encode(log.PubKey[0].Bytes()), hexutil.Encode(log.PubKey[1].Bytes()), hexutil.Encode(log.PubKey[2].Bytes()), hexutil.Encode(log.PubKey[3].Bytes())},
						Pass:        log.Pass,
					}
					if mLog.RequestId == "0xa49a38aa1c69090e9d4927535b3be2dfe027eb47190dd7809511e6e26a317934" {
						fmt.Println(len(mLog.Message))
						fmt.Println(string(mLog.Message))
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
	_models.TypeGuardianReward: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
				logs, err := filter.FilterGuardianReward(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx}, nil)
				if err != nil {
					fmt.Println("FilterGuardianReward err ", err)
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
					mLog := &_models.GuardianReward{
						Event: _models.Event{
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
	_models.TypeCallbackTriggeredFor: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
				logs, err := filter.FilterLogCallbackTriggeredFor(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogCallbackTriggeredFor err ", err)
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
					mLog := &_models.LogCallbackTriggeredFor{
						Event: _models.Event{
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
	_models.TypeError: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (chan []interface{}, chan error) {
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
				logs, err := filter.FilterLogError(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogError err ", err)
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
					mLog := &_models.LogError{
						Event: _models.Event{
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
