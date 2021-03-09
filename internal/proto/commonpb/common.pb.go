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
	ErrorCode_ERROR_CODE_SUCCESS                 ErrorCode = 0
	ErrorCode_ERROR_CODE_UNEXPECTED_ERROR        ErrorCode = 1
	ErrorCode_ERROR_CODE_CONNECT_FAILED          ErrorCode = 2
	ErrorCode_ERROR_CODE_PERMISSION_DENIED       ErrorCode = 3
	ErrorCode_ERROR_CODE_COLLECTION_NOT_EXISTS   ErrorCode = 4
	ErrorCode_ERROR_CODE_ILLEGAL_ARGUMENT        ErrorCode = 5
	ErrorCode_ERROR_CODE_ILLEGAL_DIMENSION       ErrorCode = 7
	ErrorCode_ERROR_CODE_ILLEGAL_INDEX_TYPE      ErrorCode = 8
	ErrorCode_ERROR_CODE_ILLEGAL_COLLECTION_NAME ErrorCode = 9
	ErrorCode_ERROR_CODE_ILLEGAL_TOPK            ErrorCode = 10
	ErrorCode_ERROR_CODE_ILLEGAL_ROWRECORD       ErrorCode = 11
	ErrorCode_ERROR_CODE_ILLEGAL_VECTOR_ID       ErrorCode = 12
	ErrorCode_ERROR_CODE_ILLEGAL_SEARCH_RESULT   ErrorCode = 13
	ErrorCode_ERROR_CODE_FILE_NOT_FOUND          ErrorCode = 14
	ErrorCode_ERROR_CODE_META_FAILED             ErrorCode = 15
	ErrorCode_ERROR_CODE_CACHE_FAILED            ErrorCode = 16
	ErrorCode_ERROR_CODE_CANNOT_CREATE_FOLDER    ErrorCode = 17
	ErrorCode_ERROR_CODE_CANNOT_CREATE_FILE      ErrorCode = 18
	ErrorCode_ERROR_CODE_CANNOT_DELETE_FOLDER    ErrorCode = 19
	ErrorCode_ERROR_CODE_CANNOT_DELETE_FILE      ErrorCode = 20
	ErrorCode_ERROR_CODE_BUILD_INDEX_ERROR       ErrorCode = 21
	ErrorCode_ERROR_CODE_ILLEGAL_NLIST           ErrorCode = 22
	ErrorCode_ERROR_CODE_ILLEGAL_METRIC_TYPE     ErrorCode = 23
	ErrorCode_ERROR_CODE_OUT_OF_MEMORY           ErrorCode = 24
	ErrorCode_ERROR_CODE_INDEX_NOT_EXIST         ErrorCode = 25
	// internal error code.
	ErrorCode_ERROR_CODE_DD_REQUEST_RACE ErrorCode = 1000
)

var ErrorCode_name = map[int32]string{
	0:    "ERROR_CODE_SUCCESS",
	1:    "ERROR_CODE_UNEXPECTED_ERROR",
	2:    "ERROR_CODE_CONNECT_FAILED",
	3:    "ERROR_CODE_PERMISSION_DENIED",
	4:    "ERROR_CODE_COLLECTION_NOT_EXISTS",
	5:    "ERROR_CODE_ILLEGAL_ARGUMENT",
	7:    "ERROR_CODE_ILLEGAL_DIMENSION",
	8:    "ERROR_CODE_ILLEGAL_INDEX_TYPE",
	9:    "ERROR_CODE_ILLEGAL_COLLECTION_NAME",
	10:   "ERROR_CODE_ILLEGAL_TOPK",
	11:   "ERROR_CODE_ILLEGAL_ROWRECORD",
	12:   "ERROR_CODE_ILLEGAL_VECTOR_ID",
	13:   "ERROR_CODE_ILLEGAL_SEARCH_RESULT",
	14:   "ERROR_CODE_FILE_NOT_FOUND",
	15:   "ERROR_CODE_META_FAILED",
	16:   "ERROR_CODE_CACHE_FAILED",
	17:   "ERROR_CODE_CANNOT_CREATE_FOLDER",
	18:   "ERROR_CODE_CANNOT_CREATE_FILE",
	19:   "ERROR_CODE_CANNOT_DELETE_FOLDER",
	20:   "ERROR_CODE_CANNOT_DELETE_FILE",
	21:   "ERROR_CODE_BUILD_INDEX_ERROR",
	22:   "ERROR_CODE_ILLEGAL_NLIST",
	23:   "ERROR_CODE_ILLEGAL_METRIC_TYPE",
	24:   "ERROR_CODE_OUT_OF_MEMORY",
	25:   "ERROR_CODE_INDEX_NOT_EXIST",
	1000: "ERROR_CODE_DD_REQUEST_RACE",
}

