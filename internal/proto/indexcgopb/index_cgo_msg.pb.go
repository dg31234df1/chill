// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index_cgo_msg.proto

package indexcgopb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/milvus-proto/go-api/commonpb"
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

type TypeParams struct {
	Params               []*commonpb.KeyValuePair `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TypeParams) Reset()         { *m = TypeParams{} }
func (m *TypeParams) String() string { return proto.CompactTextString(m) }
func (*TypeParams) ProtoMessage()    {}
func (*TypeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{0}
}

func (m *TypeParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeParams.Unmarshal(m, b)
}
func (m *TypeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeParams.Marshal(b, m, deterministic)
}
func (m *TypeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeParams.Merge(m, src)
}
func (m *TypeParams) XXX_Size() int {
	return xxx_messageInfo_TypeParams.Size(m)
}
func (m *TypeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeParams.DiscardUnknown(m)
}

var xxx_messageInfo_TypeParams proto.InternalMessageInfo

func (m *TypeParams) GetParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.Params
	}
	return nil
}

type IndexParams struct {
	Params               []*commonpb.KeyValuePair `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IndexParams) Reset()         { *m = IndexParams{} }
func (m *IndexParams) String() string { return proto.CompactTextString(m) }
func (*IndexParams) ProtoMessage()    {}
func (*IndexParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{1}
}

func (m *IndexParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexParams.Unmarshal(m, b)
}
func (m *IndexParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexParams.Marshal(b, m, deterministic)
}
func (m *IndexParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexParams.Merge(m, src)
}
func (m *IndexParams) XXX_Size() int {
	return xxx_messageInfo_IndexParams.Size(m)
}
func (m *IndexParams) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexParams.DiscardUnknown(m)
}

var xxx_messageInfo_IndexParams proto.InternalMessageInfo

func (m *IndexParams) GetParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.Params
	}
	return nil
}

// TypeParams & IndexParams will be replaced by MapParams later
type MapParams struct {
	Params               []*commonpb.KeyValuePair `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MapParams) Reset()         { *m = MapParams{} }
func (m *MapParams) String() string { return proto.CompactTextString(m) }
func (*MapParams) ProtoMessage()    {}
func (*MapParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{2}
}

func (m *MapParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapParams.Unmarshal(m, b)
}
func (m *MapParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapParams.Marshal(b, m, deterministic)
}
func (m *MapParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapParams.Merge(m, src)
}
func (m *MapParams) XXX_Size() int {
	return xxx_messageInfo_MapParams.Size(m)
}
func (m *MapParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MapParams.DiscardUnknown(m)
}

var xxx_messageInfo_MapParams proto.InternalMessageInfo

func (m *MapParams) GetParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.Params
	}
	return nil
}

type MapParamsV2 struct {
	Params               map[string]string `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapParamsV2) Reset()         { *m = MapParamsV2{} }
func (m *MapParamsV2) String() string { return proto.CompactTextString(m) }
func (*MapParamsV2) ProtoMessage()    {}
func (*MapParamsV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{3}
}

func (m *MapParamsV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapParamsV2.Unmarshal(m, b)
}
func (m *MapParamsV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapParamsV2.Marshal(b, m, deterministic)
}
func (m *MapParamsV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapParamsV2.Merge(m, src)
}
func (m *MapParamsV2) XXX_Size() int {
	return xxx_messageInfo_MapParamsV2.Size(m)
}
func (m *MapParamsV2) XXX_DiscardUnknown() {
	xxx_messageInfo_MapParamsV2.DiscardUnknown(m)
}

var xxx_messageInfo_MapParamsV2 proto.InternalMessageInfo

func (m *MapParamsV2) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type Binary struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Binary) Reset()         { *m = Binary{} }
func (m *Binary) String() string { return proto.CompactTextString(m) }
func (*Binary) ProtoMessage()    {}
func (*Binary) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{4}
}

func (m *Binary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Binary.Unmarshal(m, b)
}
func (m *Binary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Binary.Marshal(b, m, deterministic)
}
func (m *Binary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Binary.Merge(m, src)
}
func (m *Binary) XXX_Size() int {
	return xxx_messageInfo_Binary.Size(m)
}
func (m *Binary) XXX_DiscardUnknown() {
	xxx_messageInfo_Binary.DiscardUnknown(m)
}

var xxx_messageInfo_Binary proto.InternalMessageInfo

