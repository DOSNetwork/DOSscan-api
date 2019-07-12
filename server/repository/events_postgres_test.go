package repository

import (
	"fmt"

	"github.com/DOSNetwork/DOSscan-api/models"

	"reflect"
	"testing"
)

func TestEventsByTxAttr(t *testing.T) {
	events = []string{"LogUrl", "LogRequestUserRandom", "LogNonSupportedType", "LogNonContractCall", "LogCallbackTriggeredFor", "LogRequestFromNonExistentUC",
		"LogUpdateRandom", "LogValidationResult", "LogInsufficientPendingNode", "LogInsufficientWorkingGroup", "LogGrouping", "LogPublicKeyAccepted",
		"LogPublicKeySuggested", "LogGroupDissolve", "LogRegisteredNewPendingNode", "LogGroupingInitiated", "LogNoPendingGroup", "LogPendingGroupRemoved",
		"LogError", "UpdateGroupToPick", "UpdateGroupSize", "UpdateGroupingThreshold", "UpdateGroupMaturityPeriod", "UpdateBootstrapCommitDuration",
		"UpdateBootstrapRevealDuration", "UpdatebootstrapStartThreshold", "UpdatePendingGroupMaxLife", "GuardianReward"}
	db := Connect("postgres", "postgres", "postgres")
	repo := NewDBEventsRepository(db)
	repo.SetTxRelatedEvents(events)
	r := repo.GetEventsByTxAttr(100, 0, "sender = ?", "0xCdD9759439dF580FF183414C491F27E852Ac6240")
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestLatestEvents(t *testing.T) {
	events = []string{"LogUrl", "LogRequestUserRandom", "LogNonSupportedType", "LogNonContractCall", "LogCallbackTriggeredFor", "LogRequestFromNonExistentUC",
		"LogUpdateRandom", "LogValidationResult", "LogInsufficientPendingNode", "LogInsufficientWorkingGroup", "LogGrouping", "LogPublicKeyAccepted",
		"LogPublicKeySuggested", "LogGroupDissolve", "LogRegisteredNewPendingNode", "LogGroupingInitiated", "LogNoPendingGroup", "LogPendingGroupRemoved",
		"LogError", "UpdateGroupToPick", "UpdateGroupSize", "UpdateGroupingThreshold", "UpdateGroupMaturityPeriod", "UpdateBootstrapCommitDuration",
		"UpdateBootstrapRevealDuration", "UpdatebootstrapStartThreshold", "UpdatePendingGroupMaxLife", "GuardianReward"}
	db := Connect("postgres", "postgres", "postgres")
	repo := NewDBEventsRepository(db)
	repo.SetTxRelatedEvents(events)
	r := repo.GetLatestTxEvents("block_number desc", 20)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestGetEvent(t *testing.T) {
	db := Connect("postgres", "postgres", "postgres")
	repo := NewDBEventsRepository(db)
	r := repo.GetEvent(100, 0, "logrequestuserrandom", "request_id ILIKE ?", "0x4044d38f2d10951cf0619848e88965c1c3dca96f02bfbd14f403d59e4a25bf53")
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}
func TestGetEventCount(t *testing.T) {
	db := Connect("postgres", "postgres", "postgres")
	var count int

	db.Model(&models.LogUrl{}).Count(&count)
	fmt.Println(count)
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
