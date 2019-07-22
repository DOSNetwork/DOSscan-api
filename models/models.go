package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

const (
	TypeNewPendingNode int = iota
	TypeGrouping
	TypePublicKeySuggested
	TypePublicKeyAccepted
	TypeGroupDissolve

	TypeUpdateRandom
	TypeUrl
	TypeRequestUserRandom
	TypeValidationResult
	TypeCallbackTriggeredFor
	TypeGuardianReward

	TypeError
)

type Transaction struct {
	gorm.Model
	Hash                           string `gorm:"primary_key;unique;not null"`
	GasPrice                       uint64 `json:"gasPrice"`
	Value                          uint64 `json:"value"`
	GasLimit                       uint64 `json:"gasLimit"`
	Nonce                          uint64 `json:"nonce"`
	Sender                         string `json:"sender"`
	To                             string `json:"to"`
	BlockNumber                    uint64 `gorm:"primary_key" json:"blockNumber"`
	Data                           []byte `gorm:"type:bytea" json:"data"`
	Method                         string `gorm:"index" json:"method"`
	LogRegisteredNewPendingNodes   []LogRegisteredNewPendingNode
	LogGroupings                   []LogGrouping
	LogUrls                        []LogUrl
	LogRequestUserRandoms          []LogRequestUserRandom
	LogUpdateRandoms               []LogUpdateRandom
	LogValidationResults           []LogValidationResult
	LogNonSupportedTypes           []LogNonSupportedType
	LogNonContractCalls            []LogNonContractCall
	LogCallbackTriggeredFors       []LogCallbackTriggeredFor
	LogRequestFromNonExistentUCs   []LogRequestFromNonExistentUC
	LogInsufficientPendingNodes    []LogInsufficientPendingNode
	LogInsufficientWorkingGroups   []LogInsufficientWorkingGroup
	LogPublicKeyAccepteds          []LogPublicKeyAccepted
	LogPublicKeySuggesteds         []LogPublicKeySuggested
	LogGroupDissolves              []LogGroupDissolve
	LogGroupingInitiateds          []LogGroupingInitiated
	LogNoPendingGroups             []LogNoPendingGroup
	LogPendingGroupRemoveds        []LogPendingGroupRemoved
	LogErrors                      []LogError
	UpdateGroupToPicks             []UpdateGroupToPick
	UpdateGroupSizes               []UpdateGroupSize
	UpdateGroupingThresholds       []UpdateGroupingThreshold
	UpdateGroupMaturityPeriods     []UpdateGroupMaturityPeriod
	UpdateBootstrapCommitDurations []UpdateBootstrapCommitDuration
	UpdateBootstrapRevealDurations []UpdateBootstrapRevealDuration
	UpdatebootstrapStartThresholds []UpdatebootstrapStartThreshold
	UpdatePendingGroupMaxLifes     []UpdatePendingGroupMaxLife
	GuardianRewards                []GuardianReward
}

type Event struct {
	Method          string         `json:"method"`
	EventLog        string         `json:"eventLog"`
	TransactionHash string         `json:"txHash"`
	TxIndex         uint           `json:"-"`
	Topics          pq.StringArray `gorm:"column:topics;type:varchar(100)[]" json:"-"`
	BlockNumber     uint64         `json:"blockNumber"`
	BlockHash       string         `json:"-"`
	LogIndex        uint           `json:"-"`
	Removed         bool           `json:"-"`
	Date            time.Time      `json:"-"`
	TransactionID   uint
}

type LogRegisteredNewPendingNode struct {
	gorm.Model
	Event
	Node string `json:"node"`
}

type LogGrouping struct {
	gorm.Model
	Event
	GroupId string         `json:"groupId"`
	NodeId  pq.StringArray `gorm:"type:varchar(100)[]" json:"nodeId"`
}

type LogNonSupportedType struct {
	gorm.Model
	Event
	TransactionID   uint
	InvalidSelector string `gorm:"column:invalid_selector" json:"invalidSelector"`
}

type LogNonContractCall struct {
	gorm.Model
	Event
	TransactionID uint
	CallAddr      string `gorm:"column:call_ddr" json:"callAddr"`
}

