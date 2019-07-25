package onchain

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/DOSNetwork/DOSscan-api/models"

	"github.com/ethereum/go-ethereum/ethclient"
)

func initClient(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(url, ":Dial err ", err)
		return nil
	}
	return client
}

func TestGetBalance(t *testing.T) {
	client := initClient("wss://rinkeby.infura.io/ws/v3/3a3e5d776961418e93a8b33fef2f6642")
	r := NewGethRepo(client)
	balance, err := r.GetBalance(context.Background(), "0xba97adc6a7fd3c47579ea3540ee7e7d498cfd820")
	if err != nil {
		t.Errorf("TestGetBalance Error : %s ", err.Error())
	}
	fmt.Println(balance)
	if _, err := strconv.ParseFloat(balance, 10); err != nil || balance == "" {
		t.Errorf("TestGetBalance Error : balance not a valid balance %s ", balance)
	}
}

func TestFetchTable(t *testing.T) {
	client := initClient("wss://rinkeby.infura.io/ws/v3/3a3e5d776961418e93a8b33fef2f6642")
	r := NewGethRepo(client)
	err, logs, _ := r.FetchLogs(context.Background(), models.TypeNewPendingNode, 4468430, 4468435, 1000)
	if err != nil {
		t.Errorf("TestFetchTable Error : %s ", err.Error())
	}
	for log := range logs {
		event, ok := log.(*models.DosproxyLogRegisteredNewPendingNode)
		if !ok {
			continue
		}
		fmt.Println(event.Node)
		fmt.Println(event.Raw.BlockNumber)
	}
	fmt.Println("Done test ")
}

func TestSubscribeTable(t *testing.T) {
	client := initClient("wss://rinkeby.infura.io/ws/v3/3a3e5d776961418e93a8b33fef2f6642")
	r := NewGethRepo(client)
	err, logs, _ := r.SubscribeLogs(context.Background(), models.TypeUrl)
	if err != nil {
		t.Errorf("TestSubscribeTable Error : %s ", err.Error())
	}
	for log := range logs {
		event, ok := log.(*models.DosproxyLogUrl)
		if !ok {
			continue
		}
		fmt.Println(event.QueryId)
		fmt.Println(event.Raw.BlockNumber)
	}
	fmt.Println("Done test ")
}
