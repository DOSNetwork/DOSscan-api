package dosproxy

import (
	"fmt"
	"testing"

	//"database/sql"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
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

func TestRandomRequestJoin(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	var results []models.UserRandomRequest
	tempDb := db.Table("log_request_user_randoms").Select("log_request_user_randoms.request_id, log_request_user_randoms.dispatched_group_id,transactions.sender, transactions.block_number,transactions.hash,log_validation_results.message,log_validation_results.signature,log_validation_results.pub_key,log_validation_results.pass")
	tempDb = tempDb.Joins("inner join log_validation_results on log_validation_results.request_id = log_request_user_randoms.request_id")
	tempDb = tempDb.Joins("inner join transactions on log_validation_results.transaction_id = transactions.id").Find(&results)
	fmt.Println(len(results))

	for _, request := range results {
		db.Where(request).FirstOrCreate(&request)

		var group models.Group
		if err := db.Where(&models.Group{GroupId: request.DispatchedGroupId}).First(&group).Error; gorm.IsRecordNotFoundError(err) {
			fmt.Println("Can't find group  ", request.DispatchedGroupId)
		} else {
			res := db.Model(&group).Association("UserRandomRequests").Append(&request)
			if res.Error != nil {
				fmt.Println("res ", res.Error)
			}
			fmt.Println(group.ID, " ", group.GroupId, "-", " len ", db.Model(&group).Association("UserRandomRequests").Count())
		}
	}
}

func TestRequestJoin(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")

	var results []models.UrlRequest
	tempDb := db.Table("log_urls").Select("log_urls.request_id, log_urls.dispatched_group_id,transactions.sender, transactions.block_number,transactions.hash,log_validation_results.message,log_validation_results.signature,log_validation_results.pub_key,log_validation_results.pass,log_urls.timeout,log_urls.data_source,log_urls.selector,log_urls.randomness")
	tempDb = tempDb.Joins("inner join log_validation_results on log_validation_results.request_id = log_urls.request_id")
	tempDb = tempDb.Joins("inner join transactions on log_validation_results.transaction_id = transactions.id").Find(&results)

	fmt.Println(len(results))
	fmt.Println(results[0].RequestId)
	fmt.Println(results[0].DispatchedGroupId)
	fmt.Println(results[0].Pass)
	fmt.Println(results[0].Sender)
	fmt.Println(len(results[0].RequestResult.Signature))
	fmt.Println(len(results[0].RequestResult.PubKey))
	for _, request := range results {
		fmt.Println(request.RequestId)
		fmt.Println(request.DispatchedGroupId)
		fmt.Println(request.Pass)
		fmt.Println(request.Sender)
		fmt.Println(request.Signature[0])
		fmt.Println(request.Signature[1])
		fmt.Println(request.PubKey[0])
		fmt.Println(request.PubKey[1])
		fmt.Println(request.PubKey[2])
		fmt.Println(request.PubKey[3])
		db.Where(request).FirstOrCreate(&request)

		var group models.Group
		db.Where(&models.Group{GroupId: request.DispatchedGroupId}).First(&group)
		res := db.Model(&group).Association("UrlRequests").Append(&request)
		if res.Error != nil {
			fmt.Println("res ", res.Error)
		}
		fmt.Println(group.GroupId, "-", " len ", db.Model(&group).Association("UrlRequests").Count())
	}
}

func TestRequestRelate(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	var urls []models.UrlRequest
	var group models.Group
	db.Where(&models.Group{GroupId: "0x6eca191becad586a651b87c343f3a737cdbbeb92efec98ab91e9dbf59524b94d"}).First(&group)
	db.Model(&group).Related(&urls, "UrlRequests")
	fmt.Println(len(urls))
}

func TestNodeJoin(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")

	//1)Get node add from LogRegisteredNewPendingNode
	var addrs []string
	db.Model(&models.LogRegisteredNewPendingNode{}).Pluck("node", &addrs)
	fmt.Println(len(addrs))
	addrs = removeDuplicates(addrs)
	fmt.Println(len(addrs))

	for i := 0; i < len(addrs); i++ {
		var node models.Node
		db.Where(models.Node{Addr: addrs[i]}).FirstOrCreate(&node)

		//2)Find Group has node addr in node_id
		sel := "SELECT group_id FROM groups WHERE $1 <@ node_id"
		rows, err := db.DB().Query(sel, pq.Array(addrs[i:i+1]))
		if err != nil {
			fmt.Println(err)
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
}

func TestNodeRelate(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	var node models.Node
	db.Find(&node, 1)
	fmt.Println(node.Addr)

	var groups []models.Group
	db.Model(&node).Related(&groups, "Groups")
	fmt.Println(len(groups))

	var nodes []models.Node
	db.Model(&groups[0]).Related(&nodes, "Nodes")
	fmt.Println(len(nodes))
}

func TestInfoGroupJoin(t *testing.T) {
	var results []models.Group
	db := initDB("postgres", "postgres", "postgres")
	tempDb := db.Table("log_groupings").Select("log_groupings.group_id, log_public_key_accepteds.accepted_blk_num,log_group_dissolves.dissolved_blk_num, log_groupings.node_id, log_public_key_accepteds.pub_key")
	tempDb = tempDb.Joins("left join log_public_key_accepteds on log_public_key_accepteds.group_id = log_groupings.group_id")
	tempDb = tempDb.Joins("left join log_group_dissolves on log_group_dissolves.group_id = log_groupings.group_id")
	tempDb = tempDb.Where("log_groupings.group_id = ? ", "0x1332").Find(&results)
	fmt.Println(len(results))
	for _, group := range results {
		var existGroup models.Group
		if err := db.Where("group_id = ?", group.GroupId).First(&existGroup).Error; gorm.IsRecordNotFoundError(err) {
			db.Create(&group)
		} else {
			db.Model(&existGroup).Omit("group_id").Updates(&group)
			fmt.Println("Update group ", existGroup.GroupId, group.DissolvedBlkNum, group.AcceptedBlkNum)
		}
	}
}

func TestTxRelate(t *testing.T) {
	db := initDB("postgres", "postgres", "postgres")
	var tx models.Transaction
	db.Find(&tx, 4)
	fmt.Println(tx.Hash)
	var logurls []models.LogUrl
	db.Model(&tx).Related(&logurls)
	fmt.Println(len(logurls))
	fmt.Println(tx.Method, " ", logurls[0].Event.EventLog)
}
