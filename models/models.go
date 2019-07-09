package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

type Transaction struct {
	gorm.Model                    `json:"-"`
	Hash                          string                          `gorm:"primary_key;index"`
	GasPrice                      uint64                          `gorm:"column:gas_price" json:"gasPrice"`
	Value                         uint64                          `gorm:"column:value" json:"value"`
	GasLimit                      uint64                          `gorm:"column:gas_limit" json:"gasLimit"`
	Nonce                         uint64                          `gorm:"column:nonce" json:"nonce"`
	To                            string                          `gorm:"column:to" json:"to"`
	Data                          []byte                          `gorm:"column:data;type:bytea" json:"data"`
	Method                        string                          `gorm:"column:method;index" json:"method"`
	LogURL                        []LogURL                        `gorm:"foreignkey:TransactionHash"`
	LogRequestUserRandom          []LogRequestUserRandom          `gorm:"foreignkey:TransactionHash"`
	LogNonSupportedType           []LogNonSupportedType           `gorm:"foreignkey:TransactionHash"`
	LogNonContractCall            []LogNonContractCall            `gorm:"foreignkey:TransactionHash"`
	LogCallbackTriggeredFor       []LogCallbackTriggeredFor       `gorm:"foreignkey:TransactionHash"`
	LogRequestFromNonExistentUC   []LogRequestFromNonExistentUC   `gorm:"foreignkey:TransactionHash"`
	LogUpdateRandom               []LogUpdateRandom               `gorm:"foreignkey:TransactionHash"`
	LogValidationResult           []LogValidationResult           `gorm:"foreignkey:TransactionHash"`
	LogInsufficientPendingNode    []LogInsufficientPendingNode    `gorm:"foreignkey:TransactionHash"`
	LogInsufficientWorkingGroup   []LogInsufficientWorkingGroup   `gorm:"foreignkey:TransactionHash"`
	LogGrouping                   []LogGrouping                   `gorm:"foreignkey:TransactionHash"`
	LogPublicKeyAccepted          []LogPublicKeyAccepted          `gorm:"foreignkey:TransactionHash"`
	LogPublicKeySuggested         []LogPublicKeySuggested         `gorm:"foreignkey:TransactionHash"`
	LogGroupDissolve              []LogGroupDissolve              `gorm:"foreignkey:TransactionHash"`
	LogRegisteredNewPendingNode   []LogRegisteredNewPendingNode   `gorm:"foreignkey:TransactionHash"`
	LogGroupingInitiated          []LogGroupingInitiated          `gorm:"foreignkey:TransactionHash"`
	LogNoPendingGroup             []LogNoPendingGroup             `gorm:"foreignkey:TransactionHash"`
	LogPendingGroupRemoved        []LogPendingGroupRemoved        `gorm:"foreignkey:TransactionHash"`
	LogError                      []LogError                      `gorm:"foreignkey:TransactionHash"`
	UpdateGroupToPick             []UpdateGroupToPick             `gorm:"foreignkey:TransactionHash"`
	UpdateGroupSize               []UpdateGroupSize               `gorm:"foreignkey:TransactionHash"`
	UpdateGroupingThreshold       []UpdateGroupingThreshold       `gorm:"foreignkey:TransactionHash"`
	UpdateGroupMaturityPeriod     []UpdateGroupMaturityPeriod     `gorm:"foreignkey:TransactionHash"`
	UpdateBootstrapCommitDuration []UpdateBootstrapCommitDuration `gorm:"foreignkey:TransactionHash"`
	UpdateBootstrapRevealDuration []UpdateBootstrapRevealDuration `gorm:"foreignkey:TransactionHash"`
	UpdatebootstrapStartThreshold []UpdatebootstrapStartThreshold `gorm:"foreignkey:TransactionHash"`
	UpdatePendingGroupMaxLife     []UpdatePendingGroupMaxLife     `gorm:"foreignkey:TransactionHash"`
	GuardianReward                []GuardianReward                `gorm:"foreignkey:TransactionHash"`
}

// Set User's table name to be `transaction`
func (Transaction) TableName() string {
	return "transaction"
}

