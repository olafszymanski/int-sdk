// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: integration.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SportType int32

const (
	SportType_BASKETBALL SportType = 0
)

// Enum value maps for SportType.
var (
	SportType_name = map[int32]string{
		0: "BASKETBALL",
	}
	SportType_value = map[string]int32{
		"BASKETBALL": 0,
	}
)

func (x SportType) Enum() *SportType {
	p := new(SportType)
	*p = x
	return p
}

func (x SportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SportType) Descriptor() protoreflect.EnumDescriptor {
	return file_integration_proto_enumTypes[0].Descriptor()
}

func (SportType) Type() protoreflect.EnumType {
	return &file_integration_proto_enumTypes[0]
}

func (x SportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SportType.Descriptor instead.
func (SportType) EnumDescriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{0}
}

type MarketType int32

const (
	MarketType_MATCH_BETTING                      MarketType = 0
	MarketType_MONEYLINE                          MarketType = 1
	MarketType_MONEYLINE_1ST_QUARTER              MarketType = 2
	MarketType_MONEYLINE_1ST_HALF                 MarketType = 3
	MarketType_RACE_TO_X_POINTS                   MarketType = 4
	MarketType_OUTRIGHT                           MarketType = 5
	MarketType_TOTAL_POINTS                       MarketType = 6
	MarketType_TOTAL_POINTS_ODD_EVEN              MarketType = 7
	MarketType_TOTAL_POINTS_1ST_QUARTER           MarketType = 8
	MarketType_TOTAL_POINTS_1ST_QUARTER_ODD_EVEN  MarketType = 9
	MarketType_TOTAL_POINTS_1ST_HALF              MarketType = 10
	MarketType_TOTAL_POINTS_1ST_HALF_ODD_EVEN     MarketType = 11
	MarketType_TOTAL_POINTS_HOME_TEAM             MarketType = 12
	MarketType_TOTAL_POINTS_HOME_TEAM_1ST_QUARTER MarketType = 13
	MarketType_TOTAL_POINTS_HOME_TEAM_1ST_HALF    MarketType = 14
	MarketType_TOTAL_POINTS_AWAY_TEAM             MarketType = 15
	MarketType_TOTAL_POINTS_AWAY_TEAM_1ST_QUARTER MarketType = 16
	MarketType_TOTAL_POINTS_AWAY_TEAM_1ST_HALF    MarketType = 17
	MarketType_SPREAD_1ST_QUARTER                 MarketType = 18
	MarketType_SPREAD_1ST_HALF                    MarketType = 19
	MarketType_PLAYER_TOTAL_POINTS                MarketType = 20
	MarketType_PLAYER_TOTAL_ASSISTS               MarketType = 21
	MarketType_PLAYER_TOTAL_REBOUNDS              MarketType = 22
	MarketType_PLAYER_TOTAL_3_POINTERS            MarketType = 23
)

