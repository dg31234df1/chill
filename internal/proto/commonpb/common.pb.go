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
	MsgType_Search                  MsgType = 500
	MsgType_SearchResult            MsgType = 501
	MsgType_GetIndexState           MsgType = 502
	MsgType_GetIndexBuildProgress   MsgType = 503
	MsgType_GetCollectionStatistics MsgType = 504
	MsgType_GetPartitionStatistics  MsgType = 505
	MsgType_Retrieve                MsgType = 506
	MsgType_RetrieveResult          MsgType = 507
	MsgType_WatchDmChannels         MsgType = 508
	MsgType_RemoveDmChannels        MsgType = 509
	MsgType_WatchQueryChannels      MsgType = 510
	MsgType_RemoveQueryChannels     MsgType = 511
	// DATA SERVICE
	MsgType_SegmentInfo MsgType = 600
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
	600:  "SegmentInfo",
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
	"Undefined":               0,
	"CreateCollection":        100,
	"DropCollection":          101,
	"HasCollection":           102,
	"DescribeCollection":      103,
	"ShowCollections":         104,
	"GetSystemConfigs":        105,
	"LoadCollection":          106,
	"ReleaseCollection":       107,
	"CreateAlias":             108,
	"DropAlias":               109,
	"AlterAlias":              110,
	"CreatePartition":         200,
	"DropPartition":           201,
	"HasPartition":            202,
	"DescribePartition":       203,
	"ShowPartitions":          204,
	"LoadPartitions":          205,
	"ReleasePartitions":       206,
	"ShowSegments":            250,
	"DescribeSegment":         251,
	"LoadSegments":            252,
	"ReleaseSegments":         253,
	"HandoffSegments":         254,
	"LoadBalanceSegments":     255,
	"CreateIndex":             300,
	"DescribeIndex":           301,
	"DropIndex":               302,
	"Insert":                  400,
	"Delete":                  401,
	"Flush":                   402,
	"Search":                  500,
	"SearchResult":            501,
	"GetIndexState":           502,
	"GetIndexBuildProgress":   503,
	"GetCollectionStatistics": 504,
	"GetPartitionStatistics":  505,
	"Retrieve":                506,
	"RetrieveResult":          507,
	"WatchDmChannels":         508,
	"RemoveDmChannels":        509,
	"WatchQueryChannels":      510,
	"RemoveQueryChannels":     511,
	"SegmentInfo":             600,
	"TimeTick":                1200,
	"QueryNodeStats":          1201,
	"LoadIndex":               1202,
	"RequestID":               1203,
	"RequestTSO":              1204,
	"AllocateSegment":         1205,
	"SegmentStatistics":       1206,
	"SegmentFlushDone":        1207,
	"DataNodeTt":              1208,
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
	return fileDescriptor_555bd8c177793206, []int{2}
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
	return fileDescriptor_555bd8c177793206, []int{3}
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
	return fileDescriptor_555bd8c177793206, []int{4}
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
	return fileDescriptor_555bd8c177793206, []int{5}
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

