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

type ParticipantType int32

const (
	ParticipantType_HOME       ParticipantType = 0
	ParticipantType_AWAY       ParticipantType = 1
	ParticipantType_COMPETITOR ParticipantType = 2
)

// Enum value maps for ParticipantType.
var (
	ParticipantType_name = map[int32]string{
		0: "HOME",
		1: "AWAY",
		2: "COMPETITOR",
	}
	ParticipantType_value = map[string]int32{
		"HOME":       0,
		"AWAY":       1,
		"COMPETITOR": 2,
	}
)

func (x ParticipantType) Enum() *ParticipantType {
	p := new(ParticipantType)
	*p = x
	return p
}

func (x ParticipantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParticipantType) Descriptor() protoreflect.EnumDescriptor {
	return file_integration_proto_enumTypes[1].Descriptor()
}

func (ParticipantType) Type() protoreflect.EnumType {
	return &file_integration_proto_enumTypes[1]
}

func (x ParticipantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParticipantType.Descriptor instead.
func (ParticipantType) EnumDescriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{1}
}

type MarketType int32

const (
	MarketType_MONEYLINE                          MarketType = 0
	MarketType_MONEYLINE_1ST_QUARTER              MarketType = 1
	MarketType_MONEYLINE_1ST_HALF                 MarketType = 2
	MarketType_OUTRIGHT                           MarketType = 3
	MarketType_TOTAL_POINTS                       MarketType = 4
	MarketType_TOTAL_POINTS_ODD_EVEN              MarketType = 5
	MarketType_TOTAL_POINTS_1ST_QUARTER           MarketType = 6
	MarketType_TOTAL_POINTS_1ST_QUARTER_ODD_EVEN  MarketType = 7
	MarketType_TOTAL_POINTS_1ST_HALF              MarketType = 8
	MarketType_TOTAL_POINTS_1ST_HALF_ODD_EVEN     MarketType = 9
	MarketType_TOTAL_POINTS_HOME_TEAM             MarketType = 10
	MarketType_TOTAL_POINTS_HOME_TEAM_1ST_QUARTER MarketType = 11
	MarketType_TOTAL_POINTS_HOME_TEAM_1ST_HALF    MarketType = 12
	MarketType_TOTAL_POINTS_AWAY_TEAM             MarketType = 13
	MarketType_TOTAL_POINTS_AWAY_TEAM_1ST_QUARTER MarketType = 14
	MarketType_TOTAL_POINTS_AWAY_TEAM_1ST_HALF    MarketType = 15
	MarketType_SPREAD_1ST_QUARTER                 MarketType = 16
	MarketType_SPREAD_1ST_HALF                    MarketType = 17
)