// Enum value maps for MarketType.
var (
	MarketType_name = map[int32]string{
		0:  "MATCH_BETTING",
		1:  "MONEYLINE",
		2:  "MONEYLINE_1ST_QUARTER",
		3:  "MONEYLINE_1ST_HALF",
		4:  "RACE_TO_X_POINTS",
		5:  "OUTRIGHT",
		6:  "TOTAL_POINTS",
		7:  "TOTAL_POINTS_ODD_EVEN",
		8:  "TOTAL_POINTS_1ST_QUARTER",
		9:  "TOTAL_POINTS_1ST_QUARTER_ODD_EVEN",
		10: "TOTAL_POINTS_1ST_HALF",
		11: "TOTAL_POINTS_1ST_HALF_ODD_EVEN",
		12: "TOTAL_POINTS_HOME_TEAM",
		13: "TOTAL_POINTS_HOME_TEAM_1ST_QUARTER",
		14: "TOTAL_POINTS_HOME_TEAM_1ST_HALF",
		15: "TOTAL_POINTS_AWAY_TEAM",
		16: "TOTAL_POINTS_AWAY_TEAM_1ST_QUARTER",
		17: "TOTAL_POINTS_AWAY_TEAM_1ST_HALF",
		18: "SPREAD_1ST_QUARTER",
		19: "SPREAD_1ST_HALF",
		20: "PLAYER_TOTAL_POINTS",
		21: "PLAYER_TOTAL_ASSISTS",
		22: "PLAYER_TOTAL_REBOUNDS",
		23: "PLAYER_TOTAL_3_POINTERS",
	}
	MarketType_value = map[string]int32{
		"MATCH_BETTING":                      0,
		"MONEYLINE":                          1,
		"MONEYLINE_1ST_QUARTER":              2,
		"MONEYLINE_1ST_HALF":                 3,
		"RACE_TO_X_POINTS":                   4,
		"OUTRIGHT":                           5,
		"TOTAL_POINTS":                       6,
		"TOTAL_POINTS_ODD_EVEN":              7,
		"TOTAL_POINTS_1ST_QUARTER":           8,
		"TOTAL_POINTS_1ST_QUARTER_ODD_EVEN":  9,
		"TOTAL_POINTS_1ST_HALF":              10,
		"TOTAL_POINTS_1ST_HALF_ODD_EVEN":     11,
		"TOTAL_POINTS_HOME_TEAM":             12,
		"TOTAL_POINTS_HOME_TEAM_1ST_QUARTER": 13,
		"TOTAL_POINTS_HOME_TEAM_1ST_HALF":    14,
		"TOTAL_POINTS_AWAY_TEAM":             15,
		"TOTAL_POINTS_AWAY_TEAM_1ST_QUARTER": 16,
		"TOTAL_POINTS_AWAY_TEAM_1ST_HALF":    17,
		"SPREAD_1ST_QUARTER":                 18,
		"SPREAD_1ST_HALF":                    19,
		"PLAYER_TOTAL_POINTS":                20,
		"PLAYER_TOTAL_ASSISTS":               21,
		"PLAYER_TOTAL_REBOUNDS":              22,
		"PLAYER_TOTAL_3_POINTERS":            23,
	}
)

func (x MarketType) Enum() *MarketType {
	p := new(MarketType)
	*p = x
	return p
}

func (x MarketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MarketType) Descriptor() protoreflect.EnumDescriptor {
	return file_integration_proto_enumTypes[1].Descriptor()
}

func (MarketType) Type() protoreflect.EnumType {
	return &file_integration_proto_enumTypes[1]
}

func (x MarketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MarketType.Descriptor instead.
func (MarketType) EnumDescriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{1}
}

type Outcome_OutcomeType int32

const (
	Outcome_HOME       Outcome_OutcomeType = 0
	Outcome_AWAY       Outcome_OutcomeType = 1
	Outcome_DRAW       Outcome_OutcomeType = 2
	Outcome_COMPETITOR Outcome_OutcomeType = 3
	Outcome_OVER       Outcome_OutcomeType = 4
	Outcome_UNDER      Outcome_OutcomeType = 5
	Outcome_ODD        Outcome_OutcomeType = 6
	Outcome_EVEN       Outcome_OutcomeType = 7
)

// Enum value maps for Outcome_OutcomeType.
var (
	Outcome_OutcomeType_name = map[int32]string{
		0: "HOME",
		1: "AWAY",
		2: "DRAW",
		3: "COMPETITOR",
		4: "OVER",
		5: "UNDER",
		6: "ODD",
		7: "EVEN",
	}
	Outcome_OutcomeType_value = map[string]int32{
		"HOME":       0,
		"AWAY":       1,
		"DRAW":       2,
		"COMPETITOR": 3,
		"OVER":       4,
		"UNDER":      5,
		"ODD":        6,
		"EVEN":       7,
	}
)

func (x Outcome_OutcomeType) Enum() *Outcome_OutcomeType {
	p := new(Outcome_OutcomeType)
	*p = x
	return p
}

func (x Outcome_OutcomeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Outcome_OutcomeType) Descriptor() protoreflect.EnumDescriptor {
	return file_integration_proto_enumTypes[2].Descriptor()
}

func (Outcome_OutcomeType) Type() protoreflect.EnumType {
	return &file_integration_proto_enumTypes[2]
}

func (x Outcome_OutcomeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Outcome_OutcomeType.Descriptor instead.
func (Outcome_OutcomeType) EnumDescriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{3, 0}
}

