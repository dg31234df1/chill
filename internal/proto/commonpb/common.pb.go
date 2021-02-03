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
	ErrorCode_SUCCESS                 ErrorCode = 0
	ErrorCode_UNEXPECTED_ERROR        ErrorCode = 1
	ErrorCode_CONNECT_FAILED          ErrorCode = 2
	ErrorCode_PERMISSION_DENIED       ErrorCode = 3
	ErrorCode_COLLECTION_NOT_EXISTS   ErrorCode = 4
	ErrorCode_ILLEGAL_ARGUMENT        ErrorCode = 5
	ErrorCode_ILLEGAL_DIMENSION       ErrorCode = 7
	ErrorCode_ILLEGAL_INDEX_TYPE      ErrorCode = 8
	ErrorCode_ILLEGAL_COLLECTION_NAME ErrorCode = 9
	ErrorCode_ILLEGAL_TOPK            ErrorCode = 10
	ErrorCode_ILLEGAL_ROWRECORD       ErrorCode = 11
	ErrorCode_ILLEGAL_VECTOR_ID       ErrorCode = 12
	ErrorCode_ILLEGAL_SEARCH_RESULT   ErrorCode = 13
	ErrorCode_FILE_NOT_FOUND          ErrorCode = 14
	ErrorCode_META_FAILED             ErrorCode = 15
	ErrorCode_CACHE_FAILED            ErrorCode = 16
	ErrorCode_CANNOT_CREATE_FOLDER    ErrorCode = 17
	ErrorCode_CANNOT_CREATE_FILE      ErrorCode = 18
	ErrorCode_CANNOT_DELETE_FOLDER    ErrorCode = 19
	ErrorCode_CANNOT_DELETE_FILE      ErrorCode = 20
	ErrorCode_BUILD_INDEX_ERROR       ErrorCode = 21
	ErrorCode_ILLEGAL_NLIST           ErrorCode = 22
	ErrorCode_ILLEGAL_METRIC_TYPE     ErrorCode = 23
	ErrorCode_OUT_OF_MEMORY           ErrorCode = 24
	// internal error code.
	ErrorCode_DD_REQUEST_RACE ErrorCode = 1000
)

var ErrorCode_name = map[int32]string{
	0:    "SUCCESS",
	1:    "UNEXPECTED_ERROR",
	2:    "CONNECT_FAILED",
	3:    "PERMISSION_DENIED",
	4:    "COLLECTION_NOT_EXISTS",
	5:    "ILLEGAL_ARGUMENT",
	7:    "ILLEGAL_DIMENSION",
	8:    "ILLEGAL_INDEX_TYPE",
	9:    "ILLEGAL_COLLECTION_NAME",
	10:   "ILLEGAL_TOPK",
	11:   "ILLEGAL_ROWRECORD",
	12:   "ILLEGAL_VECTOR_ID",
	13:   "ILLEGAL_SEARCH_RESULT",
	14:   "FILE_NOT_FOUND",
	15:   "META_FAILED",
	16:   "CACHE_FAILED",
	17:   "CANNOT_CREATE_FOLDER",
	18:   "CANNOT_CREATE_FILE",
	19:   "CANNOT_DELETE_FOLDER",
	20:   "CANNOT_DELETE_FILE",
	21:   "BUILD_INDEX_ERROR",
	22:   "ILLEGAL_NLIST",
	23:   "ILLEGAL_METRIC_TYPE",
	24:   "OUT_OF_MEMORY",
	1000: "DD_REQUEST_RACE",
}

