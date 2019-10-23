package gorm

import (
	"context"
	"errors"
	"fmt"

	_models "github.com/DOSNetwork/DOSscan-api/models"
	_repository "github.com/DOSNetwork/DOSscan-api/repository"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type gormRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) _repository.DB {
	db.AutoMigrate(&_models.Group{}, &_models.Node{}, &_models.UrlRequest{}, &_models.UserRandomRequest{},
		&_models.Transaction{}, &_models.LogRegisteredNewPendingNode{}, &_models.LogUnRegisteredNewPendingNode{},
		&_models.LogGrouping{}, &_models.LogPublicKeySuggested{}, &_models.LogPublicKeyAccepted{}, &_models.LogGroupDissolve{},
		&_models.LogUpdateRandom{}, &_models.LogRequestUserRandom{}, &_models.LogUrl{}, &_models.LogValidationResult{},
		&_models.LogCallbackTriggeredFor{}, &_models.GuardianReward{}, &_models.LogMessage{})

	return &gormRepo{
		db: db,
	}
}

func (g *gormRepo) ModelsByType(ctx context.Context, modelType int, limit, offset int) (results []interface{}, err error) {
	if modelType == 0 || modelType >= len(getTable) {
		err = errors.New("Unsupported Type ")
		return
	}
	results, err = getTable[modelType](ctx, g.db, limit, offset)
	return
}

func (g *gormRepo) CountModel(ctx context.Context, modelType int) (total int, err error) {
	if modelType == 0 {
		err = errors.New("Unsupported Type ")
		return
	}
	if model := _models.TypeToStruct(modelType); model != nil {
		err = g.db.Model(model).Count(&total).Error
		if err != nil {
			fmt.Println("CountModel err ", err)
		}
	}
	return
}

func (g *gormRepo) LastBlockNum(ctx context.Context, modelType int) (lastBlkNum uint64, err error) {
	var blkNums []uint64
	if model := _models.TypeToStruct(modelType); model != nil {
		err = g.db.Model(model).Order("block_number desc").Pluck("block_number", &blkNums).Error
		if err != nil {
			fmt.Println("LastBlockNum err ", err)
		}
		if len(blkNums) != 0 {
			lastBlkNum = blkNums[0]
		}
	}
	return
}

func (g *gormRepo) SaveModel(ctx context.Context, modelType int, eventc chan []interface{}) (err error, errc chan error) {
	if modelType >= len(saveTable) {
		return errors.New("Not support model type"), nil
	}
	return nil, saveTable[modelType](ctx, g.db, eventc)
}

func (g *gormRepo) LatestEvents(ctx context.Context, limit int) (resp []interface{}, err error) {
	//block_number desc
	logs := []_models.Transaction{}

	if err := g.db.Order("block_number desc").Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		fmt.Println("LatestEvents number of tx ", len(logs))
		resp = relatedEvents(g.db, logs)
	} else if err != nil {
		fmt.Println("LatestEvents err ", err)
	}
	if limit < len(resp) {
		resp = resp[:limit]
	}
	return
}

func (g *gormRepo) NodeByAddr(ctx context.Context, addr string) (node _models.Node, err error) {
	node.Addr = addr
	err = g.db.Where(node).First(&node).Error
	if err != nil {
		return
	}
	g.db.Model(&node).Related(&node.Groups, "Groups")
	return
}

func (g *gormRepo) GroupByID(ctx context.Context, id string) (group _models.Group, err error) {
	group.GroupId = id
	err = g.db.Where(group).First(&group).Error
	if !gorm.IsRecordNotFoundError(err) {
		g.db.Model(&group).Related(&group.UrlRequests, "UrlRequests")
		g.db.Model(&group).Related(&group.UserRandomRequests, "UserRandomRequests")
		group.NumUrl = len(group.UrlRequests)
		group.NumRandom = len(group.UserRandomRequests)
	}
	return
}

func (g *gormRepo) UrlRequestByID(ctx context.Context, id string) (urlRequest _models.UrlRequest, err error) {
	urlRequest.RequestId = id
	fmt.Println("UrlRequestByID : ", id)
	//err = g.db.Where(urlRequest).First(&urlRequest).Error
	err = g.db.First(&urlRequest, "request_id = ?", id).Error
	if err != nil {
		fmt.Println("UrlRequestByID : err ", err)
	}
	if !gorm.IsRecordNotFoundError(err) {
		fmt.Println("UrlRequestByID : ", urlRequest)

		urlRequest.MessageStr = string(urlRequest.Message)
	}
	return
}

