package gorm

import (
	"context"
	"errors"
	"fmt"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/repository"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type gormRepo struct {
	db             *gorm.DB
	supportedEvent []string
	nameToType     map[string]int
}

func NewGormRepo(db *gorm.DB) repository.DB {
	var supportedEvent []string
	supportedEvent = append(supportedEvent, "LogRegisteredNewPendingNode")
	supportedEvent = append(supportedEvent, "LogGrouping")
	nameToType := make(map[string]int)
	nameToType["logregisterednewpendingnode"] = models.TypeNewPendingNode
	nameToType["loggrouping"] = models.TypeGrouping
	nameToType["logpublickeyaccepted"] = models.TypePublicKeyAccepted
	nameToType["loggroupdissolve"] = models.TypeGroupDissolve
	nameToType["logurl"] = models.TypeUrl
	nameToType["logrequestuserrandom"] = models.TypeRequestUserRandom
	nameToType["logvalidationresult"] = models.TypeValidationResult

	db.AutoMigrate(&models.Transaction{}, &models.LogRegisteredNewPendingNode{},
		&models.LogGrouping{}, &models.LogPublicKeyAccepted{}, &models.LogGroupDissolve{},
		&models.Group{}, &models.Node{}, &models.LogRequestUserRandom{}, &models.LogUrl{},
		&models.LogValidationResult{}, &models.UrlRequest{}, &models.UserRandomRequest{})

	return &gormRepo{
		db: db,
	}
}

var typeToStruct = []interface{}{
	models.TypeNewPendingNode:    &models.LogRegisteredNewPendingNode{},
	models.TypeGrouping:          &models.LogGrouping{},
	models.TypePublicKeyAccepted: &models.LogPublicKeyAccepted{},
	models.TypeGroupDissolve:     &models.LogGroupDissolve{},
	models.TypeRequestUserRandom: &models.LogRequestUserRandom{},
	models.TypeUrl:               &models.LogUrl{},
	models.TypeValidationResult:  &models.LogValidationResult{},
}

var saveTable = []func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error{
	models.TypeNewPendingNode: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					if len(event) != 2 {
						continue
					}
					tx, ok := event[0].(*models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*models.LogRegisteredNewPendingNode)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(log)
						res := db.Model(tx).Association("LogRegisteredNewPendingNodes").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
						buildNode(db, log.Node)
					}
				}
			}
		}()
		return errc
	},
	models.TypeGrouping: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					if len(event) != 2 {
						continue
					}
					tx, ok := event[0].(*models.Transaction)
					if !ok {
						fmt.Println("tx !ok")

						continue
					}
					log, ok := event[1].(*models.LogGrouping)
					if !ok {
						fmt.Println("log !ok")
						continue
					}
					fmt.Println("tx.Hash ", tx.Hash, " tx.BlockNumber ", tx.BlockNumber)
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(log)
						res := db.Model(tx).Association("LogGroupings").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
						//buildGroup(db, log.GroupId)
					}
				}
			}
		}()
		return errc
	},
	models.TypePublicKeyAccepted: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					if len(event) != 2 {
						continue
					}
					tx, ok := event[0].(*models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*models.LogPublicKeyAccepted)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(log)
						res := db.Model(tx).Association("LogPublicKeyAccepteds").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
						//buildGroup(db, log.GroupId)
					}
				}
			}
		}()
		return errc
	},
	models.TypeGroupDissolve: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					if len(event) != 2 {
						continue
					}
					tx, ok := event[0].(*models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*models.LogGroupDissolve)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(log)
						res := db.Model(tx).Association("LogGroupDissolves").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
						//buildGroup(db, log.GroupId)
					}
				}
			}
		}()
		return errc
	},
	models.TypeRequestUserRandom: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					if len(event) != 2 {
						continue
					}
					tx, ok := event[0].(*models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*models.LogRequestUserRandom)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(log)
						res := db.Model(tx).Association("LogRequestUserRandoms").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
						//buildRandomRequest(db, log.RequestId)
					}
				}
			}
		}()
		return errc
	},
	models.TypeUrl: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					if len(event) != 2 {
						continue
					}
					tx, ok := event[0].(*models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*models.LogUrl)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(log)
						res := db.Model(tx).Association("LogUrls").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
						//buildUrlRequest(db, log.RequestId)
					}
				}
			}
		}()
		return errc
	},
	models.TypeValidationResult: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
		errc := make(chan error)
		go func() {
			defer close(errc)
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					if len(event) != 2 {
						continue
					}
					tx, ok := event[0].(*models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*models.LogValidationResult)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(log)
						res := db.Model(tx).Association("LogValidationResults").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
}