// Enum value maps for MarketType.
var (
	MarketType_name = map[int32]string{
		0:  "MONEYLINE",
		1:  "MONEYLINE_1ST_QUARTER",
		2:  "MONEYLINE_1ST_HALF",
		3:  "OUTRIGHT",
		4:  "TOTAL_POINTS",
		5:  "TOTAL_POINTS_ODD_EVEN",
		6:  "TOTAL_POINTS_1ST_QUARTER",
		7:  "TOTAL_POINTS_1ST_QUARTER_ODD_EVEN",
		8:  "TOTAL_POINTS_1ST_HALF",
		9:  "TOTAL_POINTS_1ST_HALF_ODD_EVEN",
		10: "TOTAL_POINTS_HOME_TEAM",
		11: "TOTAL_POINTS_HOME_TEAM_1ST_QUARTER",
		12: "TOTAL_POINTS_HOME_TEAM_1ST_HALF",
		13: "TOTAL_POINTS_AWAY_TEAM",
		14: "TOTAL_POINTS_AWAY_TEAM_1ST_QUARTER",
		15: "TOTAL_POINTS_AWAY_TEAM_1ST_HALF",
		16: "SPREAD_1ST_QUARTER",
		17: "SPREAD_1ST_HALF",
	}
	MarketType_value = map[string]int32{
		"MONEYLINE":                          0,
		"MONEYLINE_1ST_QUARTER":              1,
		"MONEYLINE_1ST_HALF":                 2,
		"OUTRIGHT":                           3,
		"TOTAL_POINTS":                       4,
		"TOTAL_POINTS_ODD_EVEN":              5,
		"TOTAL_POINTS_1ST_QUARTER":           6,
		"TOTAL_POINTS_1ST_QUARTER_ODD_EVEN":  7,
		"TOTAL_POINTS_1ST_HALF":              8,
		"TOTAL_POINTS_1ST_HALF_ODD_EVEN":     9,
		"TOTAL_POINTS_HOME_TEAM":             10,
		"TOTAL_POINTS_HOME_TEAM_1ST_QUARTER": 11,
		"TOTAL_POINTS_HOME_TEAM_1ST_HALF":    12,
		"TOTAL_POINTS_AWAY_TEAM":             13,
		"TOTAL_POINTS_AWAY_TEAM_1ST_QUARTER": 14,
		"TOTAL_POINTS_AWAY_TEAM_1ST_HALF":    15,
		"SPREAD_1ST_QUARTER":                 16,
		"SPREAD_1ST_HALF":                    17,
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
	return file_integration_proto_enumTypes[2].Descriptor()
}

func (MarketType) Type() protoreflect.EnumType {
	return &file_integration_proto_enumTypes[2]
}

func (x MarketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MarketType.Descriptor instead.
func (MarketType) EnumDescriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{2}
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
	Numerator   string  `protobuf:"bytes,2,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator string  `protobuf:"bytes,3,opt,name=denominator,proto3" json:"denominator,omitempty"`
	American    string  `protobuf:"bytes,4,opt,name=american,proto3" json:"american,omitempty"`
	IsAvailable bool    `protobuf:"varint,5,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
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

func (x *Odds) GetAmerican() string {
	if x != nil {
		return x.American
	}
	return ""
}

func (x *Odds) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Points      float64 `protobuf:"fixed64,2,opt,name=points,proto3" json:"points,omitempty"`
	Odds        *Odds   `protobuf:"bytes,3,opt,name=odds,proto3" json:"odds,omitempty"`
	IsAvailable bool    `protobuf:"varint,4,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
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

func (x *Outcome) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Outcome) GetPoints() float64 {
	if x != nil {
		return x.Points
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

	Type        MarketType `protobuf:"varint,1,opt,name=type,proto3,enum=MarketType" json:"type,omitempty"`
	Outcomes    []*Outcome `protobuf:"bytes,2,rep,name=outcomes,proto3" json:"outcomes,omitempty"`
	IsAvailable bool       `protobuf:"varint,3,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
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
	return MarketType_MONEYLINE
}

func (x *Market) GetOutcomes() []*Outcome {
	if x != nil {
		return x.Outcomes
	}
	return nil
}

func (x *Market) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ParticipantType `protobuf:"varint,1,opt,name=type,proto3,enum=ParticipantType" json:"type,omitempty"`
	Name string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
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

func (x *Participant) GetType() ParticipantType {
	if x != nil {
		return x.Type
	}
	return ParticipantType_HOME
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
	SportType    SportType              `protobuf:"varint,2,opt,name=sport_type,json=sportType,proto3,enum=SportType" json:"sport_type,omitempty"`
	Name         string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	League       string                 `protobuf:"bytes,4,opt,name=league,proto3" json:"league,omitempty"`
	StartTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Participants []*Participant         `protobuf:"bytes,6,rep,name=participants,proto3" json:"participants,omitempty"`
	Markets      []*Market              `protobuf:"bytes,7,rep,name=markets,proto3" json:"markets,omitempty"`
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
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x04, 0x4f, 0x64, 0x64, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x61, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x73, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x19, 0x0a, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x4f, 0x64, 0x64, 0x73, 0x52, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73,
	0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x72, 0x0a,
	0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x47, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x2a, 0x1b, 0x0a, 0x09, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x41, 0x53, 0x4b,
	0x45, 0x54, 0x42, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x2a, 0x35, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x4f, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x57, 0x41, 0x59, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4d, 0x50, 0x45, 0x54, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x02, 0x2a,
	0x8c, 0x04, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d,
	0x0a, 0x09, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51,
	0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x4e, 0x45,
	0x59, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x55, 0x54, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x10, 0x04,
	0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53,
	0x5f, 0x4f, 0x44, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x54,
	0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x31, 0x53, 0x54, 0x5f,
	0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x10, 0x06, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x4f, 0x54,
	0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51, 0x55,
	0x41, 0x52, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x44, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x10, 0x07,
	0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53,
	0x5f, 0x31, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x08, 0x12, 0x22, 0x0a, 0x1e, 0x54,
	0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x31, 0x53, 0x54, 0x5f,
	0x48, 0x41, 0x4c, 0x46, 0x5f, 0x4f, 0x44, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x10, 0x09, 0x12,
	0x1a, 0x0a, 0x16, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f,
	0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x0a, 0x12, 0x26, 0x0a, 0x22, 0x54,
	0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x48, 0x4f, 0x4d, 0x45,
	0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45,
	0x52, 0x10, 0x0b, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x53, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x31, 0x53,
	0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x0c, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f, 0x54, 0x41,
	0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x45,
	0x41, 0x4d, 0x10, 0x0d, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x31,
	0x53, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x10, 0x0e, 0x12, 0x23, 0x0a, 0x1f,
	0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x57, 0x41,
	0x59, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10,
	0x0f, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x50, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x31, 0x53, 0x54, 0x5f,
	0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x10, 0x10, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x50, 0x52,
	0x45, 0x41, 0x44, 0x5f, 0x31, 0x53, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x11, 0x32, 0x55,
	0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x08, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x08,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_integration_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_integration_proto_goTypes = []interface{}{
	(SportType)(0),                // 0: SportType
	(ParticipantType)(0),          // 1: ParticipantType
	(MarketType)(0),               // 2: MarketType
	(*Request)(nil),               // 3: Request
	(*Response)(nil),              // 4: Response
	(*Odds)(nil),                  // 5: Odds
	(*Outcome)(nil),               // 6: Outcome
	(*Market)(nil),                // 7: Market
	(*Participant)(nil),           // 8: Participant
	(*Event)(nil),                 // 9: Event
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_integration_proto_depIdxs = []int32{
	0,  // 0: Request.sport_type:type_name -> SportType
	9,  // 1: Response.events:type_name -> Event
	5,  // 2: Outcome.odds:type_name -> Odds
	2,  // 3: Market.type:type_name -> MarketType
	6,  // 4: Market.outcomes:type_name -> Outcome
	1,  // 5: Participant.type:type_name -> ParticipantType
	0,  // 6: Event.sport_type:type_name -> SportType
	10, // 7: Event.start_time:type_name -> google.protobuf.Timestamp
	8,  // 8: Event.participants:type_name -> Participant
	7,  // 9: Event.markets:type_name -> Market
	3,  // 10: Integration.GetPreMatch:input_type -> Request
	3,  // 11: Integration.GetLive:input_type -> Request
	4,  // 12: Integration.GetPreMatch:output_type -> Response
	4,  // 13: Integration.GetLive:output_type -> Response
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_integration_proto_rawDesc,
			NumEnums:      3,
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