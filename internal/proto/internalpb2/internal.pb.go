// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

package internalpb2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
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
	// Definition Requests: Index
	MsgType_kCreateIndex           MsgType = 300
	MsgType_kDescribeIndex         MsgType = 301
	MsgType_kDescribeIndexProgress MsgType = 302
	// Manipulation Requests
	MsgType_kInsert MsgType = 400
	MsgType_kDelete MsgType = 401
	MsgType_kFlush  MsgType = 402
	// Query
	MsgType_kSearch       MsgType = 500
	MsgType_kSearchResult MsgType = 501
	// System Control
	MsgType_kTimeTick       MsgType = 1200
	MsgType_kQueryNodeStats MsgType = 1201
	MsgType_kLoadIndex      MsgType = 1202
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
	300:  "kCreateIndex",
	301:  "kDescribeIndex",
	302:  "kDescribeIndexProgress",
	400:  "kInsert",
	401:  "kDelete",
	402:  "kFlush",
	500:  "kSearch",
	501:  "kSearchResult",
	1200: "kTimeTick",
	1201: "kQueryNodeStats",
	1202: "kLoadIndex",
}

var MsgType_value = map[string]int32{
	"kNone":                  0,
	"kCreateCollection":      100,
	"kDropCollection":        101,
	"kHasCollection":         102,
	"kDescribeCollection":    103,
	"kShowCollections":       104,
	"kGetSysConfigs":         105,
	"kCreatePartition":       200,
	"kDropPartition":         201,
	"kHasPartition":          202,
	"kDescribePartition":     203,
	"kShowPartitions":        204,
	"kCreateIndex":           300,
	"kDescribeIndex":         301,
	"kDescribeIndexProgress": 302,
	"kInsert":                400,
	"kDelete":                401,
	"kFlush":                 402,
	"kSearch":                500,
	"kSearchResult":          501,
	"kTimeTick":              1200,
	"kQueryNodeStats":        1201,
	"kLoadIndex":             1202,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{0}
}

type StateCode int32

const (
	StateCode_INITIALIZING StateCode = 0
	StateCode_HEALTHY      StateCode = 1
	StateCode_ABNORMAL     StateCode = 2
)

var StateCode_name = map[int32]string{
	0: "INITIALIZING",
	1: "HEALTHY",
	2: "ABNORMAL",
}

var StateCode_value = map[string]int32{
	"INITIALIZING": 0,
	"HEALTHY":      1,
	"ABNORMAL":     2,
}

func (x StateCode) String() string {
	return proto.EnumName(StateCode_name, int32(x))
}

func (StateCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{1}
}

type NodeStates struct {
	NodeID               int64                    `protobuf:"varint,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Role                 string                   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	StateCode            StateCode                `protobuf:"varint,3,opt,name=state_code,json=stateCode,proto3,enum=milvus.proto.internal.StateCode" json:"state_code,omitempty"`
	ExtraInfo            []*commonpb.KeyValuePair `protobuf:"bytes,4,rep,name=extra_info,json=extraInfo,proto3" json:"extra_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NodeStates) Reset()         { *m = NodeStates{} }
func (m *NodeStates) String() string { return proto.CompactTextString(m) }
func (*NodeStates) ProtoMessage()    {}
func (*NodeStates) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{0}
}

func (m *NodeStates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStates.Unmarshal(m, b)
}
func (m *NodeStates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStates.Marshal(b, m, deterministic)
}
func (m *NodeStates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStates.Merge(m, src)
}
func (m *NodeStates) XXX_Size() int {
	return xxx_messageInfo_NodeStates.Size(m)
}
func (m *NodeStates) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStates.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStates proto.InternalMessageInfo

func (m *NodeStates) GetNodeID() int64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *NodeStates) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *NodeStates) GetStateCode() StateCode {
	if m != nil {
		return m.StateCode
	}
	return StateCode_INITIALIZING
}

func (m *NodeStates) GetExtraInfo() []*commonpb.KeyValuePair {
	if m != nil {
		return m.ExtraInfo
	}
	return nil
}

