package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/DOSNetwork/explorer-Api/listener/commitreveal"
	"github.com/DOSNetwork/explorer-Api/listener/db"
	"github.com/DOSNetwork/explorer-Api/listener/dosbridge"
	"github.com/DOSNetwork/explorer-Api/listener/dosproxy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

func newCR(crAddr common.Address, client *ethclient.Client) (*commitreveal.CommitrevealSession, error) {
	p, err := commitreveal.NewCommitreveal(crAddr, client)
	if err != nil {
		return nil, err
	}
	return &commitreveal.CommitrevealSession{Contract: p, CallOpts: bind.CallOpts{Context: context.Background()}}, nil
}

var proxyEvent = []string{
	0:  "LogUrl",
	1:  "LogRequestUserRandom",
	2:  "LogNonSupportedType",
	3:  "LogNonContractCall",
	4:  "LogCallbackTriggeredFor",
	5:  "LogRequestFromNonExistentUC",
	6:  "LogUpdateRandom",
	7:  "LogValidationResult",
	8:  "LogInsufficientPendingNode",
	9:  "LogInsufficientWorkingGroup",
	10: "Grouping",
	11: "LogPublicKeyAccepted",
	12: "LogPublicKeySuggested",
	13: "LogGroupDissolve",
	14: "LogRegisteredNewPendingNode",
	15: "LogGroupingInitiated",
	16: "LogNoPendingGroup",
	17: "LogPendingGroupRemoved",
	18: "LogError",
	19: "UpdateGroupToPick",
	20: "UpdateGroupSize",
	21: "UpdateGroupingThreshold",
	22: "UpdateGroupMaturityPeriod",
	23: "UpdateBootstrapCommitDuration",
	24: "UpdateBootstrapRevealDuration",
	25: "UpdatebootstrapStartThreshold",
	26: "UpdatePendingGroupMaxLife",
	27: "GuardianReward",
}

func main() {
	ctx := context.Background()

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	dbConn, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(":DialToEth err ", err)
		return
	}
	defer dbConn.Close()

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

	var errcList []chan error
	for idx, event := range proxyEvent {
		lastBc, errc := db.LastBlk(ctx, event, dbConn)
		errcList = append(errcList, errc)
		out, errc := dosproxy.FetchTable[idx](ctx, lastBc, &proxy.Contract.DosproxyFilterer)
		errcList = append(errcList, errc)
		errcList = append(errcList, db.ProxyTable[idx](ctx, out, dbConn))
	}
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