var getTable = []func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error){
	models.TypeNewPendingNode: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var models []models.LogRegisteredNewPendingNode
		db.Limit(limit).Offset(offset).Find(&models)
		for _, model := range models {
			results = append(results, model)
		}
		return
	},
	models.TypeGrouping: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var models []models.LogGrouping
		db.Limit(limit).Offset(offset).Find(&models)
		for _, model := range models {
			results = append(results, model)
		}
		return
	},
	models.TypePublicKeyAccepted: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var models []models.LogPublicKeyAccepted
		db.Limit(limit).Offset(offset).Find(&models)
		for _, model := range models {
			results = append(results, model)
		}
		return
	},
	models.TypeGroupDissolve: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var models []models.LogGroupDissolve
		db.Limit(limit).Offset(offset).Find(&models)
		for _, model := range models {
			results = append(results, model)
		}
		return
	},
	models.TypeUrl: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var models []models.LogUrl
		db.Limit(limit).Offset(offset).Find(&models)
		for _, model := range models {
			results = append(results, model)
		}
		return
	},
	models.TypeRequestUserRandom: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var models []models.LogRequestUserRandom
		db.Limit(limit).Offset(offset).Find(&models)
		for _, model := range models {
			results = append(results, model)
		}
		return
	},
	models.TypeValidationResult: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var models []models.LogValidationResult
		db.Limit(limit).Offset(offset).Find(&models)
		for _, model := range models {
			if model.RequestType == 2 {
				model.MessageStr = string(model.Message)
			} else {
				model.MessageStr = fmt.Sprintf("0x%x", model.Message)
			}
			results = append(results, model)
		}
		return
	},
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
	if modelType == 0 || modelType >= len(typeToStruct) {
		err = errors.New("Unsupported Type ")
		return
	}
	err = g.db.Model(typeToStruct[modelType]).Count(&total).Error
	return
}