type LogCallbackTriggeredFor struct {
	gorm.Model
	Event
	TransactionID uint
	CallbackAddr  string `gorm:"column:call_back_addr" json:"callbackAddr"`
}

type LogRequestFromNonExistentUC struct {
	gorm.Model
	Event
	TransactionID uint
}

type LogInsufficientPendingNode struct {
	Event
	NumPendingNodes uint64 `gjson:"numPendingNodes"`
}

type LogInsufficientWorkingGroup struct {
	gorm.Model
	Event
	TransactionID    uint
	NumWorkingGroups uint64 `json:"numWorkingGroups"`
	NumPendingGroups uint64 `json:"numPendingGroups"`
}

type LogGroupingInitiated struct {
	gorm.Model
	Event
	TransactionID     uint
	PendingNodePool   uint64 `json:"pendingNodePool"`
	GroupSize         uint64 `json:"groupSize"`
	GroupingThreshold uint64 `json:"groupingThreshold"`
}

type LogNoPendingGroup struct {
	gorm.Model
	Event
	TransactionID uint
	GroupId       string `json:"groupId"`
}

type LogPendingGroupRemoved struct {
	gorm.Model
	Event
	TransactionID uint
	GroupId       string `json:"groupId"`
}

type LogError struct {
	gorm.Model
	Event
	TransactionID uint
	Err           string `json:"err"`
}

type UpdateGroupToPick struct {
	gorm.Model
	Event
	TransactionID uint
	OldNum        uint64 `json:"oldNum"`
	NewNum        uint64 `json:"newNum"`
}

type UpdateGroupSize struct {
	gorm.Model
	Event
	TransactionID uint
	OldSize       uint64 `json:"oldSize"`
	NewSize       uint64 `json:"newSize"`
}

type UpdateGroupingThreshold struct {
	gorm.Model
	Event
	TransactionID uint
	OldThreshold  uint64 `json:"oldThreshold"`
	NewThreshold  uint64 `json:"newThreshold"`
}

type UpdateGroupMaturityPeriod struct {
	gorm.Model
	Event
	TransactionID uint
	OldPeriod     uint64 `json:"oldPeriod"`
	NewPeriod     uint64 `json:"newPeriod"`
}

type UpdateBootstrapCommitDuration struct {
	gorm.Model
	Event
	TransactionID uint
	OldDuration   uint64 `json:"oldDuration"`
	NewDuration   uint64 `json:"newDuration"`
}

type UpdateBootstrapRevealDuration struct {
	gorm.Model
	Event
	TransactionID uint
	OldDuration   uint64 `json:"oldDuration"`
	NewDuration   uint64 `json:"newDuration"`
}

type UpdatebootstrapStartThreshold struct {
	gorm.Model
	Event
	TransactionID uint
	OldThreshold  uint64 `json:"oldThreshold"`
	NewThreshold  uint64 `json:"newThreshold"`
}

type UpdatePendingGroupMaxLife struct {
	gorm.Model
	Event
	TransactionID uint
	OldLifeBlocks uint64 `json:"oldLifeBlocks"`
	NewLifeBlocks uint64 `json:"newLifeBlock"`
}

type GuardianReward struct {
	gorm.Model
	Event
	TransactionID uint
	BlkNum        uint64 `json:"blkNum"`
	Guardian      string `json:"guardian"`
}

type LogPublicKeyAccepted struct {
	gorm.Model
	Event
	GroupId          string         `json:"groupId"`
	AcceptedBlkNum   uint64         `json:"acceptedBlknum"`
	PubKey           pq.StringArray `gorm:"type:varchar(100)[]" json:"pubKey"`
	NumWorkingGroups uint64         `json:"numWorkingGroups"`
}

type LogPublicKeySuggested struct {
	gorm.Model
	Event
	TransactionID uint
	GroupId       string `json:"groupId"`
	PubKeyCount   uint64 `json:"pubKeyCount"`
}

type LogGroupDissolve struct {
	gorm.Model
	Event
	GroupId         string `json:"groupId"`
	DissolvedBlkNum uint64 `json:"dissolvedBlknum"`
}

type Node struct {
	gorm.Model    `json:"-"`
	Addr          string   `gorm:"unique;not null" json:"addr"`
	Balance       string   `json:"balance"`
	RegisterState bool     `json:"registerState"`
	ActiveGroups  []string `gorm:"-" json:"activeGroups"`
	ExpiredGroups int      `gorm:"-" json:"expiredGroups"`
	Groups        []Group  `gorm:"many2many:nodes_groups;" json:"-"`
}

