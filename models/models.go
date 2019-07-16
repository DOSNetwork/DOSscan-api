package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Transaction struct {
	gorm.Model
	Hash                          string                          `gorm:"column:hash;index"`
	GasPrice                      uint64                          `gorm:"column:gas_price" json:"gasPrice"`
	Value                         uint64                          `gorm:"column:value" json:"value"`
	GasLimit                      uint64                          `gorm:"column:gas_limit" json:"gasLimit"`
	Nonce                         uint64                          `gorm:"column:nonce" json:"nonce"`
	Sender                        string                          `gorm:"column:sender" json:"sender"`
	To                            string                          `gorm:"column:to" json:"to"`
	BlockNumber                   uint64                          `gorm:"column:block_number" json:"blockNumber"`
	Data                          []byte                          `gorm:"column:data;type:bytea" json:"data"`
	Method                        string                          `gorm:"column:method;index" json:"method"`
	LogUrl                        []LogUrl                        `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogRequestUserRandom          []LogRequestUserRandom          `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogUpdateRandom               []LogUpdateRandom               `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogValidationResult           []LogValidationResult           `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogNonSupportedType           []LogNonSupportedType           `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogNonContractCall            []LogNonContractCall            `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogCallbackTriggeredFor       []LogCallbackTriggeredFor       `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogRequestFromNonExistentUC   []LogRequestFromNonExistentUC   `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogInsufficientPendingNode    []LogInsufficientPendingNode    `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogInsufficientWorkingGroup   []LogInsufficientWorkingGroup   `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogGrouping                   []LogGrouping                   `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogPublicKeyAccepted          []LogPublicKeyAccepted          `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogPublicKeySuggested         []LogPublicKeySuggested         `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogGroupDissolve              []LogGroupDissolve              `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogRegisteredNewPendingNode   []LogRegisteredNewPendingNode   `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogGroupingInitiated          []LogGroupingInitiated          `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogNoPendingGroup             []LogNoPendingGroup             `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogPendingGroupRemoved        []LogPendingGroupRemoved        `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	LogError                      []LogError                      `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	UpdateGroupToPick             []UpdateGroupToPick             `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	UpdateGroupSize               []UpdateGroupSize               `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	UpdateGroupingThreshold       []UpdateGroupingThreshold       `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	UpdateGroupMaturityPeriod     []UpdateGroupMaturityPeriod     `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	UpdateBootstrapCommitDuration []UpdateBootstrapCommitDuration `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	UpdateBootstrapRevealDuration []UpdateBootstrapRevealDuration `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	UpdatebootstrapStartThreshold []UpdatebootstrapStartThreshold `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	UpdatePendingGroupMaxLife     []UpdatePendingGroupMaxLife     `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
	GuardianReward                []GuardianReward                `gorm:"foreignkey:transactionHash;association_foreignkey:hash"`
}

// Set User's table name to be `transaction`
func (Transaction) TableName() string {
	return "transaction"
}

type Event struct {
	gorm.Model
	Method          string         `gorm:"column:method" json:"method"`
	EventLog        string         `gorm:"column:event_log" json:"eventLog"`
	Transaction     Transaction    `gorm:"association_foreignkey:Hash" json:"-"`
	TransactionHash string         `gorm:"column:transaction_hash" json:"txHash"`
	TxIndex         uint           `gorm:"column:transaction_index" json:"-"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `gorm:"column:block_number" json:"blockNumber"`
	BlockHash       string         `gorm:"column:block_hash" json:"-"`
	LogIndex        uint           `gorm:"column:log_index;" json:"-"`
	Removed         bool           `gorm:"column:removed;" json:"-"`
	Date            time.Time      `gorm:"column:date;" json:"-"`
}

