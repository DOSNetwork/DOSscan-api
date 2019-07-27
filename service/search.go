package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	_models "github.com/DOSNetwork/DOSscan-api/models"
	_repository "github.com/DOSNetwork/DOSscan-api/repository"
)

type Search struct {
	onchainRepo _repository.Onchain
	dbRepo      _repository.DB
}

func NewSearch(onchainRepo _repository.Onchain, dbRepo _repository.DB) *Search {
	return &Search{
		onchainRepo: onchainRepo,
		dbRepo:      dbRepo,
	}
}

func (s *Search) Search(ctx context.Context, query string, pageSize, pageIndex int) (total int, results []interface{}, respType int, err error) {
	if query == "" {
		//Latest event
		results, err = s.dbRepo.LatestEvents(ctx, 20)
		respType = 1
		total = len(results)
		return
	} else if strings.HasPrefix(query, "0x") {
		var result interface{}
		result, respType, err = s.searchEventsByHex(ctx, query)
		results = append(results, result)
		total = len(results)
		return
	} else {
		//Search Event Name
		total, results, respType, err = s.searchEventsByName(ctx, query, pageSize, pageIndex)

	}
	return
}

func (s *Search) searchEventsByName(ctx context.Context, query string, pageSize, pageIndex int) (total int, resp []interface{}, respType int, err error) {
	if modelType := _models.StringToType(query); modelType != 0 {
		respType = modelType
		total, err = s.dbRepo.CountModel(ctx, modelType)
		if total == 0 || err != nil {
			return
		}
		limit := pageSize
		offset := pageIndex * pageSize
		if offset > total {
			offset = total / pageSize
		}
		//count
		resp, err = s.dbRepo.ModelsByType(ctx, modelType, limit, offset)
	} else {
		err = errors.New("Unsupported event name")
	}
	return
}
func (s *Search) searchEventsByHex(ctx context.Context, query string) (resp interface{}, respType int, err error) {
	// 1) text is a 66 bytes hex number that could be requestID or GroupID
	if len(query) == 66 {
		if resp, err = s.dbRepo.GroupByID(ctx, query); err == nil {
			respType = _models.TypeGroup
		} else if resp, err = s.dbRepo.UrlRequestByID(ctx, query); err == nil {
			respType = _models.TypeUrlRequest
		} else if resp, err = s.dbRepo.RandomRequestByID(ctx, query); err == nil {
			respType = _models.TypeRandomRequest
		}
	} else if len(query) == 42 {
		var node _models.Node
		if node, err = s.dbRepo.NodeByAddr(ctx, query); err == nil {
			node.Balance, err = s.onchainRepo.Balance(ctx, node.Addr)
			if err != nil {
				node.Balance = "-1"
			}
			//TODO : Chcek with staking contract
			node.RegisterState = true
			for _, group := range node.Groups {
				fmt.Println("group ", group.DissolvedBlkNum, group.AcceptedBlkNum)
				if group.DissolvedBlkNum == 0 && group.AcceptedBlkNum != 0 {
					node.ActiveGroups = append(node.ActiveGroups, group.GroupId)
					fmt.Println("Node ActiveGroup ", node.ActiveGroups, group.AcceptedBlkNum)
				} else if group.DissolvedBlkNum != 0 && group.AcceptedBlkNum != 0 {
					node.ExpiredGroups++
				}
			}
			if len(node.ActiveGroups) == 0 {
				node.ActiveGroups = append(node.ActiveGroups, "")
			}
			resp = node
			respType = _models.TypeNode
		}
	}
	return
}