type ServiceStates struct {
	StateCode            StateCode                `protobuf:"varint,1,opt,name=state_code,json=stateCode,proto3,enum=milvus.proto.internal.StateCode" json:"state_code,omitempty"`
	NodeStates           []*NodeStates            `protobuf:"bytes,2,rep,name=node_states,json=nodeStates,proto3" json:"node_states,omitempty"`
	ExtraInfo            []*commonpb.KeyValuePair `protobuf:"bytes,3,rep,name=extra_info,json=extraInfo,proto3" json:"extra_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ServiceStates) Reset()         { *m = ServiceStates{} }
func (m *ServiceStates) String() string { return proto.CompactTextString(m) }
func (*ServiceStates) ProtoMessage()    {}
func (*ServiceStates) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{1}
}

func (m *ServiceStates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceStates.Unmarshal(m, b)
}
func (m *ServiceStates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceStates.Marshal(b, m, deterministic)
}
func (m *ServiceStates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStates.Merge(m, src)
}
func (m *ServiceStates) XXX_Size() int {
	return xxx_messageInfo_ServiceStates.Size(m)
}
func (m *ServiceStates) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStates.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStates proto.InternalMessageInfo

func (m *ServiceStates) GetStateCode() StateCode {
	if m != nil {
		return m.StateCode
	}
	return StateCode_INITIALIZING
}

func (m *ServiceStates) GetNodeStates() []*NodeStates {
	if m != nil {
		return m.NodeStates
	}
	return nil
}

func (m *ServiceStates) GetExtraInfo() []*commonpb.KeyValuePair {
	if m != nil {
		return m.ExtraInfo
	}
	return nil
}

type NodeInfo struct {
	Address              *commonpb.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Role                 string            `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{2}
}

func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetAddress() *commonpb.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *NodeInfo) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type InitParams struct {
	NodeID               int64                    `protobuf:"varint,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	StartParams          []*commonpb.KeyValuePair `protobuf:"bytes,2,rep,name=start_params,json=startParams,proto3" json:"start_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *InitParams) Reset()         { *m = InitParams{} }
func (m *InitParams) String() string { return proto.CompactTextString(m) }
func (*InitParams) ProtoMessage()    {}
func (*InitParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{3}
}

func (m *InitParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitParams.Unmarshal(m, b)
}
func (m *InitParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitParams.Marshal(b, m, deterministic)
}
func (m *InitParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitParams.Merge(m, src)
}
func (m *InitParams) XXX_Size() int {
	return xxx_messageInfo_InitParams.Size(m)
}
func (m *InitParams) XXX_DiscardUnknown() {
	xxx_messageInfo_InitParams.DiscardUnknown(m)
}

var xxx_messageInfo_InitParams proto.InternalMessageInfo

func (m *InitParams) GetNodeID() int64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *InitParams) GetStartParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.StartParams
	}
	return nil
}

