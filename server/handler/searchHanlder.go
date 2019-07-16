package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	//"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/server/repository"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gin-gonic/gin"
)

const (
	eventList int = iota
	nodeInfo
	groupInfo
	requestInfo
	nodeList
)

type SearchHandler struct {
	repo         repository.EventsRepo
	sortedEvents []string
	events       map[string]string
	methods      map[string]string
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	*Body   `json:"body"`
}

type Body struct {
	Events      []interface{} `json:"events,omitempty"`
	NodeInfo    []interface{} `json:"nodeInfo,omitempty"`
	GroupInfo   []interface{} `json:"groupInfo,omitempty"`
	RequestInfo []interface{} `json:"requestInfo,omitempty"`
	NodeList    []interface{} `json:"nodelist,omitempty"`
	TotalCount  int           `json:"totalCount,omitempty"`
}

func NesSearchHandler(repo repository.EventsRepo) *SearchHandler {
	return &SearchHandler{
		repo: repo,
	}
}

func (s *SearchHandler) Init() (err error) {
	s.events, s.methods, err = getEventsAndMethodFromABI("../abi/DOSProxy.abi")
	for _, event := range s.events {
		s.sortedEvents = append(s.sortedEvents, event)
	}
	sort.Strings(s.sortedEvents)
	fmt.Println(s.sortedEvents)
	s.repo.SetTxRelatedEvents(s.sortedEvents)
	return
}

