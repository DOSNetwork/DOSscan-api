package apiv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
	if strings.HasPrefix(text, "0x") {
		//address or requestID or GroupID
	} else {
		f := searchEventTable[strings.ToLower(text)]
		if f != nil {
			fmt.Println("case 1")
			sendResponse(eventList, f(pageSize, pageIndex*pageSize), c)
			return
		}
	}
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
	//	offset := pageIndex * pageSize
	//	limit := pageSize
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
	//	offset = pageIndex*pageSize - len(resp)
	// 2) Check if text is included in tx,method or event
	//resp = append(resp, searchTx(limit, offset, text)...)
	//fmt.Println("case 3 from event total, ", len(resp))

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
