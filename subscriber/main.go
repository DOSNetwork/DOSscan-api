package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/DOSNetwork/DOSscan-api/server/repository"
	"github.com/DOSNetwork/DOSscan-api/subscriber/commitreveal"
	"github.com/DOSNetwork/DOSscan-api/subscriber/dosbridge"
	"github.com/DOSNetwork/DOSscan-api/subscriber/dosproxy"
	"github.com/jinzhu/gorm"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func newCR(crAddr common.Address, client *ethclient.Client) (*commitreveal.CommitrevealSession, error) {
	p, err := commitreveal.NewCommitreveal(crAddr, client)
	if err != nil {
		return nil, err
	}
	return &commitreveal.CommitrevealSession{Contract: p, CallOpts: bind.CallOpts{Context: context.Background()}}, nil
}

func fetchHistoricData(proxy *dosproxy.DosproxySession, db *gorm.DB, client *ethclient.Client) error {
	ctx := context.Background()
	lastBlk, err := getLatestBlock(client)
	if err != nil {
		fmt.Println("GetLatestBlock err ", err)
		return err
	}
	fmt.Println("lastBlk : ", lastBlk)
	var errcList []chan error
	for idx, event := range dosproxy.ProxyEvent {
		fromBc, errc := dosproxy.FromBlockNumber(ctx, event, db)
		errcList = append(errcList, errc)
		outc, errc := dosproxy.FetchTable[idx](ctx, fromBc, lastBlk, 1000, &proxy.Contract.DosproxyFilterer)
		errcList = append(errcList, errc)
		errc = dosproxy.ModelsTable[idx](ctx, db, outc, client)
		errcList = append(errcList, errc)
	}
	allErrc := mergeErrors(ctx, errcList...)
	for {
		select {
		case <-ctx.Done():
			return nil
		case e, ok := <-allErrc:
			if ok {
				fmt.Println("errc event err ", e)
				return e
			} else {
				return nil
			}
		}
	}
}

func main() {
	ctx := context.Background()
	db := repository.Connect("postgres", "postgres", "postgres")

	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	if err != nil {
		fmt.Println(":DialToEth err ", err)
		return
	}

	//RECONNECT:
	d, err := dosbridge.NewDosbridge(common.HexToAddress("0xf0CEFfc4209e38EA3Cd1926DDc2bC641cbFFd1cF"), client)
	if err != nil {
		fmt.Println("NewDosbridge err ", err)
		return
	}

	bridge := &dosbridge.DosbridgeSession{Contract: d, CallOpts: bind.CallOpts{Context: ctx}}
	proxyAddr, err := bridge.GetProxyAddress()
	if err != nil {
		return
	}

	proxy, err := dosproxy.NewProxy(proxyAddr, client)
	if err != nil {
		fmt.Println("newProxy err ", err)
		return
	}

	err = fetchHistoricData(proxy, db, client)
	if err != nil {
		fmt.Println("fetchHistoricData err ", err)
	}
	dosproxy.BuildRelation(db)
	/*
		fmt.Println("!!!!!!!!!!!!!!!!!!Fetch Done")

		client, err = ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
		if err != nil {
			fmt.Println(":DialToEth err ", err)
			return
		}
		goto RECONNECT
	*/
}

func getLatestBlock(client *ethclient.Client) (blknum uint64, err error) {
	var header *types.Header
	header, err = client.HeaderByNumber(context.Background(), nil)
	if err == nil {
		blknum = header.Number.Uint64()
	}
	return
}

func mergeErrors(ctx context.Context, cs ...chan error) chan error {
	var wg sync.WaitGroup
	out := make(chan error)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan error) {
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