func (m *Binary) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Binary) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type BinarySet struct {
	Datas                []*Binary `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BinarySet) Reset()         { *m = BinarySet{} }
func (m *BinarySet) String() string { return proto.CompactTextString(m) }
func (*BinarySet) ProtoMessage()    {}
func (*BinarySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{5}
}

func (m *BinarySet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BinarySet.Unmarshal(m, b)
}
func (m *BinarySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BinarySet.Marshal(b, m, deterministic)
}
func (m *BinarySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinarySet.Merge(m, src)
}
func (m *BinarySet) XXX_Size() int {
	return xxx_messageInfo_BinarySet.Size(m)
}
func (m *BinarySet) XXX_DiscardUnknown() {
	xxx_messageInfo_BinarySet.DiscardUnknown(m)
}

var xxx_messageInfo_BinarySet proto.InternalMessageInfo

func (m *BinarySet) GetDatas() []*Binary {
	if m != nil {
		return m.Datas
	}
	return nil
}

func init() {
	proto.RegisterType((*TypeParams)(nil), "milvus.proto.indexcgo.TypeParams")
	proto.RegisterType((*IndexParams)(nil), "milvus.proto.indexcgo.IndexParams")
	proto.RegisterType((*MapParams)(nil), "milvus.proto.indexcgo.MapParams")
	proto.RegisterType((*MapParamsV2)(nil), "milvus.proto.indexcgo.MapParamsV2")
	proto.RegisterMapType((map[string]string)(nil), "milvus.proto.indexcgo.MapParamsV2.ParamsEntry")
	proto.RegisterType((*Binary)(nil), "milvus.proto.indexcgo.Binary")
	proto.RegisterType((*BinarySet)(nil), "milvus.proto.indexcgo.BinarySet")
}

func init() { proto.RegisterFile("index_cgo_msg.proto", fileDescriptor_c009bd9544a7343c) }

var fileDescriptor_c009bd9544a7343c = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xc6, 0x0a, 0x7d, 0xdd, 0x41, 0xa2, 0x42, 0x19, 0x08, 0xb3, 0xa7, 0x5d, 0x4c,
	0x65, 0x43, 0x74, 0x9e, 0x64, 0xe0, 0x54, 0x44, 0x18, 0x55, 0x76, 0xf0, 0x32, 0xd2, 0x2e, 0xd4,
	0x60, 0x9b, 0x94, 0x34, 0x1d, 0xf6, 0x5b, 0xf8, 0x91, 0xa5, 0x49, 0x2b, 0x53, 0x94, 0x1d, 0x76,
	0xfb, 0xe7, 0xcf, 0xfb, 0xfd, 0xda, 0xf7, 0xe0, 0x90, 0xf1, 0x35, 0xfd, 0x58, 0xc5, 0x89, 0x58,
	0x65, 0x45, 0x82, 0x73, 0x29, 0x94, 0x40, 0xc7, 0x19, 0x4b, 0x37, 0x65, 0x61, 0x5e, 0x58, 0x4f,
	0xc4, 0x89, 0x18, 0xf4, 0x63, 0x91, 0x65, 0x82, 0x9b, 0xda, 0xbf, 0x03, 0x78, 0xa9, 0x72, 0xba,
	0x20, 0x92, 0x64, 0x05, 0x9a, 0x82, 0x9d, 0xeb, 0xe4, 0x59, 0xc3, 0xee, 0xc8, 0x1d, 0x9f, 0xe2,
	0x1f, 0x8e, 0x86, 0x7c, 0xa4, 0xd5, 0x92, 0xa4, 0x25, 0x5d, 0x10, 0x26, 0xc3, 0x06, 0xf0, 0xef,
	0xc1, 0x7d, 0xa8, 0x3f, 0xb1, 0xbf, 0x69, 0x0e, 0xce, 0x13, 0xc9, 0xf7, 0xf7, 0x7c, 0x5a, 0xe0,
	0x7e, 0x8b, 0x96, 0x63, 0x34, 0xff, 0xa5, 0xc2, 0xf8, 0xcf, 0x03, 0xe1, 0x2d, 0x06, 0x9b, 0x70,
	0xcb, 0x95, 0xac, 0x5a, 0xef, 0x60, 0x0a, 0xee, 0x56, 0x8d, 0x0e, 0xa0, 0xfb, 0x4e, 0x2b, 0xcf,
	0x1a, 0x5a, 0x23, 0x27, 0xac, 0x23, 0x3a, 0x82, 0xde, 0xa6, 0xfe, 0x1b, 0xaf, 0xa3, 0x3b, 0xf3,
	0xb8, 0xee, 0x5c, 0x59, 0xfe, 0x39, 0xd8, 0x33, 0xc6, 0xc9, 0x6e, 0xaa, 0xdf, 0x50, 0xfe, 0x0d,
	0x38, 0x86, 0x78, 0xa6, 0x0a, 0x4d, 0xa0, 0xb7, 0x26, 0x8a, 0xb4, 0x0b, 0x9c, 0xfc, 0xb3, 0x80,
	0x01, 0x42, 0x33, 0x3b, 0xbb, 0x7c, 0xbd, 0x48, 0x98, 0x7a, 0x2b, 0xa3, 0xfa, 0x58, 0x81, 0x21,
	0xce, 0x98, 0x68, 0x52, 0xc0, 0xb8, 0xa2, 0x92, 0x93, 0x34, 0xd0, 0x92, 0xa0, 0x95, 0xe4, 0x51,
	0x64, 0xeb, 0x66, 0xf2, 0x15, 0x00, 0x00, 0xff, 0xff, 0x83, 0x12, 0x13, 0xfb, 0x5d, 0x02, 0x00,
	0x00,
}
