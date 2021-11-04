// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package commonpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_Success               ErrorCode = 0
	ErrorCode_UnexpectedError       ErrorCode = 1
	ErrorCode_ConnectFailed         ErrorCode = 2
	ErrorCode_PermissionDenied      ErrorCode = 3
	ErrorCode_CollectionNotExists   ErrorCode = 4
	ErrorCode_IllegalArgument       ErrorCode = 5
	ErrorCode_IllegalDimension      ErrorCode = 7
	ErrorCode_IllegalIndexType      ErrorCode = 8
	ErrorCode_IllegalCollectionName ErrorCode = 9
	ErrorCode_IllegalTOPK           ErrorCode = 10
	ErrorCode_IllegalRowRecord      ErrorCode = 11
	ErrorCode_IllegalVectorID       ErrorCode = 12
	ErrorCode_IllegalSearchResult   ErrorCode = 13
	ErrorCode_FileNotFound          ErrorCode = 14
	ErrorCode_MetaFailed            ErrorCode = 15
	ErrorCode_CacheFailed           ErrorCode = 16
	ErrorCode_CannotCreateFolder    ErrorCode = 17
	ErrorCode_CannotCreateFile      ErrorCode = 18
	ErrorCode_CannotDeleteFolder    ErrorCode = 19
	ErrorCode_CannotDeleteFile      ErrorCode = 20
	ErrorCode_BuildIndexError       ErrorCode = 21
	ErrorCode_IllegalNLIST          ErrorCode = 22
	ErrorCode_IllegalMetricType     ErrorCode = 23
	ErrorCode_OutOfMemory           ErrorCode = 24
	ErrorCode_IndexNotExist         ErrorCode = 25
	ErrorCode_EmptyCollection       ErrorCode = 26
	// internal error code.
	ErrorCode_DDRequestRace ErrorCode = 1000
)

var ErrorCode_name = map[int32]string{
	0:    "Success",
	1:    "UnexpectedError",
	2:    "ConnectFailed",
	3:    "PermissionDenied",
	4:    "CollectionNotExists",
	5:    "IllegalArgument",
	7:    "IllegalDimension",
	8:    "IllegalIndexType",
	9:    "IllegalCollectionName",
	10:   "IllegalTOPK",
	11:   "IllegalRowRecord",
	12:   "IllegalVectorID",
	13:   "IllegalSearchResult",
	14:   "FileNotFound",
	15:   "MetaFailed",
	16:   "CacheFailed",
	17:   "CannotCreateFolder",
	18:   "CannotCreateFile",
	19:   "CannotDeleteFolder",
	20:   "CannotDeleteFile",
	21:   "BuildIndexError",
	22:   "IllegalNLIST",
	23:   "IllegalMetricType",
	24:   "OutOfMemory",
	25:   "IndexNotExist",
	26:   "EmptyCollection",
	1000: "DDRequestRace",
}