var ErrorCode_value = map[string]int32{
	"ERROR_CODE_SUCCESS":                 0,
	"ERROR_CODE_UNEXPECTED_ERROR":        1,
	"ERROR_CODE_CONNECT_FAILED":          2,
	"ERROR_CODE_PERMISSION_DENIED":       3,
	"ERROR_CODE_COLLECTION_NOT_EXISTS":   4,
	"ERROR_CODE_ILLEGAL_ARGUMENT":        5,
	"ERROR_CODE_ILLEGAL_DIMENSION":       7,
	"ERROR_CODE_ILLEGAL_INDEX_TYPE":      8,
	"ERROR_CODE_ILLEGAL_COLLECTION_NAME": 9,
	"ERROR_CODE_ILLEGAL_TOPK":            10,
	"ERROR_CODE_ILLEGAL_ROWRECORD":       11,
	"ERROR_CODE_ILLEGAL_VECTOR_ID":       12,
	"ERROR_CODE_ILLEGAL_SEARCH_RESULT":   13,
	"ERROR_CODE_FILE_NOT_FOUND":          14,
	"ERROR_CODE_META_FAILED":             15,
	"ERROR_CODE_CACHE_FAILED":            16,
	"ERROR_CODE_CANNOT_CREATE_FOLDER":    17,
	"ERROR_CODE_CANNOT_CREATE_FILE":      18,
	"ERROR_CODE_CANNOT_DELETE_FOLDER":    19,
	"ERROR_CODE_CANNOT_DELETE_FILE":      20,
	"ERROR_CODE_BUILD_INDEX_ERROR":       21,
	"ERROR_CODE_ILLEGAL_NLIST":           22,
	"ERROR_CODE_ILLEGAL_METRIC_TYPE":     23,
	"ERROR_CODE_OUT_OF_MEMORY":           24,
	"ERROR_CODE_INDEX_NOT_EXIST":         25,
	"ERROR_CODE_DD_REQUEST_RACE":         1000,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type IndexState int32

const (
	IndexState_INDEX_STATE_NONE        IndexState = 0
	IndexState_INDEX_STATE_UNISSUED    IndexState = 1
	IndexState_INDEX_STATE_IN_PROGRESS IndexState = 2
	IndexState_INDEX_STATE_FINISHED    IndexState = 3
	IndexState_INDEX_STATE_FAILED      IndexState = 4
	IndexState_INDEX_STATE_DELETED     IndexState = 5
)

var IndexState_name = map[int32]string{
	0: "INDEX_STATE_NONE",
	1: "INDEX_STATE_UNISSUED",
	2: "INDEX_STATE_IN_PROGRESS",
	3: "INDEX_STATE_FINISHED",
	4: "INDEX_STATE_FAILED",
	5: "INDEX_STATE_DELETED",
}

var IndexState_value = map[string]int32{
	"INDEX_STATE_NONE":        0,
	"INDEX_STATE_UNISSUED":    1,
	"INDEX_STATE_IN_PROGRESS": 2,
	"INDEX_STATE_FINISHED":    3,
	"INDEX_STATE_FAILED":      4,
	"INDEX_STATE_DELETED":     5,
}

func (x IndexState) String() string {
	return proto.EnumName(IndexState_name, int32(x))
}

func (IndexState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type SegmentState int32

const (
	SegmentState_SegmentNone     SegmentState = 0
	SegmentState_SegmentNotExist SegmentState = 1
	SegmentState_SegmentGrowing  SegmentState = 2
	SegmentState_SegmentSealed   SegmentState = 3
	SegmentState_SegmentFlushed  SegmentState = 4
)

var SegmentState_name = map[int32]string{
	0: "SegmentNone",
	1: "SegmentNotExist",
	2: "SegmentGrowing",
	3: "SegmentSealed",
	4: "SegmentFlushed",
}

var SegmentState_value = map[string]int32{
	"SegmentNone":     0,
	"SegmentNotExist": 1,
	"SegmentGrowing":  2,
	"SegmentSealed":   3,
	"SegmentFlushed":  4,
}

func (x SegmentState) String() string {
	return proto.EnumName(SegmentState_name, int32(x))
}

func (SegmentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type MsgType int32

const (
	MsgType_kNone MsgType = 0
	// Definition Requests: collection
	MsgType_kCreateCollection   MsgType = 100
	MsgType_kDropCollection     MsgType = 101
	MsgType_kHasCollection      MsgType = 102
	MsgType_kDescribeCollection MsgType = 103
	MsgType_kShowCollections    MsgType = 104
	MsgType_kGetSysConfigs      MsgType = 105
	MsgType_kLoadCollection     MsgType = 106
	MsgType_kReleaseCollection  MsgType = 107
	// Definition Requests: partition
	MsgType_kCreatePartition   MsgType = 200
	MsgType_kDropPartition     MsgType = 201
	MsgType_kHasPartition      MsgType = 202
	MsgType_kDescribePartition MsgType = 203
	MsgType_kShowPartitions    MsgType = 204
	MsgType_kLoadPartition     MsgType = 205
	MsgType_kReleasePartition  MsgType = 206
	// Define Requests: segment
	MsgType_kShowSegment     MsgType = 250
	MsgType_kDescribeSegment MsgType = 251
	// Definition Requests: Index
	MsgType_kCreateIndex   MsgType = 300
	MsgType_kDescribeIndex MsgType = 301
	MsgType_kDropIndex     MsgType = 302
	// Manipulation Requests
	MsgType_kInsert MsgType = 400
	MsgType_kDelete MsgType = 401
	MsgType_kFlush  MsgType = 402
	// Query
	MsgType_kSearch                  MsgType = 500
	MsgType_kSearchResult            MsgType = 501
	MsgType_kGetIndexState           MsgType = 502
	MsgType_kGetCollectionStatistics MsgType = 503
	MsgType_kGetPartitionStatistics  MsgType = 504
	// Data Service
	MsgType_kSegmentInfo MsgType = 600
	// System Control
	MsgType_kTimeTick          MsgType = 1200
	MsgType_kQueryNodeStats    MsgType = 1201
	MsgType_kLoadIndex         MsgType = 1202
	MsgType_kRequestID         MsgType = 1203
	MsgType_kRequestTSO        MsgType = 1204
	MsgType_kAllocateSegment   MsgType = 1205
	MsgType_kSegmentStatistics MsgType = 1206
	MsgType_kSegmentFlushDone  MsgType = 1207
)

var MsgType_name = map[int32]string{
	0:    "kNone",
	100:  "kCreateCollection",
	101:  "kDropCollection",
	102:  "kHasCollection",
	103:  "kDescribeCollection",
	104:  "kShowCollections",
	105:  "kGetSysConfigs",
	106:  "kLoadCollection",
	107:  "kReleaseCollection",
	200:  "kCreatePartition",
	201:  "kDropPartition",
	202:  "kHasPartition",
	203:  "kDescribePartition",
	204:  "kShowPartitions",
	205:  "kLoadPartition",
	206:  "kReleasePartition",
	250:  "kShowSegment",
	251:  "kDescribeSegment",
	300:  "kCreateIndex",
	301:  "kDescribeIndex",
	302:  "kDropIndex",
	400:  "kInsert",
	401:  "kDelete",
	402:  "kFlush",
	500:  "kSearch",
	501:  "kSearchResult",
	502:  "kGetIndexState",
	503:  "kGetCollectionStatistics",
	504:  "kGetPartitionStatistics",
	600:  "kSegmentInfo",
	1200: "kTimeTick",
	1201: "kQueryNodeStats",
	1202: "kLoadIndex",
	1203: "kRequestID",
	1204: "kRequestTSO",
	1205: "kAllocateSegment",
	1206: "kSegmentStatistics",
	1207: "kSegmentFlushDone",
}

var MsgType_value = map[string]int32{
	"kNone":                    0,
	"kCreateCollection":        100,
	"kDropCollection":          101,
	"kHasCollection":           102,
	"kDescribeCollection":      103,
	"kShowCollections":         104,
	"kGetSysConfigs":           105,
	"kLoadCollection":          106,
	"kReleaseCollection":       107,
	"kCreatePartition":         200,
	"kDropPartition":           201,
	"kHasPartition":            202,
	"kDescribePartition":       203,
	"kShowPartitions":          204,
	"kLoadPartition":           205,
	"kReleasePartition":        206,
	"kShowSegment":             250,
	"kDescribeSegment":         251,
	"kCreateIndex":             300,
	"kDescribeIndex":           301,
	"kDropIndex":               302,
	"kInsert":                  400,
	"kDelete":                  401,
	"kFlush":                   402,
	"kSearch":                  500,
	"kSearchResult":            501,
	"kGetIndexState":           502,
	"kGetCollectionStatistics": 503,
	"kGetPartitionStatistics":  504,
	"kSegmentInfo":             600,
	"kTimeTick":                1200,
	"kQueryNodeStats":          1201,
	"kLoadIndex":               1202,
	"kRequestID":               1203,
	"kRequestTSO":              1204,
	"kAllocateSegment":         1205,
	"kSegmentStatistics":       1206,
	"kSegmentFlushDone":        1207,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

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
	return fileDescriptor_555bd8c177793206, []int{1}
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
	return ErrorCode_ERROR_CODE_SUCCESS
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
	return fileDescriptor_555bd8c177793206, []int{2}
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
	return MsgType_kNone
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

func init() {
	proto.RegisterEnum("milvus.proto.common.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("milvus.proto.common.IndexState", IndexState_name, IndexState_value)
	proto.RegisterEnum("milvus.proto.common.SegmentState", SegmentState_name, SegmentState_value)
	proto.RegisterEnum("milvus.proto.common.MsgType", MsgType_name, MsgType_value)
	proto.RegisterType((*Empty)(nil), "milvus.proto.common.Empty")
	proto.RegisterType((*Status)(nil), "milvus.proto.common.Status")
	proto.RegisterType((*KeyValuePair)(nil), "milvus.proto.common.KeyValuePair")
	proto.RegisterType((*Blob)(nil), "milvus.proto.common.Blob")
	proto.RegisterType((*Address)(nil), "milvus.proto.common.Address")
	proto.RegisterType((*MsgBase)(nil), "milvus.proto.common.MsgBase")
	proto.RegisterType((*MsgHeader)(nil), "milvus.proto.common.MsgHeader")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xcb, 0x6e, 0xdb, 0x46,
	0x1b, 0x0d, 0x25, 0xd9, 0x0a, 0x3f, 0x2b, 0xf6, 0x78, 0xec, 0xd8, 0x4a, 0xe2, 0x24, 0xfe, 0xf5,
	0x17, 0x45, 0x60, 0x20, 0x76, 0xd1, 0x02, 0xed, 0x2a, 0x40, 0x65, 0x72, 0x64, 0x13, 0xa1, 0x48,
	0x65, 0x48, 0xa5, 0x49, 0x37, 0x04, 0x25, 0x4d, 0x64, 0x56, 0x94, 0xa8, 0x72, 0xa8, 0x24, 0xca,
	0x53, 0xb4, 0x01, 0xfa, 0x06, 0x5d, 0xb6, 0x45, 0xef, 0x7d, 0x85, 0xde, 0xd7, 0x7d, 0x84, 0x02,
	0xdd, 0xf6, 0x86, 0x76, 0x53, 0xcc, 0x50, 0x17, 0x5a, 0x50, 0x76, 0xfc, 0xce, 0xf9, 0xe6, 0xcc,
	0xf9, 0xce, 0x70, 0x48, 0x28, 0xb5, 0xa3, 0x7e, 0x3f, 0x1a, 0x1c, 0x0e, 0xe3, 0x28, 0x89, 0xf0,
	0x56, 0x3f, 0x08, 0x1f, 0x8f, 0x78, 0x5a, 0x1d, 0xa6, 0x54, 0xa5, 0x08, 0x2b, 0xa4, 0x3f, 0x4c,
	0xc6, 0x15, 0x0f, 0x56, 0x9d, 0xc4, 0x4f, 0x46, 0x1c, 0xdf, 0x01, 0x60, 0x71, 0x1c, 0xc5, 0x5e,
	0x3b, 0xea, 0xb0, 0xb2, 0xb2, 0xaf, 0xdc, 0x5a, 0x7f, 0xf5, 0xc6, 0xe1, 0x92, 0xc5, 0x87, 0x44,
	0xb4, 0x69, 0x51, 0x87, 0x51, 0x95, 0x4d, 0x1f, 0xf1, 0x0e, 0xac, 0xc6, 0xcc, 0xe7, 0xd1, 0xa0,
	0x9c, 0xdb, 0x57, 0x6e, 0xa9, 0x74, 0x52, 0x55, 0x5e, 0x87, 0xd2, 0x5d, 0x36, 0xbe, 0xef, 0x87,
	0x23, 0xd6, 0xf0, 0x83, 0x18, 0x23, 0xc8, 0xf7, 0xd8, 0x58, 0xea, 0xab, 0x54, 0x3c, 0xe2, 0x6d,
	0x58, 0x79, 0x2c, 0xe8, 0xc9, 0xc2, 0xb4, 0xa8, 0xec, 0x41, 0xe1, 0x38, 0x8c, 0x5a, 0x73, 0x56,
	0xac, 0x28, 0x4d, 0xd9, 0xdb, 0x50, 0xac, 0x76, 0x3a, 0x31, 0xe3, 0x1c, 0xaf, 0x43, 0x2e, 0x18,
	0x4e, 0xf4, 0x72, 0xc1, 0x10, 0x63, 0x28, 0x0c, 0xa3, 0x38, 0x91, 0x6a, 0x79, 0x2a, 0x9f, 0x2b,
	0xcf, 0x15, 0x28, 0xd6, 0x79, 0xf7, 0xd8, 0xe7, 0x0c, 0xbf, 0x01, 0x17, 0xfb, 0xbc, 0xeb, 0x25,
	0xe3, 0xe1, 0x74, 0xca, 0xbd, 0xa5, 0x53, 0xd6, 0x79, 0xd7, 0x1d, 0x0f, 0x19, 0x2d, 0xf6, 0xd3,
	0x07, 0xe1, 0xa4, 0xcf, 0xbb, 0x86, 0x3e, 0x51, 0x4e, 0x0b, 0xbc, 0x07, 0x6a, 0x12, 0xf4, 0x19,
	0x4f, 0xfc, 0xfe, 0xb0, 0x9c, 0xdf, 0x57, 0x6e, 0x15, 0xe8, 0x1c, 0xc0, 0x57, 0xe1, 0x22, 0x8f,
	0x46, 0x71, 0x9b, 0x19, 0x7a, 0xb9, 0x20, 0x97, 0xcd, 0xea, 0xca, 0x1d, 0x50, 0xeb, 0xbc, 0x7b,
	0xca, 0xfc, 0x0e, 0x8b, 0xf1, 0x2b, 0x50, 0x68, 0xf9, 0x3c, 0x75, 0xb4, 0xf6, 0x62, 0x47, 0x62,
	0x02, 0x2a, 0x3b, 0x0f, 0x7e, 0x5b, 0x05, 0x95, 0x64, 0xe2, 0xc7, 0x84, 0x52, 0x9b, 0x7a, 0x9a,
	0xad, 0x13, 0xcf, 0x69, 0x6a, 0x1a, 0x71, 0x1c, 0x74, 0x01, 0xdf, 0x84, 0x6b, 0x19, 0xbc, 0x69,
	0x91, 0x07, 0x0d, 0xa2, 0xb9, 0x44, 0xf7, 0x24, 0x8a, 0x14, 0x7c, 0x1d, 0xae, 0x64, 0x1a, 0x34,
	0xdb, 0xb2, 0x88, 0xe6, 0x7a, 0xb5, 0xaa, 0x61, 0x12, 0x1d, 0xe5, 0xf0, 0x3e, 0xec, 0x65, 0xe8,
	0x06, 0xa1, 0x75, 0xc3, 0x71, 0x0c, 0xdb, 0xf2, 0x74, 0x62, 0x19, 0x44, 0x47, 0x79, 0xfc, 0x12,
	0xec, 0x9f, 0x13, 0x30, 0x4d, 0xa2, 0xb9, 0xa2, 0xc3, 0xb2, 0x5d, 0x8f, 0x3c, 0x30, 0x1c, 0xd7,
	0x41, 0x85, 0x05, 0x1f, 0x86, 0x69, 0x92, 0x93, 0xaa, 0xe9, 0x55, 0xe9, 0x49, 0xb3, 0x4e, 0x2c,
	0x17, 0xad, 0x2c, 0x6c, 0x34, 0x6d, 0xd0, 0x8d, 0x3a, 0xb1, 0xc4, 0x7e, 0xa8, 0x88, 0xff, 0x07,
	0xd7, 0x97, 0x74, 0x18, 0x96, 0x4e, 0x1e, 0x78, 0xee, 0xc3, 0x06, 0x41, 0x17, 0xf1, 0xcb, 0x50,
	0x59, 0xd2, 0x92, 0xf5, 0x54, 0xad, 0x13, 0xa4, 0xe2, 0x6b, 0xb0, 0xbb, 0xa4, 0xcf, 0xb5, 0x1b,
	0x77, 0x11, 0xbc, 0xc0, 0x09, 0xb5, 0xdf, 0xa2, 0x44, 0xb3, 0xa9, 0x8e, 0xd6, 0x5e, 0xd0, 0x71,
	0x9f, 0x68, 0xae, 0x4d, 0x3d, 0x43, 0x47, 0xa5, 0x85, 0x50, 0xa6, 0x1d, 0x0e, 0xa9, 0x52, 0xed,
	0xd4, 0xa3, 0xc4, 0x69, 0x9a, 0x2e, 0xba, 0xb4, 0x90, 0x7d, 0xcd, 0x30, 0x89, 0x0c, 0xad, 0x66,
	0x37, 0x2d, 0x1d, 0xad, 0xe3, 0xab, 0xb0, 0x93, 0xa1, 0xeb, 0xc4, 0xad, 0x4e, 0xcf, 0x65, 0x63,
	0x61, 0x02, 0xad, 0xaa, 0x9d, 0x92, 0x29, 0x89, 0xf0, 0xff, 0xe1, 0xe6, 0x39, 0xd2, 0x12, 0xaa,
	0x1a, 0x25, 0x55, 0x97, 0x78, 0x35, 0xdb, 0xd4, 0x09, 0x45, 0x9b, 0x0b, 0x71, 0x2e, 0x34, 0x19,
	0x26, 0x41, 0x78, 0xb9, 0x8e, 0x4e, 0x4c, 0x32, 0xd7, 0xd9, 0x5a, 0xae, 0x33, 0x6d, 0x12, 0x3a,
	0xdb, 0x0b, 0x79, 0x1d, 0x37, 0x0d, 0x53, 0x9f, 0x9c, 0x5b, 0xfa, 0x16, 0x5e, 0xc6, 0x7b, 0x50,
	0x5e, 0x92, 0x97, 0x65, 0x1a, 0x8e, 0x8b, 0x76, 0x70, 0x05, 0x6e, 0x2c, 0x61, 0xeb, 0xc4, 0xa5,
	0x86, 0x96, 0x1e, 0xfd, 0xee, 0x82, 0x82, 0xdd, 0x74, 0x3d, 0xbb, 0xe6, 0xd5, 0x49, 0xdd, 0xa6,
	0x0f, 0x51, 0x19, 0xdf, 0x80, 0xab, 0x59, 0x05, 0xb9, 0xf7, 0xec, 0xfd, 0x44, 0x57, 0xf0, 0xcd,
	0x73, 0xbc, 0xae, 0x7b, 0x94, 0xdc, 0x6b, 0x12, 0xc7, 0xf5, 0x68, 0x55, 0x23, 0xe8, 0xd7, 0xe2,
	0xc1, 0x87, 0x0a, 0x80, 0x31, 0xe8, 0xb0, 0xa7, 0xe2, 0x6b, 0x29, 0xbe, 0x05, 0x28, 0x15, 0x71,
	0x5c, 0x91, 0x97, 0x65, 0x5b, 0x04, 0x5d, 0xc0, 0x65, 0xd8, 0xce, 0xa2, 0x4d, 0xcb, 0x70, 0x9c,
	0x26, 0xd1, 0x91, 0x22, 0x8e, 0x2b, 0xcb, 0x18, 0x96, 0xd7, 0xa0, 0xf6, 0x09, 0x15, 0x77, 0x34,
	0xb7, 0xb8, 0xac, 0x66, 0x58, 0x86, 0x73, 0x2a, 0xef, 0xd6, 0x0e, 0xe0, 0x73, 0x4c, 0x7a, 0xc0,
	0x05, 0xbc, 0x0b, 0x5b, 0x59, 0x3c, 0x4d, 0x5b, 0x47, 0x2b, 0x07, 0x11, 0x94, 0x1c, 0xd6, 0xed,
	0xb3, 0x41, 0x92, 0xfa, 0xdc, 0x80, 0xb5, 0x49, 0x6d, 0x45, 0x03, 0x86, 0x2e, 0xe0, 0x2d, 0xd8,
	0x98, 0x01, 0x09, 0x79, 0x1a, 0xf0, 0x04, 0x29, 0x18, 0xc3, 0xfa, 0x04, 0x3c, 0x89, 0xa3, 0x27,
	0xc1, 0xa0, 0x8b, 0x72, 0x78, 0x13, 0x2e, 0x4d, 0x95, 0x98, 0x1f, 0xb2, 0x0e, 0xca, 0x67, 0xda,
	0x6a, 0xe1, 0x88, 0x9f, 0xb1, 0x0e, 0x2a, 0x1c, 0x7c, 0xb0, 0x2a, 0xbf, 0xac, 0xf2, 0x03, 0xa9,
	0xc2, 0x4a, 0x6f, 0xb2, 0xcd, 0x65, 0xd8, 0xec, 0x69, 0x31, 0xf3, 0x13, 0xa6, 0x45, 0x61, 0xc8,
	0xda, 0x49, 0x10, 0x0d, 0x50, 0x47, 0xec, 0xde, 0xd3, 0xe3, 0x68, 0x98, 0x01, 0x99, 0x90, 0xed,
	0x9d, 0xfa, 0x3c, 0x83, 0x3d, 0x12, 0x03, 0xf6, 0x74, 0xc6, 0xdb, 0x71, 0xd0, 0xca, 0x2a, 0x74,
	0x45, 0xf0, 0x3d, 0xe7, 0x2c, 0x7a, 0x32, 0x07, 0x39, 0x3a, 0x93, 0x12, 0x27, 0x2c, 0x71, 0xc6,
	0x5c, 0x8b, 0x06, 0x8f, 0x82, 0x2e, 0x47, 0x81, 0xdc, 0xcb, 0x8c, 0xfc, 0x4e, 0x66, 0xf9, 0x3b,
	0x22, 0xd0, 0x1e, 0x65, 0x21, 0xf3, 0x79, 0x56, 0xb6, 0x87, 0x2f, 0x03, 0x9a, 0xf8, 0x6d, 0xf8,
	0x71, 0x12, 0x48, 0xf4, 0x5b, 0x05, 0x6f, 0xc1, 0xba, 0xf4, 0x3b, 0x07, 0xbf, 0x13, 0x69, 0x5d,
	0x12, 0x7e, 0xe7, 0xd8, 0xf7, 0x0a, 0xde, 0x05, 0x3c, 0xf3, 0x3b, 0x27, 0x7e, 0x50, 0xf0, 0x36,
	0x6c, 0x48, 0xbf, 0x33, 0x90, 0xa3, 0x1f, 0x53, 0x5d, 0xe1, 0x6d, 0xde, 0xfa, 0x93, 0x82, 0x77,
	0x60, 0x73, 0xea, 0x6d, 0x8e, 0xff, 0xac, 0xe0, 0x4d, 0x28, 0x49, 0x89, 0x49, 0xf6, 0xe8, 0x1f,
	0x45, 0xda, 0x9d, 0x6e, 0x37, 0x85, 0xff, 0x4d, 0x3b, 0xd3, 0x29, 0xe4, 0xab, 0x8a, 0x3e, 0xca,
	0xa5, 0x13, 0x4c, 0x3a, 0x53, 0xf0, 0xe3, 0x1c, 0xde, 0x00, 0x90, 0x63, 0xa5, 0xc0, 0x27, 0x39,
	0x5c, 0x82, 0x62, 0xcf, 0x18, 0x70, 0x16, 0x27, 0xe8, 0xbd, 0xbc, 0xac, 0x74, 0x16, 0xb2, 0x84,
	0xa1, 0xf7, 0xf3, 0x78, 0x0d, 0x56, 0x7b, 0xf2, 0xbc, 0xd1, 0xf3, 0x94, 0x72, 0x98, 0x1f, 0xb7,
	0xcf, 0xd0, 0xef, 0x79, 0x99, 0x44, 0x5a, 0x51, 0xc6, 0x47, 0x61, 0x82, 0xfe, 0xc8, 0xcb, 0x0d,
	0x4f, 0x58, 0x32, 0xbf, 0x2b, 0xe8, 0xcf, 0x3c, 0xbe, 0x0e, 0x65, 0x01, 0xce, 0x23, 0x17, 0x4c,
	0xc0, 0x93, 0xa0, 0xcd, 0xd1, 0x5f, 0x79, 0xbc, 0x07, 0xbb, 0x82, 0x9e, 0x4d, 0x9d, 0x61, 0xff,
	0xce, 0xa7, 0xf3, 0xa7, 0x43, 0x1a, 0x83, 0x47, 0x11, 0xfa, 0xa5, 0x80, 0xd7, 0x41, 0xed, 0xb9,
	0x41, 0x9f, 0xb9, 0x41, 0xbb, 0x87, 0x3e, 0x55, 0x65, 0xca, 0xf7, 0x46, 0x2c, 0x1e, 0x5b, 0x51,
	0x87, 0x89, 0xd5, 0x1c, 0x7d, 0xa6, 0xca, 0x31, 0x45, 0xca, 0xe9, 0x98, 0x9f, 0xa7, 0x00, 0x65,
	0xef, 0x8e, 0x18, 0x4f, 0x0c, 0x1d, 0x7d, 0xa1, 0x62, 0x04, 0x6b, 0x53, 0xc0, 0x75, 0x6c, 0xf4,
	0xa5, 0x2a, 0x93, 0xad, 0x86, 0x61, 0xd4, 0xf6, 0x93, 0x59, 0xb2, 0x5f, 0xa9, 0xf2, 0x7c, 0x33,
	0x17, 0x6b, 0x62, 0xee, 0x6b, 0x55, 0x1e, 0x5a, 0xf6, 0x52, 0xe8, 0xe2, 0xfd, 0xff, 0x46, 0x3d,
	0x3e, 0x7e, 0xfb, 0xcd, 0x6e, 0x90, 0x9c, 0x8d, 0x5a, 0xe2, 0xe7, 0x7d, 0xf4, 0x2c, 0x08, 0xc3,
	0xe0, 0x59, 0xc2, 0xda, 0x67, 0x47, 0xe9, 0x8f, 0xfd, 0x76, 0x27, 0xe0, 0x49, 0x1c, 0xb4, 0x46,
	0x09, 0xeb, 0x1c, 0x05, 0x83, 0x84, 0xc5, 0x03, 0x3f, 0x3c, 0x92, 0x7f, 0xfb, 0xa3, 0xf4, 0x6f,
	0x3f, 0x6c, 0xb5, 0x56, 0x65, 0xfd, 0xda, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x2b, 0xf4,
	0x53, 0xd0, 0x09, 0x00, 0x00,
}
