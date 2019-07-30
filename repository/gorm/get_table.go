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
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeGrouping: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogGrouping
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypePublicKeySuggested: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogPublicKeySuggested
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypePublicKeyAccepted: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogPublicKeyAccepted
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeGroupDissolve: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogGroupDissolve
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeUpdateRandom: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogUpdateRandom
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeUrl: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogUrl
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeRequestUserRandom: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogRequestUserRandom
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeValidationResult: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogValidationResult
		db.Limit(limit).Offset(offset).Find(&_models)
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
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeCallbackTriggeredFor: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogCallbackTriggeredFor
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
	_models.TypeError: func(ctx context.Context, db *gorm.DB, limit, offset int) (results []interface{}, err error) {
		var _models []_models.LogError
		db.Limit(limit).Offset(offset).Find(&_models)
		for _, model := range _models {
			results = append(results, model)
		}
		return
	},
}
