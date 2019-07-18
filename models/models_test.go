package models

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
)

func initDB(user, password, dbName string) *gorm.DB {
	postgres_url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbName)
	var db *gorm.DB
	db, err := gorm.Open("postgres", postgres_url)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

type Result struct {
	GroupId     string
	BlockNumber uint64
}

func TestJoin(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	sel := "SELECT group_id FROM groups WHERE $1 <@ node_id"
	node_id := []string{"0xc4B5086bB6896352A84b0359a00228D67A9c4c2c", "0x62C7184cDA33cd61D7725783b62103D07f6b0FC3"} // <-- two elements now
	rows, err := db.Query(sel, pq.Array(tags))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
}

func TestRelate(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	var groupInfo GroupInfo
	db.Find(&groupInfo, GroupInfo{GroupId: "0x51bc658ca734ba7e6c2fda2df44f8a17c1324435a374a90c760310bb2b7b6fb4"})
	db.Model(&groupInfo).Related(&groupInfo.LogUrl)
	fmt.Println(len(groupInfo.LogUrl))
	for _, url := range groupInfo.LogUrl {
		fmt.Println(url.DispatchedGroupId)
		//fmt.Println(url.GroupInfoID)
	}
}
func TestBlockNum(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")

	var event interface{}
	event = LogPublicKeyAccepted{}
	var blkNums []uint64
	db.Limit(1).Order("block_number desc").Find(&event).Pluck("block_number", &blkNums)
	fmt.Println(blkNums[0])
}
func Test123(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	var count int
	db.Model(&LogPublicKeyAccepted{}).Count(&count)
	fmt.Println("LogPublicKeyAccepted : ", count)
	accepted := &LogPublicKeyAccepted{}
	db.Limit(1).Order("block_number desc").Find(accepted)
	fmt.Println("LogPublicKeyAccepted : ", accepted.BlockNumber)

	db.Model(&LogGrouping{}).Count(&count)
	fmt.Println("LogGrouping : ", count)
	grouping := &LogGrouping{}
	db.Limit(1).Order("block_number desc").Find(grouping)
	fmt.Println("LogGrouping : ", grouping.BlockNumber)

	db.Model(&LogGroupDissolve{}).Count(&count)
	fmt.Println("LogGroupDissolve : ", count)
	dissolve := &LogGroupDissolve{}
	db.Limit(1).Order("block_number desc").Find(dissolve)
	fmt.Println("LogGroupDissolve : ", dissolve.BlockNumber)

	keyAcceptedLogs := []LogPublicKeyAccepted{}
	if err := db.Where("block_number <= ?", grouping.BlockNumber).Find(&keyAcceptedLogs).Error; gorm.IsRecordNotFoundError(err) {
		fmt.Println(err)
	}
	for _, keyAccepted := range keyAcceptedLogs {
		groupingLogs := []LogGrouping{}
		if err := db.Where("group_id = ?", keyAccepted.GroupId).Find(&groupingLogs).Error; gorm.IsRecordNotFoundError(err) {
			t.Errorf("Test123 Error : %s.", err.Error())
		}
		if len(groupingLogs) != 1 {
			t.Errorf("Test123 Error : %d %s %v", len(groupingLogs), keyAccepted.GroupId, keyAccepted.Event.BlockNumber)
		}
	}

}