type Participant_ParticipantType int32

const (
	Participant_HOME       Participant_ParticipantType = 0
	Participant_AWAY       Participant_ParticipantType = 1
	Participant_COMPETITOR Participant_ParticipantType = 2
)

// Enum value maps for Participant_ParticipantType.
var (
	Participant_ParticipantType_name = map[int32]string{
		0: "HOME",
		1: "AWAY",
		2: "COMPETITOR",
	}
	Participant_ParticipantType_value = map[string]int32{
		"HOME":       0,
		"AWAY":       1,
		"COMPETITOR": 2,
	}
)

func (x Participant_ParticipantType) Enum() *Participant_ParticipantType {
	p := new(Participant_ParticipantType)
	*p = x
	return p
}

func (x Participant_ParticipantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Participant_ParticipantType) Descriptor() protoreflect.EnumDescriptor {
	return file_integration_proto_enumTypes[3].Descriptor()
}

func (Participant_ParticipantType) Type() protoreflect.EnumType {
	return &file_integration_proto_enumTypes[3]
}

func (x Participant_ParticipantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Participant_ParticipantType.Descriptor instead.
func (Participant_ParticipantType) EnumDescriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{5, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportType SportType `protobuf:"varint,1,opt,name=sport_type,json=sportType,proto3,enum=SportType" json:"sport_type,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetSportType() SportType {
	if x != nil {
		return x.SportType
	}
	return SportType_BASKETBALL
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Odds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decimal     float64 `protobuf:"fixed64,1,opt,name=decimal,proto3" json:"decimal,omitempty"`
	American    string  `protobuf:"bytes,2,opt,name=american,proto3" json:"american,omitempty"`
	Numerator   string  `protobuf:"bytes,3,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator string  `protobuf:"bytes,4,opt,name=denominator,proto3" json:"denominator,omitempty"`
	IsFixed     bool    `protobuf:"varint,5,opt,name=is_fixed,json=isFixed,proto3" json:"is_fixed,omitempty"`
}

func (x *Odds) Reset() {
	*x = Odds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Odds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Odds) ProtoMessage() {}

func (x *Odds) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Odds.ProtoReflect.Descriptor instead.
func (*Odds) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{2}
}

func (x *Odds) GetDecimal() float64 {
	if x != nil {
		return x.Decimal
	}
	return 0
}

func (x *Odds) GetAmerican() string {
	if x != nil {
		return x.American
	}
	return ""
}

func (x *Odds) GetNumerator() string {
	if x != nil {
		return x.Numerator
	}
	return ""
}

func (x *Odds) GetDenominator() string {
	if x != nil {
		return x.Denominator
	}
	return ""
}

func (x *Odds) GetIsFixed() bool {
	if x != nil {
		return x.IsFixed
	}
	return false
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        Outcome_OutcomeType `protobuf:"varint,1,opt,name=type,proto3,enum=Outcome_OutcomeType" json:"type,omitempty"`
	ExternalId  string              `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Name        *string             `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`       // optionally used in case type is COMPETITOR
	Points      *float64            `protobuf:"fixed64,4,opt,name=points,proto3,oneof" json:"points,omitempty"` // optionally used for points markets like total points or spread
	Odds        *Odds               `protobuf:"bytes,5,opt,name=odds,proto3" json:"odds,omitempty"`
	IsAvailable bool                `protobuf:"varint,6,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{3}
}

func (x *Outcome) GetType() Outcome_OutcomeType {
	if x != nil {
		return x.Type
	}
	return Outcome_HOME
}

func (x *Outcome) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Outcome) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Outcome) GetPoints() float64 {
	if x != nil && x.Points != nil {
		return *x.Points
	}
	return 0
}

func (x *Outcome) GetOdds() *Odds {
	if x != nil {
		return x.Odds
	}
	return nil
}

func (x *Outcome) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type Market struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       MarketType `protobuf:"varint,1,opt,name=type,proto3,enum=MarketType" json:"type,omitempty"`
	ExternalId string     `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Name       *string    `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"` // optionally used in case of a player market
	Outcomes   []*Outcome `protobuf:"bytes,4,rep,name=outcomes,proto3" json:"outcomes,omitempty"`
}

