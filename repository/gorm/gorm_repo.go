package gorm

import (
	"context"
	"fmt"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/DOSNetwork/DOSscan-api/repository"
	"github.com/jinzhu/gorm"
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
					tx := event[0].(*models.Transaction)
					if tx == nil {
						continue
					}
					log, ok := event[1].(*models.LogGrouping)
					if !ok {
						continue
					}

					if err := db.Where("block_number = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(&log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&log)
						res := db.Model(&tx).Association("LogGrouping").Append(&log)
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

func (g *gormRepo) SaveModel(ctx context.Context, modelType int, eventc chan []interface{}) chan error {
	return saveTable[modelType](ctx, g.db, eventc)
}
func (g *gormRepo) GetEventsByName(ctx context.Context, name string) []interface{} {
	var result []interface{}
	return result
}

func (g *gormRepo) GetGroupByID(ctx context.Context, id string) interface{} {
	return nil
}

func (g *gormRepo) GetRequestByID(ctx context.Context, id string) interface{} {
	return nil
}

func (g *gormRepo) GetNodeByID(ctx context.Context, id string) interface{} {
	return nil
}