func (s *SearchHandler) Search(c *gin.Context) {
	var err error
	var resp string
	var pageSize, pageIndex int
	text := c.Query("text")
	pageSize, err = strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		return
	}
	pageIndex, err = strconv.Atoi(c.Query("pageIndex"))
	if err != nil {
		return
	}

	var events []interface{}
	if text == "" {
		//pageSize = 1
		text = "LogPublicKeyAccepted"
		events = s.repo.GetLatestTxEvents("block_number desc", pageSize)
		if pageSize >= len(events) {
			resp, err = setResponse(0, "success", eventList, len(events), events)
		} else {
			resp, err = setResponse(0, "success", eventList, pageSize, events[:pageSize])
		}
		c.String(http.StatusOK, resp)
		return
	} else if strings.HasPrefix(text, "0x") {
		events, err = searchEventsByHex(s.repo, text, 100, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		events, err = searchEventsByEventName(s.repo, s.events, s.sortedEvents, text, 100, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	offset := pageIndex * pageSize
	limit := pageSize
	if offset > len(events) {
		offset = len(events) - (len(events) % limit)
	}
	if offset+limit >= len(events) {
		resp, err = setResponse(0, "success", eventList, len(events), events[offset:])
	} else {
		resp, err = setResponse(0, "success", eventList, len(events), events[offset:(offset+limit)])
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	c.String(http.StatusOK, resp)
	return
}

func eventsByRequest(repo repository.EventsRepo, limit, offset int, text string) (resp []interface{}) {
	resp = append(resp, repo.GetEvent(limit, offset, "logurl", "query_id ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logrequestuserrandom", "request_id ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logupdaterandom", "last_randomness ILIKE ?", text)...)
	return
}

func eventsByGroup(repo repository.EventsRepo, limit, offset int, text string) (resp []interface{}) {
	resp = append(resp, repo.GetEvent(limit, offset, "logurl", "dispatched_groupid ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logrequestuserrandom", "dispatched_groupid ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logupdaterandom", "dispatched_groupid ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logvalidationresult", "traffic_id ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "loggrouping", "group_id ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logpublickeyaccepted", "group_id ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logpublickeysuggested", "group_id ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "loggroupdissolve", "group_id ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "lognopendinggroup", "group_id ILIKE ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logpendinggroupremoved", "group_id ILIKE ?", text)...)
	return
}

func eventsByAddr(repo repository.EventsRepo, limit, offset int, text string) (resp []interface{}) {
	resp = append(resp, repo.GetEvent(limit, offset, "logregisterednewpendingnode", "node = ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "logcallbacktriggeredfor", "call_back_addr = ?", text)...)
	resp = append(resp, repo.GetEvent(limit, offset, "lognoncontractcall", "call_ddr = ?", text)...)
	resp = append(resp, repo.GetEventsByTxAttr(100, 0, "sender = ?", text)...)
	return
}

func searchEventsByEventName(repo repository.EventsRepo, eventMap map[string]string, sortedEvent []string, text string, pageSize, pageIndex int) ([]interface{}, error) {
	var resp []interface{}
	limit := 100
	offset := 0
	if eventMap[strings.ToLower(text)] != "" {
		fmt.Println("searchEventsByEventName 1")
		resp = append(resp, repo.GetEvent(limit, offset, strings.ToLower(text), nil)...)
	} else {
		fmt.Println("searchEventsByEventName 2")
		for _, event := range sortedEvent {
			fmt.Println("searchEventsByEventName 2 ", event)
			if caseInsensitiveContains(event, text) {
				resp = append(resp, repo.GetEvent(limit, offset, strings.ToLower(event), nil)...)
			}
		}
	}
	fmt.Println("searchEventsByEventName ", text, len(resp))
	return resp, nil
}

func searchEventsByHex(repo repository.EventsRepo, text string, pageSize, pageIndex int) ([]interface{}, error) {
	var resp []interface{}
	limit := 100
	offset := 0
	fmt.Println("searchEventsByHex ", text)

	// 1) text is a 66 bytes hex number that could be requestID or GroupID
	if len(text) == 66 {
		resp = append(resp, eventsByRequest(repo, limit, offset, text)...)
		resp = append(resp, eventsByGroup(repo, limit, offset, text)...)
	} else if len(text) <= 66 {
		// 1) text is a 42 bytes hex number that could be requestID or GroupID
		if len(text) == 42 {
			resp = append(resp, eventsByAddr(repo, limit, offset, text)...)
		}
		if len(resp) == 0 {

			resp = append(resp, eventsByRequest(repo, limit, offset, "%"+text+"%")...)
			resp = append(resp, eventsByGroup(repo, limit, offset, "%"+text+"%")...)
			resp = append(resp, eventsByAddr(repo, limit, offset, "%"+text+"%")...)
			fmt.Println("search hex case3 ", len(resp))
		} else {
			fmt.Println("search hex case4 ")
		}
	}
	fmt.Println("searchEventsByHex ", text, len(resp))

	return resp, nil
}

func setResponse(code int, msg string, rType, totalCount int, logs []interface{}) (string, error) {
	var resp Response
	fmt.Println("setResponse ", len(logs))
	resp = Response{
		Code:    code,
		Message: msg,
	}
	switch rType {
	case eventList:
		resp.Body.Events = logs
	case nodeInfo:
		resp.Body = &Body{NodeInfo: logs, TotalCount: totalCount}
	case groupInfo:
		resp.Body = &Body{GroupInfo: logs, TotalCount: totalCount}
	case requestInfo:
		resp.Body = &Body{RequestInfo: logs, TotalCount: totalCount}
	case nodeList:
		resp.Body = &Body{NodeList: logs, TotalCount: totalCount}

	}
	fmt.Println("resp ", resp.Body)

	var jsonData []byte
	jsonData, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return string(jsonData), err
}

func getEventsAndMethodFromABI(abiPath string) (map[string]string, map[string]string, error) {
	events := make(map[string]string)

	methods := make(map[string]string)
	jsonFile, err := os.Open(abiPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return events, methods, err
	}

	abiJsonByte, _ := ioutil.ReadAll(jsonFile)
	proxyAbi, err := abi.JSON(strings.NewReader(string(abiJsonByte)))
	if err != nil {
		fmt.Println(err)
		return events, methods, err
	}

	for _, event := range proxyAbi.Events {
		if event.Name != "LogUnRegisteredNewPendingNode" &&
			event.Name != "OwnershipRenounced" &&
			event.Name != "OwnershipTransferred" {
			events[strings.ToLower(event.Name)] = event.Name
		}
	}
	for _, method := range proxyAbi.Methods {
		methods[strings.ToLower(method.Name)] = method.Name
	}
	return events, methods, err
}

func caseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}
