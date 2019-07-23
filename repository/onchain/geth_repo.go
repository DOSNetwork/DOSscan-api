package onchain

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/repository"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type gethRepo struct {
	client *ethclient.Client
	proxy  *models.DosproxySession
	bridge *models.DosbridgeSession
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

	return &gethRepo{
		client: client,
		proxy:  proxy,
		bridge: bridge,
	}
}

var fetchTable = []func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer) (chan interface{}, chan error){
	models.TypeNewPendingNode: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
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
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = nextBlk + 1

			}
		}()
		return out, errc
	},
	models.TypeGrouping: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
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
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = nextBlk + 1
			}
			fmt.Println("LogGrouping Done fetch")

		}()
		return out, errc
	},
	models.TypePublicKeyAccepted: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
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
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = nextBlk + 1
			}
		}()
		return out, errc
	},
	models.TypeGroupDissolve: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
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
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = nextBlk + 1

			}

		}()
		return out, errc
	},
	models.TypeRequestUserRandom: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
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
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = nextBlk + 1

			}
		}()
		return out, errc
	},
	models.TypeUrl: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
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
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = nextBlk + 1
			}
			fmt.Println("LogUrl Done")

		}()
		return out, errc
	},
	models.TypeValidationResult: func(ctx context.Context, fromBlock, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
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
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = nextBlk + 1
			}
		}()
		return out, errc
	},
}

var subscriptionTable = []func(ctx context.Context, filter *models.DosproxyFilterer) (error, chan interface{}, <-chan error){
	models.TypeNewPendingNode: func(ctx context.Context, filter *models.DosproxyFilterer) (error, chan interface{}, <-chan error) {
		out := make(chan interface{})
		ch := make(chan *models.DosproxyLogRegisteredNewPendingNode)
		sub, err := filter.WatchLogRegisteredNewPendingNode(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for event := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- event:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeGrouping: func(ctx context.Context, filter *models.DosproxyFilterer) (error, chan interface{}, <-chan error) {
		out := make(chan interface{})
		ch := make(chan *models.DosproxyLogGrouping)
		sub, err := filter.WatchLogGrouping(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for event := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- event:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypePublicKeyAccepted: func(ctx context.Context, filter *models.DosproxyFilterer) (error, chan interface{}, <-chan error) {
		out := make(chan interface{})
		ch := make(chan *models.DosproxyLogPublicKeyAccepted)
		sub, err := filter.WatchLogPublicKeyAccepted(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for event := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- event:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeGroupDissolve: func(ctx context.Context, filter *models.DosproxyFilterer) (error, chan interface{}, <-chan error) {
		out := make(chan interface{})
		ch := make(chan *models.DosproxyLogGroupDissolve)
		sub, err := filter.WatchLogGroupDissolve(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for event := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- event:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeRequestUserRandom: func(ctx context.Context, filter *models.DosproxyFilterer) (error, chan interface{}, <-chan error) {
		out := make(chan interface{})
		ch := make(chan *models.DosproxyLogRequestUserRandom)
		sub, err := filter.WatchLogRequestUserRandom(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for event := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- event:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeUrl: func(ctx context.Context, filter *models.DosproxyFilterer) (error, chan interface{}, <-chan error) {
		out := make(chan interface{})
		ch := make(chan *models.DosproxyLogUrl)
		sub, err := filter.WatchLogUrl(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for event := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- event:
				}
			}
		}()
		return err, out, sub.Err()
	},
	models.TypeValidationResult: func(ctx context.Context, filter *models.DosproxyFilterer) (error, chan interface{}, <-chan error) {
		out := make(chan interface{})
		ch := make(chan *models.DosproxyLogValidationResult)
		sub, err := filter.WatchLogValidationResult(&bind.WatchOpts{Context: ctx}, ch)
		if err != nil {
			return err, nil, nil
		}
		go func() {
			defer close(out)
			for event := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- event:
				}
			}
		}()
		return err, out, sub.Err()
	},
}

func (g *gethRepo) GetBalance(ctx context.Context, hexAddr string) (string, error) {
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

//ctx context.Context, fromBlockc chan uint64, toBlock uint64, blockLimit uint64, filter *models.DosproxyFilterer
//FetchLogs(ctx context.Context, logType int, eventc chan []interface{}) (err error, errc chan error)
func (g *gethRepo) FetchLogs(ctx context.Context, logType int, fromBlock, toBlock uint64, blockLimit uint64) (err error, eventc chan interface{}, errc chan error) {
	if logType >= len(fetchTable) {
		return errors.New("Not support model type"), nil, nil
	}
	out, errc := fetchTable[logType](ctx, fromBlock, toBlock, blockLimit, &g.proxy.Contract.DosproxyFilterer)
	return nil, out, errc
}

func (g *gethRepo) SubscribeLogs(ctx context.Context, logType int) (err error, eventc chan interface{}, errc <-chan error) {
	if logType >= len(subscriptionTable) {
		return errors.New("Not support model type"), nil, nil
	}
	return subscriptionTable[logType](ctx, &g.proxy.Contract.DosproxyFilterer)
}
