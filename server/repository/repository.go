package repository

import (
	//	"context"
	"fmt"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/jinzhu/gorm"
)

func Connect(user, password, dbName string) *gorm.DB {
	postgres_url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbName)
	var db *gorm.DB
	db, err := gorm.Open("postgres", postgres_url)
	if err != nil {
		fmt.Println(err)
	}

	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON’T change existing column’s type or delete unused columns to protect your data.
	db.AutoMigrate(&models.Transaction{}, &models.LogURL{}, &models.LogRequestUserRandom{}, &models.LogNonSupportedType{}, &models.LogNonContractCall{}, &models.LogCallbackTriggeredFor{}, &models.LogRequestFromNonExistentUC{}, &models.LogUpdateRandom{}, &models.LogValidationResult{}, &models.LogInsufficientPendingNode{}, &models.LogInsufficientWorkingGroup{}, &models.LogGrouping{}, &models.LogPublicKeyAccepted{}, &models.LogPublicKeySuggested{}, &models.LogGroupDissolve{}, &models.LogRegisteredNewPendingNode{}, &models.LogGroupingInitiated{}, &models.LogNoPendingGroup{}, &models.LogPendingGroupRemoved{}, &models.LogError{}, &models.UpdateGroupToPick{}, &models.UpdateGroupSize{}, &models.UpdateGroupingThreshold{}, &models.UpdateGroupMaturityPeriod{}, &models.UpdateBootstrapCommitDuration{}, &models.UpdateBootstrapRevealDuration{}, &models.UpdatebootstrapStartThreshold{}, &models.UpdatePendingGroupMaxLife{}, &models.GuardianReward{})
	//db.AutoMigrate(&LogURL{}, &LogRequestUserRandom{}, &LogNonSupportedType{}, &LogNonContractCall{}, &LogCallbackTriggeredFor{}, &LogRequestFromNonExistentUC{}, &LogUpdateRandom{}, &LogValidationResult{}, &LogInsufficientPendingNode{}, &LogInsufficientWorkingGroup{}, &LogGrouping{}, &LogPublicKeyAccepted{}, &LogGroupDissolve{}, &LogRegisteredNewPendingNode{}, &LogGroupingInitiated{}, &LogNoPendingGroup{}, &LogPendingGroupRemoved{}, &LogError{}, &UpdateGroupToPick{}, &UpdateGroupSize{}, &UpdateGroupingThreshold{}, &UpdateGroupMaturityPeriod{}, &UpdateBootstrapCommitDuration{}, &UpdateBootstrapRevealDuration{}, &UpdatebootstrapStartThreshold{}, &UpdatePendingGroupMaxLife{}, &GuardianReward{})
	fmt.Println("DB Connected")
	return db
}

// PostRepo explain...
type EventsRepo interface {
	LoadEvent(limit, offset int, event string) []interface{}
	SearchRelatedEvents(limit int, field, condition string) []interface{}
}