// `LogURL` belongs to `Transaction`, `TransactionID` is the foreign key
type LogUrl struct {
	Event
	QueryId             string              `gorm:"column:query_id" json:"queryId"`
	Timeout             string              `gorm:"column:time_out" json:"timeOut"`
	DataSource          string              `gorm:"column:data_source" json:"dataSource"`
	Selector            string              `gorm:"column:selector" json:"selector"`
	Randomness          string              `gorm:"column:randomness" json:"randomness"`
	DispatchedGroupId   string              `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
	LogValidationResult LogValidationResult `gorm:"foreignkey:traffic_id;association_foreignkey:query_id"`
}

// Set User's table name to be `logurl`
func (LogUrl) TableName() string {
	return "logurl"
}

type LogRequestUserRandom struct {
	Event
	RequestId            string              `gorm:"column:request_id" json:"requestId"`
	LastSystemRandomness string              `gorm:"column:last_system_randomness" json:"lastSystemRandomness"`
	UserSeed             string              `gorm:"column:user_seed" json:"userSeed"`
	DispatchedGroupId    string              `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
	LogValidationResult  LogValidationResult `gorm:"foreignkey:traffic_id;association_foreignkey:request_id"`
}

// Set User's table name to be `profiles`
func (LogRequestUserRandom) TableName() string {
	return "logrequestuserrandom"
}

type LogNonSupportedType struct {
	Event
	InvalidSelector string `gorm:"column:invalid_selector" json:"invalidSelector"`
}

// Set User's table name to be `profiles`
func (LogNonSupportedType) TableName() string {
	return "lognonsupportedtype"
}

type LogNonContractCall struct {
	Event
	CallAddr string `gorm:"column:call_ddr" json:"callAddr"`
}

// Set User's table name to be `profiles`
func (LogNonContractCall) TableName() string {
	return "lognoncontractcall"
}

type LogCallbackTriggeredFor struct {
	Event
	CallbackAddr string `gorm:"column:call_back_addr" json:"callbackAddr"`
}

// Set User's table name to be `profiles`
func (LogCallbackTriggeredFor) TableName() string {
	return "logcallbacktriggeredfor"
}

type LogRequestFromNonExistentUC struct {
	Event
}

// Set User's table name to be `profiles`
func (LogRequestFromNonExistentUC) TableName() string {
	return "logrequestfromnonexistentuc"
}

type LogUpdateRandom struct {
	Event
	LastRandomness    string `gorm:"column:last_randomness" json:"lastRandomness"`
	DispatchedGroupId string `gorm:"column:dispatched_groupid" json:"dispatchedGroupId"`
}

// Set User's table name to be `profiles`
func (LogUpdateRandom) TableName() string {
	return "logupdaterandom"
}

type LogValidationResult struct {
	Event
	TrafficId   string         `gorm:"column:traffic_id" json:"trafficId"`
	TrafficType uint8          `gorm:"column:traffic_type" json:"trafficType"`
	Message     string         `gorm:"column:message;type:bytea" json:"message"`
	Signature   pq.StringArray `gorm:"column:signature;type:varchar(100)[]" json:"signature"`
	PubKey      pq.StringArray `gorm:"column:pub_key;type:varchar(100)[]" json:"pubKey"`
	Pass        bool           `gorm:"column:pass" json:"pass"`
}

// Set User's table name to be `profiles`
func (LogValidationResult) TableName() string {
	return "logvalidationresult"
}

type LogInsufficientPendingNode struct {
	Event
	NumPendingNodes uint64 `gorm:"column:num_pending_nodes" json:"numPendingNodes"`
}

// Set User's table name to be `profiles`
func (LogInsufficientPendingNode) TableName() string {
	return "loginsufficientpendingnode"
}

type LogInsufficientWorkingGroup struct {
	Event
	NumWorkingGroups uint64 `gorm:"column:num_working_groups" json:"numWorkingGroups"`
	NumPendingGroups uint64 `gorm:"column:num_pending_groups" json:"numPendingGroups"`
}

// Set User's table name to be `profiles`
func (LogInsufficientWorkingGroup) TableName() string {
	return "loginsufficientworkinggroup"
}

type LogGrouping struct {
	Event
	GroupId string         `gorm:"column:group_id" json:"groupId"`
	NodeId  pq.StringArray `gorm:"column:node_id;type:varchar(100)[]" json:"nodeId"`
}

// Set User's table name to be `profiles`
func (LogGrouping) TableName() string {
	return "loggrouping"
}