var ErrorCode_value = map[string]int32{
	"SUCCESS":                 0,
	"UNEXPECTED_ERROR":        1,
	"CONNECT_FAILED":          2,
	"PERMISSION_DENIED":       3,
	"COLLECTION_NOT_EXISTS":   4,
	"ILLEGAL_ARGUMENT":        5,
	"ILLEGAL_DIMENSION":       7,
	"ILLEGAL_INDEX_TYPE":      8,
	"ILLEGAL_COLLECTION_NAME": 9,
	"ILLEGAL_TOPK":            10,
	"ILLEGAL_ROWRECORD":       11,
	"ILLEGAL_VECTOR_ID":       12,
	"ILLEGAL_SEARCH_RESULT":   13,
	"FILE_NOT_FOUND":          14,
	"META_FAILED":             15,
	"CACHE_FAILED":            16,
	"CANNOT_CREATE_FOLDER":    17,
	"CANNOT_CREATE_FILE":      18,
	"CANNOT_DELETE_FOLDER":    19,
	"CANNOT_DELETE_FILE":      20,
	"BUILD_INDEX_ERROR":       21,
	"ILLEGAL_NLIST":           22,
	"ILLEGAL_METRIC_TYPE":     23,
	"OUT_OF_MEMORY":           24,
	"DD_REQUEST_RACE":         1000,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type IndexState int32

const (
	IndexState_NONE       IndexState = 0
	IndexState_UNISSUED   IndexState = 1
	IndexState_INPROGRESS IndexState = 2
	IndexState_FINISHED   IndexState = 3
	IndexState_FAILED     IndexState = 4
)

var IndexState_name = map[int32]string{
	0: "NONE",
	1: "UNISSUED",
	2: "INPROGRESS",
	3: "FINISHED",
	4: "FAILED",
}

var IndexState_value = map[string]int32{
	"NONE":       0,
	"UNISSUED":   1,
	"INPROGRESS": 2,
	"FINISHED":   3,
	"FAILED":     4,
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
	// Definition Requests: partition
	MsgType_kCreatePartition   MsgType = 200
	MsgType_kDropPartition     MsgType = 201
	MsgType_kHasPartition      MsgType = 202
	MsgType_kDescribePartition MsgType = 203
	MsgType_kShowPartitions    MsgType = 204
	// Define Requests: segment
	MsgType_kShowSegment     MsgType = 250
	MsgType_kDescribeSegment MsgType = 251
	// Definition Requests: Index
	MsgType_kCreateIndex   MsgType = 300
	MsgType_kDescribeIndex MsgType = 301
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
	200:  "kCreatePartition",
	201:  "kDropPartition",
	202:  "kHasPartition",
	203:  "kDescribePartition",
	204:  "kShowPartitions",
	250:  "kShowSegment",
	251:  "kDescribeSegment",
	300:  "kCreateIndex",
	301:  "kDescribeIndex",
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
	"kCreatePartition":         200,
	"kDropPartition":           201,
	"kHasPartition":            202,
	"kDescribePartition":       203,
	"kShowPartitions":          204,
	"kShowSegment":             250,
	"kDescribeSegment":         251,
	"kCreateIndex":             300,
	"kDescribeIndex":           301,
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
	return ErrorCode_SUCCESS
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
	// 1175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x8e, 0x2c, 0xd9, 0x32, 0x47, 0x8a, 0xbc, 0x5e, 0x9f, 0xf4, 0xff, 0x75, 0x8b, 0xc0, 0x57,
	0x81, 0x81, 0xd8, 0x45, 0x0b, 0xb4, 0x57, 0x01, 0x2a, 0x93, 0x2b, 0x9b, 0x08, 0x45, 0x2a, 0x4b,
	0x2a, 0x4d, 0x7a, 0x43, 0x50, 0xe2, 0x46, 0x26, 0x48, 0x89, 0x2a, 0x77, 0x95, 0x44, 0x79, 0x8a,
	0x36, 0x8f, 0x51, 0xb4, 0x40, 0xcf, 0xed, 0x23, 0xf4, 0xf4, 0x00, 0x7d, 0x84, 0x3e, 0x40, 0x4f,
	0x68, 0x6f, 0x8a, 0x5d, 0x92, 0x92, 0x50, 0xa4, 0x77, 0x9c, 0x6f, 0x76, 0xbe, 0xfd, 0xbe, 0x99,
	0xe1, 0x42, 0x73, 0x94, 0x4e, 0x26, 0xe9, 0xf4, 0x6c, 0x96, 0xa5, 0x22, 0xc5, 0x7b, 0x93, 0x28,
	0x79, 0x32, 0xe7, 0x79, 0x74, 0x96, 0xa7, 0x4e, 0xea, 0xb0, 0x49, 0x26, 0x33, 0xb1, 0x38, 0xf1,
	0x61, 0xcb, 0x15, 0x81, 0x98, 0x73, 0x7c, 0x17, 0x80, 0x65, 0x59, 0x9a, 0xf9, 0xa3, 0x34, 0x64,
	0xed, 0xca, 0xad, 0xca, 0xed, 0xd6, 0x1b, 0xaf, 0x9d, 0xbd, 0xa4, 0xf8, 0x8c, 0xc8, 0x63, 0x7a,
	0x1a, 0x32, 0xaa, 0xb1, 0xf2, 0x13, 0x1f, 0xc2, 0x56, 0xc6, 0x02, 0x9e, 0x4e, 0xdb, 0x1b, 0xb7,
	0x2a, 0xb7, 0x35, 0x5a, 0x44, 0x27, 0x6f, 0x41, 0xf3, 0x1e, 0x5b, 0x3c, 0x08, 0x92, 0x39, 0xeb,
	0x07, 0x51, 0x86, 0x11, 0x54, 0x63, 0xb6, 0x50, 0xfc, 0x1a, 0x95, 0x9f, 0x78, 0x1f, 0x36, 0x9f,
	0xc8, 0x74, 0x51, 0x98, 0x07, 0x27, 0xc7, 0x50, 0xbb, 0x48, 0xd2, 0xe1, 0x2a, 0x2b, 0x2b, 0x9a,
	0x65, 0xf6, 0x0e, 0xd4, 0x3b, 0x61, 0x98, 0x31, 0xce, 0x71, 0x0b, 0x36, 0xa2, 0x59, 0xc1, 0xb7,
	0x11, 0xcd, 0x30, 0x86, 0xda, 0x2c, 0xcd, 0x84, 0x62, 0xab, 0x52, 0xf5, 0x7d, 0xf2, 0xa2, 0x02,
	0xf5, 0x1e, 0x1f, 0x5f, 0x04, 0x9c, 0xe1, 0xb7, 0x61, 0x7b, 0xc2, 0xc7, 0xbe, 0x58, 0xcc, 0x4a,
	0x97, 0xc7, 0x2f, 0x75, 0xd9, 0xe3, 0x63, 0x6f, 0x31, 0x63, 0xb4, 0x3e, 0xc9, 0x3f, 0xa4, 0x92,
	0x09, 0x1f, 0x9b, 0x46, 0xc1, 0x9c, 0x07, 0xf8, 0x18, 0x34, 0x11, 0x4d, 0x18, 0x17, 0xc1, 0x64,
	0xd6, 0xae, 0xde, 0xaa, 0xdc, 0xae, 0xd1, 0x15, 0x80, 0xff, 0x0f, 0xdb, 0x3c, 0x9d, 0x67, 0x23,
	0x66, 0x1a, 0xed, 0x9a, 0x2a, 0x5b, 0xc6, 0x27, 0x77, 0x41, 0xeb, 0xf1, 0xf1, 0x15, 0x0b, 0x42,
	0x96, 0xe1, 0xd7, 0xa1, 0x36, 0x0c, 0x78, 0xae, 0xa8, 0xf1, 0xdf, 0x8a, 0xa4, 0x03, 0xaa, 0x4e,
	0x9e, 0x7e, 0x5b, 0x03, 0x6d, 0x39, 0x09, 0xdc, 0x80, 0xba, 0x3b, 0xd0, 0x75, 0xe2, 0xba, 0xe8,
	0x06, 0xde, 0x07, 0x34, 0xb0, 0xc9, 0xc3, 0x3e, 0xd1, 0x3d, 0x62, 0xf8, 0x84, 0x52, 0x87, 0xa2,
	0x0a, 0xc6, 0xd0, 0xd2, 0x1d, 0xdb, 0x26, 0xba, 0xe7, 0x77, 0x3b, 0xa6, 0x45, 0x0c, 0xb4, 0x81,
	0x0f, 0x60, 0xb7, 0x4f, 0x68, 0xcf, 0x74, 0x5d, 0xd3, 0xb1, 0x7d, 0x83, 0xd8, 0x26, 0x31, 0x50,
	0x15, 0xff, 0x0f, 0x0e, 0x74, 0xc7, 0xb2, 0x88, 0xee, 0x49, 0xd8, 0x76, 0x3c, 0x9f, 0x3c, 0x34,
	0x5d, 0xcf, 0x45, 0x35, 0xc9, 0x6d, 0x5a, 0x16, 0xb9, 0xec, 0x58, 0x7e, 0x87, 0x5e, 0x0e, 0x7a,
	0xc4, 0xf6, 0xd0, 0xa6, 0xe4, 0x29, 0x51, 0xc3, 0xec, 0x11, 0x5b, 0xd2, 0xa1, 0x3a, 0x3e, 0x04,
	0x5c, 0xc2, 0xa6, 0x6d, 0x90, 0x87, 0xbe, 0xf7, 0xa8, 0x4f, 0xd0, 0x36, 0x7e, 0x05, 0x8e, 0x4a,
	0x7c, 0xfd, 0x9e, 0x4e, 0x8f, 0x20, 0x0d, 0x23, 0x68, 0x96, 0x49, 0xcf, 0xe9, 0xdf, 0x43, 0xb0,
	0xce, 0x4e, 0x9d, 0x77, 0x29, 0xd1, 0x1d, 0x6a, 0xa0, 0xc6, 0x3a, 0xfc, 0x80, 0xe8, 0x9e, 0x43,
	0x7d, 0xd3, 0x40, 0x4d, 0x29, 0xbe, 0x84, 0x5d, 0xd2, 0xa1, 0xfa, 0x95, 0x4f, 0x89, 0x3b, 0xb0,
	0x3c, 0x74, 0x53, 0xb6, 0xa0, 0x6b, 0x5a, 0x44, 0x39, 0xea, 0x3a, 0x03, 0xdb, 0x40, 0x2d, 0xbc,
	0x03, 0x8d, 0x1e, 0xf1, 0x3a, 0x65, 0x4f, 0x76, 0xe4, 0xfd, 0x7a, 0x47, 0xbf, 0x22, 0x25, 0x82,
	0x70, 0x1b, 0xf6, 0xf5, 0x8e, 0x2d, 0x8b, 0x74, 0x4a, 0x3a, 0x1e, 0xf1, 0xbb, 0x8e, 0x65, 0x10,
	0x8a, 0x76, 0xa5, 0xc1, 0x7f, 0x65, 0x4c, 0x8b, 0x20, 0xbc, 0x56, 0x61, 0x10, 0x8b, 0xac, 0x2a,
	0xf6, 0xd6, 0x2a, 0xca, 0x8c, 0xac, 0xd8, 0x97, 0x66, 0x2e, 0x06, 0xa6, 0x65, 0x14, 0x8d, 0xca,
	0x87, 0x76, 0x80, 0x77, 0xe1, 0x66, 0x69, 0xc6, 0xb6, 0x4c, 0xd7, 0x43, 0x87, 0xf8, 0x08, 0xf6,
	0x4a, 0xa8, 0x47, 0x3c, 0x6a, 0xea, 0x79, 0x57, 0x8f, 0xe4, 0x59, 0x67, 0xe0, 0xf9, 0x4e, 0xd7,
	0xef, 0x91, 0x9e, 0x43, 0x1f, 0xa1, 0x36, 0xde, 0x87, 0x1d, 0xc3, 0xf0, 0x29, 0xb9, 0x3f, 0x20,
	0xae, 0xe7, 0xd3, 0x8e, 0x4e, 0xd0, 0x2f, 0xf5, 0x53, 0x1b, 0xc0, 0x9c, 0x86, 0xec, 0x99, 0xfc,
	0xf3, 0x19, 0xde, 0x86, 0x9a, 0xed, 0xd8, 0x04, 0xdd, 0xc0, 0x4d, 0xd8, 0x1e, 0xd8, 0xa6, 0xeb,
	0x0e, 0x88, 0x81, 0x2a, 0xb8, 0x05, 0x60, 0xda, 0x7d, 0xea, 0x5c, 0x52, 0xb9, 0x55, 0x1b, 0x32,
	0xdb, 0x35, 0x6d, 0xd3, 0xbd, 0x52, 0x2b, 0x02, 0xb0, 0x55, 0xf4, 0xa7, 0x76, 0x9a, 0x42, 0xd3,
	0x65, 0xe3, 0x09, 0x9b, 0x8a, 0x9c, 0x71, 0x07, 0x1a, 0x45, 0x6c, 0xa7, 0x53, 0x86, 0x6e, 0xe0,
	0x3d, 0xd8, 0x59, 0x02, 0x82, 0x3c, 0x8b, 0xb8, 0xc8, 0xf7, 0xb1, 0x00, 0x2f, 0xb3, 0xf4, 0x69,
	0x34, 0x1d, 0xa3, 0x0d, 0x69, 0xa1, 0x64, 0x62, 0x41, 0xc2, 0x42, 0x54, 0x5d, 0x3b, 0xd6, 0x4d,
	0xe6, 0xfc, 0x9a, 0x85, 0xa8, 0x76, 0xfa, 0xd1, 0xa6, 0xfa, 0x9f, 0xd5, 0x6f, 0xa9, 0xc1, 0x66,
	0x5c, 0x5c, 0x73, 0x00, 0xbb, 0xb1, 0x9e, 0xb1, 0x40, 0x30, 0x3d, 0x4d, 0x12, 0x36, 0x12, 0x51,
	0x3a, 0x45, 0xa1, 0xbc, 0x3d, 0x36, 0xb2, 0x74, 0xb6, 0x06, 0x32, 0x49, 0x1b, 0x5f, 0x05, 0x7c,
	0x0d, 0x7b, 0x2c, 0x3b, 0x1b, 0x1b, 0x8c, 0x8f, 0xb2, 0x68, 0xb8, 0xce, 0x30, 0x96, 0x4b, 0x1f,
	0xbb, 0xd7, 0xe9, 0xd3, 0x15, 0xc8, 0xd1, 0xb5, 0xa2, 0xb8, 0x64, 0xc2, 0x5d, 0x70, 0x3d, 0x9d,
	0x3e, 0x8e, 0xc6, 0x1c, 0x45, 0xf8, 0x00, 0x50, 0x21, 0xa1, 0x1f, 0x64, 0x22, 0x52, 0xf5, 0xdf,
	0x55, 0xf0, 0x1e, 0xb4, 0x94, 0x84, 0x15, 0xf8, 0xbd, 0x6c, 0xc0, 0x4d, 0x29, 0x61, 0x85, 0xfd,
	0x50, 0xc1, 0x47, 0x80, 0x97, 0x12, 0x56, 0x89, 0x1f, 0x2b, 0x72, 0x92, 0x4a, 0xc2, 0x12, 0xe4,
	0xe8, 0xa7, 0x0a, 0xde, 0x85, 0xa6, 0x42, 0x8b, 0x0e, 0xa1, 0xbf, 0x2a, 0x4a, 0x41, 0xc9, 0x50,
	0xc2, 0x7f, 0xe7, 0x27, 0x73, 0x61, 0x6a, 0xf4, 0xe8, 0xe3, 0x8d, 0x5c, 0x54, 0x71, 0x32, 0x07,
	0x3f, 0x91, 0x53, 0xae, 0xc7, 0xe6, 0x94, 0xb3, 0x4c, 0xa0, 0x0f, 0xaa, 0x2a, 0x32, 0x58, 0xc2,
	0x04, 0x43, 0x1f, 0x56, 0x71, 0x03, 0xb6, 0x62, 0x35, 0x04, 0xf4, 0x22, 0x4f, 0xb9, 0x2c, 0xc8,
	0x46, 0xd7, 0xe8, 0xd7, 0xaa, 0xf2, 0x92, 0x47, 0x94, 0xf1, 0x79, 0x22, 0xd0, 0x6f, 0x55, 0xc5,
	0x7f, 0xc9, 0xc4, 0x6a, 0xd5, 0xd0, 0xef, 0x55, 0xfc, 0x2a, 0xb4, 0x25, 0xb8, 0xea, 0xa4, 0xcc,
	0x44, 0x5c, 0x44, 0x23, 0x8e, 0xfe, 0xa8, 0xe2, 0x63, 0x38, 0x92, 0xe9, 0xa5, 0xcb, 0xb5, 0xec,
	0x9f, 0xd5, 0xdc, 0x6e, 0xee, 0xc9, 0x9c, 0x3e, 0x4e, 0xd1, 0xcf, 0x35, 0xdc, 0x02, 0x2d, 0xf6,
	0xa2, 0x09, 0xf3, 0xa2, 0x51, 0x8c, 0x3e, 0xd5, 0x54, 0x9f, 0xee, 0xcf, 0x59, 0xb6, 0xb0, 0xd3,
	0x90, 0xc9, 0x6a, 0x8e, 0x3e, 0xd3, 0xf0, 0x0e, 0x40, 0x6c, 0xa5, 0x41, 0x98, 0xdb, 0xfc, 0x3c,
	0x07, 0x28, 0x7b, 0x7f, 0xce, 0xb8, 0x30, 0x0d, 0xf4, 0x85, 0x7c, 0x75, 0x1a, 0x25, 0xe0, 0xb9,
	0x0e, 0xfa, 0x52, 0x53, 0x8d, 0xec, 0x24, 0x49, 0x3a, 0x0a, 0xc4, 0xb2, 0x91, 0x5f, 0x69, 0x6a,
	0x42, 0x6b, 0xdb, 0x5e, 0x88, 0xfb, 0x5a, 0xc3, 0x87, 0xb0, 0x1b, 0xaf, 0x6f, 0xaa, 0x21, 0x97,
	0xf2, 0x1b, 0xed, 0xe2, 0xe2, 0xbd, 0x77, 0xc6, 0x91, 0xb8, 0x9e, 0x0f, 0xe5, 0x3b, 0x7e, 0xfe,
	0x3c, 0x4a, 0x92, 0xe8, 0xb9, 0x60, 0xa3, 0xeb, 0xf3, 0xfc, 0x8d, 0xbf, 0x13, 0x46, 0x5c, 0x64,
	0xd1, 0x70, 0x2e, 0x58, 0x78, 0x1e, 0x4d, 0x05, 0xcb, 0xa6, 0x41, 0x72, 0xae, 0x1e, 0xfe, 0xf3,
	0xfc, 0xe1, 0x9f, 0x0d, 0x87, 0x5b, 0x2a, 0x7e, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2,
	0xc3, 0x62, 0xa9, 0xdb, 0x07, 0x00, 0x00,
}