// `LogURL` belongs to `Transaction`, `TransactionID` is the foreign key
type LogURL struct {
	gorm.Model        `json:"-"`
	Method            string         `gorm:"column:method" json:"method"`
	EventLog          string         `gorm:"column:event_log" json:"eventLog"`
	Transaction       Transaction    `gorm:"association_foreignkey:Hash" json:"-"`
	TransactionHash   string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex           uint           `gorm:"column:transaction_index" json:"-"`
	Topics            pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber       uint64         `gorm:"column:block_number" json:"-"`
	BlockHash         string         `gorm:"column:block_hash" json:"-"`
	LogIndex          uint           `gorm:"column:log_index;" json:"-"`
	Removed           bool           `gorm:"column:removed;" json:"-"`
	Date              time.Time      `gorm:"column:date;" json:"-"`
	QueryId           string         `gorm:"column:query_id" json:"queryId"`
	Timeout           string         `gorm:"column:time_out" json:"timeOut"`
	DataSource        string         `gorm:"column:data_source" json:"dataSource"`
	Selector          string         `gorm:"column:selector" json:"selector"`
	Randomness        string         `gorm:"column:randomness" json:"randomness"`
	DispatchedGroupId string         `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
}

// Set User's table name to be `logurl`
func (LogURL) TableName() string {
	return "logurl"
}

type LogRequestUserRandom struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`

	RequestId            string `gorm:"column:request_id" json:"requestId"`
	LastSystemRandomness string `gorm:"column:last_system_randomness" json:"lastSystemRandomness"`
	UserSeed             string `gorm:"column:user_seed" json:"userSeed"`
	DispatchedGroupId    string `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
}

// Set User's table name to be `profiles`
func (LogRequestUserRandom) TableName() string {
	return "logrequestuserrandom"
}

type LogNonSupportedType struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index"json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	InvalidSelector string         `gorm:"column:invalid_selector" json:"invalidSelector"`
}

// Set User's table name to be `profiles`
func (LogNonSupportedType) TableName() string {
	return "lognonsupportedtype"
}

type LogNonContractCall struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	From            string         `gorm:"column:from" json:"from"`
}

// Set User's table name to be `profiles`
func (LogNonContractCall) TableName() string {
	return "lognoncontractcall"
}

type LogCallbackTriggeredFor struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	CallbackAddr    string         `gorm:"column:call_back_addr" json:"callbackAddr"`
}

// Set User's table name to be `profiles`
func (LogCallbackTriggeredFor) TableName() string {
	return "logcallbacktriggeredfor"
}

type LogRequestFromNonExistentUC struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
}

// Set User's table name to be `profiles`
func (LogRequestFromNonExistentUC) TableName() string {
	return "logrequestfromnonexistentuc"
}

type LogUpdateRandom struct {
	gorm.Model        `json:"-"`
	Method            string         `gorm:"column:method" json:"method"`
	EventLog          string         `gorm:"column:event_log" json:"eventLog"`
	Topics            pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber       uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash         string         `gorm:"column:block_hash" json:"-"`
	TransactionHash   string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex           uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex          uint           `gorm:"column:log_index;" json:"-"`
	Removed           bool           `gorm:"column:removed;" json:"-"`
	Date              time.Time      `gorm:"column:date;" json:"-"`
	LastRandomness    string         `gorm:"column:last_randomness" json:"lastRandomness"`
	DispatchedGroupId string         `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
}

// Set User's table name to be `profiles`
func (LogUpdateRandom) TableName() string {
	return "logupdaterandom"
}

type LogValidationResult struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	TrafficType     uint8          `gorm:"column:traffic_type"`
	TrafficId       string         `gorm:"column:traffic_id"`
	Message         []byte         `gorm:"column:message;type:bytea" json:"message"`
	Signature       pq.StringArray `gorm:"column:signature;type:varchar(100)[]" json:"signature"`
	PubKey          pq.StringArray `gorm:"column:pub_key;type:varchar(100)[]" json:"pubKey"`
	Pass            bool           `gorm:"column:pass;" json:"pass"`
}

// Set User's table name to be `profiles`
func (LogValidationResult) TableName() string {
	return "logvalidationresult"
}

type LogInsufficientPendingNode struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	NumPendingNodes uint64         `gorm:"column:num_pending_nodes" json:"numPendingNodes"`
}

// Set User's table name to be `profiles`
func (LogInsufficientPendingNode) TableName() string {
	return "loginsufficientpendingnode"
}

type LogInsufficientWorkingGroup struct {
	gorm.Model       `json:"-"`
	Method           string         `gorm:"column:method" json:"method"`
	EventLog         string         `gorm:"column:event_log" json:"eventLog"`
	Topics           pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber      uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash        string         `gorm:"column:block_hash" json:"-"`
	TransactionHash  string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex          uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex         uint           `gorm:"column:log_index;" json:"-"`
	Removed          bool           `gorm:"column:removed;" json:"-"`
	Date             time.Time      `gorm:"column:date;" json:"-"`
	NumWorkingGroups uint64         `gorm:"column:num_working_groups" json:"numWorkingGroups"`
	NumPendingGroups uint64         `gorm:"column:num_pending_groups" json:"numPendingGroups"`
}

