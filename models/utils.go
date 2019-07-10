package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type loadEventFunc func(int, int, *gorm.DB) []interface{}

//TODO : Change to Method to Events
var events []string

func init() {
	events = []string{"LogURL", "LogRequestUserRandom", "LogNonSupportedType", "LogNonContractCall", "LogCallbackTriggeredFor", "LogRequestFromNonExistentUC",
		"LogUpdateRandom", "LogValidationResult", "LogInsufficientPendingNode", "LogInsufficientWorkingGroup", "LogGrouping", "LogPublicKeyAccepted",
		"LogPublicKeySuggested", "LogGroupDissolve", "LogRegisteredNewPendingNode", "LogGroupingInitiated", "LogNoPendingGroup", "LogPendingGroupRemoved",
		"LogError", "UpdateGroupToPick", "UpdateGroupSize", "UpdateGroupingThreshold", "UpdateGroupMaturityPeriod", "UpdateBootstrapCommitDuration",
		"UpdateBootstrapRevealDuration", "UpdatebootstrapStartThreshold", "UpdatePendingGroupMaxLife", "GuardianReward"}
}

var LoadEventTable = map[string]loadEventFunc{
	"logurl":                        loadLogURL,
	"logrequestuserrandom":          loadLogRequestUserRandom,
	"lognonsupportedtype":           loadLogNonSupportedType,
	"lognoncontractcall":            loadLogNonContractCall,
	"logcallbacktriggeredfor":       loadLogCallbackTriggeredFor,
	"logrequestfromnonexistentuc":   loadLogRequestFromNonExistentUC,
	"logupdaterandom":               loadLogUpdateRandom,
	"logvalidationresult":           loadLogValidationResult,
	"loginsufficientpendingnode":    loadLogInsufficientPendingNode,
	"loginsufficientworkinggroup":   loadLogInsufficientWorkingGroup,
	"loggrouping":                   loadLogGrouping,
	"logpublickeyaccepted":          loadLogPublicKeyAccepted,
	"logpublickeysuggested":         loadLogPublicKeySuggested,
	"loggroupdissolve":              loadLogGroupDissolve,
	"logregisterednewpendingnode":   loadLogRegisteredNewPendingNode,
	"loggroupinginitiated":          loadLogGroupingInitiated,
	"lognopendinggroup":             loadLogNoPendingGroup,
	"logpendinggroupremoved":        loadLogPendingGroupRemoved,
	"logerror":                      loadLogError,
	"updategrouptopick":             loadUpdateGroupToPick,
	"updategroupsize":               loadUpdateGroupSize,
	"updategroupingthreshold":       loadUpdateGroupingThreshold,
	"updategroupmaturityperiod":     loadUpdateGroupMaturityPeriod,
	"updatebootstrapcommitduration": loadUpdateBootstrapCommitDuration,
	"updatebootstraprevealduration": loadUpdateBootstrapRevealDuration,
	"updatebootstrapstartthreshold": loadUpdatebootstrapStartThreshold,
	"updatependinggroupmaxlife":     loadUpdatePendingGroupMaxLife,
	"guardianreward":                loadGuardianReward,
}