type Group struct {
	gorm.Model         `json:"-"`
	GroupId            string              `gorm:"unique;not null" json:"groupId"`
	AcceptedBlkNum     uint64              `json:"acceptedBlknum"`
	DissolvedBlkNum    uint64              `json:"dissolvedBlknum"`
	NodeId             pq.StringArray      `gorm:"type:varchar(100)[]" json:"nodeId"`
	PubKey             pq.StringArray      `gorm:"type:varchar(100)[]" json:"pubKey"`
	Nodes              []Node              `gorm:"many2many:nodes_groups;" json:"-"`
	NumUrl             int                 `gorm:"-" json:"urlRequests"`
	NumRandom          int                 `gorm:"-" json:"randomRequests"`
	UrlRequests        []UrlRequest        `json:"-"`
	UserRandomRequests []UserRandomRequest `json:"-"`
}

type LogRequestUserRandom struct {
	gorm.Model           `json:"-"`
	Event                `json:"-"`
	TransactionID        uint   `json:"-"`
	RequestId            string `gorm:"unique;not null" json:"requestId"`
	LastSystemRandomness string `json:"lastSystemRandomness"`
	UserSeed             string `json:"userSeed"`
	DispatchedGroupId    string `json:"dispatchedGroupId"`
}

type LogUpdateRandom struct {
	gorm.Model        `json:"-"`
	Event             `json:"-"`
	TransactionID     uint   `json:"-"`
	LastRandomness    string `json:"lastRandomness"`
	DispatchedGroupId string `json:"dispatchedGroupId"`
}

type LogUrl struct {
	gorm.Model        `json:"-"`
	Event             `json:"-"`
	TransactionID     uint   `json:"-"`
	RequestId         string `gorm:"unique;not null" json:"queryId"`
	Timeout           string `json:"timeOut"`
	DataSource        string `json:"dataSource"`
	Selector          string `json:"selector"`
	Randomness        string `json:"randomness"`
	DispatchedGroupId string `json:"dispatchedGroupId"`
}

type LogValidationResult struct {
	gorm.Model             `json:"-"`
	Event                  `json:"-"`
	TransactionID          uint           `json:"-"`
	LogUrlID               uint           `json:"-"`
	LogRequestUserRandomID uint           `json:"-"`
	LogUpdateRandomID      uint           `json:"-"`
	RequestId              string         `json:"requestId"`
	RequestType            uint8          `json:"requestType"`
	Message                string         `gorm:"type:bytea" json:"message"`
	Signature              pq.StringArray `gorm:"type:varchar(100)[]" json:"signature"`
	PubKey                 pq.StringArray `gorm:"type:varchar(100)[]" json:"pubKey"`
	Pass                   bool           `json:"pass"`
}

type Request struct {
	RequestId         string `json:"requestId"`
	DispatchedGroupId string `json:"dispatchedGroupId"`
	Sender            string `json:"submitter"`
	BlockNumber       uint64 `json:"submittedBlkNum"`
	Hash              string `json:"submittedTxHash"`
}

type RequestResult struct {
	Message   string         `gorm:"type:bytea" json:"message"`
	Signature pq.StringArray `gorm:"type:varchar(100)[]" json:"signature"`
	PubKey    pq.StringArray `gorm:"type:varchar(100)[]" json:"pubKey"`
	Pass      bool           `json:"pass"`
}

type UrlRequest struct {
	gorm.Model `json:"-"`
	Request
	RequestResult
	GroupID    uint   `json:"-"`
	Timeout    string `json:"timeOut"`
	DataSource string `json:"dataSource"`
	Selector   string `json:"selector"`
	Randomness string `json:"randomness"`
}

type UserRandomRequest struct {
	gorm.Model `json:"-"`
	Request
	RequestResult
	GroupID              uint   `json:"-"`
	LastSystemRandomness string `json:"lastSystemRandomness"`
	UserSeed             string `json:"userSeed"`
}

type SysRandomRequest struct {
	gorm.Model
	Request
	RequestResult
	GroupID uint
}