// Set User's table name to be `profiles`
func (LogInsufficientWorkingGroup) TableName() string {
	return "loginsufficientworkinggroup"
}

type LogGrouping struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	GroupId         string         `gorm:"column:group_id" json:"groupId"`
	NodeId          pq.StringArray `gorm:"column:node_id;type:varchar(100)[]" json:"nodeId"`
}

// Set User's table name to be `profiles`
func (LogGrouping) TableName() string {
	return "loggrouping"
}

type LogPublicKeyAccepted struct {
	gorm.Model       `json:"-"`
	Method           string         `gorm:"column:method" json:"method"`
	EventLog         string         `gorm:"column:event_log" json:"eventLog"`
	Topics           pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber      uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash        string         `gorm:"column:block_hash" json:"-"`
	TransactionHash  string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex          uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex         uint           `gorm:"column:log_index;" json:"-"`
	Removed          bool           `gorm:"column:removed;" json:"-"`
	Date             time.Time      `gorm:"column:date;" json:"-"`
	GroupId          string         `gorm:"column:group_id" json:"groupId"`
	PubKey           pq.StringArray `gorm:"column:pub_key;type:varchar(100)[]" json:"pubKey"`
	NumWorkingGroups uint64         `gorm:"column:num_working_groups;" json:"numWorkingGroups"`
}

// Set User's table name to be `profiles`
func (LogPublicKeyAccepted) TableName() string {
	return "logpublickeyaccepted"
}

type LogPublicKeySuggested struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	GroupId         string         `gorm:"column:group_id" json:"groupId"`
	PubKeyCount     uint64         `gorm:"column:pub_key_count" json:"pubKeyCount"`
}

// Set User's table name to be `profiles`
func (LogPublicKeySuggested) TableName() string {
	return "logpublickeysuggested"
}

type LogGroupDissolve struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	GroupId         string         `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogGroupDissolve) TableName() string {
	return "loggroupdissolve"
}

type LogRegisteredNewPendingNode struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	Node            string         `gorm:"column:node" json:"node"`
}

// Set User's table name to be `profiles`
func (LogRegisteredNewPendingNode) TableName() string {
	return "logregisterednewpendingnode"
}

type LogGroupingInitiated struct {
	gorm.Model        `json:"-"`
	Method            string         `gorm:"column:method" json:"method"`
	EventLog          string         `gorm:"column:event_log" json:"eventLog"`
	Topics            pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber       uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash         string         `gorm:"column:block_hash" json:"-"`
	TransactionHash   string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex           uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex          uint           `gorm:"column:log_index;" json:"-"`
	Removed           bool           `gorm:"column:removed;" json:"-"`
	Date              time.Time      `gorm:"column:date;" json:"-"`
	PendingNodePool   uint64         `gorm:"column:pending_node_pool" json:"pendingNodePool"`
	Groupsize         uint64         `gorm:"column:group_size" json:"groupSize"`
	Groupingthreshold uint64         `gorm:"column:grouping_threshold" json:"groupingThreshold"`
}

// Set User's table name to be `profiles`
func (LogGroupingInitiated) TableName() string {
	return "loggroupinginitiated"
}

type LogNoPendingGroup struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	GroupId         string         `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogNoPendingGroup) TableName() string {
	return "lognopendinggroup"
}

