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

	_models "github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/models/dosbridge"
	"github.com/DOSNetwork/DOSscan-api/models/dosproxy"
	_repository "github.com/DOSNetwork/DOSscan-api/repository"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type gethRepo struct {
	client   *ethclient.Client
	proxy    *dosproxy.DosproxySession
	bridge   *dosbridge.DosbridgeSession
	proxyAbi abi.ABI
}

func NewGethRepo(client *ethclient.Client, bridgeAddress string) (_repository.Onchain, error) {
	ctx := context.Background()
	d, err := dosbridge.NewDosbridge(common.HexToAddress(bridgeAddress), client)
	if err != nil {
		fmt.Println("NewDosbridge err ", err)
		return nil, err
	}

	bridge := &dosbridge.DosbridgeSession{Contract: d, CallOpts: bind.CallOpts{Context: ctx}}
	proxyAddr, err := bridge.GetProxyAddress()
	if err != nil {
		return nil, err
	}

	p, err := dosproxy.NewDosproxy(proxyAddr, client)
	if err != nil {
		return nil, err
	}
	proxy := &dosproxy.DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: ctx}}

	jsonFile, err := os.Open("./abi/DOSProxy.abi")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	abiJsonByte, _ := ioutil.ReadAll(jsonFile)
	proxyAbi, err := abi.JSON(strings.NewReader(string(abiJsonByte)))
	if err != nil {
		return nil, err
	}

	return &gethRepo{
		client:   client,
		proxy:    proxy,
		bridge:   bridge,
		proxyAbi: proxyAbi,
	}, nil
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

//ctx context.Context, fromBlockc chan uint64, toBlock uint64, blockLimit uint64, filter *_models.DosproxyFilterer
//FetchLogs(ctx context.Context, logType int, eventc chan []interface{}) (err error, errc chan error)
func (g *gethRepo) FetchLogs(ctx context.Context, logType int, fromBlock, toBlock uint64, blockLimit uint64) (err error, eventc chan []interface{}, errc chan error) {
	if logType >= len(fetchTable) {
		return errors.New("Not support model type"), nil, nil
	}
	out, errc := fetchTable[logType](ctx, fromBlock, toBlock, blockLimit, &g.proxy.Contract.DosproxyFilterer, g.proxyAbi, g.client)
	return nil, out, errc
}

/*
func (g *gethRepo) SubscribeLogs(ctx context.Context, logType int) (err error, eventc chan []interface{}, errc <-chan error) {
	if logType >= len(subscriptionTable) {
		return errors.New("Not support model type"), nil, nil
	}
	return subscriptionTable[logType](ctx, &g.proxy.Contract.DosproxyFilterer, g.proxyAbi, g.client)
}
*/
func getTx(txHash common.Hash, blockNum uint64, blockhash common.Hash, index uint, proxyAbi abi.ABI, client *ethclient.Client) (*_models.Transaction, error) {
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Println("TransactionByHash err", err)
		return nil, err
	}
	sender, err := client.TransactionSender(context.Background(), tx, blockhash, index)
	if err != nil {
		fmt.Println("GetTransactionSender err", err)
		return nil, err
	}
	var methodName string
	if method, err := proxyAbi.MethodById(tx.Data()[:4]); err == nil {
		methodName = method.Name
	} else {
		methodName = fmt.Sprintf("ExternalCall 0x%x", tx.Data()[:4])
	}
	mtx := _models.Transaction{
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
	return &mtx, nil
}
