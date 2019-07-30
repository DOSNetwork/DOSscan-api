package onchain

import (
	"context"
	"fmt"

	_models "github.com/DOSNetwork/DOSscan-api/models"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var subscriptionTable = []func(ctx context.Context, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error){
	_models.TypeNewPendingNode: func(ctx context.Context, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *_models.DosproxyLogRegisteredNewPendingNode)
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
	_models.TypeGrouping: func(ctx context.Context, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *_models.DosproxyLogGrouping)
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
	_models.TypePublicKeyAccepted: func(ctx context.Context, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *_models.DosproxyLogPublicKeyAccepted)
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
	_models.TypeGroupDissolve: func(ctx context.Context, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *_models.DosproxyLogGroupDissolve)
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
	_models.TypeRequestUserRandom: func(ctx context.Context, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *_models.DosproxyLogRequestUserRandom)
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
	_models.TypeUrl: func(ctx context.Context, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *_models.DosproxyLogUrl)
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
	_models.TypeValidationResult: func(ctx context.Context, filter *_models.DosproxyFilterer, proxyAbi abi.ABI, client *ethclient.Client) (error, chan []interface{}, <-chan error) {
		out := make(chan []interface{})
		ch := make(chan *_models.DosproxyLogValidationResult)
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