func relatedEvents(txs []Transaction) []interface{} {
	var resp []interface{}
	for _, tx := range txs {
		for _, event := range tx.LogURL {
			resp = append(resp, event)
		}
		for _, event := range tx.LogRequestUserRandom {
			resp = append(resp, event)
		}
		for _, event := range tx.LogNonSupportedType {
			resp = append(resp, event)
		}
		for _, event := range tx.LogNonContractCall {
			resp = append(resp, event)
		}
		for _, event := range tx.LogCallbackTriggeredFor {
			resp = append(resp, event)
		}
		for _, event := range tx.LogRequestFromNonExistentUC {
			resp = append(resp, event)
		}
		for _, event := range tx.LogUpdateRandom {
			resp = append(resp, event)
		}
		for _, event := range tx.LogValidationResult {
			resp = append(resp, event)
		}
		for _, event := range tx.LogInsufficientPendingNode {
			resp = append(resp, event)
		}
		for _, event := range tx.LogInsufficientWorkingGroup {
			resp = append(resp, event)
		}
		for _, event := range tx.LogGrouping {
			resp = append(resp, event)
		}
		for _, event := range tx.LogPublicKeyAccepted {
			resp = append(resp, event)
		}
		for _, event := range tx.LogPublicKeySuggested {
			resp = append(resp, event)
		}
		for _, event := range tx.LogGroupDissolve {
			resp = append(resp, event)
		}
		for _, event := range tx.LogRegisteredNewPendingNode {
			resp = append(resp, event)
		}
		for _, event := range tx.LogGroupingInitiated {
			resp = append(resp, event)
		}
		for _, event := range tx.LogNoPendingGroup {
			resp = append(resp, event)
		}
		for _, event := range tx.LogPendingGroupRemoved {
			resp = append(resp, event)
		}
		for _, event := range tx.LogError {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateGroupToPick {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateGroupSize {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateGroupingThreshold {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateGroupMaturityPeriod {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateBootstrapCommitDuration {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateBootstrapRevealDuration {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdatebootstrapStartThreshold {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdatePendingGroupMaxLife {
			resp = append(resp, event)
		}
		for _, event := range tx.GuardianReward {
			resp = append(resp, event)
		}
	}
	return resp
}

//TODO : Should check method name and load corresponging event only
func SearchRelatedEvents(limit int, field, condition string, db *gorm.DB) []interface{} {
	logs := []Transaction{}
	var resp []interface{}
	if field == "sender" || field == "hash" || field == "method" {
		for _, event := range events {
			db = db.Preload(event)
		}

		if err := db.Where(field+" ILIKE ?", "%"+condition+"%").Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
			fmt.Println("searchTx ", len(logs))
			resp = relatedEvents(logs)
		}
	}
	return resp
}

func loadLogURL(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogURL{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogRequestUserRandom(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogRequestUserRandom{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogNonSupportedType(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogNonSupportedType{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogNonContractCall(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogNonContractCall{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogCallbackTriggeredFor(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogCallbackTriggeredFor{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogRequestFromNonExistentUC(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogRequestFromNonExistentUC{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogUpdateRandom(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogUpdateRandom{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadLogValidationResult(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogValidationResult{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogInsufficientPendingNode(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogInsufficientPendingNode{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadLogInsufficientWorkingGroup(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogInsufficientWorkingGroup{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogGrouping(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogGrouping{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogPublicKeyAccepted(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogPublicKeyAccepted{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogPublicKeySuggested(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogPublicKeySuggested{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadLogGroupDissolve(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogGroupDissolve{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogRegisteredNewPendingNode(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogRegisteredNewPendingNode{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadLogGroupingInitiated(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogGroupingInitiated{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogNoPendingGroup(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogNoPendingGroup{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadLogPendingGroupRemoved(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogPendingGroupRemoved{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadLogError(limit, offset int, db *gorm.DB) []interface{} {
	logs := []LogError{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadUpdateGroupToPick(limit, offset int, db *gorm.DB) []interface{} {
	logs := []UpdateGroupToPick{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadUpdateGroupSize(limit, offset int, db *gorm.DB) []interface{} {
	logs := []UpdateGroupSize{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadUpdateGroupingThreshold(limit, offset int, db *gorm.DB) []interface{} {
	logs := []UpdateGroupingThreshold{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadUpdateGroupMaturityPeriod(limit, offset int, db *gorm.DB) []interface{} {
	logs := []UpdateGroupMaturityPeriod{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadUpdateBootstrapCommitDuration(limit, offset int, db *gorm.DB) []interface{} {
	logs := []UpdateBootstrapCommitDuration{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadUpdateBootstrapRevealDuration(limit, offset int, db *gorm.DB) []interface{} {
	logs := []UpdateBootstrapRevealDuration{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadUpdatebootstrapStartThreshold(limit, offset int, db *gorm.DB) []interface{} {
	logs := []UpdatebootstrapStartThreshold{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
func loadUpdatePendingGroupMaxLife(limit, offset int, db *gorm.DB) []interface{} {
	logs := []UpdatePendingGroupMaxLife{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}

func loadGuardianReward(limit, offset int, db *gorm.DB) []interface{} {
	logs := []GuardianReward{}
	var resp []interface{}
	if err := db.Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		for _, log := range logs {
			resp = append(resp, log)
		}
	}
	return resp
}
