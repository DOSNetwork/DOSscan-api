package gorm

import (
	"context"
	"fmt"

	_models "github.com/DOSNetwork/DOSscan-api/models"
	"github.com/jinzhu/gorm"
)

var saveTable = []func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error{
	_models.TypeNewPendingNode: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogRegisteredNewPendingNode)
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
	_models.TypeGrouping: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						fmt.Println("tx !ok")

						continue
					}
					log, ok := event[1].(*_models.LogGrouping)
					if !ok {
						fmt.Println("log !ok")
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
						res := db.Model(tx).Association("LogGroupings").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
						buildGroup(db, log.GroupId)
						for _, node := range log.NodeId {
							buildNode(db, node)
						}
					}
				}
			}
		}()
		return errc
	},
	_models.TypePublicKeySuggested: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogPublicKeySuggested)
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
						res := db.Model(tx).Association("LogPublicKeySuggesteds").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
						buildGroup(db, log.GroupId)
					}
				}
			}
		}()
		return errc
	},
	_models.TypePublicKeyAccepted: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogPublicKeyAccepted)
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
						buildGroup(db, log.GroupId)
					}
				}
			}
		}()
		return errc
	},
	_models.TypeGroupDissolve: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogGroupDissolve)
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
						buildGroup(db, log.GroupId)
					}
				}
			}
		}()
		return errc
	},
	_models.TypeRequestUserRandom: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogRequestUserRandom)
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
						buildRandomRequest(db, log.RequestId)
					}
				}
			}
		}()
		return errc
	},
	_models.TypeUpdateRandom: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogUpdateRandom)
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
						res := db.Model(tx).Association("LogUpdateRandoms").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	_models.TypeUrl: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogUrl)
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
						buildUrlRequest(db, log.RequestId)
					}
				}
			}
		}()
		return errc
	},
	_models.TypeValidationResult: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogValidationResult)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						if err := db.Create(tx).Error; err != nil {
							fmt.Println("Create tx err ", err)
						}
					} else if err != nil {
						fmt.Println("save TypeValidationResult err ", err)
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
	_models.TypeGuardianReward: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.GuardianReward)
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
						res := db.Model(tx).Association("GuardianRewards").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	_models.TypeCallbackTriggeredFor: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogCallbackTriggeredFor)
					if !ok {
						continue
					}
					if tx.Hash != log.TransactionHash || tx.BlockNumber != log.Event.BlockNumber {
						continue
					}
					if err := db.Where("hash = ?", tx.Hash).First(tx).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(tx)
					} else if err != nil {
						fmt.Println("save TypeCallbackTriggeredFor err ", err)
					}
					if err := db.Where("transaction_hash = ? AND log_index = ?", tx.Hash, log.Event.LogIndex).First(log).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(log)
						res := db.Model(tx).Association("LogCallbackTriggeredFors").Append(log)
						if res.Error != nil {
							fmt.Println("res ", res.Error)
						}
					}
				}
			}
		}()
		return errc
	},
	_models.TypeError: func(ctx context.Context, db *gorm.DB, eventc chan []interface{}) chan error {
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
					tx, ok := event[0].(*_models.Transaction)
					if !ok {
						continue
					}
					log, ok := event[1].(*_models.LogError)
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
						res := db.Model(tx).Association("LogErrors").Append(log)
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
