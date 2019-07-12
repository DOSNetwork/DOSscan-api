package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/DOSNetwork/DOSscan-api/server/repository"
	"github.com/DOSNetwork/DOSscan-api/subscriber/commitreveal"
	"github.com/DOSNetwork/DOSscan-api/subscriber/dosbridge"
	"github.com/DOSNetwork/DOSscan-api/subscriber/dosproxy"

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

var proxyEvent = []string{
	0:  "logurl",
	1:  "logrequestuserrandom",
	2:  "lognonsupportedtype",
	3:  "lognoncontractcall",
	4:  "logcallbacktriggeredfor",
	5:  "logrequestfromnonexistentuc",
	6:  "logupdaterandom",
	7:  "logvalidationresult",
	8:  "loginsufficientpendingnode",
	9:  "loginsufficientworkinggroup",
	10: "loggrouping",
	11: "logpublickeyaccepted",
	12: "logpublickeySuggested",
	13: "loggroupdissolve",
	14: "logregisterednewpendingnode",
	15: "loggroupinginitiated",
	16: "lognopendinggroup",
	17: "logpendinggroupremoved",
	18: "logerror",
	19: "updategrouptopick",
	20: "updategroupsize",
	21: "updategroupingthreshold",
	22: "updategroupmaturityperiod",
	23: "updatebootstrapcommitduration",
	24: "updatebootstraprevealduration",
	25: "updatebootstrapstartthreshold",
	26: "updatependinggroupmaxLife",
	27: "guardianreward",
}

func main() {
	ctx := context.Background()
	db := repository.Connect("postgres", "postgres", "postgres")

	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	if err != nil {
		fmt.Println(":DialToEth err ", err)
		return
	}

RECONNECT:
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
	lastBlk, err := getLatestBlock(client)
	if err != nil {
		fmt.Println("GetLatestBlock err ", err)
		return
	}
	var errcList []chan error
	for idx, event := range proxyEvent {
		fromBc, errc := dosproxy.FromBlockNumber(ctx, event, db)
		errcList = append(errcList, errc)
		outc, errc := dosproxy.FetchTable[idx](ctx, fromBc, lastBlk, &proxy.Contract.DosproxyFilterer)
		errcList = append(errcList, errc)
		dosproxy.ModelsTable[idx](ctx, db, outc, client)

	} //0x919840ad
	//TODO Add CommitReveal Event
	//TODO Add Stacking Event
	//TODO Add AddressBridge Event

	allErrc := mergeErrors(ctx, errcList...)
L:
	for {
		select {
		case <-ctx.Done():
			return
		case e, ok := <-allErrc:
			if ok {
				fmt.Println("errc event err ", e)
			} else {
				break L
			}
		}
	}
	client, err = ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	if err != nil {
		fmt.Println(":DialToEth err ", err)
		return
	}
	goto RECONNECT
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
