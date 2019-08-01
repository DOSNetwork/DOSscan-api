package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_models "github.com/DOSNetwork/DOSscan-api/models"
	_repository "github.com/DOSNetwork/DOSscan-api/repository"
	_service "github.com/DOSNetwork/DOSscan-api/service"

	"github.com/gin-gonic/gin"
)

const (
	eventList int = iota
	addressType
	groupType
	urlType
	randomType
	nodeList
)

type SearchHandler struct {
	search *_service.Search
	cache  _repository.Cache
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	*Body   `json:"body"`
}

type Body struct {
	Events     []interface{} `json:"events,omitempty"`
	Address    []interface{} `json:"address,omitempty"`
	Group      []interface{} `json:"group,omitempty"`
	Url        []interface{} `json:"url,omitempty"`
	Random     []interface{} `json:"random,omitempty"`
	NodeList   []interface{} `json:"nodelist,omitempty"`
	TotalCount int           `json:"totalCount,omitempty"`
}

func NesSearchHandler(search *_service.Search, cache _repository.Cache) *SearchHandler {
	return &SearchHandler{
		search: search,
		cache:  cache,
	}
}

func (s *SearchHandler) SupportedEvents(c *gin.Context) {
	fmt.Println("SupportedEvents")
	jsonData, err := json.MarshalIndent(_models.SupportedEvents(), "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))
	c.String(http.StatusOK, string(jsonData))
}

func (s *SearchHandler) Search(c *gin.Context) {
	var err error
	var resp string
	var pageSize, pageIndex int
	query := c.Query("text")
	pageSize, err = strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		return
	}
	pageIndex, err = strconv.Atoi(c.Query("pageIndex"))
	if err != nil {
		return
	}
	fmt.Println("Search ", query, " ", pageSize, " ", pageIndex)
	key := query + c.Query("pageSize") + c.Query("pageIndex")
	ctx := context.Background()

	if resp, err := s.cache.Get(ctx, key); err == nil {
		fmt.Println("Get result from server")
		fmt.Println(resp)
		c.String(http.StatusOK, resp)
		return
	}

	if total, results, resultType, err := s.search.Search(ctx, query, pageSize, pageIndex); err != nil {
		//TODO Add error code
		resp, err = setResponse(0, "fail", resultType, total, results)
	} else {
		resp, err = setResponse(0, "success", resultType, total, results)
	}
	s.cache.Set(ctx, key, resp)
	c.String(http.StatusOK, resp)
	return
}

func setResponse(code int, msg string, rType, totalCount int, logs []interface{}) (string, error) {
	var resp Response
	fmt.Println("setResponse type = ", rType, _models.TypeLatestEvents)
	resp = Response{
		Code:    code,
		Message: msg,
	}
	fmt.Println(_models.TypeLatestEvents <= rType)
	if _models.TypeLatestEvents <= rType {
		fmt.Println("!!!!setResponse type = ", rType, len(logs))
		resp.Body = &Body{Events: logs, TotalCount: totalCount}
	} else {
		switch rType {
		case _models.TypeNode:
			resp.Body = &Body{Address: logs, TotalCount: totalCount}
		case _models.TypeGroup:
			resp.Body = &Body{Group: logs, TotalCount: totalCount}
		case _models.TypeUrlRequest:
			resp.Body = &Body{Url: logs, TotalCount: totalCount}
		case _models.TypeRandomRequest:
			resp.Body = &Body{Random: logs, TotalCount: totalCount}
		default:
			resp.Body = &Body{}
		}
	}

	var jsonData []byte
	jsonData, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("resp ", string(jsonData))
	return string(jsonData), err
}
