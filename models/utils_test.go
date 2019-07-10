package models

import (
	"fmt"

	"reflect"
	"testing"
)

func TestLoadEventTable(t *testing.T) {
	db := Connect()
	r := LoadEventTable["logurl"](2, 0, db)
	fmt.Println(reflect.TypeOf(r[0]))
	if reflect.TypeOf(r[0]).String() != "models.LogURL" {
		t.Errorf("TestLoadEventTable Error : %s.", reflect.TypeOf(r[0]))
	}
	r = LoadEventTable["guardianreward"](2, 0, db)
	fmt.Println(reflect.TypeOf(r[0]))
	if reflect.TypeOf(r[0]).String() != "models.GuardianReward" {
		t.Errorf("TestLoadEventTable Error : %s.", reflect.TypeOf(r[0]))
	}
}

func TestSearchTx(t *testing.T) {
	db := Connect()
	events := []string{"LogURL", "LogRequestUserRandom", "LogNonSupportedType", "LogNonContractCall", "LogCallbackTriggeredFor", "LogRequestFromNonExistentUC",
		"LogUpdateRandom", "LogValidationResult", "LogInsufficientPendingNode", "LogInsufficientWorkingGroup", "LogGrouping", "LogPublicKeyAccepted",
		"LogPublicKeySuggested", "LogGroupDissolve", "LogRegisteredNewPendingNode", "LogGroupingInitiated", "LogNoPendingGroup", "LogPendingGroupRemoved",
		"LogError", "UpdateGroupToPick", "UpdateGroupSize", "UpdateGroupingThreshold", "UpdateGroupMaturityPeriod", "UpdateBootstrapCommitDuration",
		"UpdateBootstrapRevealDuration", "UpdatebootstrapStartThreshold", "UpdatePendingGroupMaxLife", "GuardianReward"}
	r := SearchEventsByTx(200, events, "ab", db)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestSearchMethod(t *testing.T) {
	db := Connect()
	events := []string{"LogURL", "LogRequestUserRandom", "LogNonSupportedType", "LogNonContractCall", "LogCallbackTriggeredFor", "LogRequestFromNonExistentUC",
		"LogUpdateRandom", "LogValidationResult", "LogInsufficientPendingNode", "LogInsufficientWorkingGroup", "LogGrouping", "LogPublicKeyAccepted",
		"LogPublicKeySuggested", "LogGroupDissolve", "LogRegisteredNewPendingNode", "LogGroupingInitiated", "LogNoPendingGroup", "LogPendingGroupRemoved",
		"LogError", "UpdateGroupToPick", "UpdateGroupSize", "UpdateGroupingThreshold", "UpdateGroupMaturityPeriod", "UpdateBootstrapCommitDuration",
		"UpdateBootstrapRevealDuration", "UpdatebootstrapStartThreshold", "UpdatePendingGroupMaxLife", "GuardianReward"}
	r := SearchEventsByMethod(200, events, "signalGroupFormation", db)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}