func (x *Market) Reset() {
	*x = Market{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Market) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Market) ProtoMessage() {}

func (x *Market) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Market.ProtoReflect.Descriptor instead.
func (*Market) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{4}
}

func (x *Market) GetType() MarketType {
	if x != nil {
		return x.Type
	}
	return MarketType_MATCH_BETTING
}

func (x *Market) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Market) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Market) GetOutcomes() []*Outcome {
	if x != nil {
		return x.Outcomes
	}
	return nil
}

type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Participant_ParticipantType `protobuf:"varint,1,opt,name=type,proto3,enum=Participant_ParticipantType" json:"type,omitempty"`
	Name string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Participant) Reset() {
	*x = Participant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant.ProtoReflect.Descriptor instead.
func (*Participant) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{5}
}

func (x *Participant) GetType() Participant_ParticipantType {
	if x != nil {
		return x.Type
	}
	return Participant_HOME
}

func (x *Participant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExternalId   string                 `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	SportType    SportType              `protobuf:"varint,3,opt,name=sport_type,json=sportType,proto3,enum=SportType" json:"sport_type,omitempty"`
	Name         string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	League       string                 `protobuf:"bytes,5,opt,name=league,proto3" json:"league,omitempty"`
	StartTime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	IsLive       bool                   `protobuf:"varint,7,opt,name=is_live,json=isLive,proto3" json:"is_live,omitempty"`
	Participants []*Participant         `protobuf:"bytes,8,rep,name=participants,proto3" json:"participants,omitempty"`
	Markets      []*Market              `protobuf:"bytes,9,rep,name=markets,proto3" json:"markets,omitempty"`
	Link         string                 `protobuf:"bytes,10,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Event) GetSportType() SportType {
	if x != nil {
		return x.SportType
	}
	return SportType_BASKETBALL
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetLeague() string {
	if x != nil {
		return x.League
	}
	return ""
}

func (x *Event) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Event) GetIsLive() bool {
	if x != nil {
		return x.IsLive
	}
	return false
}

func (x *Event) GetParticipants() []*Participant {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *Event) GetMarkets() []*Market {
	if x != nil {
		return x.Markets
	}
	return nil
}

func (x *Event) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_integration_proto protoreflect.FileDescriptor

var file_integration_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2a, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x04, 0x4f, 0x64, 0x64, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x61, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x22, 0xc1, 0x02, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x04, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4f, 0x64,
	0x64, 0x73, 0x52, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x63, 0x0a, 0x0b, 0x4f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f,
	0x4d, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x57, 0x41, 0x59, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x52, 0x41, 0x57, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4d, 0x50,
	0x45, 0x54, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x56, 0x45, 0x52,
	0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x05, 0x12, 0x07, 0x0a,
	0x03, 0x4f, 0x44, 0x44, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x56, 0x45, 0x4e, 0x10, 0x07,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12,
	0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x08, 0x6f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x73,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x0b, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x35, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x41, 0x57, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4d, 0x50, 0x45, 0x54,
	0x49, 0x54, 0x4f, 0x52, 0x10, 0x02, 0x22, 0xcc, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x29, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x30, 0x0a, 0x0c,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x07, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x2a, 0x1b, 0x0a, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x41, 0x53, 0x4b, 0x45, 0x54, 0x42, 0x41, 0x4c, 0x4c,
	0x10, 0x00, 0x2a, 0xa0, 0x05, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x42, 0x45, 0x54, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x4c, 0x49, 0x4e,
	0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x4c, 0x49, 0x4e, 0x45,
	0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x31, 0x53, 0x54, 0x5f,
	0x48, 0x41, 0x4c, 0x46, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x54,
	0x4f, 0x5f, 0x58, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x4f, 0x55, 0x54, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x4f,
	0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15,
	0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x4f, 0x44, 0x44,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x4f, 0x54, 0x41, 0x4c,
	0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52,
	0x54, 0x45, 0x52, 0x10, 0x08, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45,
	0x52, 0x5f, 0x4f, 0x44, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x10, 0x09, 0x12, 0x19, 0x0a, 0x15,
	0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x31, 0x53, 0x54,
	0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x0a, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x4f, 0x54, 0x41, 0x4c,
	0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46,
	0x5f, 0x4f, 0x44, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x10, 0x0b, 0x12, 0x1a, 0x0a, 0x16, 0x54,
	0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x48, 0x4f, 0x4d, 0x45,
	0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x0c, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x4f, 0x54, 0x41, 0x4c,
	0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x54, 0x45, 0x41,
	0x4d, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x10, 0x0d, 0x12,
	0x23, 0x0a, 0x1f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f,
	0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x48, 0x41,
	0x4c, 0x46, 0x10, 0x0e, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x0f,
	0x12, 0x26, 0x0a, 0x22, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53,
	0x5f, 0x41, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51,
	0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x10, 0x10, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x4f, 0x54, 0x41,
	0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x45,
	0x41, 0x4d, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x11, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x50, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52,
	0x54, 0x45, 0x52, 0x10, 0x12, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x50, 0x52, 0x45, 0x41, 0x44, 0x5f,
	0x31, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x13, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4c,
	0x41, 0x59, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54,
	0x53, 0x10, 0x14, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x54, 0x4f,
	0x54, 0x41, 0x4c, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54, 0x53, 0x10, 0x15, 0x12, 0x19, 0x0a,
	0x15, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x52, 0x45,
	0x42, 0x4f, 0x55, 0x4e, 0x44, 0x53, 0x10, 0x16, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4c, 0x41, 0x59,
	0x45, 0x52, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x33, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x53, 0x10, 0x17, 0x32, 0x55, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integration_proto_rawDescOnce sync.Once
	file_integration_proto_rawDescData = file_integration_proto_rawDesc
)