type LogPendingGroupRemoved struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	GroupId         string         `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogPendingGroupRemoved) TableName() string {
	return "logpendinggroupremoved"
}

type LogError struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	Err             string         `gorm:"column:err" json:"err"`
}

func (LogError) TableName() string {
	return "logerror"
}

type UpdateGroupToPick struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	OldNum          uint64         `gorm:"column:old_num" json:"oldNum"`
	NewNum          uint64         `gorm:"column:new_num" json:"newNum"`
}

func (UpdateGroupToPick) TableName() string {
	return "updategrouptopick"
}

type UpdateGroupSize struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	OldSize         uint64         `gorm:"column:old_size" json:"oldSize"`
	NewSize         uint64         `gorm:"column:new_size" json:"newSize"`
}

func (UpdateGroupSize) TableName() string {
	return "updategroupsize"
}

type UpdateGroupingThreshold struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	OldThreshold    uint64         `gorm:"column:old_threshold" json:"oldThreshold"`
	NewThreshold    uint64         `gorm:"column:new_threshold" json:"newThreshold"`
}

func (UpdateGroupingThreshold) TableName() string {
	return "updategroupingthreshold"
}

type UpdateGroupMaturityPeriod struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	OldPeriod       uint64         `gorm:"column:old_period" json:"oldPeriod"`
	NewPeriod       uint64         `gorm:"column:new_period" json:"newPeriod"`
}

func (UpdateGroupMaturityPeriod) TableName() string {
	return "updategroupmaturityperiod"
}

type UpdateBootstrapCommitDuration struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	OldDuration     uint64         `gorm:"column:old_duration" json:"oldDuration"`
	NewDuration     uint64         `gorm:"column:new_duration" json:"newDuration"`
}

func (UpdateBootstrapCommitDuration) TableName() string {
	return "updatebootstrapcommitduration"
}

type UpdateBootstrapRevealDuration struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	OldDuration     uint64         `gorm:"column:old_duration" json:"oldDuration"`
	NewDuration     uint64         `gorm:"column:new_duration" json:"newDuration"`
}

func (UpdateBootstrapRevealDuration) TableName() string {
	return "updatebootstraprevealduration"
}

type UpdatebootstrapStartThreshold struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	OldThreshold    uint64         `gorm:"column:old_threshold" json:"oldThreshold"`
	NewThreshold    uint64         `gorm:"column:new_threshold" json:"newThreshold"`
}

func (UpdatebootstrapStartThreshold) TableName() string {
	return "updatebootstrapstartthreshold"
}

type UpdatePendingGroupMaxLife struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	OldLifeBlocks   uint64         `gorm:"column:old_life_blocks" json:"oldLifeBlocks"`
	NewLifeBlocks   uint64         `gorm:"column:new_life_blocks" json:"newLifeBlock"`
}

func (UpdatePendingGroupMaxLife) TableName() string {
	return "updatependinggroupmaxlife"
}

type GuardianReward struct {
	gorm.Model      `json:"-"`
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number;" json:"-"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
	BlkNum          uint64         `gorm:"column:blk_num" json:"blkNum"`
	Guardian        string         `gorm:"column:guardian" json:"guardian"`
}

func (GuardianReward) TableName() string {
	return "guardianreward"
}

var DB *gorm.DB

func Connect() *gorm.DB {
	postgres_url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	var db *gorm.DB
	db, err := gorm.Open("postgres", postgres_url)
	if err != nil {
		log.Fatal(err)
	}
	//&LogPublicKeySuggested{},

	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON’T change existing column’s type or delete unused columns to protect your data.
	db.AutoMigrate(&Transaction{}, &LogURL{}, &LogRequestUserRandom{}, &LogNonSupportedType{}, &LogNonContractCall{}, &LogCallbackTriggeredFor{}, &LogRequestFromNonExistentUC{}, &LogUpdateRandom{}, &LogValidationResult{}, &LogInsufficientPendingNode{}, &LogInsufficientWorkingGroup{}, &LogGrouping{}, &LogPublicKeyAccepted{}, &LogPublicKeySuggested{}, &LogGroupDissolve{}, &LogRegisteredNewPendingNode{}, &LogGroupingInitiated{}, &LogNoPendingGroup{}, &LogPendingGroupRemoved{}, &LogError{}, &UpdateGroupToPick{}, &UpdateGroupSize{}, &UpdateGroupingThreshold{}, &UpdateGroupMaturityPeriod{}, &UpdateBootstrapCommitDuration{}, &UpdateBootstrapRevealDuration{}, &UpdatebootstrapStartThreshold{}, &UpdatePendingGroupMaxLife{}, &GuardianReward{})
	//db.AutoMigrate(&LogURL{}, &LogRequestUserRandom{}, &LogNonSupportedType{}, &LogNonContractCall{}, &LogCallbackTriggeredFor{}, &LogRequestFromNonExistentUC{}, &LogUpdateRandom{}, &LogValidationResult{}, &LogInsufficientPendingNode{}, &LogInsufficientWorkingGroup{}, &LogGrouping{}, &LogPublicKeyAccepted{}, &LogGroupDissolve{}, &LogRegisteredNewPendingNode{}, &LogGroupingInitiated{}, &LogNoPendingGroup{}, &LogPendingGroupRemoved{}, &LogError{}, &UpdateGroupToPick{}, &UpdateGroupSize{}, &UpdateGroupingThreshold{}, &UpdateGroupMaturityPeriod{}, &UpdateBootstrapCommitDuration{}, &UpdateBootstrapRevealDuration{}, &UpdatebootstrapStartThreshold{}, &UpdatePendingGroupMaxLife{}, &GuardianReward{})

	// DB.LogMode(true)
	log.Info("DB Connected")
	DB = db
	return db
}
