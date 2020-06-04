package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

const (
	_ int = iota
	TypeNode
	TypeGroup
	TypeUrlRequest
	TypeRandomRequest
	TypeLatestEvents
	TypeUnregisterPendingNode
	TypeNewPendingNode
	TypeGrouping
	TypePublicKeySuggested
	TypePublicKeyAccepted
	TypeGroupDissolve
	TypeUpdateRandom
	TypeUrl
	TypeRequestUserRandom
	TypeValidationResult

	TypeGuardianReward
	TypeCallbackTriggeredFor
	TypeMessage
)

var supportedEvents []string
var stringToType map[string]int
var typeToStruct []interface{}

func init() {
	supportedEvents = append(supportedEvents, "LogRegisteredNewPendingNode")
	supportedEvents = append(supportedEvents, "LogUnRegisteredNewPendingNode")
	supportedEvents = append(supportedEvents, "LogGrouping")
	supportedEvents = append(supportedEvents, "LogPublicKeySuggested")
	supportedEvents = append(supportedEvents, "LogPublicKeyAccepted")
	supportedEvents = append(supportedEvents, "LogGroupDissolve")
	supportedEvents = append(supportedEvents, "LogUpdateRandom")
	supportedEvents = append(supportedEvents, "LogUrl")
	supportedEvents = append(supportedEvents, "LogRequestUserRandom")
	supportedEvents = append(supportedEvents, "LogValidationResult")
	supportedEvents = append(supportedEvents, "LogCallbackTriggeredFor")
	supportedEvents = append(supportedEvents, "GuardianReward")
	supportedEvents = append(supportedEvents, "LogMessage")

	stringToType = make(map[string]int)
	stringToType["node"] = TypeNode
	stringToType["group"] = TypeGroup
	stringToType["urlrequest"] = TypeUrlRequest
	stringToType["randomreques"] = TypeRandomRequest
	stringToType["logregisterednewpendingnode"] = TypeNewPendingNode
	stringToType["logunregisterednewpendingnode"] = TypeUnregisterPendingNode
	stringToType["loggrouping"] = TypeGrouping
	stringToType["logpublickeysuggested"] = TypePublicKeySuggested
	stringToType["logpublickeyaccepted"] = TypePublicKeyAccepted
	stringToType["loggroupdissolve"] = TypeGroupDissolve
	stringToType["logupdaterandom"] = TypeUpdateRandom
	stringToType["logurl"] = TypeUrl
	stringToType["logrequestuserrandom"] = TypeRequestUserRandom
	stringToType["logvalidationresult"] = TypeValidationResult
	stringToType["logcallbacktriggeredfor"] = TypeCallbackTriggeredFor
	stringToType["guardianreward"] = TypeGuardianReward
	stringToType["logmessage"] = TypeMessage

	typeToStruct = []interface{}{
		TypeNewPendingNode:        &LogRegisteredNewPendingNode{},
		TypeUnregisterPendingNode: &LogUnRegisteredNewPendingNode{},
		TypeGrouping:              &LogGrouping{},
		TypePublicKeySuggested:    &LogPublicKeySuggested{},
		TypePublicKeyAccepted:     &LogPublicKeyAccepted{},
		TypeGroupDissolve:         &LogGroupDissolve{},
		TypeUpdateRandom:          &LogUpdateRandom{},
		TypeRequestUserRandom:     &LogRequestUserRandom{},
		TypeUrl:                   &LogUrl{},
		TypeValidationResult:      &LogValidationResult{},
		TypeCallbackTriggeredFor:  &LogCallbackTriggeredFor{},
		TypeGuardianReward:        &GuardianReward{},
		TypeMessage:               &LogMessage{},
		TypeNode:                  &Node{},
		TypeGroup:                 &Group{},
		TypeUrlRequest:            &UrlRequest{},
		TypeRandomRequest:         &UserRandomRequest{},
	}
}

func SupportedEvents() []string {
	return supportedEvents
}

func StringToType(s string) int {
	return stringToType[strings.ToLower(s)]
}

func TypeToStruct(t int) interface{} {
	if t < len(typeToStruct) {
		return typeToStruct[t]
	}
	return nil
}

