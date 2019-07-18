package repository

import (
	"fmt"

	"context"
	"math"
	"math/big"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
)

//TODO : Change to Method to Events

type dbEventsRepo struct {
	db     *gorm.DB
	client *ethclient.Client
	events []string
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewDBEventsRepository(db *gorm.DB, client *ethclient.Client) EventsRepo {
	return &dbEventsRepo{
		db:     db,
		events: []string{},
		client: client,
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
func (d *dbEventsRepo) CountEvent(event interface{}) (count int) {
	d.db.Model(event).Count(&count)
	return
}

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

func getBalance(client *ethclient.Client, addr common.Address) string {
	wei, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return ""
	}

	balance := new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))

	return balance.String()
}

func (d *dbEventsRepo) GetNode(addr string) []interface{} {
	var node models.Node
	var resp []interface{}
	if err := d.db.Where(models.Node{Addr: addr}).First(&node).Error; !gorm.IsRecordNotFoundError(err) {
		node.Balance = getBalance(d.client, common.HexToAddress(addr))
		node.RegisterState = true
		d.db.Model(&node).Related(&node.Groups, "Groups")
		for _, group := range node.Groups {
			if group.DissolvedBlkNum == 0 && group.AcceptedBlkNum != 0 {
				node.ActiveGroups = append(node.ActiveGroups, group.GroupId)
				fmt.Println("Node ActiveGroup ", node.ActiveGroups, group.AcceptedBlkNum)
			} else if group.DissolvedBlkNum != 0 && group.AcceptedBlkNum != 0 {
				node.ExpiredGroups++
			}
		}
		resp = append(resp, node)
	}
	fmt.Println("Node balance ", node.Balance)
	fmt.Println("Node ActiveGroup ", node.ActiveGroups)
	return resp
}

func (d *dbEventsRepo) GetGroup(groupId string) []interface{} {
	var group models.Group
	var resp []interface{}
	if err := d.db.Where(models.Group{GroupId: groupId}).First(&group).Error; !gorm.IsRecordNotFoundError(err) {

		fmt.Println("Group Id ", group.GroupId)
		fmt.Println("Group AcceptedBlkNum ", group.AcceptedBlkNum)
		fmt.Println("Group DissolvedBlkNum ", group.DissolvedBlkNum)
		fmt.Println("Group NodeId ", group.NodeId)
		d.db.Model(&group).Related(&group.UrlRequests, "UrlRequests")
		group.NumUrl = len(group.UrlRequests)
		d.db.Model(&group).Related(&group.UserRandomRequests, "UserRandomRequests")
		group.NumRandom = len(group.UserRandomRequests)

		resp = append(resp, group)
	}
	return resp
}

func (d *dbEventsRepo) GetRequest(requestId string) []interface{} {
	urlRequest := models.UrlRequest{}
	urlRequest.RequestId = requestId
	var resp []interface{}
	if err := d.db.Where(urlRequest).First(&urlRequest).Error; !gorm.IsRecordNotFoundError(err) {
		fmt.Println("urlRequest Id ", urlRequest.RequestId)
		resp = append(resp, urlRequest)
	}

	return resp
}
func (d *dbEventsRepo) GetRandomRequest(requestId string) []interface{} {
	randomRequest := models.UserRandomRequest{}
	randomRequest.RequestId = requestId
	var resp []interface{}
	if err := d.db.Where(randomRequest).First(&randomRequest).Error; !gorm.IsRecordNotFoundError(err) {
		fmt.Println("randomRequest Id ", randomRequest.RequestId)
		resp = append(resp, randomRequest)
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
		fmt.Println("blockNum ", tx.BlockNumber, tx.Method)
		for _, event := range tx.LogUrls {
			resp = append(resp, event)
		}
		for _, event := range tx.LogRequestUserRandoms {
			resp = append(resp, event)
		}
		for _, event := range tx.LogNonSupportedTypes {
			resp = append(resp, event)
		}
		for _, event := range tx.LogNonContractCalls {
			resp = append(resp, event)
		}
		for _, event := range tx.LogCallbackTriggeredFors {
			resp = append(resp, event)
		}
		for _, event := range tx.LogRequestFromNonExistentUCs {
			resp = append(resp, event)
		}
		for _, event := range tx.LogUpdateRandoms {
			resp = append(resp, event)
		}
		for _, event := range tx.LogValidationResults {
			resp = append(resp, event)
		}
		for _, event := range tx.LogInsufficientPendingNodes {
			resp = append(resp, event)
		}
		for _, event := range tx.LogInsufficientWorkingGroups {
			resp = append(resp, event)
		}
		for _, event := range tx.LogGroupings {
			resp = append(resp, event)
		}
		for _, event := range tx.LogPublicKeyAccepteds {
			resp = append(resp, event)
		}
		for _, event := range tx.LogPublicKeySuggesteds {
			resp = append(resp, event)
		}
		for _, event := range tx.LogGroupDissolves {
			resp = append(resp, event)
		}
		for _, event := range tx.LogRegisteredNewPendingNodes {
			resp = append(resp, event)
		}
		for _, event := range tx.LogGroupingInitiateds {
			resp = append(resp, event)
		}
		for _, event := range tx.LogNoPendingGroups {
			resp = append(resp, event)
		}
		for _, event := range tx.LogPendingGroupRemoveds {
			resp = append(resp, event)
		}
		for _, event := range tx.LogErrors {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateGroupToPicks {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateGroupSizes {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateGroupingThresholds {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateGroupMaturityPeriods {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateBootstrapCommitDurations {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdateBootstrapRevealDurations {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdatebootstrapStartThresholds {
			resp = append(resp, event)
		}
		for _, event := range tx.UpdatePendingGroupMaxLifes {
			resp = append(resp, event)
		}
		for _, event := range tx.GuardianRewards {
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
