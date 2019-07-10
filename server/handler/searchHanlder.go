package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	//"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/server/repository"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gin-gonic/gin"
)

const (
	eventList int = iota
	nodeList
)

type SearchHandler struct {
	repo    repository.EventsRepo
	events  []string
	methods []string
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	*Body   `json:"body"`
}

type Body struct {
	Events   []interface{} `json:"events,omitempty"`
	Nodelist []interface{} `json:"nodelist,omitempty"`
}

func NesSearchHandler(repo repository.EventsRepo) *SearchHandler {
	return &SearchHandler{
		repo: repo,
	}
}

func (s *SearchHandler) Init() (err error) {
	s.events, s.methods, err = getEventsAndMethodFromABI("./abi/DOSProxy.abi")
	return
}

func (s *SearchHandler) Search(c *gin.Context) {
	text := c.Query("text")
	_ = text
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		//sendError(c, 0, err.Error())
		return
	}
	_ = pageSize
	pageIndex, err := strconv.Atoi(c.Query("pageIndex"))
	if err != nil {
		//sendError(c, 0, err.Error())
		return
	}
	_ = pageIndex

	if strings.HasPrefix(text, "0x") {
		// 1) tx is hex number that could be address requestID or GroupID
		//a lengeth of full address is 42 byte
		if len(text) == 66 {
			s.repo.SearchRelatedEvents(50, "sender", text)
			s.repo.SearchRelatedEvents(50, "sender", text)
		} else {
			if len(text) == 42 {
				s.repo.SearchRelatedEvents(50, "sender", text)
			}
		}
		//a lengeth of requestID or GroupID is 66 byte

	} else {
		// 2) text is a full event name

	}

	return
}

func setResponse(code int, msg string, rType int, logs []interface{}) (string, error) {
	var resp Response
	switch rType {
	case eventList:
		resp = Response{
			Code:    code,
			Message: msg,
			Body:    &Body{Events: []interface{}{}},
		}
		resp.Events = logs
	case nodeList:
		resp = Response{
			Code:    code,
			Message: msg,
			Body:    &Body{Nodelist: []interface{}{}},
		}
		resp.Nodelist = logs
	}

	var jsonData []byte
	jsonData, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return string(jsonData), err
}

func getEventsAndMethodFromABI(abiPath string) ([]string, []string, error) {
	var events []string
	var methods []string
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
		events = append(events, event.Name)
	}
	for _, method := range proxyAbi.Methods {
		methods = append(methods, method.Name)
	}
	return events, methods, err
}

func caseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}