func (g *gormRepo) RandomRequestByID(ctx context.Context, id string) (randRequest _models.UserRandomRequest, err error) {
	randRequest.RequestId = id
	err = g.db.Where(randRequest).First(&randRequest).Error
	if !gorm.IsRecordNotFoundError(err) {
		randRequest.MessageStr = fmt.Sprintf("0x%x", randRequest.Message)
	}
	return
}

func (g *gormRepo) BuildRelation(ctx context.Context) {
	buildGroup(g.db, "")
	buildUrlRequest(g.db, "")
	buildRandomRequest(g.db, "")
}
func buildNode(db *gorm.DB, addr string) {
	var node _models.Node
	db.Where(_models.Node{Addr: addr}).FirstOrCreate(&node)
	//Find Group has node addr in node_id
	sel := "SELECT group_id FROM groups WHERE $1 <@ node_id"
	rows, err := db.DB().Query(sel, pq.Array([]string{addr}))
	if err != nil {
		fmt.Println("Query err ", err)
	}
	for rows.Next() {
		var group _models.Group
		rows.Scan(&group.GroupId)
		if err := db.Where("group_id = ?", group.GroupId).First(&group).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println("Can't find group ", group.GroupId, " ", node.Addr)
		} else {
			res := db.Model(&node).Association("Groups").Append(&group)
			if res.Error != nil {
				fmt.Println("res ", res.Error)
			}
			//fmt.Println("len ", db.Model(&node).Association("Groups").Count())
			//fmt.Println("len ", node.Groups[0].GroupId)
		}
	}
}

func buildGroup(db *gorm.DB, grouId string) {
	var results []_models.Group

	selectStr := "log_groupings.group_id,log_groupings.node_id,"
	selectStr = selectStr + "log_public_key_accepteds.accepted_blk_num,"
	selectStr = selectStr + "log_public_key_accepteds.pub_key,"
	selectStr = selectStr + "log_group_dissolves.dissolved_blk_num"
	jStr := "left join log_public_key_accepteds on "
	jStr = jStr + "log_public_key_accepteds.group_id = log_groupings.group_id"
	jStr2 := "left join log_group_dissolves on "
	jStr2 = jStr2 + "log_group_dissolves.group_id = log_groupings.group_id"
	tempDb := db.Table("log_groupings").Select(selectStr)
	tempDb = tempDb.Joins(jStr)
	tempDb = tempDb.Joins(jStr2)

	if grouId == "" {
		tempDb.Find(&results)
	} else {
		tempDb.Where(&_models.LogGrouping{GroupId: grouId}, grouId).Find(&results)
	}
	//fmt.Println("-", " len buildGroup", len(results))

	for _, group := range results {
		var existGroup _models.Group
		if err := db.Where("group_id = ?", group.GroupId).First(&existGroup).Error; gorm.IsRecordNotFoundError(err) {
			db.Create(&group)
			for _, addr := range group.NodeId {
				var node _models.Node
				err = db.Where(_models.Node{Addr: addr}).First(&node).Error
				if err != nil {
					fmt.Println("Can't find addr ", addr)
				}
				res := db.Model(&node).Association("Groups").Append(&group)
				if res.Error != nil {
					fmt.Println("res ", res.Error)
				}
			}
		} else {
			db.Model(&existGroup).Omit("group_id").Updates(&group)
		}
	}
}

func buildUrlRequest(db *gorm.DB, requestId string) {
	var results []_models.UrlRequest
	tempDb := db.Table("log_urls").Select("log_urls.request_id, log_urls.dispatched_group_id,transactions.sender, transactions.block_number,transactions.hash,log_validation_results.message,log_validation_results.signature,log_validation_results.pub_key,log_validation_results.pass,log_urls.timeout,log_urls.data_source,log_urls.selector,log_urls.randomness")
	tempDb = tempDb.Joins("left join log_validation_results on log_validation_results.request_id = log_urls.request_id")
	tempDb = tempDb.Joins("left join transactions on log_validation_results.transaction_id = transactions.id")
	if requestId == "" {
		tempDb.Find(&results)
	} else {
		tempDb.Where("log_urls.request_id = ? ", requestId).Find(&results)
	}
	//fmt.Println("-", " len buildUrlRequest", len(results))

	for _, request := range results {
		db.Save(&request)
		var group _models.Group
		if err := db.Where(&_models.Group{GroupId: request.DispatchedGroupId}).First(&group).Error; gorm.IsRecordNotFoundError(err) {
			continue
		}
		res := db.Model(&group).Association("UrlRequests").Append(&request)
		if res.Error != nil {
			fmt.Println("res ", res.Error)
			continue
		}
	}
}