var ErrorCode_value = map[string]int32{
	"Success":               0,
	"UnexpectedError":       1,
	"ConnectFailed":         2,
	"PermissionDenied":      3,
	"CollectionNotExists":   4,
	"IllegalArgument":       5,
	"IllegalDimension":      7,
	"IllegalIndexType":      8,
	"IllegalCollectionName": 9,
	"IllegalTOPK":           10,
	"IllegalRowRecord":      11,
	"IllegalVectorID":       12,
	"IllegalSearchResult":   13,
	"FileNotFound":          14,
	"MetaFailed":            15,
	"CacheFailed":           16,
	"CannotCreateFolder":    17,
	"CannotCreateFile":      18,
	"CannotDeleteFolder":    19,
	"CannotDeleteFile":      20,
	"BuildIndexError":       21,
	"IllegalNLIST":          22,
	"IllegalMetricType":     23,
	"OutOfMemory":           24,
	"IndexNotExist":         25,
	"EmptyCollection":       26,
	"DDRequestRace":         1000,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type IndexState int32

const (
	IndexState_IndexStateNone IndexState = 0
	IndexState_Unissued       IndexState = 1
	IndexState_InProgress     IndexState = 2
	IndexState_Finished       IndexState = 3
	IndexState_Failed         IndexState = 4
)

var IndexState_name = map[int32]string{
	0: "IndexStateNone",
	1: "Unissued",
	2: "InProgress",
	3: "Finished",
	4: "Failed",
}

var IndexState_value = map[string]int32{
	"IndexStateNone": 0,
	"Unissued":       1,
	"InProgress":     2,
	"Finished":       3,
	"Failed":         4,
}

func (x IndexState) String() string {
	return proto.EnumName(IndexState_name, int32(x))
}

func (IndexState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type SegmentState int32

const (
	SegmentState_SegmentStateNone SegmentState = 0
	SegmentState_NotExist         SegmentState = 1
	SegmentState_Growing          SegmentState = 2
	SegmentState_Sealed           SegmentState = 3
	SegmentState_Flushed          SegmentState = 4
	SegmentState_Flushing         SegmentState = 5
)

var SegmentState_name = map[int32]string{
	0: "SegmentStateNone",
	1: "NotExist",
	2: "Growing",
	3: "Sealed",
	4: "Flushed",
	5: "Flushing",
}

var SegmentState_value = map[string]int32{
	"SegmentStateNone": 0,
	"NotExist":         1,
	"Growing":          2,
	"Sealed":           3,
	"Flushed":          4,
	"Flushing":         5,
}

func (x SegmentState) String() string {
	return proto.EnumName(SegmentState_name, int32(x))
}

func (SegmentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type MsgType int32

const (
	MsgType_Undefined MsgType = 0
	// DEFINITION REQUESTS: COLLECTION
	MsgType_CreateCollection   MsgType = 100
	MsgType_DropCollection     MsgType = 101
	MsgType_HasCollection      MsgType = 102
	MsgType_DescribeCollection MsgType = 103
	MsgType_ShowCollections    MsgType = 104
	MsgType_GetSystemConfigs   MsgType = 105
	MsgType_LoadCollection     MsgType = 106
	MsgType_ReleaseCollection  MsgType = 107
	MsgType_CreateAlias        MsgType = 108
	MsgType_DropAlias          MsgType = 109
	MsgType_AlterAlias         MsgType = 110
	// DEFINITION REQUESTS: PARTITION
	MsgType_CreatePartition   MsgType = 200
	MsgType_DropPartition     MsgType = 201
	MsgType_HasPartition      MsgType = 202
	MsgType_DescribePartition MsgType = 203
	MsgType_ShowPartitions    MsgType = 204
	MsgType_LoadPartitions    MsgType = 205
	MsgType_ReleasePartitions MsgType = 206
	// DEFINE REQUESTS: SEGMENT
	MsgType_ShowSegments        MsgType = 250
	MsgType_DescribeSegment     MsgType = 251
	MsgType_LoadSegments        MsgType = 252
	MsgType_ReleaseSegments     MsgType = 253
	MsgType_HandoffSegments     MsgType = 254
	MsgType_LoadBalanceSegments MsgType = 255
	// DEFINITION REQUESTS: INDEX
	MsgType_CreateIndex   MsgType = 300
	MsgType_DescribeIndex MsgType = 301
	MsgType_DropIndex     MsgType = 302
	// MANIPULATION REQUESTS
	MsgType_Insert MsgType = 400
	MsgType_Delete MsgType = 401
	MsgType_Flush  MsgType = 402
	// QUERY
	MsgType_Search                   MsgType = 500
	MsgType_SearchResult             MsgType = 501
	MsgType_GetIndexState            MsgType = 502
	MsgType_GetIndexBuildProgress    MsgType = 503
	MsgType_GetCollectionStatistics  MsgType = 504
	MsgType_GetPartitionStatistics   MsgType = 505
	MsgType_Retrieve                 MsgType = 506
	MsgType_RetrieveResult           MsgType = 507
	MsgType_WatchDmChannels          MsgType = 508
	MsgType_RemoveDmChannels         MsgType = 509
	MsgType_WatchQueryChannels       MsgType = 510
	MsgType_RemoveQueryChannels      MsgType = 511
	MsgType_SealedSegmentsChangeInfo MsgType = 512
	// DATA SERVICE
	MsgType_SegmentInfo MsgType = 600
	MsgType_SystemInfo  MsgType = 601
	// SYSTEM CONTROL
	MsgType_TimeTick          MsgType = 1200
	MsgType_QueryNodeStats    MsgType = 1201
	MsgType_LoadIndex         MsgType = 1202
	MsgType_RequestID         MsgType = 1203
	MsgType_RequestTSO        MsgType = 1204
	MsgType_AllocateSegment   MsgType = 1205
	MsgType_SegmentStatistics MsgType = 1206
	MsgType_SegmentFlushDone  MsgType = 1207
	MsgType_DataNodeTt        MsgType = 1208
)

var MsgType_name = map[int32]string{
	0:    "Undefined",
	100:  "CreateCollection",
	101:  "DropCollection",
	102:  "HasCollection",
	103:  "DescribeCollection",
	104:  "ShowCollections",
	105:  "GetSystemConfigs",
	106:  "LoadCollection",
	107:  "ReleaseCollection",
	108:  "CreateAlias",
	109:  "DropAlias",
	110:  "AlterAlias",
	200:  "CreatePartition",
	201:  "DropPartition",
	202:  "HasPartition",
	203:  "DescribePartition",
	204:  "ShowPartitions",
	205:  "LoadPartitions",
	206:  "ReleasePartitions",
	250:  "ShowSegments",
	251:  "DescribeSegment",
	252:  "LoadSegments",
	253:  "ReleaseSegments",
	254:  "HandoffSegments",
	255:  "LoadBalanceSegments",
	300:  "CreateIndex",
	301:  "DescribeIndex",
	302:  "DropIndex",
	400:  "Insert",
	401:  "Delete",
	402:  "Flush",
	500:  "Search",
	501:  "SearchResult",
	502:  "GetIndexState",
	503:  "GetIndexBuildProgress",
	504:  "GetCollectionStatistics",
	505:  "GetPartitionStatistics",
	506:  "Retrieve",
	507:  "RetrieveResult",
	508:  "WatchDmChannels",
	509:  "RemoveDmChannels",
	510:  "WatchQueryChannels",
	511:  "RemoveQueryChannels",
	512:  "SealedSegmentsChangeInfo",
	600:  "SegmentInfo",
	601:  "SystemInfo",
	1200: "TimeTick",
	1201: "QueryNodeStats",
	1202: "LoadIndex",
	1203: "RequestID",
	1204: "RequestTSO",
	1205: "AllocateSegment",
	1206: "SegmentStatistics",
	1207: "SegmentFlushDone",
	1208: "DataNodeTt",
}

var MsgType_value = map[string]int32{
	"Undefined":                0,
	"CreateCollection":         100,
	"DropCollection":           101,
	"HasCollection":            102,
	"DescribeCollection":       103,
	"ShowCollections":          104,
	"GetSystemConfigs":         105,
	"LoadCollection":           106,
	"ReleaseCollection":        107,
	"CreateAlias":              108,
	"DropAlias":                109,
	"AlterAlias":               110,
	"CreatePartition":          200,
	"DropPartition":            201,
	"HasPartition":             202,
	"DescribePartition":        203,
	"ShowPartitions":           204,
	"LoadPartitions":           205,
	"ReleasePartitions":        206,
	"ShowSegments":             250,
	"DescribeSegment":          251,
	"LoadSegments":             252,
	"ReleaseSegments":          253,
	"HandoffSegments":          254,
	"LoadBalanceSegments":      255,
	"CreateIndex":              300,
	"DescribeIndex":            301,
	"DropIndex":                302,
	"Insert":                   400,
	"Delete":                   401,
	"Flush":                    402,
	"Search":                   500,
	"SearchResult":             501,
	"GetIndexState":            502,
	"GetIndexBuildProgress":    503,
	"GetCollectionStatistics":  504,
	"GetPartitionStatistics":   505,
	"Retrieve":                 506,
	"RetrieveResult":           507,
	"WatchDmChannels":          508,
	"RemoveDmChannels":         509,
	"WatchQueryChannels":       510,
	"RemoveQueryChannels":      511,
	"SealedSegmentsChangeInfo": 512,
	"SegmentInfo":              600,
	"SystemInfo":               601,
	"TimeTick":                 1200,
	"QueryNodeStats":           1201,
	"LoadIndex":                1202,
	"RequestID":                1203,
	"RequestTSO":               1204,
	"AllocateSegment":          1205,
	"SegmentStatistics":        1206,
	"SegmentFlushDone":         1207,
	"DataNodeTt":               1208,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

type DslType int32

const (
	DslType_Dsl        DslType = 0
	DslType_BoolExprV1 DslType = 1
)

var DslType_name = map[int32]string{
	0: "Dsl",
	1: "BoolExprV1",
}

var DslType_value = map[string]int32{
	"Dsl":        0,
	"BoolExprV1": 1,
}

func (x DslType) String() string {
	return proto.EnumName(DslType_name, int32(x))
}

func (DslType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

type Status struct {
	ErrorCode            ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=milvus.proto.common.ErrorCode" json:"error_code,omitempty"`
	Reason               string    `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_Success
}

func (m *Status) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type KeyValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KeyDataPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyDataPair) Reset()         { *m = KeyDataPair{} }
func (m *KeyDataPair) String() string { return proto.CompactTextString(m) }
func (*KeyDataPair) ProtoMessage()    {}
func (*KeyDataPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *KeyDataPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyDataPair.Unmarshal(m, b)
}
func (m *KeyDataPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyDataPair.Marshal(b, m, deterministic)
}
func (m *KeyDataPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyDataPair.Merge(m, src)
}
func (m *KeyDataPair) XXX_Size() int {
	return xxx_messageInfo_KeyDataPair.Size(m)
}
func (m *KeyDataPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyDataPair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyDataPair proto.InternalMessageInfo

func (m *KeyDataPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyDataPair) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Blob struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blob) Reset()         { *m = Blob{} }
func (m *Blob) String() string { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()    {}
func (*Blob) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *Blob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blob.Unmarshal(m, b)
}
func (m *Blob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blob.Marshal(b, m, deterministic)
}
func (m *Blob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blob.Merge(m, src)
}
func (m *Blob) XXX_Size() int {
	return xxx_messageInfo_Blob.Size(m)
}
func (m *Blob) XXX_DiscardUnknown() {
	xxx_messageInfo_Blob.DiscardUnknown(m)
}

var xxx_messageInfo_Blob proto.InternalMessageInfo

func (m *Blob) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Address struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Address) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type MsgBase struct {
	MsgType              MsgType  `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=milvus.proto.common.MsgType" json:"msg_type,omitempty"`
	MsgID                int64    `protobuf:"varint,2,opt,name=msgID,proto3" json:"msgID,omitempty"`
	Timestamp            uint64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SourceID             int64    `protobuf:"varint,4,opt,name=sourceID,proto3" json:"sourceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBase) Reset()         { *m = MsgBase{} }
func (m *MsgBase) String() string { return proto.CompactTextString(m) }
func (*MsgBase) ProtoMessage()    {}
func (*MsgBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *MsgBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBase.Unmarshal(m, b)
}
func (m *MsgBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBase.Marshal(b, m, deterministic)
}
func (m *MsgBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBase.Merge(m, src)
}
func (m *MsgBase) XXX_Size() int {
	return xxx_messageInfo_MsgBase.Size(m)
}
func (m *MsgBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBase.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBase proto.InternalMessageInfo

func (m *MsgBase) GetMsgType() MsgType {
	if m != nil {
		return m.MsgType
	}
	return MsgType_Undefined
}

func (m *MsgBase) GetMsgID() int64 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *MsgBase) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MsgBase) GetSourceID() int64 {
	if m != nil {
		return m.SourceID
	}
	return 0
}

// Don't Modify This. @czs
type MsgHeader struct {
	Base                 *MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgHeader) Reset()         { *m = MsgHeader{} }
func (m *MsgHeader) String() string { return proto.CompactTextString(m) }
func (*MsgHeader) ProtoMessage()    {}
func (*MsgHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *MsgHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgHeader.Unmarshal(m, b)
}
func (m *MsgHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgHeader.Marshal(b, m, deterministic)
}
func (m *MsgHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHeader.Merge(m, src)
}
func (m *MsgHeader) XXX_Size() int {
	return xxx_messageInfo_MsgHeader.Size(m)
}
func (m *MsgHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHeader proto.InternalMessageInfo

func (m *MsgHeader) GetBase() *MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

// Don't Modify This. @czs
type DMLMsgHeader struct {
	Base                 *MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ShardName            string   `protobuf:"bytes,2,opt,name=shardName,proto3" json:"shardName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DMLMsgHeader) Reset()         { *m = DMLMsgHeader{} }
func (m *DMLMsgHeader) String() string { return proto.CompactTextString(m) }
func (*DMLMsgHeader) ProtoMessage()    {}
func (*DMLMsgHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

func (m *DMLMsgHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DMLMsgHeader.Unmarshal(m, b)
}
func (m *DMLMsgHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DMLMsgHeader.Marshal(b, m, deterministic)
}
func (m *DMLMsgHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DMLMsgHeader.Merge(m, src)
}
func (m *DMLMsgHeader) XXX_Size() int {
	return xxx_messageInfo_DMLMsgHeader.Size(m)
}
func (m *DMLMsgHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DMLMsgHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DMLMsgHeader proto.InternalMessageInfo

func (m *DMLMsgHeader) GetBase() *MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *DMLMsgHeader) GetShardName() string {
	if m != nil {
		return m.ShardName
	}
	return ""
}

func init() {
	proto.RegisterEnum("milvus.proto.common.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("milvus.proto.common.IndexState", IndexState_name, IndexState_value)
	proto.RegisterEnum("milvus.proto.common.SegmentState", SegmentState_name, SegmentState_value)
	proto.RegisterEnum("milvus.proto.common.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("milvus.proto.common.DslType", DslType_name, DslType_value)
	proto.RegisterType((*Status)(nil), "milvus.proto.common.Status")
	proto.RegisterType((*KeyValuePair)(nil), "milvus.proto.common.KeyValuePair")
	proto.RegisterType((*KeyDataPair)(nil), "milvus.proto.common.KeyDataPair")
	proto.RegisterType((*Blob)(nil), "milvus.proto.common.Blob")
	proto.RegisterType((*Address)(nil), "milvus.proto.common.Address")
	proto.RegisterType((*MsgBase)(nil), "milvus.proto.common.MsgBase")
	proto.RegisterType((*MsgHeader)(nil), "milvus.proto.common.MsgHeader")
	proto.RegisterType((*DMLMsgHeader)(nil), "milvus.proto.common.DMLMsgHeader")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4b, 0x73, 0x1b, 0x37,
	0x12, 0xd6, 0x70, 0x28, 0x51, 0x84, 0x28, 0x09, 0x82, 0x1e, 0x96, 0xbd, 0xda, 0x2d, 0x17, 0x4f,
	0x2e, 0x55, 0x59, 0xda, 0x5d, 0xd7, 0xee, 0x9e, 0x7c, 0x90, 0x38, 0x7a, 0xb0, 0x6c, 0x3d, 0x76,
	0x28, 0x3b, 0xa9, 0x1c, 0xe2, 0x82, 0x66, 0x9a, 0x24, 0xe2, 0x19, 0x80, 0x01, 0x40, 0x59, 0xbc,
	0xe5, 0x27, 0x24, 0xfe, 0x1d, 0x49, 0x2a, 0xef, 0xa4, 0xf2, 0x0b, 0xf2, 0x3e, 0x27, 0xf7, 0x1c,
	0xf2, 0x03, 0xf2, 0xf4, 0x33, 0xd5, 0x98, 0x21, 0x39, 0xae, 0xb2, 0x4f, 0xb9, 0xa1, 0x3f, 0x34,
	0xbe, 0xee, 0xfe, 0xba, 0x81, 0x19, 0x52, 0x8b, 0x54, 0x9a, 0x2a, 0xb9, 0xd1, 0xd3, 0xca, 0x2a,
	0xb6, 0x98, 0x8a, 0xe4, 0xac, 0x6f, 0x32, 0x6b, 0x23, 0xdb, 0xaa, 0xdf, 0x21, 0x53, 0x2d, 0xcb,
	0x6d, 0xdf, 0xb0, 0xeb, 0x84, 0x80, 0xd6, 0x4a, 0xdf, 0x89, 0x54, 0x0c, 0xab, 0xde, 0x65, 0xef,
	0xca, 0xdc, 0xbf, 0xff, 0xb1, 0xf1, 0x9c, 0x33, 0x1b, 0x3b, 0xe8, 0xd6, 0x50, 0x31, 0x84, 0x55,
	0x18, 0x2e, 0xd9, 0x0a, 0x99, 0xd2, 0xc0, 0x8d, 0x92, 0xab, 0xa5, 0xcb, 0xde, 0x95, 0x6a, 0x98,
	0x5b, 0xf5, 0xff, 0x92, 0xda, 0x0d, 0x18, 0xdc, 0xe6, 0x49, 0x1f, 0x8e, 0xb9, 0xd0, 0x8c, 0x12,
	0xff, 0x2e, 0x0c, 0x1c, 0x7f, 0x35, 0xc4, 0x25, 0x5b, 0x22, 0x93, 0x67, 0xb8, 0x9d, 0x1f, 0xcc,
	0x8c, 0xfa, 0x35, 0x32, 0x73, 0x03, 0x06, 0x01, 0xb7, 0xfc, 0x05, 0xc7, 0x18, 0x29, 0xc7, 0xdc,
	0x72, 0x77, 0xaa, 0x16, 0xba, 0x75, 0x7d, 0x8d, 0x94, 0xb7, 0x13, 0x75, 0x3a, 0xa6, 0xf4, 0xdc,
	0x66, 0x4e, 0x79, 0x95, 0x54, 0xb6, 0xe2, 0x58, 0x83, 0x31, 0x6c, 0x8e, 0x94, 0x44, 0x2f, 0x67,
	0x2b, 0x89, 0x1e, 0x92, 0xf5, 0x94, 0xb6, 0x8e, 0xcc, 0x0f, 0xdd, 0xba, 0x7e, 0xdf, 0x23, 0x95,
	0x03, 0xd3, 0xd9, 0xe6, 0x06, 0xd8, 0xff, 0xc8, 0x74, 0x6a, 0x3a, 0x77, 0xec, 0xa0, 0x37, 0x94,
	0x66, 0xed, 0xb9, 0xd2, 0x1c, 0x98, 0xce, 0xc9, 0xa0, 0x07, 0x61, 0x25, 0xcd, 0x16, 0x98, 0x49,
	0x6a, 0x3a, 0xcd, 0x20, 0x67, 0xce, 0x0c, 0xb6, 0x46, 0xaa, 0x56, 0xa4, 0x60, 0x2c, 0x4f, 0x7b,
	0xab, 0xfe, 0x65, 0xef, 0x4a, 0x39, 0x1c, 0x03, 0xec, 0x12, 0x99, 0x36, 0xaa, 0xaf, 0x23, 0x68,
	0x06, 0xab, 0x65, 0x77, 0x6c, 0x64, 0xd7, 0xaf, 0x93, 0xea, 0x81, 0xe9, 0xec, 0x03, 0x8f, 0x41,
	0xb3, 0x7f, 0x92, 0xf2, 0x29, 0x37, 0x59, 0x46, 0x33, 0x2f, 0xce, 0x08, 0x2b, 0x08, 0x9d, 0x67,
	0xfd, 0x55, 0x52, 0x0b, 0x0e, 0x6e, 0xfe, 0x05, 0x06, 0x4c, 0xdd, 0x74, 0xb9, 0x8e, 0x0f, 0x79,
	0x3a, 0xec, 0xd8, 0x18, 0x58, 0xff, 0xbc, 0x4c, 0xaa, 0xa3, 0xf1, 0x60, 0x33, 0xa4, 0xd2, 0xea,
	0x47, 0x11, 0x18, 0x43, 0x27, 0xd8, 0x22, 0x99, 0xbf, 0x25, 0xe1, 0xbc, 0x07, 0x91, 0x85, 0xd8,
	0xf9, 0x50, 0x8f, 0x2d, 0x90, 0xd9, 0x86, 0x92, 0x12, 0x22, 0xbb, 0xcb, 0x45, 0x02, 0x31, 0x2d,
	0xb1, 0x25, 0x42, 0x8f, 0x41, 0xa7, 0xc2, 0x18, 0xa1, 0x64, 0x00, 0x52, 0x40, 0x4c, 0x7d, 0x76,
	0x81, 0x2c, 0x36, 0x54, 0x92, 0x40, 0x64, 0x85, 0x92, 0x87, 0xca, 0xee, 0x9c, 0x0b, 0x63, 0x0d,
	0x2d, 0x23, 0x6d, 0x33, 0x49, 0xa0, 0xc3, 0x93, 0x2d, 0xdd, 0xe9, 0xa7, 0x20, 0x2d, 0x9d, 0x44,
	0x8e, 0x1c, 0x0c, 0x44, 0x0a, 0x12, 0x99, 0x68, 0xa5, 0x80, 0x36, 0x65, 0x0c, 0xe7, 0xd8, 0x1f,
	0x3a, 0xcd, 0x2e, 0x92, 0xe5, 0x1c, 0x2d, 0x04, 0xe0, 0x29, 0xd0, 0x2a, 0x9b, 0x27, 0x33, 0xf9,
	0xd6, 0xc9, 0xd1, 0xf1, 0x0d, 0x4a, 0x0a, 0x0c, 0xa1, 0xba, 0x17, 0x42, 0xa4, 0x74, 0x4c, 0x67,
	0x0a, 0x29, 0xdc, 0x86, 0xc8, 0x2a, 0xdd, 0x0c, 0x68, 0x0d, 0x13, 0xce, 0xc1, 0x16, 0x70, 0x1d,
	0x75, 0x43, 0x30, 0xfd, 0xc4, 0xd2, 0x59, 0x46, 0x49, 0x6d, 0x57, 0x24, 0x70, 0xa8, 0xec, 0xae,
	0xea, 0xcb, 0x98, 0xce, 0xb1, 0x39, 0x42, 0x0e, 0xc0, 0xf2, 0x5c, 0x81, 0x79, 0x0c, 0xdb, 0xe0,
	0x51, 0x17, 0x72, 0x80, 0xb2, 0x15, 0xc2, 0x1a, 0x5c, 0x4a, 0x65, 0x1b, 0x1a, 0xb8, 0x85, 0x5d,
	0x95, 0xc4, 0xa0, 0xe9, 0x02, 0xa6, 0xf3, 0x0c, 0x2e, 0x12, 0xa0, 0x6c, 0xec, 0x1d, 0x40, 0x02,
	0x23, 0xef, 0xc5, 0xb1, 0x77, 0x8e, 0xa3, 0xf7, 0x12, 0x26, 0xbf, 0xdd, 0x17, 0x49, 0xec, 0x24,
	0xc9, 0xda, 0xb2, 0x8c, 0x39, 0xe6, 0xc9, 0x1f, 0xde, 0x6c, 0xb6, 0x4e, 0xe8, 0x0a, 0x5b, 0x26,
	0x0b, 0x39, 0x72, 0x00, 0x56, 0x8b, 0xc8, 0x89, 0x77, 0x01, 0x53, 0x3d, 0xea, 0xdb, 0xa3, 0xf6,
	0x01, 0xa4, 0x4a, 0x0f, 0xe8, 0x2a, 0x36, 0xd4, 0x31, 0x0d, 0x5b, 0x44, 0x2f, 0x62, 0x84, 0x9d,
	0xb4, 0x67, 0x07, 0x63, 0x79, 0xe9, 0x25, 0xc6, 0xc8, 0x6c, 0x10, 0x84, 0xf0, 0x7a, 0x1f, 0x8c,
	0x0d, 0x79, 0x04, 0xf4, 0xa7, 0xca, 0xfa, 0xcb, 0x84, 0xb8, 0xb3, 0xf8, 0x20, 0x01, 0x63, 0x64,
	0x6e, 0x6c, 0x1d, 0x2a, 0x09, 0x74, 0x82, 0xd5, 0xc8, 0xf4, 0x2d, 0x29, 0x8c, 0xe9, 0x43, 0x4c,
	0x3d, 0xd4, 0xad, 0x29, 0x8f, 0xb5, 0xea, 0xe0, 0x95, 0xa6, 0x25, 0xdc, 0xdd, 0x15, 0x52, 0x98,
	0xae, 0x9b, 0x18, 0x42, 0xa6, 0x72, 0x01, 0xcb, 0xeb, 0x6d, 0x52, 0x6b, 0x41, 0x07, 0x87, 0x23,
	0xe3, 0x5e, 0x22, 0xb4, 0x68, 0x8f, 0xd9, 0x47, 0x69, 0x7b, 0x38, 0xbc, 0x7b, 0x5a, 0xdd, 0x13,
	0xb2, 0x43, 0x4b, 0x48, 0xd6, 0x02, 0x9e, 0x38, 0xe2, 0x19, 0x52, 0xd9, 0x4d, 0xfa, 0x2e, 0x4a,
	0xd9, 0xc5, 0x44, 0x03, 0xdd, 0x26, 0xd7, 0x7f, 0x9c, 0x76, 0x4f, 0x86, 0xbb, 0xf9, 0xb3, 0xa4,
	0x7a, 0x4b, 0xc6, 0xd0, 0x16, 0x12, 0x62, 0x3a, 0xe1, 0xd4, 0x77, 0x5d, 0x2a, 0xc8, 0x10, 0x63,
	0x91, 0x81, 0x56, 0xbd, 0x02, 0x06, 0x28, 0xe1, 0x3e, 0x37, 0x05, 0xa8, 0x8d, 0x2d, 0x0d, 0xc0,
	0x44, 0x5a, 0x9c, 0x16, 0x8f, 0x77, 0x50, 0xda, 0x56, 0x57, 0xdd, 0x1b, 0x63, 0x86, 0x76, 0x31,
	0xd2, 0x1e, 0xd8, 0xd6, 0xc0, 0x58, 0x48, 0x1b, 0x4a, 0xb6, 0x45, 0xc7, 0x50, 0x81, 0x91, 0x6e,
	0x2a, 0x1e, 0x17, 0x8e, 0xbf, 0x86, 0x4d, 0x0d, 0x21, 0x01, 0x6e, 0x8a, 0xac, 0x77, 0xdd, 0xfc,
	0xb9, 0x54, 0xb7, 0x12, 0xc1, 0x0d, 0x4d, 0xb0, 0x14, 0xcc, 0x32, 0x33, 0x53, 0xd4, 0x7d, 0x2b,
	0xb1, 0xa0, 0x33, 0x5b, 0xb2, 0x25, 0x32, 0x9f, 0xf9, 0x1f, 0x73, 0x6d, 0x85, 0x23, 0xf9, 0xc2,
	0x73, 0x1d, 0xd6, 0xaa, 0x37, 0xc6, 0xbe, 0xc4, 0xeb, 0x5e, 0xdb, 0xe7, 0x66, 0x0c, 0x7d, 0xe5,
	0xb1, 0x15, 0xb2, 0x30, 0x2c, 0x6d, 0x8c, 0x7f, 0xed, 0xb1, 0x45, 0x32, 0x87, 0xa5, 0x8d, 0x30,
	0x43, 0xbf, 0x71, 0x20, 0x16, 0x51, 0x00, 0xbf, 0x75, 0x0c, 0x79, 0x15, 0x05, 0xfc, 0x3b, 0x17,
	0x0c, 0x19, 0xf2, 0x46, 0x1b, 0xfa, 0xc0, 0xc3, 0x4c, 0x87, 0xc1, 0x72, 0x98, 0x3e, 0x74, 0x8e,
	0xc8, 0x3a, 0x72, 0x7c, 0xe4, 0x1c, 0x73, 0xce, 0x11, 0xfa, 0xd8, 0xa1, 0xfb, 0x5c, 0xc6, 0xaa,
	0xdd, 0x1e, 0xa1, 0x4f, 0x3c, 0xb6, 0x4a, 0x16, 0xf1, 0xf8, 0x36, 0x4f, 0xb8, 0x8c, 0xc6, 0xfe,
	0x4f, 0x3d, 0x46, 0x87, 0x42, 0xba, 0x41, 0xa6, 0x6f, 0x97, 0x9c, 0x28, 0x79, 0x02, 0x19, 0xf6,
	0x4e, 0x89, 0xcd, 0x65, 0xea, 0x66, 0xf6, 0xbb, 0x25, 0x36, 0x43, 0xa6, 0x9a, 0xd2, 0x80, 0xb6,
	0xf4, 0x4d, 0x1c, 0xb6, 0xa9, 0xec, 0xba, 0xd2, 0xb7, 0x70, 0xa4, 0x27, 0xdd, 0xb0, 0xd1, 0xfb,
	0x6e, 0x23, 0x7b, 0x58, 0xe8, 0xcf, 0xbe, 0x2b, 0xb5, 0xf8, 0xca, 0xfc, 0xe2, 0x63, 0xa4, 0x3d,
	0xb0, 0xe3, 0x1b, 0x44, 0x7f, 0xf5, 0xd9, 0x25, 0xb2, 0x3c, 0xc4, 0xdc, 0x9d, 0x1f, 0xdd, 0x9d,
	0xdf, 0x7c, 0xb6, 0x46, 0x2e, 0xec, 0x81, 0x1d, 0xcf, 0x01, 0x1e, 0x12, 0xc6, 0x8a, 0xc8, 0xd0,
	0xdf, 0x7d, 0xf6, 0x37, 0xb2, 0xb2, 0x07, 0x76, 0xa4, 0x6f, 0x61, 0xf3, 0x0f, 0x9f, 0xcd, 0x92,
	0xe9, 0x10, 0x1f, 0x05, 0x38, 0x03, 0xfa, 0xc0, 0xc7, 0x26, 0x0d, 0xcd, 0x3c, 0x9d, 0x87, 0x3e,
	0x4a, 0xf7, 0x12, 0xb7, 0x51, 0x37, 0x48, 0x1b, 0x5d, 0x2e, 0x25, 0x24, 0x86, 0x3e, 0xf2, 0xd9,
	0x32, 0xa1, 0x21, 0xa4, 0xea, 0x0c, 0x0a, 0xf0, 0x63, 0x7c, 0xec, 0x99, 0x73, 0xfe, 0x7f, 0x1f,
	0xf4, 0x60, 0xb4, 0xf1, 0xc4, 0x47, 0xa9, 0x33, 0xff, 0x67, 0x77, 0x9e, 0xfa, 0xec, 0xef, 0x64,
	0x35, 0xbb, 0xa0, 0x43, 0xfd, 0x71, 0xb3, 0x03, 0x4d, 0xd9, 0x56, 0xf4, 0x8d, 0x32, 0x76, 0x22,
	0xdf, 0x70, 0xc8, 0xf7, 0x65, 0x36, 0x4f, 0x48, 0x76, 0x45, 0x1c, 0xf0, 0x43, 0x19, 0xab, 0x38,
	0x11, 0x29, 0x9c, 0x88, 0xe8, 0x2e, 0x7d, 0xaf, 0x8a, 0x55, 0xb8, 0x20, 0x87, 0x2a, 0x06, 0x2c,
	0xd7, 0xd0, 0xf7, 0xab, 0xd8, 0x2a, 0x6c, 0x75, 0xd6, 0xaa, 0x0f, 0x9c, 0x9d, 0xbf, 0x61, 0xcd,
	0x80, 0x7e, 0x88, 0x1f, 0x0c, 0x92, 0xdb, 0x27, 0xad, 0x23, 0xfa, 0x51, 0x15, 0xcb, 0xde, 0x4a,
	0x12, 0x15, 0x71, 0x3b, 0x1a, 0xb8, 0x8f, 0xab, 0x38, 0xb1, 0x85, 0xe7, 0x27, 0x17, 0xf2, 0x93,
	0x2a, 0xca, 0x91, 0xe3, 0xae, 0xcd, 0x01, 0x3e, 0x4b, 0x9f, 0x3a, 0x56, 0xfc, 0x0f, 0xc2, 0x4c,
	0x4e, 0x2c, 0xfd, 0xac, 0xba, 0x5e, 0x27, 0x95, 0xc0, 0x24, 0xee, 0x95, 0xa9, 0x10, 0x3f, 0x30,
	0x09, 0x9d, 0xc0, 0x4b, 0xb9, 0xad, 0x54, 0xb2, 0x73, 0xde, 0xd3, 0xb7, 0xff, 0x45, 0xbd, 0xed,
	0xff, 0xbc, 0x72, 0xad, 0x23, 0x6c, 0xb7, 0x7f, 0x8a, 0x9f, 0xf1, 0xcd, 0xec, 0xbb, 0x7e, 0x55,
	0xa8, 0x7c, 0xb5, 0x29, 0xa4, 0x05, 0x2d, 0x79, 0xb2, 0xe9, 0x3e, 0xf5, 0x9b, 0xd9, 0xa7, 0xbe,
	0x77, 0x7a, 0x3a, 0xe5, 0xec, 0x6b, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x69, 0x88, 0x56,
	0x3b, 0x0a, 0x00, 0x00,
}