type StringList struct {
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringList) Reset()         { *m = StringList{} }
func (m *StringList) String() string { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()    {}
func (*StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{4}
}

func (m *StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringList.Unmarshal(m, b)
}
func (m *StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringList.Marshal(b, m, deterministic)
}
func (m *StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringList.Merge(m, src)
}
func (m *StringList) XXX_Size() int {
	return xxx_messageInfo_StringList.Size(m)
}
func (m *StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_StringList proto.InternalMessageInfo

func (m *StringList) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type MsgBase struct {
	MsgType              MsgType  `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=milvus.proto.internal.MsgType" json:"msg_type,omitempty"`
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
	return fileDescriptor_41f4a519b878ee3b, []int{5}
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

type TimeTickMsg struct {
	Base                 *MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeTickMsg) Reset()         { *m = TimeTickMsg{} }
func (m *TimeTickMsg) String() string { return proto.CompactTextString(m) }
func (*TimeTickMsg) ProtoMessage()    {}
func (*TimeTickMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{6}
}

func (m *TimeTickMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeTickMsg.Unmarshal(m, b)
}
func (m *TimeTickMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeTickMsg.Marshal(b, m, deterministic)
}
func (m *TimeTickMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeTickMsg.Merge(m, src)
}
func (m *TimeTickMsg) XXX_Size() int {
	return xxx_messageInfo_TimeTickMsg.Size(m)
}
func (m *TimeTickMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeTickMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TimeTickMsg proto.InternalMessageInfo

func (m *TimeTickMsg) GetBase() *MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func init() {
	proto.RegisterEnum("milvus.proto.internal.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("milvus.proto.internal.StateCode", StateCode_name, StateCode_value)
	proto.RegisterType((*NodeStates)(nil), "milvus.proto.internal.NodeStates")
	proto.RegisterType((*ServiceStates)(nil), "milvus.proto.internal.ServiceStates")
	proto.RegisterType((*NodeInfo)(nil), "milvus.proto.internal.NodeInfo")
	proto.RegisterType((*InitParams)(nil), "milvus.proto.internal.InitParams")
	proto.RegisterType((*StringList)(nil), "milvus.proto.internal.StringList")
	proto.RegisterType((*MsgBase)(nil), "milvus.proto.internal.MsgBase")
	proto.RegisterType((*TimeTickMsg)(nil), "milvus.proto.internal.TimeTickMsg")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor_41f4a519b878ee3b) }

var fileDescriptor_41f4a519b878ee3b = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xde, 0xf1, 0x78, 0xe3, 0x4c, 0xd9, 0xeb, 0xf4, 0x56, 0x92, 0x5d, 0x6b, 0x59, 0x21, 0x63,
	0x71, 0x88, 0x56, 0x22, 0x91, 0x8c, 0x84, 0xe0, 0x04, 0x4e, 0xbc, 0x6c, 0x46, 0x38, 0x26, 0x8c,
	0xad, 0x95, 0xd8, 0x8b, 0x35, 0x9e, 0xa9, 0x8c, 0x9b, 0xf9, 0x69, 0xab, 0xbb, 0x1d, 0xd6, 0xfb,
	0x14, 0x80, 0x78, 0x0c, 0x40, 0xc0, 0x95, 0x17, 0xe0, 0xf7, 0xca, 0x93, 0xc0, 0x1d, 0x4d, 0x8f,
	0x7f, 0x62, 0xc9, 0x41, 0x82, 0x5b, 0xd7, 0xd7, 0x55, 0x5f, 0x7f, 0x5f, 0x55, 0xcd, 0x40, 0x9d,
	0x67, 0x9a, 0x64, 0xe6, 0x27, 0xc7, 0x53, 0x29, 0xb4, 0xc0, 0xc3, 0x94, 0x27, 0xd7, 0x33, 0x55,
	0x44, 0xc7, 0xcb, 0xcb, 0x47, 0xb5, 0x40, 0xa4, 0xa9, 0xc8, 0x0a, 0xb8, 0xf5, 0x93, 0x05, 0xd0,
	0x17, 0x21, 0x0d, 0xb4, 0xaf, 0x49, 0xe1, 0x03, 0xd8, 0xc9, 0x44, 0x48, 0x6e, 0xb7, 0x61, 0x35,
	0xad, 0x23, 0xdb, 0x5b, 0x44, 0x88, 0x50, 0x96, 0x22, 0xa1, 0x46, 0xa9, 0x69, 0x1d, 0x39, 0x9e,
	0x39, 0xe3, 0xfb, 0x00, 0x2a, 0xaf, 0x1a, 0x05, 0x22, 0xa4, 0x86, 0xdd, 0xb4, 0x8e, 0xea, 0xed,
	0xe6, 0xf1, 0xd6, 0x47, 0x8f, 0x0d, 0xfd, 0x99, 0x08, 0xc9, 0x73, 0xd4, 0xf2, 0x88, 0x1f, 0x00,
	0xd0, 0x4b, 0x2d, 0xfd, 0x11, 0xcf, 0xae, 0x44, 0xa3, 0xdc, 0xb4, 0x8f, 0xaa, 0xed, 0x37, 0x36,
	0x09, 0x16, 0x5a, 0x3f, 0xa2, 0xf9, 0x73, 0x3f, 0x99, 0xd1, 0xa5, 0xcf, 0xa5, 0xe7, 0x98, 0x22,
	0x37, 0xbb, 0x12, 0xad, 0x3f, 0x2d, 0xb8, 0x37, 0x20, 0x79, 0xcd, 0x83, 0xa5, 0x81, 0x4d, 0x51,
	0xd6, 0x7f, 0x17, 0x75, 0x0a, 0xd5, 0xdc, 0xf3, 0xc8, 0x20, 0xaa, 0x51, 0xda, 0xa6, 0x6a, 0xc5,
	0xb0, 0xee, 0x9c, 0x07, 0xd9, 0xba, 0x8b, 0x9b, 0xc6, 0xec, 0xff, 0x61, 0xec, 0x39, 0xec, 0xe6,
	0xdc, 0xf9, 0x19, 0xdf, 0x81, 0x8a, 0x1f, 0x86, 0x92, 0x94, 0x32, 0x7e, 0xaa, 0xed, 0xc7, 0x5b,
	0xa9, 0x3a, 0x45, 0x8e, 0xb7, 0x4c, 0xde, 0x36, 0xb3, 0xd6, 0x67, 0x00, 0x6e, 0xc6, 0xf5, 0xa5,
	0x2f, 0xfd, 0xf4, 0xf6, 0x69, 0x77, 0xa1, 0xa6, 0xb4, 0x2f, 0xf5, 0x68, 0x6a, 0xf2, 0xb6, 0x37,
	0x61, 0x9b, 0x83, 0xaa, 0x29, 0x2b, 0xd8, 0x5b, 0x6f, 0x02, 0x0c, 0xb4, 0xe4, 0x59, 0xd4, 0xe3,
	0x4a, 0xe7, 0x6f, 0x5d, 0xe7, 0x79, 0x05, 0x9b, 0xe3, 0x2d, 0xa2, 0xd6, 0xd7, 0x16, 0x54, 0x2e,
	0x54, 0x74, 0xea, 0x2b, 0xc2, 0xf7, 0x60, 0x37, 0x55, 0xd1, 0x48, 0xcf, 0xa7, 0xcb, 0xd1, 0xbd,
	0x7e, 0x4b, 0xe3, 0x2f, 0x54, 0x34, 0x9c, 0x4f, 0xc9, 0xab, 0xa4, 0xc5, 0x01, 0x0f, 0xe0, 0x6e,
	0xaa, 0x22, 0xb7, 0x6b, 0xdc, 0xda, 0x5e, 0x11, 0xe0, 0x63, 0x70, 0x34, 0x4f, 0x49, 0x69, 0x3f,
	0x9d, 0x9a, 0x0d, 0x2d, 0x7b, 0x6b, 0x00, 0x1f, 0xc1, 0xae, 0x12, 0x33, 0x19, 0xe4, 0x0d, 0x28,
	0x9b, 0xb2, 0x55, 0xdc, 0xea, 0x40, 0x75, 0xc8, 0x53, 0x1a, 0xf2, 0x20, 0xbe, 0x50, 0x11, 0xb6,
	0xa1, 0x3c, 0xf6, 0x15, 0x2d, 0x06, 0xf0, 0x2f, 0xaa, 0x72, 0x1f, 0x9e, 0xc9, 0x7d, 0xf2, 0x87,
	0x6d, 0x9c, 0x19, 0x79, 0x0e, 0xdc, 0x8d, 0xfb, 0x22, 0x23, 0x76, 0x07, 0x0f, 0xe1, 0x7e, 0x7c,
	0x26, 0xc9, 0xec, 0x5b, 0x92, 0x50, 0xa0, 0xb9, 0xc8, 0x58, 0x88, 0xfb, 0xb0, 0x17, 0x77, 0xa5,
	0x98, 0xde, 0x00, 0x09, 0x11, 0xea, 0xf1, 0xb9, 0xaf, 0x6e, 0x60, 0x57, 0xf8, 0x10, 0xf6, 0xe3,
	0x2e, 0xa9, 0x40, 0xf2, 0xf1, 0x4d, 0x86, 0x08, 0x0f, 0x80, 0xc5, 0x83, 0x89, 0xf8, 0x7c, 0x0d,
	0x2a, 0x36, 0x31, 0x14, 0xcf, 0x48, 0x0f, 0xe6, 0xea, 0x4c, 0x64, 0x57, 0x3c, 0x52, 0x8c, 0xe3,
	0x21, 0xb0, 0x85, 0x84, 0x4b, 0x5f, 0x6a, 0x6e, 0xea, 0x7f, 0xb6, 0x70, 0x1f, 0xea, 0x46, 0xc2,
	0x1a, 0xfc, 0xc5, 0x42, 0x84, 0x7b, 0xb9, 0x84, 0x35, 0xf6, 0xab, 0x85, 0x0f, 0x01, 0x57, 0x12,
	0xd6, 0x17, 0xbf, 0x59, 0x78, 0x00, 0x7b, 0x46, 0xc2, 0x0a, 0x54, 0xec, 0x77, 0x0b, 0xef, 0x43,
	0x6d, 0xf1, 0x9c, 0x9b, 0x85, 0xf4, 0x92, 0x7d, 0x53, 0x2a, 0x9e, 0x5a, 0x30, 0x14, 0xe0, 0xb7,
	0x25, 0x7c, 0x0d, 0x1e, 0x6c, 0x82, 0x97, 0x52, 0x44, 0xf9, 0x2a, 0xb3, 0xef, 0x4a, 0x58, 0x83,
	0x4a, 0xec, 0x66, 0x8a, 0xa4, 0x66, 0x5f, 0xd8, 0x26, 0xea, 0x52, 0x42, 0x9a, 0xd8, 0x97, 0x36,
	0x56, 0x61, 0x27, 0xfe, 0x30, 0x99, 0xa9, 0x09, 0xfb, 0xaa, 0xb8, 0x1a, 0x90, 0x2f, 0x83, 0x09,
	0xfb, 0xcb, 0x36, 0xf2, 0x8b, 0xc8, 0x23, 0x35, 0x4b, 0x34, 0xfb, 0xdb, 0xc6, 0x3a, 0x38, 0xf1,
	0x72, 0xb8, 0xec, 0x7b, 0xc7, 0xa8, 0xfe, 0x64, 0x46, 0x72, 0xbe, 0xfc, 0x9c, 0x15, 0xfb, 0xc1,
	0xc1, 0x3d, 0x80, 0xb8, 0x27, 0xfc, 0xb0, 0x90, 0xf7, 0xa3, 0xf3, 0xe4, 0x5d, 0x70, 0x56, 0x7f,
	0x0c, 0x64, 0x50, 0x73, 0xfb, 0xee, 0xd0, 0xed, 0xf4, 0xdc, 0x17, 0x6e, 0xff, 0x19, 0xbb, 0x83,
	0x55, 0xa8, 0x9c, 0x3f, 0xed, 0xf4, 0x86, 0xe7, 0x9f, 0x32, 0x0b, 0x6b, 0xb0, 0xdb, 0x39, 0xed,
	0x7f, 0xec, 0x5d, 0x74, 0x7a, 0xac, 0x74, 0xfa, 0xf4, 0xc5, 0x59, 0xc4, 0xf5, 0x64, 0x36, 0xce,
	0x3f, 0x9a, 0x93, 0x57, 0x3c, 0x49, 0xf8, 0x2b, 0x4d, 0xc1, 0xe4, 0xa4, 0x58, 0xa3, 0xb7, 0x42,
	0xae, 0xb4, 0xe4, 0xe3, 0x99, 0xa6, 0xf0, 0x64, 0xb9, 0x4c, 0x27, 0x66, 0xb7, 0x56, 0xe1, 0x74,
	0xdc, 0x1e, 0xef, 0x18, 0xe8, 0xed, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x75, 0x95, 0xa0,
	0xe9, 0x05, 0x00, 0x00,
}
