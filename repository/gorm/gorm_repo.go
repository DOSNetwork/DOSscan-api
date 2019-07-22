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
	db *gorm.DB
}

func NewGethRepo(db *gorm.DB) repository.DB {
	return &gormRepo{
		db: db,
	}
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
					tx, ok := event[0].(models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(models.LogRegisteredNewPendingNode)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(&tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(&log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&log)
						res := db.Model(&tx).Association("LogRegisteredNewPendingNodes").Append(&log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
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
					tx, ok := event[0].(models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(models.LogGrouping)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(&tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(&log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&log)
						res := db.Model(&tx).Association("LogGroupings").Append(&log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
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
					tx, ok := event[0].(models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(models.LogPublicKeyAccepted)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(&tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(&log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&log)
						res := db.Model(&tx).Association("LogPublicKeyAccepteds").Append(&log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
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
					tx, ok := event[0].(models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(models.LogGroupDissolve)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(&tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&tx)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(&log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&log)
						res := db.Model(&tx).Association("LogGroupDissolves").Append(&log)
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
}

func (g *gormRepo) EventsByModelType(ctx context.Context, modelType int, limit, offset int) (result []interface{}, err error) {
	if modelType >= len(getTable) {
		err = errors.New("Not support")
		return
	}
	result, err = getTable[modelType](ctx, g.db, limit, offset)
	return
}

func (g *gormRepo) SaveModel(ctx context.Context, modelType int, eventc chan []interface{}) (err error, errc chan error) {
	if modelType >= len(saveTable) {
		return errors.New("Not support model type"), nil
	}
	return nil, saveTable[modelType](ctx, g.db, eventc)
}

func (g *gormRepo) NodeByAddr(ctx context.Context, addr string) (node models.Node, err error) {
	node.Addr = addr
	err = g.db.Where(node).First(&node).Error
	return
}

func (g *gormRepo) GroupByID(ctx context.Context, id string) (group models.Group, err error) {
	group.GroupId = id
	err = g.db.Where(group).First(&group).Error
	return
}

func (g *gormRepo) UrlRequestByID(ctx context.Context, id string) (urlRequest models.UrlRequest, err error) {
	urlRequest.RequestId = id
	err = g.db.Where(urlRequest).First(&urlRequest).Error
	return
}

func (g *gormRepo) RandomRequestByID(ctx context.Context, id string) (randRequest models.UserRandomRequest, err error) {
	randRequest.RequestId = id
	err = g.db.Where(randRequest).First(&randRequest).Error
	return
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