func init() {
	proto.RegisterEnum("milvus.proto.common.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("milvus.proto.common.IndexState", IndexState_name, IndexState_value)
	proto.RegisterEnum("milvus.proto.common.SegmentState", SegmentState_name, SegmentState_value)
	proto.RegisterEnum("milvus.proto.common.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("milvus.proto.common.DslType", DslType_name, DslType_value)
	proto.RegisterType((*Status)(nil), "milvus.proto.common.Status")
	proto.RegisterType((*KeyValuePair)(nil), "milvus.proto.common.KeyValuePair")
	proto.RegisterType((*Blob)(nil), "milvus.proto.common.Blob")
	proto.RegisterType((*Address)(nil), "milvus.proto.common.Address")
	proto.RegisterType((*MsgBase)(nil), "milvus.proto.common.MsgBase")
	proto.RegisterType((*MsgHeader)(nil), "milvus.proto.common.MsgHeader")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdb, 0x6e, 0x1b, 0xb7,
	0x16, 0xf5, 0x68, 0x64, 0xcb, 0xa2, 0x65, 0x9b, 0xa6, 0x2f, 0x71, 0x72, 0x8c, 0x83, 0xc0, 0x4f,
	0x81, 0x81, 0xd8, 0xe7, 0x9c, 0xe0, 0xb4, 0x4f, 0x79, 0xb0, 0x35, 0xbe, 0x08, 0x89, 0x2f, 0x1d,
	0x39, 0x69, 0xd1, 0x97, 0x80, 0x9e, 0xd9, 0x92, 0xd8, 0xcc, 0x90, 0x2a, 0xc9, 0x71, 0xac, 0xbf,
	0x68, 0xf3, 0x1d, 0x6d, 0xd1, 0x4b, 0x7a, 0x41, 0xbf, 0xa0, 0xf7, 0xe7, 0x7e, 0x42, 0x3f, 0xa0,
	0xd7, 0x5c, 0x8b, 0xcd, 0x19, 0x49, 0x13, 0x20, 0x7d, 0x9b, 0xbd, 0xb8, 0xb9, 0xb8, 0xf6, 0xda,
	0x9b, 0x1c, 0xd2, 0x88, 0x54, 0x9a, 0x2a, 0xb9, 0xd9, 0xd7, 0xca, 0x2a, 0xb6, 0x98, 0x8a, 0xe4,
	0x3c, 0x33, 0x79, 0xb4, 0x99, 0x2f, 0xad, 0xdf, 0x23, 0x53, 0x6d, 0xcb, 0x6d, 0x66, 0xd8, 0x4d,
	0x42, 0x40, 0x6b, 0xa5, 0xef, 0x45, 0x2a, 0x86, 0x55, 0xef, 0xaa, 0x77, 0x6d, 0xee, 0x7f, 0xff,
	0xde, 0x7c, 0xc5, 0x9e, 0xcd, 0x5d, 0x4c, 0x6b, 0xaa, 0x18, 0xc2, 0x3a, 0x0c, 0x3f, 0xd9, 0x0a,
	0x99, 0xd2, 0xc0, 0x8d, 0x92, 0xab, 0x95, 0xab, 0xde, 0xb5, 0x7a, 0x58, 0x44, 0xeb, 0xaf, 0x91,
	0xc6, 0x2d, 0x18, 0xdc, 0xe5, 0x49, 0x06, 0x27, 0x5c, 0x68, 0x46, 0x89, 0x7f, 0x1f, 0x06, 0x8e,
	0xbf, 0x1e, 0xe2, 0x27, 0x5b, 0x22, 0x93, 0xe7, 0xb8, 0x5c, 0x6c, 0xcc, 0x83, 0xf5, 0x35, 0x52,
	0xdd, 0x49, 0xd4, 0xd9, 0x78, 0x15, 0x77, 0x34, 0x86, 0xab, 0xd7, 0x49, 0x6d, 0x3b, 0x8e, 0x35,
	0x18, 0xc3, 0xe6, 0x48, 0x45, 0xf4, 0x0b, 0xbe, 0x8a, 0xe8, 0x33, 0x46, 0xaa, 0x7d, 0xa5, 0xad,
	0x63, 0xf3, 0x43, 0xf7, 0xbd, 0xfe, 0xd0, 0x23, 0xb5, 0x43, 0xd3, 0xdd, 0xe1, 0x06, 0xd8, 0xeb,
	0x64, 0x3a, 0x35, 0xdd, 0x7b, 0x76, 0xd0, 0x1f, 0x56, 0xb9, 0xf6, 0xca, 0x2a, 0x0f, 0x4d, 0xf7,
	0x74, 0xd0, 0x87, 0xb0, 0x96, 0xe6, 0x1f, 0xa8, 0x24, 0x35, 0xdd, 0x56, 0x50, 0x30, 0xe7, 0x01,
	0x5b, 0x23, 0x75, 0x2b, 0x52, 0x30, 0x96, 0xa7, 0xfd, 0x55, 0xff, 0xaa, 0x77, 0xad, 0x1a, 0x8e,
	0x01, 0x76, 0x85, 0x4c, 0x1b, 0x95, 0xe9, 0x08, 0x5a, 0xc1, 0x6a, 0xd5, 0x6d, 0x1b, 0xc5, 0xeb,
	0x37, 0x49, 0xfd, 0xd0, 0x74, 0x0f, 0x80, 0xc7, 0xa0, 0xd9, 0x7f, 0x48, 0xf5, 0x8c, 0x9b, 0x5c,
	0xd1, 0xcc, 0x3f, 0x2b, 0xc2, 0x0a, 0x42, 0x97, 0xb9, 0xf1, 0x75, 0x95, 0xd4, 0x47, 0x9d, 0x60,
	0x33, 0xa4, 0xd6, 0xce, 0xa2, 0x08, 0x8c, 0xa1, 0x13, 0x6c, 0x91, 0xcc, 0xdf, 0x91, 0x70, 0xd1,
	0x87, 0xc8, 0x42, 0xec, 0x72, 0xa8, 0xc7, 0x16, 0xc8, 0x6c, 0x53, 0x49, 0x09, 0x91, 0xdd, 0xe3,
	0x22, 0x81, 0x98, 0x56, 0xd8, 0x12, 0xa1, 0x27, 0xa0, 0x53, 0x61, 0x8c, 0x50, 0x32, 0x00, 0x29,
	0x20, 0xa6, 0x3e, 0xbb, 0x44, 0x16, 0x9b, 0x2a, 0x49, 0x20, 0xb2, 0x42, 0xc9, 0x23, 0x65, 0x77,
	0x2f, 0x84, 0xb1, 0x86, 0x56, 0x91, 0xb6, 0x95, 0x24, 0xd0, 0xe5, 0xc9, 0xb6, 0xee, 0x66, 0x29,
	0x48, 0x4b, 0x27, 0x91, 0xa3, 0x00, 0x03, 0x91, 0x82, 0x44, 0x26, 0x5a, 0x2b, 0xa1, 0x2d, 0x19,
	0xc3, 0x05, 0xfa, 0x47, 0xa7, 0xd9, 0x65, 0xb2, 0x5c, 0xa0, 0xa5, 0x03, 0x78, 0x0a, 0xb4, 0xce,
	0xe6, 0xc9, 0x4c, 0xb1, 0x74, 0x7a, 0x7c, 0x72, 0x8b, 0x92, 0x12, 0x43, 0xa8, 0x1e, 0x84, 0x10,
	0x29, 0x1d, 0xd3, 0x99, 0x92, 0x84, 0xbb, 0x10, 0x59, 0xa5, 0x5b, 0x01, 0x6d, 0xa0, 0xe0, 0x02,
	0x6c, 0x03, 0xd7, 0x51, 0x2f, 0x04, 0x93, 0x25, 0x96, 0xce, 0x32, 0x4a, 0x1a, 0x7b, 0x22, 0x81,
	0x23, 0x65, 0xf7, 0x54, 0x26, 0x63, 0x3a, 0xc7, 0xe6, 0x08, 0x39, 0x04, 0xcb, 0x0b, 0x07, 0xe6,
	0xf1, 0xd8, 0x26, 0x8f, 0x7a, 0x50, 0x00, 0x94, 0xad, 0x10, 0xd6, 0xe4, 0x52, 0x2a, 0xdb, 0xd4,
	0xc0, 0x2d, 0xec, 0xa9, 0x24, 0x06, 0x4d, 0x17, 0x50, 0xce, 0x4b, 0xb8, 0x48, 0x80, 0xb2, 0x71,
	0x76, 0x00, 0x09, 0x8c, 0xb2, 0x17, 0xc7, 0xd9, 0x05, 0x8e, 0xd9, 0x4b, 0x28, 0x7e, 0x27, 0x13,
	0x49, 0xec, 0x2c, 0xc9, 0xdb, 0xb2, 0x8c, 0x1a, 0x0b, 0xf1, 0x47, 0xb7, 0x5b, 0xed, 0x53, 0xba,
	0xc2, 0x96, 0xc9, 0x42, 0x81, 0x1c, 0x82, 0xd5, 0x22, 0x72, 0xe6, 0x5d, 0x42, 0xa9, 0xc7, 0x99,
	0x3d, 0xee, 0x1c, 0x42, 0xaa, 0xf4, 0x80, 0xae, 0x62, 0x43, 0x1d, 0xd3, 0xb0, 0x45, 0xf4, 0x32,
	0x9e, 0xb0, 0x9b, 0xf6, 0xed, 0x60, 0x6c, 0x2f, 0xbd, 0xc2, 0x18, 0x99, 0x0d, 0x82, 0x10, 0xde,
	0xcd, 0xc0, 0xd8, 0x90, 0x47, 0x40, 0x7f, 0xa9, 0x6d, 0xbc, 0x45, 0x88, 0xdb, 0x8b, 0x77, 0x1f,
	0x18, 0x23, 0x73, 0xe3, 0xe8, 0x48, 0x49, 0xa0, 0x13, 0xac, 0x41, 0xa6, 0xef, 0x48, 0x61, 0x4c,
	0x06, 0x31, 0xf5, 0xd0, 0xb7, 0x96, 0x3c, 0xd1, 0xaa, 0x8b, 0x57, 0x8e, 0x56, 0x70, 0x75, 0x4f,
	0x48, 0x61, 0x7a, 0x6e, 0x62, 0x08, 0x99, 0x2a, 0x0c, 0xac, 0x6e, 0x74, 0x48, 0xa3, 0x0d, 0x5d,
	0x1c, 0x8e, 0x9c, 0x7b, 0x89, 0xd0, 0x72, 0x3c, 0x66, 0x1f, 0xc9, 0xf6, 0x70, 0x78, 0xf7, 0xb5,
	0x7a, 0x20, 0x64, 0x97, 0x56, 0x90, 0xac, 0x0d, 0x3c, 0x71, 0xc4, 0x33, 0xa4, 0xb6, 0x97, 0x64,
	0xee, 0x94, 0xaa, 0x3b, 0x13, 0x03, 0x4c, 0x9b, 0xdc, 0x78, 0x34, 0xed, 0xae, 0xb4, 0xbb, 0x99,
	0xb3, 0xa4, 0x7e, 0x47, 0xc6, 0xd0, 0x11, 0x12, 0x62, 0x3a, 0xe1, 0xdc, 0x77, 0x5d, 0x2a, 0xd9,
	0x10, 0x63, 0x91, 0x81, 0x56, 0xfd, 0x12, 0x06, 0x68, 0xe1, 0x01, 0x37, 0x25, 0xa8, 0x83, 0x2d,
	0x0d, 0xc0, 0x44, 0x5a, 0x9c, 0x95, 0xb7, 0x77, 0xd1, 0xda, 0x76, 0x4f, 0x3d, 0x18, 0x63, 0x86,
	0xf6, 0xf0, 0xa4, 0x7d, 0xb0, 0xed, 0x81, 0xb1, 0x90, 0x36, 0x95, 0xec, 0x88, 0xae, 0xa1, 0x02,
	0x4f, 0xba, 0xad, 0x78, 0x5c, 0xda, 0xfe, 0x0e, 0x36, 0x35, 0x84, 0x04, 0xb8, 0x29, 0xb3, 0xde,
	0x77, 0xf3, 0xe7, 0xa4, 0x6e, 0x27, 0x82, 0x1b, 0x9a, 0x60, 0x29, 0xa8, 0x32, 0x0f, 0x53, 0xf4,
	0x7d, 0x3b, 0xb1, 0xa0, 0xf3, 0x58, 0xb2, 0x25, 0x32, 0x9f, 0xe7, 0x9f, 0x70, 0x6d, 0x85, 0x23,
	0xf9, 0xc6, 0x73, 0x1d, 0xd6, 0xaa, 0x3f, 0xc6, 0xbe, 0xc5, 0xeb, 0xde, 0x38, 0xe0, 0x66, 0x0c,
	0x7d, 0xe7, 0xb1, 0x15, 0xb2, 0x30, 0x2c, 0x6d, 0x8c, 0x7f, 0xef, 0xb1, 0x45, 0x32, 0x87, 0xa5,
	0x8d, 0x30, 0x43, 0x7f, 0x70, 0x20, 0x16, 0x51, 0x02, 0x7f, 0x74, 0x0c, 0x45, 0x15, 0x25, 0xfc,
	0x27, 0x77, 0x18, 0x32, 0x14, 0x8d, 0x36, 0xf4, 0xb1, 0x87, 0x4a, 0x87, 0x87, 0x15, 0x30, 0x7d,
	0xe2, 0x12, 0x91, 0x75, 0x94, 0xf8, 0xd4, 0x25, 0x16, 0x9c, 0x23, 0xf4, 0x99, 0x43, 0x0f, 0xb8,
	0x8c, 0x55, 0xa7, 0x33, 0x42, 0x9f, 0x7b, 0x6c, 0x95, 0x2c, 0xe2, 0xf6, 0x1d, 0x9e, 0x70, 0x19,
	0x8d, 0xf3, 0x5f, 0x78, 0x8c, 0x0e, 0x8d, 0x74, 0x83, 0x4c, 0x3f, 0xa8, 0x38, 0x53, 0x0a, 0x01,
	0x39, 0xf6, 0x61, 0x85, 0xcd, 0xe5, 0xee, 0xe6, 0xf1, 0x47, 0x15, 0x36, 0x43, 0xa6, 0x5a, 0xd2,
	0x80, 0xb6, 0xf4, 0x3d, 0x1c, 0xb6, 0xa9, 0xfc, 0xba, 0xd2, 0xf7, 0x71, 0xa4, 0x27, 0xdd, 0xb0,
	0xd1, 0x87, 0x6e, 0x21, 0x7f, 0x58, 0xe8, 0xaf, 0xbe, 0x2b, 0xb5, 0xfc, 0xca, 0xfc, 0xe6, 0xe3,
	0x49, 0xfb, 0x60, 0xc7, 0x37, 0x88, 0xfe, 0xee, 0xb3, 0x2b, 0x64, 0x79, 0x88, 0xb9, 0x3b, 0x3f,
	0xba, 0x3b, 0x7f, 0xf8, 0x6c, 0x8d, 0x5c, 0xda, 0x07, 0x3b, 0x9e, 0x03, 0xdc, 0x24, 0x8c, 0x15,
	0x91, 0xa1, 0x7f, 0xfa, 0xec, 0x5f, 0x64, 0x65, 0x1f, 0xec, 0xc8, 0xdf, 0xd2, 0xe2, 0x5f, 0x3e,
	0x9b, 0x25, 0xd3, 0x21, 0x3e, 0x0a, 0x70, 0x0e, 0xf4, 0xb1, 0x8f, 0x4d, 0x1a, 0x86, 0x85, 0x9c,
	0x27, 0x3e, 0x5a, 0xf7, 0x26, 0xb7, 0x51, 0x2f, 0x48, 0x9b, 0x3d, 0x2e, 0x25, 0x24, 0x86, 0x3e,
	0xf5, 0xd9, 0x32, 0xa1, 0x21, 0xa4, 0xea, 0x1c, 0x4a, 0xf0, 0x33, 0x7c, 0xec, 0x99, 0x4b, 0x7e,
	0x23, 0x03, 0x3d, 0x18, 0x2d, 0x3c, 0xf7, 0xd1, 0xea, 0x3c, 0xff, 0xe5, 0x95, 0x17, 0x3e, 0x5a,
	0x5d, 0x38, 0xdf, 0x92, 0x1d, 0x45, 0x7f, 0xae, 0xa2, 0xaa, 0x53, 0x91, 0xc2, 0xa9, 0x88, 0xee,
	0xd3, 0x8f, 0xeb, 0xa8, 0xca, 0x6d, 0x3a, 0x52, 0x31, 0xa0, 0x7c, 0x43, 0x3f, 0xa9, 0xa3, 0xf5,
	0xd8, 0xba, 0xdc, 0xfa, 0x4f, 0x5d, 0x5c, 0xbc, 0x49, 0xad, 0x80, 0x7e, 0x86, 0x3f, 0x00, 0x52,
	0xc4, 0xa7, 0xed, 0x63, 0xfa, 0xa8, 0x8e, 0x65, 0x6c, 0x27, 0x89, 0x8a, 0xb8, 0x1d, 0x0d, 0xd0,
	0xe7, 0x75, 0x9c, 0xc0, 0xd2, 0x73, 0x52, 0x18, 0xf3, 0x45, 0x1d, 0xcb, 0x2b, 0x70, 0xd7, 0xb6,
	0x00, 0x9f, 0x99, 0x2f, 0x1d, 0x6b, 0xc0, 0x2d, 0x47, 0x25, 0xa7, 0x96, 0x7e, 0x55, 0xdf, 0x58,
	0x27, 0xb5, 0xc0, 0x24, 0xee, 0xd5, 0xa8, 0x11, 0x3f, 0x30, 0x09, 0x9d, 0xc0, 0x4b, 0xb6, 0xa3,
	0x54, 0xb2, 0x7b, 0xd1, 0xd7, 0x77, 0xff, 0x4b, 0xbd, 0x9d, 0xff, 0xbf, 0x7d, 0xa3, 0x2b, 0x6c,
	0x2f, 0x3b, 0xc3, 0x1f, 0xef, 0x56, 0xfe, 0x27, 0xbe, 0x2e, 0x54, 0xf1, 0xb5, 0x25, 0xa4, 0x05,
	0x2d, 0x79, 0xb2, 0xe5, 0x7e, 0xce, 0x5b, 0xf9, 0xcf, 0xb9, 0x7f, 0x76, 0x36, 0xe5, 0xe2, 0x1b,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x10, 0x52, 0xd9, 0x76, 0x09, 0x00, 0x00,
}
