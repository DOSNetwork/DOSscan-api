package apiv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	eventList int = iota
	nodeList
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	*Body   `json:"body"`
}

type Body struct {
	Events   []interface{} `json:"events,omitempty"`
	Nodelist []interface{} `json:"nodelist,omitempty"`
}

type searchEventFunc func(int, int, *gin.Context) bool

var searchEventTable = map[string]searchEventFunc{
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

func searchLogURL(limit, offset int, c *gin.Context) bool {
	logs := []models.LogURL{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogRequestUserRandom(limit, offset int, c *gin.Context) bool {
	logs := []models.LogRequestUserRandom{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchLogNonSupportedType(limit, offset int, c *gin.Context) bool {
	logs := []models.LogNonSupportedType{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogNonContractCall(limit, offset int, c *gin.Context) bool {
	logs := []models.LogNonContractCall{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogCallbackTriggeredFor(limit, offset int, c *gin.Context) bool {
	logs := []models.LogCallbackTriggeredFor{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogRequestFromNonExistentUC(limit, offset int, c *gin.Context) bool {
	logs := []models.LogRequestFromNonExistentUC{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogUpdateRandom(limit, offset int, c *gin.Context) bool {
	logs := []models.LogUpdateRandom{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchLogValidationResult(limit, offset int, c *gin.Context) bool {
	logs := []models.LogValidationResult{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogInsufficientPendingNode(limit, offset int, c *gin.Context) bool {
	logs := []models.LogInsufficientPendingNode{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchLogInsufficientWorkingGroup(limit, offset int, c *gin.Context) bool {
	logs := []models.LogInsufficientWorkingGroup{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogGrouping(limit, offset int, c *gin.Context) bool {
	logs := []models.LogGrouping{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchLogPublicKeyAccepted(limit, offset int, c *gin.Context) bool {
	logs := []models.LogPublicKeyAccepted{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogPublicKeySuggested(limit, offset int, c *gin.Context) bool {
	logs := []models.LogPublicKeySuggested{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchLogGroupDissolve(limit, offset int, c *gin.Context) bool {
	logs := []models.LogGroupDissolve{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogRegisteredNewPendingNode(limit, offset int, c *gin.Context) bool {
	logs := []models.LogRegisteredNewPendingNode{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchLogGroupingInitiated(limit, offset int, c *gin.Context) bool {
	logs := []models.LogGroupingInitiated{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogNoPendingGroup(limit, offset int, c *gin.Context) bool {
	logs := []models.LogNoPendingGroup{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchLogPendingGroupRemoved(limit, offset int, c *gin.Context) bool {
	logs := []models.LogPendingGroupRemoved{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchLogError(limit, offset int, c *gin.Context) bool {
	logs := []models.LogError{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchUpdateGroupToPick(limit, offset int, c *gin.Context) bool {
	logs := []models.UpdateGroupToPick{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchUpdateGroupSize(limit, offset int, c *gin.Context) bool {
	logs := []models.UpdateGroupSize{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchUpdateGroupingThreshold(limit, offset int, c *gin.Context) bool {
	logs := []models.UpdateGroupingThreshold{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchUpdateGroupMaturityPeriod(limit, offset int, c *gin.Context) bool {
	logs := []models.UpdateGroupMaturityPeriod{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchUpdateBootstrapCommitDuration(limit, offset int, c *gin.Context) bool {
	logs := []models.UpdateBootstrapCommitDuration{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchUpdateBootstrapRevealDuration(limit, offset int, c *gin.Context) bool {
	logs := []models.UpdateBootstrapRevealDuration{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchUpdatebootstrapStartThreshold(limit, offset int, c *gin.Context) bool {
	logs := []models.UpdatebootstrapStartThreshold{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}
func searchUpdatePendingGroupMaxLife(limit, offset int, c *gin.Context) bool {
	logs := []models.UpdatePendingGroupMaxLife{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchGuardianReward(limit, offset int, c *gin.Context) bool {
	logs := []models.GuardianReward{}
	if err := models.DB.Offset(offset).Limit(limit).Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			resp = append(resp, log)
		}
		sendResponse(eventList, resp, c)
		return true
	}
}

func searchTx(limit, offset int, condition string, c *gin.Context) bool {
	logs := []models.Transaction{}
	db := models.DB.Preload("LogURL").Preload("LogRequestUserRandom").Preload("LogNonSupportedType")
	db = db.Preload("LogNonContractCall").Preload("LogCallbackTriggeredFor").Preload("LogRequestFromNonExistentUC")
	db = db.Preload("LogUpdateRandom").Preload("LogValidationResult").Preload("LogInsufficientPendingNode")
	db = db.Preload("LogInsufficientWorkingGroup").Preload("LogGrouping").Preload("LogPublicKeyAccepted")
	db = db.Preload("LogPublicKeySuggested").Preload("LogGroupDissolve").Preload("LogRegisteredNewPendingNode")
	db = db.Preload("LogGroupingInitiated").Preload("LogNoPendingGroup").Preload("LogPendingGroupRemoved")
	db = db.Preload("LogError").Preload("UpdateGroupToPick").Preload("UpdateGroupSize")
	db = db.Preload("UpdateGroupingThreshold").Preload("UpdateGroupMaturityPeriod").Preload("UpdateBootstrapCommitDuration")
	db = db.Preload("UpdateBootstrapRevealDuration").Preload("UpdatebootstrapStartThreshold").Preload("UpdatePendingGroupMaxLife")
	db = db.Preload("GuardianReward")

	if err := db.Where("hash ILIKE ?", "%"+condition+"%").Or("method ILIKE ?", "%"+condition+"%").Find(&logs).Error; gorm.IsRecordNotFoundError(err) {
		fmt.Println(err)
		return false
	} else {
		var resp []interface{}
		for _, log := range logs {
			for _, l := range log.LogURL {
				resp = append(resp, l)
			}
			for _, l := range log.LogRequestUserRandom {
				resp = append(resp, l)
			}
			for _, l := range log.LogNonSupportedType {
				resp = append(resp, l)
			}
			for _, l := range log.LogNonContractCall {
				resp = append(resp, l)
			}
			for _, l := range log.LogCallbackTriggeredFor {
				resp = append(resp, l)
			}
			for _, l := range log.LogRequestFromNonExistentUC {
				resp = append(resp, l)
			}
			for _, l := range log.LogUpdateRandom {
				resp = append(resp, l)
			}
			for _, l := range log.LogValidationResult {
				resp = append(resp, l)
			}
			for _, l := range log.LogInsufficientPendingNode {
				resp = append(resp, l)
			}
			for _, l := range log.LogInsufficientWorkingGroup {
				resp = append(resp, l)
			}
			for _, l := range log.LogGrouping {
				resp = append(resp, l)
			}
			for _, l := range log.LogPublicKeyAccepted {
				resp = append(resp, l)
			}
			for _, l := range log.LogPublicKeySuggested {
				resp = append(resp, l)
			}
			for _, l := range log.LogGroupDissolve {
				resp = append(resp, l)
			}
			for _, l := range log.LogRegisteredNewPendingNode {
				resp = append(resp, l)
			}
			for _, l := range log.LogGroupingInitiated {
				resp = append(resp, l)
			}
			for _, l := range log.LogNoPendingGroup {
				resp = append(resp, l)
			}
			for _, l := range log.LogPendingGroupRemoved {
				resp = append(resp, l)
			}
			for _, l := range log.LogError {
				resp = append(resp, l)
			}
			for _, l := range log.UpdateGroupToPick {
				resp = append(resp, l)
			}
			for _, l := range log.UpdateGroupSize {
				resp = append(resp, l)
			}
			for _, l := range log.UpdateGroupingThreshold {
				resp = append(resp, l)
			}
			for _, l := range log.UpdateGroupMaturityPeriod {
				resp = append(resp, l)
			}
			for _, l := range log.UpdateBootstrapCommitDuration {
				resp = append(resp, l)
			}
			for _, l := range log.UpdateBootstrapRevealDuration {
				resp = append(resp, l)
			}
			for _, l := range log.UpdatebootstrapStartThreshold {
				resp = append(resp, l)
			}
			for _, l := range log.UpdatePendingGroupMaxLife {
				resp = append(resp, l)
			}
			for _, l := range log.GuardianReward {
				resp = append(resp, l)
			}
		}
		if offset >= len(resp) {
			offset = len(resp) - (len(resp) % limit)
		}

		if offset+limit >= len(resp) {
			sendResponse(eventList, resp[offset:], c)
		} else {
			sendResponse(eventList, resp[offset:(offset+limit)], c)
		}

		return true
	}
}
func sendError(c *gin.Context, code int, err string) {
	resp := Response{
		Code:    0,
		Message: "err",
	}
	c.JSON(404, resp)
}
func sendResponse(rType int, logs []interface{}, c *gin.Context) {
	var resp Response
	switch rType {
	case eventList:
		resp = Response{
			Code:    1,
			Message: "success",
			Body:    &Body{Events: []interface{}{}},
		}
		resp.Events = logs
	case nodeList:
		resp = Response{
			Code:    1,
			Message: "success",
			Body:    &Body{Nodelist: []interface{}{}},
		}
		resp.Nodelist = logs
	}

	var jsonData []byte
	jsonData, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, string(jsonData))
}

func search(c *gin.Context) {

	text := strings.Replace(c.Query("text"), "0x", "", -1)
	if text == "" {
		sendError(c, 0, "Empty search")
		return
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		sendError(c, 0, err.Error())
		return
	}
	pageIndex, err := strconv.Atoi(c.Query("pageIndex"))
	if err != nil {
		sendError(c, 0, err.Error())
		return
	}

	fmt.Println("search", text, pageSize, pageIndex)

	f := searchEventTable[strings.ToLower(text)]
	if f != nil {
		if f(pageSize, pageIndex*pageSize, c) {
			return
		}
	}

	if !searchTx(pageSize, pageIndex*pageSize, text, c) {
		sendError(c, 0, "There are no results that match your search")
	}
}

///test http://localhost:8080/api/v1/validationResult/bb14823effa49c05f3b3f970aec6ffcab4da4cb1c6596044bfc6ba95b83a79b
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/search", search)
	}
}
