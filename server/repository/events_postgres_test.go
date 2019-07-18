package repository

import (
	"fmt"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/ethereum/go-ethereum/ethclient"

	"reflect"
	"testing"
)

func TestEventsByTxAttr(t *testing.T) {
	events := []string{"LogUrl", "LogRequestUserRandom", "LogNonSupportedType", "LogNonContractCall", "LogCallbackTriggeredFor", "LogRequestFromNonExistentUC",
		"LogUpdateRandom", "LogValidationResult", "LogInsufficientPendingNode", "LogInsufficientWorkingGroup", "LogGrouping", "LogPublicKeyAccepted",
		"LogPublicKeySuggested", "LogGroupDissolve", "LogRegisteredNewPendingNode", "LogGroupingInitiated", "LogNoPendingGroup", "LogPendingGroupRemoved",
		"LogError", "UpdateGroupToPick", "UpdateGroupSize", "UpdateGroupingThreshold", "UpdateGroupMaturityPeriod", "UpdateBootstrapCommitDuration",
		"UpdateBootstrapRevealDuration", "UpdatebootstrapStartThreshold", "UpdatePendingGroupMaxLife", "GuardianReward"}
	db := Connect("postgres", "postgres", "postgres")
	client, _ := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	repo := NewDBEventsRepository(db, client)
	repo.SetTxRelatedEvents(events)
	r := repo.GetEventsByTxAttr(100, 0, "sender = ?", "0xCdD9759439dF580FF183414C491F27E852Ac6240")
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestLatestEvents(t *testing.T) {
	events := []string{"LogUrls", "LogRequestUserRandoms", "LogCallbackTriggeredFors",
		"LogUpdateRandoms", "LogValidationResults", "LogGroupings", "LogPublicKeyAccepteds",
		"LogPublicKeySuggesteds", "LogGroupDissolves", "LogRegisteredNewPendingNodes", "GuardianRewards"}
	db := Connect("postgres", "postgres", "postgres")
	client, _ := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	repo := NewDBEventsRepository(db, client)
	repo.SetTxRelatedEvents(events)
	r := repo.GetLatestTxEvents("block_number desc", 20)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestGetEvent(t *testing.T) {
	db := Connect("postgres", "postgres", "postgres")
	client, _ := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	repo := NewDBEventsRepository(db, client)
	r := repo.GetEvent(100, 0, "logrequestuserrandom", "request_id ILIKE ?", "0x4044d38f2d10951cf0619848e88965c1c3dca96f02bfbd14f403d59e4a25bf53")
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestGetEventCount(t *testing.T) {
	db := Connect("postgres", "postgres", "postgres")
	client, _ := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	repo := NewDBEventsRepository(db, client)
	fmt.Println(repo.CountEvent(&models.LogRegisteredNewPendingNode{}))
}

func TestGetNode(t *testing.T) {
	db := Connect("postgres", "postgres", "postgres")
	client, _ := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	repo := NewDBEventsRepository(db, client)
	resp := repo.GetNode("0xfEa0E757b1C3c2Cd7169CC1889069B26806F2b45")
	fmt.Println(len(resp))
}
func TestGetGroup(t *testing.T) {
	db := Connect("postgres", "postgres", "postgres")
	client, _ := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	repo := NewDBEventsRepository(db, client)
	resp := repo.GetGroup("0x888ad3cbe70d83127dd6fdedb1a157e2d9cc44bdaa15d61803081ac25bf392f6")
	fmt.Println(len(resp))
}
func TestGetRequest(t *testing.T) {
	db := Connect("postgres", "postgres", "postgres")
	client, _ := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	repo := NewDBEventsRepository(db, client)
	resp := repo.GetRequest("0x5c42e9d87d2406a855081c0524a9d1fb50e915a7becf6ade793e2433a122339d")
	fmt.Println(len(resp))
}

/*
func TestSearchTx(t *testing.T) {
	db := Connect()
	r := SearchRelatedEvents(200, "hash", "385f702381b6a19b0c5e636b", db)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestSearchMethod(t *testing.T) {
	db := Connect()
	r := SearchRelatedEvents(200, "method", "signalGroupFormation", db)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestSearchAddr(t *testing.T) {
	db := Connect()
	r := SearchRelatedEvents(200, "sender", "0xCdD9759439dF580FF183414C491F27E852Ac6240", db)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}
*/