func file_integration_proto_rawDescGZIP() []byte {
	file_integration_proto_rawDescOnce.Do(func() {
		file_integration_proto_rawDescData = protoimpl.X.CompressGZIP(file_integration_proto_rawDescData)
	})
	return file_integration_proto_rawDescData
}

var file_integration_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_integration_proto_goTypes = []interface{}{
	(SportType)(0),                   // 0: SportType
	(MarketType)(0),                  // 1: MarketType
	(Outcome_OutcomeType)(0),         // 2: Outcome.OutcomeType
	(Participant_ParticipantType)(0), // 3: Participant.ParticipantType
	(*Request)(nil),                  // 4: Request
	(*Response)(nil),                 // 5: Response
	(*Odds)(nil),                     // 6: Odds
	(*Outcome)(nil),                  // 7: Outcome
	(*Market)(nil),                   // 8: Market
	(*Participant)(nil),              // 9: Participant
	(*Event)(nil),                    // 10: Event
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
}
var file_integration_proto_depIdxs = []int32{
	0,  // 0: Request.sport_type:type_name -> SportType
	10, // 1: Response.events:type_name -> Event
	2,  // 2: Outcome.type:type_name -> Outcome.OutcomeType
	6,  // 3: Outcome.odds:type_name -> Odds
	1,  // 4: Market.type:type_name -> MarketType
	7,  // 5: Market.outcomes:type_name -> Outcome
	3,  // 6: Participant.type:type_name -> Participant.ParticipantType
	0,  // 7: Event.sport_type:type_name -> SportType
	11, // 8: Event.start_time:type_name -> google.protobuf.Timestamp
	9,  // 9: Event.participants:type_name -> Participant
	8,  // 10: Event.markets:type_name -> Market
	4,  // 11: Integration.GetPreMatch:input_type -> Request
	4,  // 12: Integration.GetLive:input_type -> Request
	5,  // 13: Integration.GetPreMatch:output_type -> Response
	5,  // 14: Integration.GetLive:output_type -> Response
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_integration_proto_init() }
func file_integration_proto_init() {
	if File_integration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_integration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_integration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Odds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_integration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_integration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Market); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_integration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participant); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_integration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_integration_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_integration_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_integration_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integration_proto_goTypes,
		DependencyIndexes: file_integration_proto_depIdxs,
		EnumInfos:         file_integration_proto_enumTypes,
		MessageInfos:      file_integration_proto_msgTypes,
	}.Build()
	File_integration_proto = out.File
	file_integration_proto_rawDesc = nil
	file_integration_proto_goTypes = nil
	file_integration_proto_depIdxs = nil
}
