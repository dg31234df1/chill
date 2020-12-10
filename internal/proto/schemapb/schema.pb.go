// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package schemapb

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

//*
// @brief Field data type
type DataType int32

const (
	DataType_NONE          DataType = 0
	DataType_BOOL          DataType = 1
	DataType_INT8          DataType = 2
	DataType_INT16         DataType = 3
	DataType_INT32         DataType = 4
	DataType_INT64         DataType = 5
	DataType_FLOAT         DataType = 10
	DataType_DOUBLE        DataType = 11
	DataType_STRING        DataType = 20
	DataType_VECTOR_BINARY DataType = 100
	DataType_VECTOR_FLOAT  DataType = 101
)

var DataType_name = map[int32]string{
	0:   "NONE",
	1:   "BOOL",
	2:   "INT8",
	3:   "INT16",
	4:   "INT32",
	5:   "INT64",
	10:  "FLOAT",
	11:  "DOUBLE",
	20:  "STRING",
	100: "VECTOR_BINARY",
	101: "VECTOR_FLOAT",
}

var DataType_value = map[string]int32{
	"NONE":          0,
	"BOOL":          1,
	"INT8":          2,
	"INT16":         3,
	"INT32":         4,
	"INT64":         5,
	"FLOAT":         10,
	"DOUBLE":        11,
	"STRING":        20,
	"VECTOR_BINARY": 100,
	"VECTOR_FLOAT":  101,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

//*
// @brief Field schema
type FieldSchema struct {
	FieldID              int64                    `protobuf:"varint,1,opt,name=fieldID,proto3" json:"fieldID,omitempty"`
	Name                 string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsPrimaryKey         bool                     `protobuf:"varint,3,opt,name=is_primary_key,json=isPrimaryKey,proto3" json:"is_primary_key,omitempty"`
	Description          string                   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DataType             DataType                 `protobuf:"varint,5,opt,name=data_type,json=dataType,proto3,enum=milvus.proto.schema.DataType" json:"data_type,omitempty"`
	TypeParams           []*commonpb.KeyValuePair `protobuf:"bytes,6,rep,name=type_params,json=typeParams,proto3" json:"type_params,omitempty"`
	IndexParams          []*commonpb.KeyValuePair `protobuf:"bytes,7,rep,name=index_params,json=indexParams,proto3" json:"index_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FieldSchema) Reset()         { *m = FieldSchema{} }
func (m *FieldSchema) String() string { return proto.CompactTextString(m) }
func (*FieldSchema) ProtoMessage()    {}
func (*FieldSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *FieldSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldSchema.Unmarshal(m, b)
}
func (m *FieldSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldSchema.Marshal(b, m, deterministic)
}
func (m *FieldSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldSchema.Merge(m, src)
}
func (m *FieldSchema) XXX_Size() int {
	return xxx_messageInfo_FieldSchema.Size(m)
}
func (m *FieldSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldSchema.DiscardUnknown(m)
}

var xxx_messageInfo_FieldSchema proto.InternalMessageInfo

func (m *FieldSchema) GetFieldID() int64 {
	if m != nil {
		return m.FieldID
	}
	return 0
}

func (m *FieldSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldSchema) GetIsPrimaryKey() bool {
	if m != nil {
		return m.IsPrimaryKey
	}
	return false
}

func (m *FieldSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FieldSchema) GetDataType() DataType {
	if m != nil {
		return m.DataType
	}
	return DataType_NONE
}

func (m *FieldSchema) GetTypeParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.TypeParams
	}
	return nil
}

func (m *FieldSchema) GetIndexParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.IndexParams
	}
	return nil
}

//*
// @brief Collection schema
type CollectionSchema struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AutoID               bool           `protobuf:"varint,3,opt,name=autoID,proto3" json:"autoID,omitempty"`
	Fields               []*FieldSchema `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CollectionSchema) Reset()         { *m = CollectionSchema{} }
func (m *CollectionSchema) String() string { return proto.CompactTextString(m) }
func (*CollectionSchema) ProtoMessage()    {}
func (*CollectionSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}

func (m *CollectionSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionSchema.Unmarshal(m, b)
}
func (m *CollectionSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionSchema.Marshal(b, m, deterministic)
}
func (m *CollectionSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionSchema.Merge(m, src)
}
func (m *CollectionSchema) XXX_Size() int {
	return xxx_messageInfo_CollectionSchema.Size(m)
}
func (m *CollectionSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionSchema.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionSchema proto.InternalMessageInfo

func (m *CollectionSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectionSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CollectionSchema) GetAutoID() bool {
	if m != nil {
		return m.AutoID
	}
	return false
}

func (m *CollectionSchema) GetFields() []*FieldSchema {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterEnum("milvus.proto.schema.DataType", DataType_name, DataType_value)
	proto.RegisterType((*FieldSchema)(nil), "milvus.proto.schema.FieldSchema")
	proto.RegisterType((*CollectionSchema)(nil), "milvus.proto.schema.CollectionSchema")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x8a, 0xd3, 0x40,
	0x18, 0xc5, 0x4d, 0x9b, 0x66, 0xdb, 0x2f, 0x71, 0x19, 0x47, 0x91, 0x20, 0x08, 0x71, 0xf1, 0x22,
	0x08, 0x36, 0xd8, 0x95, 0x65, 0xf1, 0xca, 0xcd, 0xa6, 0x2b, 0x61, 0x4b, 0x52, 0xb2, 0x71, 0x41,
	0x6f, 0xc2, 0x34, 0x19, 0xed, 0x60, 0xfe, 0x91, 0x4c, 0xc4, 0xec, 0x5b, 0x78, 0xeb, 0x1b, 0xf9,
	0x56, 0x92, 0x64, 0x0a, 0x55, 0x7b, 0xb1, 0x77, 0xe7, 0x7c, 0x33, 0xe7, 0x63, 0xce, 0x6f, 0x40,
	0xab, 0xe3, 0x2d, 0xcd, 0xc8, 0xbc, 0xac, 0x0a, 0x5e, 0xe0, 0xc7, 0x19, 0x4b, 0xbf, 0x37, 0xf5,
	0xe0, 0xe6, 0xc3, 0xd1, 0x33, 0x2d, 0x2e, 0xb2, 0xac, 0xc8, 0x87, 0xe1, 0xc9, 0xef, 0x11, 0xa8,
	0x57, 0x8c, 0xa6, 0xc9, 0x4d, 0x7f, 0x8a, 0x75, 0x38, 0xfa, 0xd2, 0x59, 0xd7, 0xd1, 0x25, 0x43,
	0x32, 0xc7, 0xc1, 0xce, 0x62, 0x0c, 0x72, 0x4e, 0x32, 0xaa, 0x8f, 0x0c, 0xc9, 0x9c, 0x05, 0xbd,
	0xc6, 0x2f, 0xe1, 0x98, 0xd5, 0x51, 0x59, 0xb1, 0x8c, 0x54, 0x6d, 0xf4, 0x8d, 0xb6, 0xfa, 0xd8,
	0x90, 0xcc, 0x69, 0xa0, 0xb1, 0x7a, 0x3d, 0x0c, 0xaf, 0x69, 0x8b, 0x0d, 0x50, 0x13, 0x5a, 0xc7,
	0x15, 0x2b, 0x39, 0x2b, 0x72, 0x5d, 0xee, 0x17, 0xec, 0x8f, 0xf0, 0x3b, 0x98, 0x25, 0x84, 0x93,
	0x88, 0xb7, 0x25, 0xd5, 0x27, 0x86, 0x64, 0x1e, 0x2f, 0x9e, 0xcf, 0x0f, 0x3c, 0x7e, 0xee, 0x10,
	0x4e, 0xc2, 0xb6, 0xa4, 0xc1, 0x34, 0x11, 0x0a, 0xdb, 0xa0, 0x76, 0xb1, 0xa8, 0x24, 0x15, 0xc9,
	0x6a, 0x5d, 0x31, 0xc6, 0xa6, 0xba, 0x78, 0xf1, 0x77, 0x5a, 0x54, 0xbe, 0xa6, 0xed, 0x2d, 0x49,
	0x1b, 0xba, 0x26, 0xac, 0x0a, 0xa0, 0x4b, 0xad, 0xfb, 0x10, 0x76, 0x40, 0x63, 0x79, 0x42, 0x7f,
	0xec, 0x96, 0x1c, 0xdd, 0x77, 0x89, 0xda, 0xc7, 0x86, 0x2d, 0x27, 0xbf, 0x24, 0x40, 0x97, 0x45,
	0x9a, 0xd2, 0xb8, 0x2b, 0x25, 0x80, 0xee, 0xb0, 0x49, 0x7b, 0xd8, 0xfe, 0x01, 0x32, 0xfa, 0x1f,
	0xc8, 0x53, 0x50, 0x48, 0xc3, 0x0b, 0xd7, 0x11, 0x40, 0x85, 0xc3, 0xe7, 0xa0, 0xf4, 0xff, 0x51,
	0xeb, 0x72, 0xff, 0x44, 0xe3, 0x20, 0xa5, 0xbd, 0x0f, 0x0d, 0xc4, 0xfd, 0x57, 0x3f, 0x25, 0x98,
	0xee, 0xe8, 0xe1, 0x29, 0xc8, 0x9e, 0xef, 0x2d, 0xd1, 0x83, 0x4e, 0xd9, 0xbe, 0xbf, 0x42, 0x52,
	0xa7, 0x5c, 0x2f, 0x3c, 0x47, 0x23, 0x3c, 0x83, 0x89, 0xeb, 0x85, 0x6f, 0xce, 0xd0, 0x58, 0xc8,
	0xd3, 0x05, 0x92, 0x85, 0x3c, 0x7b, 0x8b, 0x26, 0x9d, 0xbc, 0x5a, 0xf9, 0x17, 0x21, 0x02, 0x0c,
	0xa0, 0x38, 0xfe, 0x47, 0x7b, 0xb5, 0x44, 0x6a, 0xa7, 0x6f, 0xc2, 0xc0, 0xf5, 0x3e, 0xa0, 0x27,
	0xf8, 0x11, 0x3c, 0xbc, 0x5d, 0x5e, 0x86, 0x7e, 0x10, 0xd9, 0xae, 0x77, 0x11, 0x7c, 0x42, 0x09,
	0x46, 0xa0, 0x89, 0xd1, 0x10, 0xa6, 0xb6, 0xfd, 0xf9, 0xfd, 0x57, 0xc6, 0xb7, 0xcd, 0xa6, 0x63,
	0x6b, 0xdd, 0xb1, 0x34, 0x65, 0x77, 0x9c, 0xc6, 0x5b, 0x6b, 0x28, 0xf5, 0x3a, 0x61, 0x35, 0xaf,
	0xd8, 0xa6, 0xe1, 0x34, 0xb1, 0x58, 0xce, 0x69, 0x95, 0x93, 0xd4, 0xea, 0x9b, 0x5a, 0x43, 0xd3,
	0x72, 0xb3, 0x51, 0x7a, 0x7f, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xaf, 0x4d, 0x07, 0xfa,
	0x02, 0x00, 0x00,
}
