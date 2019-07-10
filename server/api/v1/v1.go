package apiv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
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

func searchTx(limit, offset int, condition string) []interface{} {
	logs := []models.Transaction{}
	var resp []interface{}
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

	if err := db.Where("hash ILIKE ?", "%"+condition+"%").Or("method ILIKE ?", "%"+condition+"%").Offset(offset).Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		fmt.Println("searchTx ", len(logs))
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
	}
	return resp
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
	// 1) text is a full event name
	/*
		f := searchEventTable[strings.ToLower(text)]
		if f != nil {
			fmt.Println("case 1")
			sendResponse(eventList, f(pageSize, pageIndex*pageSize), c)
			return
		}
	*/
	// 2)
	offset := pageIndex * pageSize
	limit := pageSize
	var resp []interface{}
	/*
		keys := make([]string, 0)
		for k, _ := range searchEventTable {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			if caseInsensitiveContains(key, text) {
				resp = append(resp, searchEventTable[key](limit, offset)...)
			}
			fmt.Println("case 2 from event ", key, "  ", len(resp))
			if len(resp) >= pageSize {
				fmt.Println("case 2 total, ", len(resp[:pageSize]))
				sendResponse(eventList, resp[:pageSize], c)
				return
			}
			offset = pageIndex*pageSize - len(resp)
		}
	*/
	offset = pageIndex*pageSize - len(resp)
	// 2) Check if text is included in tx,method or event
	resp = append(resp, searchTx(limit, offset, text)...)
	fmt.Println("case 3 from event total, ", len(resp))

	if len(resp) >= pageSize {
		sendResponse(eventList, resp[:pageSize], c)
	} else {
		sendResponse(eventList, resp, c)
	}
}

func caseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}

///test http://localhost:8080/api/v1/validationResult/bb14823effa49c05f3b3f970aec6ffcab4da4cb1c6596044bfc6ba95b83a79b
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/search", search)
	}
}
