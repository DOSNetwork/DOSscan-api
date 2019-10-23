package gorm

import (
	"context"
	"fmt"

	_models "github.com/DOSNetwork/DOSscan-api/models"
	"github.com/jinzhu/gorm"
)

var getTable = []func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error){
	_models.TypeNewPendingNode: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogRegisteredNewPendingNode
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeUnregisterPendingNode: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogUnRegisteredNewPendingNode
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeGrouping: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogGrouping
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypePublicKeySuggested: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogPublicKeySuggested
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypePublicKeyAccepted: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogPublicKeyAccepted
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeGroupDissolve: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogGroupDissolve
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeUpdateRandom: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogUpdateRandom
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeUrl: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogUrl
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeRequestUserRandom: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogRequestUserRandom
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeValidationResult: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogValidationResult
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			if model.RequestType == 2 {
				model.MessageStr = string(model.Message)
			} else {
				model.MessageStr = fmt.Sprintf("0x%x", model.Message)
			}
			results = append(results, model)
		}
		return
	},
	_models.TypeGuardianReward: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.GuardianReward
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeCallbackTriggeredFor: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogCallbackTriggeredFor
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeMessage: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogMessage
		db.Order("block_number desc").Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
}
