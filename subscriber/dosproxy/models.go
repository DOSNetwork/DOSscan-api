package dosproxy

import (
	"context"
	"fmt"

	"github.com/DOSNetwork/explorer-Api/models"
	"github.com/jinzhu/gorm"
)

func FromBlockNumber(ctx context.Context, event string, db *gorm.DB) (chan uint64, chan error) {
	out := make(chan uint64)
	errc := make(chan error)
	go func() {
		var lastBlkNum uint64

		latestRecord := fmt.Sprintf("SELECT block_number FROM %s ORDER BY block_number DESC LIMIT 1;", event)
		rows, err := db.Raw(latestRecord).Rows() // (*sql.Rows, error)
		if err != nil {
			fmt.Println(event, " : lastblock err", err)
			lastBlkNum = 0
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&lastBlkNum)
		}
		if lastBlkNum < 4468402 {
			lastBlkNum = 4468402
		}
		fmt.Println(event, " : lastblock ", lastBlkNum)
		select {
		case <-ctx.Done():
		case out <- lastBlkNum:
		}
	}()
	return out, errc
}

var ModelsTable = []func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error{
	0: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					fmt.Println("DosproxyLogUrl got event ")

					log, ok := event.(*DosproxyLogUrl)
					if !ok {
						continue
					}
					fmt.Println("DosproxyLogUrl got event ", log.Raw.BlockNumber)

					mLog := models.LogURL{
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:            fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:           log.Raw.TxIndex,
						LogIndex:          log.Raw.Index,
						Removed:           log.Raw.Removed,
						QueryId:           fmt.Sprintf("%x", log.QueryId),
						Timeout:           fmt.Sprintf("%x", log.Timeout),
						DataSource:        fmt.Sprintf("%x", log.DataSource),
						Selector:          fmt.Sprintf("%x", log.Selector),
						Randomness:        fmt.Sprintf("%x", log.Randomness),
						DispatchedGroupId: fmt.Sprintf("%x", log.DispatchedGroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	1: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogRequestUserRandom)
					if !ok {
						continue
					}

					mLog := models.LogRequestUserRandom{
						BlockNumber:          log.Raw.BlockNumber,
						BlockHash:            fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:               fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:              log.Raw.TxIndex,
						LogIndex:             log.Raw.Index,
						Removed:              log.Raw.Removed,
						RequestId:            fmt.Sprintf("%x", log.RequestId),
						LastSystemRandomness: fmt.Sprintf("%x", log.LastSystemRandomness),
						UserSeed:             fmt.Sprintf("%x", log.UserSeed),
						DispatchedGroupId:    fmt.Sprintf("%x", log.DispatchedGroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	2: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogNonSupportedType)
					if !ok {
						continue
					}

					mLog := models.LogNonSupportedType{
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:          fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						InvalidSelector: fmt.Sprintf("%x", log.InvalidSelector),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	3: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogNonContractCall)
					if !ok {
						continue
					}

					mLog := models.LogNonContractCall{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						From:        fmt.Sprintf("%x", log.From.Big()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	4: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogCallbackTriggeredFor)
					if !ok {
						continue
					}

					mLog := models.LogCallbackTriggeredFor{
						BlockNumber:  log.Raw.BlockNumber,
						BlockHash:    fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:       fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:      log.Raw.TxIndex,
						LogIndex:     log.Raw.Index,
						Removed:      log.Raw.Removed,
						CallbackAddr: fmt.Sprintf("%x", log.CallbackAddr.Big()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	5: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogRequestFromNonExistentUC)
					if !ok {
						continue
					}

					mLog := models.LogRequestFromNonExistentUC{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	6: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogUpdateRandom)
					if !ok {
						continue
					}

					mLog := models.LogUpdateRandom{
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:            fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:           log.Raw.TxIndex,
						LogIndex:          log.Raw.Index,
						Removed:           log.Raw.Removed,
						LastRandomness:    fmt.Sprintf("%x", log.LastRandomness),
						DispatchedGroupId: fmt.Sprintf("%x", log.DispatchedGroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	7: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogValidationResult)
					if !ok {
						continue
					}

					mLog := models.LogValidationResult{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						TrafficType: log.TrafficType,
						TrafficId:   fmt.Sprintf("%x", log.TrafficId),
						Message:     log.Message,
						Signature:   []string{fmt.Sprintf("%x", log.Signature[0]), fmt.Sprintf("%x", log.Signature[1])},
						PubKey:      []string{fmt.Sprintf("%x", log.PubKey[0]), fmt.Sprintf("%x", log.PubKey[1]), fmt.Sprintf("%x", log.PubKey[2]), fmt.Sprintf("%x", log.PubKey[3])},
						Pass:        log.Pass,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	8: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogInsufficientPendingNode)
					if !ok {
						continue
					}

					mLog := models.LogInsufficientPendingNode{
						BlockNumber:     log.Raw.BlockNumber,
						BlockHash:       fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:          fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:         log.Raw.TxIndex,
						LogIndex:        log.Raw.Index,
						Removed:         log.Raw.Removed,
						NumPendingNodes: log.NumPendingNodes.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	9: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					fmt.Println("DosproxyLogInsufficientWorkingGroup")
					log, ok := event.(*DosproxyLogInsufficientWorkingGroup)
					if !ok {
						continue
					}

					mLog := models.LogInsufficientWorkingGroup{
						BlockNumber:      log.Raw.BlockNumber,
						BlockHash:        fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:           fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:          log.Raw.TxIndex,
						LogIndex:         log.Raw.Index,
						Removed:          log.Raw.Removed,
						NumWorkingGroups: log.NumWorkingGroups.Uint64(),
						NumPendingGroups: log.NumPendingGroups.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	10: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogGrouping)
					if !ok {
						continue
					}
					var nodeIdstr []string
					for _, n := range log.NodeId {
						nodeIdstr = append(nodeIdstr, fmt.Sprintf("%x", n.Big()))
					}
					mLog := models.LogGrouping{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						GroupId:     fmt.Sprintf("%x", log.GroupId),
						NodeId:      nodeIdstr,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	11: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogPublicKeyAccepted)
					if !ok {
						continue
					}

					mLog := models.LogPublicKeyAccepted{
						BlockNumber:      log.Raw.BlockNumber,
						BlockHash:        fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:           fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:          log.Raw.TxIndex,
						LogIndex:         log.Raw.Index,
						Removed:          log.Raw.Removed,
						GroupId:          fmt.Sprintf("%x", log.GroupId),
						PubKey:           []string{fmt.Sprintf("%x", log.PubKey[0]), fmt.Sprintf("%x", log.PubKey[1]), fmt.Sprintf("%x", log.PubKey[2]), fmt.Sprintf("%x", log.PubKey[3])},
						NumWorkingGroups: log.NumWorkingGroups.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	12: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogPublicKeySuggested)
					if !ok {
						continue
					}

					mLog := models.LogPublicKeySuggested{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						GroupId:     fmt.Sprintf("%x", log.GroupId),
						PubKeyCount: log.PubKeyCount.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	13: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogGroupDissolve)
					if !ok {
						continue
					}

					mLog := models.LogGroupDissolve{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						GroupId:     fmt.Sprintf("%x", log.GroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	14: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogRegisteredNewPendingNode)
					if !ok {
						continue
					}

					mLog := models.LogRegisteredNewPendingNode{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						Node:        fmt.Sprintf("%x", log.Node),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	15: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogGroupingInitiated)
					if !ok {
						continue
					}

					mLog := models.LogGroupingInitiated{
						BlockNumber:       log.Raw.BlockNumber,
						BlockHash:         fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:            fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:           log.Raw.TxIndex,
						LogIndex:          log.Raw.Index,
						Removed:           log.Raw.Removed,
						PendingNodePool:   log.PendingNodePool.Uint64(),
						Groupsize:         log.Groupsize.Uint64(),
						Groupingthreshold: log.Groupingthreshold.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	16: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogNoPendingGroup)
					if !ok {
						continue
					}

					mLog := models.LogNoPendingGroup{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						GroupId:     fmt.Sprintf("%x", log.GroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	17: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogPendingGroupRemoved)
					if !ok {
						continue
					}

					mLog := models.LogPendingGroupRemoved{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						GroupId:     fmt.Sprintf("%x", log.GroupId),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	18: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyLogError)
					if !ok {
						continue
					}

					mLog := models.LogError{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						Err:         log.Err,
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	19: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateGroupToPick)
					if !ok {
						continue
					}

					mLog := models.UpdateGroupToPick{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						OldNum:      log.OldNum.Uint64(),
						NewNum:      log.NewNum.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	20: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateGroupSize)
					if !ok {
						continue
					}

					mLog := models.UpdateGroupSize{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						OldSize:     log.OldSize.Uint64(),
						NewSize:     log.NewSize.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	21: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateGroupingThreshold)
					if !ok {
						continue
					}

					mLog := models.UpdateGroupingThreshold{
						BlockNumber:  log.Raw.BlockNumber,
						BlockHash:    fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:       fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:      log.Raw.TxIndex,
						LogIndex:     log.Raw.Index,
						Removed:      log.Raw.Removed,
						OldThreshold: log.OldThreshold.Uint64(),
						NewThreshold: log.NewThreshold.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	22: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateGroupMaturityPeriod)
					if !ok {
						continue
					}

					mLog := models.UpdateGroupMaturityPeriod{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						OldPeriod:   log.OldPeriod.Uint64(),
						NewPeriod:   log.NewPeriod.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	23: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateBootstrapCommitDuration)
					if !ok {
						continue
					}

					mLog := models.UpdateBootstrapCommitDuration{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						OldDuration: log.OldDuration.Uint64(),
						NewDuration: log.NewDuration.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	24: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdateBootstrapRevealDuration)
					if !ok {
						continue
					}

					mLog := models.UpdateBootstrapRevealDuration{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						OldDuration: log.OldDuration.Uint64(),
						NewDuration: log.NewDuration.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	25: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdatebootstrapStartThreshold)
					if !ok {
						continue
					}

					mLog := models.UpdatebootstrapStartThreshold{
						BlockNumber:  log.Raw.BlockNumber,
						BlockHash:    fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:       fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:      log.Raw.TxIndex,
						LogIndex:     log.Raw.Index,
						Removed:      log.Raw.Removed,
						OldThreshold: log.OldThreshold.Uint64(),
						NewThreshold: log.NewThreshold.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	26: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyUpdatePendingGroupMaxLife)
					if !ok {
						continue
					}

					mLog := models.UpdatePendingGroupMaxLife{
						BlockNumber:   log.Raw.BlockNumber,
						BlockHash:     fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:        fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:       log.Raw.TxIndex,
						LogIndex:      log.Raw.Index,
						Removed:       log.Raw.Removed,
						OldLifeBlocks: log.OldLifeBlocks.Uint64(),
						NewLifeBlocks: log.NewLifeBlocks.Uint64(),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
	27: func(ctx context.Context, db *gorm.DB, eventc chan interface{}) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case event, ok := <-eventc:
					if !ok {
						return
					}
					log, ok := event.(*DosproxyGuardianReward)
					if !ok {
						continue
					}

					mLog := models.GuardianReward{
						BlockNumber: log.Raw.BlockNumber,
						BlockHash:   fmt.Sprintf("%x", log.Raw.BlockHash.Big()),
						TxHash:      fmt.Sprintf("%x", log.Raw.TxHash.Big()),
						TxIndex:     log.Raw.TxIndex,
						LogIndex:    log.Raw.Index,
						Removed:     log.Raw.Removed,
						BlkNum:      log.BlkNum.Uint64(),
						Guardian:    fmt.Sprintf("%x", log.Guardian.Big()),
					}
					if err := db.Where("block_number = ? AND log_index = ?", log.Raw.BlockNumber, log.Raw.Index).First(&mLog).Error; gorm.IsRecordNotFoundError(err) {
						db.Create(&mLog)
						fmt.Println("Saved Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					} else {
						fmt.Println("duplicate Event Log: ", log.Raw.BlockNumber, log.Raw.Index)
					}
				}
			}
		}()
		return errc
	},
}