type LogPublicKeyAccepted struct {
	Event
	GroupId          string         `gorm:"column:group_id" json:"groupId"`
	PubKey           pq.StringArray `gorm:"column:pub_key;type:varchar(100)[]" json:"pubKey"`
	NumWorkingGroups uint64         `gorm:"column:num_working_groups;" json:"numWorkingGroups"`
}

// Set User's table name to be `profiles`
func (LogPublicKeyAccepted) TableName() string {
	return "logpublickeyaccepted"
}

type LogPublicKeySuggested struct {
	Event
	GroupId     string `gorm:"column:group_id" json:"groupId"`
	PubKeyCount uint64 `gorm:"column:pub_key_count" json:"pubKeyCount"`
}

// Set User's table name to be `profiles`
func (LogPublicKeySuggested) TableName() string {
	return "logpublickeysuggested"
}

type LogGroupDissolve struct {
	Event
	GroupId string `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogGroupDissolve) TableName() string {
	return "loggroupdissolve"
}

type LogRegisteredNewPendingNode struct {
	Event
	Node string `gorm:"column:node" json:"node"`
}

// Set User's table name to be `profiles`
func (LogRegisteredNewPendingNode) TableName() string {
	return "logregisterednewpendingnode"
}

type LogGroupingInitiated struct {
	Event
	PendingNodePool   uint64 `gorm:"column:pending_node_pool" json:"pendingNodePool"`
	Groupsize         uint64 `gorm:"column:group_size" json:"groupSize"`
	Groupingthreshold uint64 `gorm:"column:grouping_threshold" json:"groupingThreshold"`
}

// Set User's table name to be `profiles`
func (LogGroupingInitiated) TableName() string {
	return "loggroupinginitiated"
}

type LogNoPendingGroup struct {
	Event
	GroupId string `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogNoPendingGroup) TableName() string {
	return "lognopendinggroup"
}

type LogPendingGroupRemoved struct {
	Event
	GroupId string `gorm:"column:group_id" json:"groupId"`
}

// Set User's table name to be `profiles`
func (LogPendingGroupRemoved) TableName() string {
	return "logpendinggroupremoved"
}

type LogError struct {
	Event
	Err string `gorm:"column:err" json:"err"`
}

func (LogError) TableName() string {
	return "logerror"
}

type UpdateGroupToPick struct {
	Event
	OldNum uint64 `gorm:"column:old_num" json:"oldNum"`
	NewNum uint64 `gorm:"column:new_num" json:"newNum"`
}

func (UpdateGroupToPick) TableName() string {
	return "updategrouptopick"
}

type UpdateGroupSize struct {
	Event
	OldSize uint64 `gorm:"column:old_size" json:"oldSize"`
	NewSize uint64 `gorm:"column:new_size" json:"newSize"`
}

func (UpdateGroupSize) TableName() string {
	return "updategroupsize"
}

type UpdateGroupingThreshold struct {
	Event
	OldThreshold uint64 `gorm:"column:old_threshold" json:"oldThreshold"`
	NewThreshold uint64 `gorm:"column:new_threshold" json:"newThreshold"`
}

func (UpdateGroupingThreshold) TableName() string {
	return "updategroupingthreshold"
}

type UpdateGroupMaturityPeriod struct {
	Event
	OldPeriod uint64 `gorm:"column:old_period" json:"oldPeriod"`
	NewPeriod uint64 `gorm:"column:new_period" json:"newPeriod"`
}

func (UpdateGroupMaturityPeriod) TableName() string {
	return "updategroupmaturityperiod"
}

type UpdateBootstrapCommitDuration struct {
	Event
	OldDuration uint64 `gorm:"column:old_duration" json:"oldDuration"`
	NewDuration uint64 `gorm:"column:new_duration" json:"newDuration"`
}

func (UpdateBootstrapCommitDuration) TableName() string {
	return "updatebootstrapcommitduration"
}

type UpdateBootstrapRevealDuration struct {
	Event
	OldDuration uint64 `gorm:"column:old_duration" json:"oldDuration"`
	NewDuration uint64 `gorm:"column:new_duration" json:"newDuration"`
}

func (UpdateBootstrapRevealDuration) TableName() string {
	return "updatebootstraprevealduration"
}

type UpdatebootstrapStartThreshold struct {
	Event
	OldThreshold uint64 `gorm:"column:old_threshold" json:"oldThreshold"`
	NewThreshold uint64 `gorm:"column:new_threshold" json:"newThreshold"`
}