type Transaction struct {
	gorm.Model
	Hash                           string                          `gorm:"primary_key;unique;not null"`
	GasPrice                       uint64                          `json:"gasPrice"`
	Value                          uint64                          `json:"value"`
	GasLimit                       uint64                          `json:"gasLimit"`
	Nonce                          uint64                          `json:"nonce"`
	Sender                         string                          `json:"sender"`
	To                             string                          `json:"to"`
	BlockNumber                    uint64                          `gorm:"primary_key" json:"blockNumber"`
	Data                           []byte                          `gorm:"type:bytea" json:"data"`
	Method                         string                          `gorm:"index" json:"method"`
	LogRegisteredNewPendingNodes   []LogRegisteredNewPendingNode   `gorm:"auto_preload"`
	LogUnRegisteredNewPendingNodes []LogUnRegisteredNewPendingNode `gorm:"auto_preload"`
	LogGroupings                   []LogGrouping                   `gorm:"auto_preload"`
	LogPublicKeySuggesteds         []LogPublicKeySuggested         `gorm:"auto_preload"`
	LogPublicKeyAccepteds          []LogPublicKeyAccepted          `gorm:"auto_preload"`
	LogUpdateRandoms               []LogUpdateRandom               `gorm:"auto_preload"`
	LogUrls                        []LogUrl                        `gorm:"auto_preload"`
	LogRequestUserRandoms          []LogRequestUserRandom          `gorm:"auto_preload"`
	LogValidationResults           []LogValidationResult           `gorm:"auto_preload"`
	LogCallbackTriggeredFors       []LogCallbackTriggeredFor       `gorm:"auto_preload"`
	GuardianRewards                []GuardianReward                `gorm:"auto_preload"`
	LogMessages                    []LogMessage                    `gorm:"auto_preload"`

	LogNonSupportedTypes           []LogNonSupportedType           `gorm:"auto_preload"`
	LogNonContractCalls            []LogNonContractCall            `gorm:"auto_preload"`
	LogRequestFromNonExistentUCs   []LogRequestFromNonExistentUC   `gorm:"auto_preload"`
	LogInsufficientPendingNodes    []LogInsufficientPendingNode    `gorm:"auto_preload"`
	LogInsufficientWorkingGroups   []LogInsufficientWorkingGroup   `gorm:"auto_preload"`
	LogGroupDissolves              []LogGroupDissolve              `gorm:"auto_preload"`
	LogGroupingInitiateds          []LogGroupingInitiated          `gorm:"auto_preload"`
	LogNoPendingGroups             []LogNoPendingGroup             `gorm:"auto_preload"`
	LogPendingGroupRemoveds        []LogPendingGroupRemoved        `gorm:"auto_preload"`
	UpdateGroupToPicks             []UpdateGroupToPick             `gorm:"auto_preload"`
	UpdateGroupSizes               []UpdateGroupSize               `gorm:"auto_preload"`
	UpdateGroupingThresholds       []UpdateGroupingThreshold       `gorm:"auto_preload"`
	UpdateGroupMaturityPeriods     []UpdateGroupMaturityPeriod     `gorm:"auto_preload"`
	UpdateBootstrapCommitDurations []UpdateBootstrapCommitDuration `gorm:"auto_preload"`
	UpdateBootstrapRevealDurations []UpdateBootstrapRevealDuration `gorm:"auto_preload"`
	UpdatebootstrapStartThresholds []UpdatebootstrapStartThreshold `gorm:"auto_preload"`
	UpdatePendingGroupMaxLifes     []UpdatePendingGroupMaxLife     `gorm:"auto_preload"`
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

type LogUnRegisteredNewPendingNode struct {
	gorm.Model
	Event
	Node           string `json:"node"`
	UnregisterFrom uint8  `json:"unregisterFrom"`
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
	InvalidSelector string `gorm:"column:invalid_selector" json:"invalidSelector"`
}

type LogNonContractCall struct {
	gorm.Model
	Event
	CallAddr string `json:"callAddr"`
}

type LogCallbackTriggeredFor struct {
	gorm.Model
	Event
	CallbackAddr string `json:"callbackAddr"`
}

type LogRequestFromNonExistentUC struct {
	gorm.Model
	Event
}

type LogInsufficientPendingNode struct {
	Event
	NumPendingNodes uint64 `gjson:"numPendingNodes"`
}

type LogInsufficientWorkingGroup struct {
	gorm.Model
	Event
	NumWorkingGroups uint64 `json:"numWorkingGroups"`
	NumPendingGroups uint64 `json:"numPendingGroups"`
}

type LogGroupingInitiated struct {
	gorm.Model
	Event
	PendingNodePool uint64 `json:"pendingNodePool"`
	GroupSize       uint64 `json:"groupSize"`
}

type LogNoPendingGroup struct {
	gorm.Model
	Event
	GroupId string `json:"groupId"`
}

type LogPendingGroupRemoved struct {
	gorm.Model
	Event
	GroupId string `json:"groupId"`
}

type LogMessage struct {
	gorm.Model
	Event
	Info string `json:"info"`
}

type UpdateGroupToPick struct {
	gorm.Model
	Event
	OldNum uint64 `json:"oldNum"`
	NewNum uint64 `json:"newNum"`
}

type UpdateGroupSize struct {
	gorm.Model
	Event
	OldSize uint64 `json:"oldSize"`
	NewSize uint64 `json:"newSize"`
}

type UpdateGroupingThreshold struct {
	gorm.Model
	Event
	OldThreshold uint64 `json:"oldThreshold"`
	NewThreshold uint64 `json:"newThreshold"`
}

type UpdateGroupMaturityPeriod struct {
	gorm.Model
	Event
	OldPeriod uint64 `json:"oldPeriod"`
	NewPeriod uint64 `json:"newPeriod"`
}

type UpdateBootstrapCommitDuration struct {
	gorm.Model
	Event
	OldDuration uint64 `json:"oldDuration"`
	NewDuration uint64 `json:"newDuration"`
}

type UpdateBootstrapRevealDuration struct {
	gorm.Model
	Event
	OldDuration uint64 `json:"oldDuration"`
	NewDuration uint64 `json:"newDuration"`
}

type UpdatebootstrapStartThreshold struct {
	gorm.Model
	Event
	OldThreshold uint64 `json:"oldThreshold"`
	NewThreshold uint64 `json:"newThreshold"`
}

type UpdatePendingGroupMaxLife struct {
	gorm.Model
	Event
	OldLifeBlocks uint64 `json:"oldLifeBlocks"`
	NewLifeBlocks uint64 `json:"newLifeBlock"`
}

type GuardianReward struct {
	gorm.Model
	Event
	BlkNum   uint64 `json:"blkNum"`
	Guardian string `json:"guardian"`
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
	GroupId     string `json:"groupId"`
	PubKeyCount uint64 `json:"pubKeyCount"`
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
	gorm.Model
	Event
	RequestId            string `gorm:"unique;not null" json:"requestId"`
	LastSystemRandomness string `json:"lastSystemRandomness"`
	UserSeed             string `json:"userSeed"`
	DispatchedGroupId    string `json:"dispatchedGroupId"`
}

type LogUpdateRandom struct {
	gorm.Model
	Event
	LastRandomness    string `json:"lastRandomness"`
	DispatchedGroupId string `json:"dispatchedGroupId"`
}

type LogUrl struct {
	gorm.Model
	Event
	RequestId         string `gorm:"unique;not null" json:"queryId"`
	Timeout           string `json:"timeOut"`
	DataSource        string `json:"dataSource"`
	Selector          string `json:"selector"`
	Randomness        string `json:"randomness"`
	DispatchedGroupId string `json:"dispatchedGroupId"`
}

type LogValidationResult struct {
	gorm.Model
	Event

	LogUrlID               uint           `json:"-"`
	LogRequestUserRandomID uint           `json:"-"`
	LogUpdateRandomID      uint           `json:"-"`
	RequestId              string         `gorm:"not null" json:"trafficId"`
	RequestType            uint8          `json:"trafficType"`
	Message                []byte         `gorm:"type:bytea" json:"-"`
	MessageStr             string         `json:"message"`
	Signature              pq.StringArray `gorm:"type:varchar(100)[]" json:"signature"`
	PubKey                 pq.StringArray `gorm:"type:varchar(100)[]" json:"pubKey"`
	Pass                   bool           `json:"pass"`
}

type Request struct {
	RequestId         string `gorm:"unique" json:"requestId"`
	DispatchedGroupId string `json:"dispatchedGroupId"`
	Sender            string `json:"submitter"`
	BlockNumber       uint64 `json:"submittedBlkNum"`
	Hash              string `json:"submittedTxHash"`
}

type RequestResult struct {
	Message    []byte         `gorm:"type:bytea" json:"-"`
	MessageStr string         `json:"message"`
	Signature  pq.StringArray `gorm:"type:varchar(100)[]" json:"signature"`
	PubKey     pq.StringArray `gorm:"type:varchar(100)[]" json:"pubKey"`
	Pass       bool           `json:"pass"`
}

type UrlRequest struct {
	gorm.Model    `json:"-"`
	Request       `gorm:"embedded"`
	RequestResult `gorm:"embedded"`
	GroupID       uint   `json:"-"`
	Timeout       string `json:"timeOut"`
	DataSource    string `json:"dataSource"`
	Selector      string `json:"selector"`
	Randomness    string `json:"randomness"`
}

type UserRandomRequest struct {
	gorm.Model           `json:"-"`
	Request              `gorm:"embedded"`
	RequestResult        `gorm:"embedded"`
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