func (g *gormRepo) LastBlockNum(ctx context.Context, modelType int) (lastBlkNum uint64, err error) {
	var blkNums []uint64
	err = g.db.Limit(1).Order("block_number desc").Find(typeToStruct[modelType]).Pluck("block_number", &blkNums).Error
	if len(blkNums) == 0 {
		lastBlkNum = 4468402
	} else {
		lastBlkNum = blkNums[0]
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
	logs := []models.Transaction{}

	if err := g.db.Order("block_number desc").Limit(limit).Find(&logs).Error; !gorm.IsRecordNotFoundError(err) {
		resp = relatedEvents(g.db, logs)
	}
	if limit < len(resp) {
		resp = resp[:limit]
	}
	return
}

func (g *gormRepo) NodeByAddr(ctx context.Context, addr string) (node models.Node, err error) {
	node.Addr = addr
	err = g.db.Where(node).First(&node).Error
	if err != nil {
		return
	}
	g.db.Model(&node).Related(&node.Groups, "Groups")
	return
}

func (g *gormRepo) GroupByID(ctx context.Context, id string) (group models.Group, err error) {
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

func (g *gormRepo) UrlRequestByID(ctx context.Context, id string) (urlRequest models.UrlRequest, err error) {
	urlRequest.RequestId = id
	err = g.db.Where(urlRequest).First(&urlRequest).Error
	if !gorm.IsRecordNotFoundError(err) {
		fmt.Println("UrlRequestByID : ", len(urlRequest.Message))

		urlRequest.MessageStr = string(urlRequest.Message)
	}
	return
}

func (g *gormRepo) RandomRequestByID(ctx context.Context, id string) (randRequest models.UserRandomRequest, err error) {
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
	var node models.Node
	db.Where(models.Node{Addr: addr}).FirstOrCreate(&node)
	//Find Group has node addr in node_id
	sel := "SELECT group_id FROM groups WHERE $1 <@ node_id"
	rows, err := db.DB().Query(sel, pq.Array([]string{addr}))
	if err != nil {
		fmt.Println("Query err ", err)
	}
	for rows.Next() {
		var group models.Group
		rows.Scan(&group.GroupId)
		if err := db.Where("group_id = ?", group.GroupId).First(&group).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println("Can't find group ", group.GroupId, " ", node.Addr)
		} else {
			res := db.Model(&node).Association("Groups").Append(&group)
			if res.Error != nil {
				fmt.Println("res ", res.Error)
			}
			fmt.Println("len ", db.Model(&node).Association("Groups").Count())
			fmt.Println("len ", node.Groups[0].GroupId)
		}
	}
}

func buildGroup(db *gorm.DB, grouId string) {
	var results []models.Group

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
		tempDb.Where(&models.LogGrouping{GroupId: grouId}, grouId).Find(&results)
	}
	fmt.Println(len(results))
	for _, group := range results {
		var existGroup models.Group
		if err := db.Where("group_id = ?", group.GroupId).First(&existGroup).Error; gorm.IsRecordNotFoundError(err) {
			db.Create(&group)
			for _, addr := range group.NodeId {
				var node models.Node
				err = db.Where(models.Node{Addr: addr}).First(&node).Error
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
			fmt.Println("Update group ", existGroup.GroupId, group.DissolvedBlkNum, group.AcceptedBlkNum)
		}
	}
}

func buildUrlRequest(db *gorm.DB, requestId string) {
	var results []models.UrlRequest
	tempDb := db.Table("log_urls").Select("log_urls.request_id, log_urls.dispatched_group_id,transactions.sender, transactions.block_number,transactions.hash,log_validation_results.message,log_validation_results.signature,log_validation_results.pub_key,log_validation_results.pass,log_urls.timeout,log_urls.data_source,log_urls.selector,log_urls.randomness")
	tempDb = tempDb.Joins("left join log_validation_results on log_validation_results.request_id = log_urls.request_id")
	tempDb = tempDb.Joins("left join transactions on log_validation_results.transaction_id = transactions.id")
	if requestId == "" {
		tempDb.Find(&results)
	} else {
		tempDb.Where("log_urls.request_id = ? ", requestId).Find(&results)
	}

	for _, request := range results {
		db.Where(request).FirstOrCreate(&request)
		var group models.Group
		if err := db.Where(&models.Group{GroupId: request.DispatchedGroupId}).First(&group).Error; gorm.IsRecordNotFoundError(err) {
			continue
		}
		res := db.Model(&group).Association("UrlRequests").Append(&request)
		if res.Error != nil {
			fmt.Println("res ", res.Error)
			continue
		}
		fmt.Println(group.GroupId, "-", " len ", db.Model(&group).Association("UrlRequests").Count())
	}
}

func buildRandomRequest(db *gorm.DB, requestId string) {

	var results []models.UserRandomRequest
	tempDb := db.Table("log_request_user_randoms").Select("log_request_user_randoms.request_id, log_request_user_randoms.dispatched_group_id,log_request_user_randoms.last_system_randomness,log_request_user_randoms.user_seed,transactions.sender, transactions.block_number,transactions.hash,log_validation_results.message,log_validation_results.signature,log_validation_results.pub_key,log_validation_results.pass")
	tempDb = tempDb.Joins("left join log_validation_results on log_validation_results.request_id = log_request_user_randoms.request_id")
	tempDb = tempDb.Joins("left join transactions on log_validation_results.transaction_id = transactions.id")
	if requestId == "" {
		tempDb.Find(&results)
	} else {
		tempDb.Where("log_request_user_randoms.request_id = ? ", requestId).Find(&results)
	}
	fmt.Println("-", " len buildRandomRequest", len(results))

	for _, request := range results {
		db.Where(request).FirstOrCreate(&request)
		var group models.Group

		if err := db.Where(&models.Group{GroupId: request.DispatchedGroupId}).First(&group).Error; gorm.IsRecordNotFoundError(err) {
			continue
		}
		res := db.Model(&group).Association("UserRandomRequests").Append(&request)
		if res.Error != nil {
			fmt.Println("res ", res.Error)
		}
		fmt.Println(group.GroupId, "-", " len buildRandomRequest", db.Model(&group).Association("UserRandomRequests").Count())
	}
}

func relatedEvents(db *gorm.DB, txs []models.Transaction) []interface{} {
	var resp []interface{}
	for _, tx := range txs {
		fmt.Println("blockNum ", tx.BlockNumber, tx.Method)
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
		db.Model(&tx).Related(&tx.LogGroupings, "LogGroupings")
		for _, event := range tx.LogGroupings {
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
		db.Model(&tx).Related(&tx.LogRegisteredNewPendingNodes, "LogRegisteredNewPendingNodes")
		for _, event := range tx.LogRegisteredNewPendingNodes {
			resp = append(resp, event)
		}
	}
	return resp
}