func (UpdatebootstrapStartThreshold) TableName() string {
	return "updatebootstrapstartthreshold"
}

type UpdatePendingGroupMaxLife struct {
	Event
	OldLifeBlocks uint64 `gorm:"column:old_life_blocks" json:"oldLifeBlocks"`
	NewLifeBlocks uint64 `gorm:"column:new_life_blocks" json:"newLifeBlock"`
}

func (UpdatePendingGroupMaxLife) TableName() string {
	return "updatependinggroupmaxlife"
}

type GuardianReward struct {
	Event
	BlkNum   uint64 `gorm:"column:blk_num" json:"blkNum"`
	Guardian string `gorm:"column:guardian" json:"guardian"`
}

func (GuardianReward) TableName() string {
	return "guardianreward"
}

type NodeInfo struct {
	NodeAddr      string         `gorm:"column:node_addr;primary_key" json:"nodeAddr"`
	Balance       string         `json:"balance"`
	RegisterState bool           `json:"registerState"`
	GroupingIds   pq.StringArray `gorm:"column:all_group_ids;type:varchar(100)[]"`
	DissolveIds   pq.StringArray `gorm:"column:all_dissolve_ids;type:varchar(100)[]"`
}

func (NodeInfo) TableName() string {
	return "nodeinfo"
}

type GroupInfo struct {
	GroupId              string                 `gorm:"column:group_id;primary_key" json:"groupId"`
	GroupMembers         pq.StringArray         `gorm:"column:group_members;type:varchar(100)[]" json:"groupMembers"`
	GroupPubKey          pq.StringArray         `gorm:"column:group_pubkey;type:varchar(100)[]" json:"groupPubkey"`
	LogUrl               []LogUrl               `gorm:"foreignkey:dispatchedGroupId;association_foreignkey:group_id"`
	LogRequestUserRandom []LogRequestUserRandom `gorm:"foreignkey:dispatchedGroupId;association_foreignkey:group_id"`
	LogUpdateRandom      []LogUpdateRandom      `gorm:"foreignkey:dispatchedGroupId;association_foreignkey:group_id"`
}

func (GroupInfo) TableName() string {
	return "groupinfo"
}

type RequestInfo struct {
	RequestId       string `gorm:"column:request_id;primary_key" json:"requestId"`
	GroupId         string `gorm:"column:group_id" json:"groupId"`
	Submitter       string `gorm:"column:submitter;type:varchar(100)[]" json:"submitter"`
	SubmittedBlk    uint64 `gorm:"column:blk_num" json:"submittedBlkNum"`
	SubmittedTxHash string `gorm:"column:transaction_hash" json:"submittedTxHash"`
}

type RequestResult struct {
	Message   string         `gorm:"column:message;type:bytea" json:"message"`
	Signature pq.StringArray `gorm:"column:signature;type:varchar(100)[]" json:"signature"`
	PubKey    pq.StringArray `gorm:"column:pub_key;type:varchar(100)[]" json:"pubKey"`
	Pass      bool           `gorm:"column:pass" json:"pass"`
}

type RequestUrlInfo struct {
	RequestInfo
	Timeout    string `gorm:"column:time_out" json:"timeOut"`
	DataSource string `gorm:"column:data_source" json:"dataSource"`
	Selector   string `gorm:"column:selector" json:"selector"`
	Randomness string `gorm:"column:randomness" json:"randomness"`
	RequestResult
}

func (RequestUrlInfo) TableName() string {
	return "requesturlinfo"
}

type RequestURandomInfo struct {
	RequestInfo
	LastSystemRandomness string `gorm:"column:last_system_randomness" json:"lastSystemRandomness"`
	UserSeed             string `gorm:"column:user_seed" json:"userSeed"`
	RequestResult
}

func (RequestURandomInfo) TableName() string {
	return "requesturandominfo"
}

type RequestSRandomInfo struct {
	RequestInfo
	LastSystemRandomness string `gorm:"column:last_system_randomness" json:"lastSystemRandomness"`
	RequestResult
}

func (RequestSRandomInfo) TableName() string {
	return "requestsrandominfo"
}
