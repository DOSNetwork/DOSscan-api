package repository

import (
	"fmt"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/jinzhu/gorm"
)

//TODO : Change to Method to Events
var events []string

func init() {
	events = []string{"LogURL", "LogRequestUserRandom", "LogNonSupportedType", "LogNonContractCall", "LogCallbackTriggeredFor", "LogRequestFromNonExistentUC",
		"LogUpdateRandom", "LogValidationResult", "LogInsufficientPendingNode", "LogInsufficientWorkingGroup", "LogGrouping", "LogPublicKeyAccepted",
		"LogPublicKeySuggested", "LogGroupDissolve", "LogRegisteredNewPendingNode", "LogGroupingInitiated", "LogNoPendingGroup", "LogPendingGroupRemoved",
		"LogError", "UpdateGroupToPick", "UpdateGroupSize", "UpdateGroupingThreshold", "UpdateGroupMaturityPeriod", "UpdateBootstrapCommitDuration",
		"UpdateBootstrapRevealDuration", "UpdatebootstrapStartThreshold", "UpdatePendingGroupMaxLife", "GuardianReward"}
}

type dbEventsRepo struct {
	db     *gorm.DB
	events []string
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewDBEventsRepository(db *gorm.DB) EventsRepo {
	return &dbEventsRepo{
		db:     db,
		events: []string{},
	}
}

func (d *dbEventsRepo) SetTxRelatedEvents(r []string) {
	d.events = r
}

func checkQuery(query interface{}, args ...interface{}) bool {
	_ = query
	_ = args
	return true
}

//
func (d *dbEventsRepo) GetEvent(limit, offset int, event string, query interface{}, args ...interface{}) []interface{} {
	var resp []interface{}
	if checkQuery(query, args) {
		f := loadEventTable[event]
		if f != nil {
			resp = loadEventTable[event](d.db, limit, offset, query, args...)
		} else {
			fmt.Println("cant find", event)
		}
	}
	return resp
}

//TODO : Should check method name and load corresponging event only
func (d *dbEventsRepo) GetLatestTxEvents(order string, limit int) []interface{} {
	logs := []models.Transaction{}
	var resp []interface{}
	db := d.db
	for _, event := range d.events {
		db = db.Preload(event)
	}

	if err := db.Order(order).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		resp = relatedEvents(logs)
	}

	return resp
}

//TODO : Should check method name and load corresponging event only
func (d *dbEventsRepo) GetEventsByTxAttr(limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.Transaction{}
	var resp []interface{}
	db := d.db
	if checkQuery(query, args) {
		for _, event := range d.events {
			db = db.Preload(event)
		}

		if err := db.Offset(offset).Limit(limit).Where(query, args).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
			resp = relatedEvents(logs)
		}
	}
	return resp
}

type searchEventFunc func(*gorm.DB, int, int, interface{}, ...interface{}) []interface{}

var loadEventTable = map[string]searchEventFunc{
	"logurl":                        searchLogURL,
	"logrequestuserrandom":          searchLogRequestUserRandom,
	"lognonsupportedtype":           searchLogNonSupportedType,
	"lognoncontractcall":            searchLogNonContractCall,
	"logcallbacktriggeredfor":       searchLogCallbackTriggeredFor,
	"logrequestfromnonexistentuc":   searchLogRequestFromNonExistentUC,
	"logupdaterandom":               searchLogUpdateRandom,
	"logvalidationresult":           searchLogValidationResult,
	"loginsufficientpendingnode":    searchLogInsufficientPendingNode,
	"loginsufficientworkinggroup":   searchLogInsufficientWorkingGroup,
	"loggrouping":                   searchLogGrouping,
	"logpublickeyaccepted":          searchLogPublicKeyAccepted,
	"logpublickeysuggested":         searchLogPublicKeySuggested,
	"loggroupdissolve":              searchLogGroupDissolve,
	"logregisterednewpendingnode":   searchLogRegisteredNewPendingNode,
	"loggroupinginitiated":          searchLogGroupingInitiated,
	"lognopendinggroup":             searchLogNoPendingGroup,
	"logpendinggroupremoved":        searchLogPendingGroupRemoved,
	"logerror":                      searchLogError,
	"updategrouptopick":             searchUpdateGroupToPick,
	"updategroupsize":               searchUpdateGroupSize,
	"updategroupingthreshold":       searchUpdateGroupingThreshold,
	"updategroupmaturityperiod":     searchUpdateGroupMaturityPeriod,
	"updatebootstrapcommitduration": searchUpdateBootstrapCommitDuration,
	"updatebootstraprevealduration": searchUpdateBootstrapRevealDuration,
	"updatebootstrapstartthreshold": searchUpdatebootstrapStartThreshold,
	"updatependinggroupmaxlife":     searchUpdatePendingGroupMaxLife,
	"guardianreward":                searchGuardianReward,
}

func relatedEvents(txs []models.Transaction) []interface{} {
	var resp []interface{}
	for _, tx := range txs {
		fmt.Println("blockNum ", tx.BlockNumber)
		for _, event := range tx.LogUrl {
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

func searchLogURL(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogUrl{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
		fmt.Println(len(logs))
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogRequestUserRandom(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogRequestUserRandom{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogNonSupportedType(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogNonSupportedType{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogNonContractCall(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogNonContractCall{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogCallbackTriggeredFor(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogCallbackTriggeredFor{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogRequestFromNonExistentUC(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogRequestFromNonExistentUC{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogUpdateRandom(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogUpdateRandom{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchLogValidationResult(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogValidationResult{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogInsufficientPendingNode(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogInsufficientPendingNode{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchLogInsufficientWorkingGroup(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogInsufficientWorkingGroup{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogGrouping(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogGrouping{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogPublicKeyAccepted(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogPublicKeyAccepted{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogPublicKeySuggested(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogPublicKeySuggested{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchLogGroupDissolve(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogGroupDissolve{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogRegisteredNewPendingNode(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogRegisteredNewPendingNode{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchLogGroupingInitiated(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogGroupingInitiated{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogNoPendingGroup(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogNoPendingGroup{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchLogPendingGroupRemoved(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogPendingGroupRemoved{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchLogError(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.LogError{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchUpdateGroupToPick(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.UpdateGroupToPick{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchUpdateGroupSize(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.UpdateGroupSize{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchUpdateGroupingThreshold(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.UpdateGroupingThreshold{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchUpdateGroupMaturityPeriod(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.UpdateGroupMaturityPeriod{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchUpdateBootstrapCommitDuration(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.UpdateBootstrapCommitDuration{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchUpdateBootstrapRevealDuration(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.UpdateBootstrapRevealDuration{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchUpdatebootstrapStartThreshold(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.UpdatebootstrapStartThreshold{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
func searchUpdatePendingGroupMaxLife(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.UpdatePendingGroupMaxLife{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}

func searchGuardianReward(db *gorm.DB, limit, offset int, query interface{}, args ...interface{}) []interface{} {
	logs := []models.GuardianReward{}
	var resp []interface{}
	if query == nil {
		if err := db.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println(err)
		}
	} else {
		if checkQuery(query, args) {
			if err := db.Offset(offset).Limit(limit).Where(query, args).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}
	}
	for _, log := range logs {
		resp = append(resp, log)
	}
	return resp
}