func buildRandomRequest(db *gorm.DB, requestId string) {

	var results []_models.UserRandomRequest
	tempDb := db.Table("log_request_user_randoms").Select("log_request_user_randoms.request_id, log_request_user_randoms.dispatched_group_id,log_request_user_randoms.last_system_randomness,log_request_user_randoms.user_seed,transactions.sender, transactions.block_number,transactions.hash,log_validation_results.message,log_validation_results.signature,log_validation_results.pub_key,log_validation_results.pass")
	tempDb = tempDb.Joins("left join log_validation_results on log_validation_results.request_id = log_request_user_randoms.request_id")
	tempDb = tempDb.Joins("left join transactions on log_validation_results.transaction_id = transactions.id")
	if requestId == "" {
		tempDb.Find(&results)
	} else {
		tempDb.Where("log_request_user_randoms.request_id = ? ", requestId).Find(&results)
	}
	//fmt.Println("-", " len buildRandomRequest", len(results))

	for _, request := range results {
		db.Save(&request)
		var group _models.Group

		if err := db.Where(&_models.Group{GroupId: request.DispatchedGroupId}).First(&group).Error; gorm.IsRecordNotFoundError(err) {
			continue
		}
		res := db.Model(&group).Association("UserRandomRequests").Append(&request)
		if res.Error != nil {
			fmt.Println("res ", res.Error)
		}
	}
}

func relatedEvents(db *gorm.DB, txs []_models.Transaction) []interface{} {
	var resp []interface{}
	for _, tx := range txs {
		db.Model(&tx).Related(&tx.LogRegisteredNewPendingNodes, "LogRegisteredNewPendingNodes")
		for _, event := range tx.LogRegisteredNewPendingNodes {
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.LogUnRegisteredNewPendingNodes, "LogUnRegisteredNewPendingNode")
		for _, event := range tx.LogUnRegisteredNewPendingNodes {
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.LogGroupings, "LogGroupings")
		for _, event := range tx.LogGroupings {
			resp = append(resp, event)
		}

		db.Model(&tx).Related(&tx.LogPublicKeySuggesteds, "LogPublicKeySuggesteds")
		for _, event := range tx.LogPublicKeySuggesteds {
			resp = append(resp, event)
		}

		db.Model(&tx).Related(&tx.LogPublicKeyAccepteds, "LogPublicKeyAccepteds")
		for _, event := range tx.LogPublicKeyAccepteds {
			resp = append(resp, event)
		}

		db.Model(&tx).Related(&tx.LogGroupDissolves, "LogGroupDissolves")
		for _, event := range tx.LogGroupDissolves {
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.LogUpdateRandoms, "LogUpdateRandoms")
		for _, event := range tx.LogUpdateRandoms {
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.LogUrls, "LogUrls")
		for _, event := range tx.LogUrls {
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.LogRequestUserRandoms, "LogRequestUserRandoms")
		for _, event := range tx.LogRequestUserRandoms {
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.LogValidationResults, "LogValidationResults")
		for _, event := range tx.LogValidationResults {
			if event.RequestType == 2 {
				event.MessageStr = string(event.Message)
			} else {
				event.MessageStr = fmt.Sprintf("0x%x", event.Message)
			}
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.LogCallbackTriggeredFors, "LogCallbackTriggeredFors")
		for _, event := range tx.LogCallbackTriggeredFors {
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.GuardianRewards, "GuardianRewards")
		for _, event := range tx.GuardianRewards {
			resp = append(resp, event)
		}
		db.Model(&tx).Related(&tx.LogMessages, "LogMessages")
		for _, event := range tx.LogMessages {
			resp = append(resp, event)
		}
	}
	return resp
}
