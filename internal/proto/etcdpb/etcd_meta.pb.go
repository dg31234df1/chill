// Code generated by protoc-gen-go. DO NOT EDIT.
// source: etcd_meta.proto

package etcdpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/milvus/internal/proto/commonpb"
	schemapb "github.com/milvus-io/milvus/internal/proto/schemapb"
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

type IndexInfo struct {
	IndexName            string                   `protobuf:"bytes,1,opt,name=index_name,json=indexName,proto3" json:"index_name,omitempty"`
	IndexID              int64                    `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	IndexParams          []*commonpb.KeyValuePair `protobuf:"bytes,3,rep,name=index_params,json=indexParams,proto3" json:"index_params,omitempty"`
	Deleted              bool                     `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	CreateTime           uint64                   `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IndexInfo) Reset()         { *m = IndexInfo{} }
func (m *IndexInfo) String() string { return proto.CompactTextString(m) }
func (*IndexInfo) ProtoMessage()    {}
func (*IndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{0}
}

func (m *IndexInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexInfo.Unmarshal(m, b)
}
func (m *IndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexInfo.Marshal(b, m, deterministic)
}
func (m *IndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexInfo.Merge(m, src)
}
func (m *IndexInfo) XXX_Size() int {
	return xxx_messageInfo_IndexInfo.Size(m)
}
func (m *IndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IndexInfo proto.InternalMessageInfo

func (m *IndexInfo) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *IndexInfo) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *IndexInfo) GetIndexParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.IndexParams
	}
	return nil
}

func (m *IndexInfo) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *IndexInfo) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type FieldIndexInfo struct {
	FiledID              int64    `protobuf:"varint,1,opt,name=filedID,proto3" json:"filedID,omitempty"`
	IndexID              int64    `protobuf:"varint,2,opt,name=indexID,proto3" json:"indexID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldIndexInfo) Reset()         { *m = FieldIndexInfo{} }
func (m *FieldIndexInfo) String() string { return proto.CompactTextString(m) }
func (*FieldIndexInfo) ProtoMessage()    {}
func (*FieldIndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{1}
}

func (m *FieldIndexInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldIndexInfo.Unmarshal(m, b)
}
func (m *FieldIndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldIndexInfo.Marshal(b, m, deterministic)
}
func (m *FieldIndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldIndexInfo.Merge(m, src)
}
func (m *FieldIndexInfo) XXX_Size() int {
	return xxx_messageInfo_FieldIndexInfo.Size(m)
}
func (m *FieldIndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldIndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FieldIndexInfo proto.InternalMessageInfo

func (m *FieldIndexInfo) GetFiledID() int64 {
	if m != nil {
		return m.FiledID
	}
	return 0
}

func (m *FieldIndexInfo) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

type CollectionInfo struct {
	ID         int64                      `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Schema     *schemapb.CollectionSchema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	CreateTime uint64                     `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// deprecate
	PartitionIDs []int64 `protobuf:"varint,4,rep,packed,name=partitionIDs,proto3" json:"partitionIDs,omitempty"`
	// deprecate
	PartitionNames       []string          `protobuf:"bytes,5,rep,name=partitionNames,proto3" json:"partitionNames,omitempty"`
	FieldIndexes         []*FieldIndexInfo `protobuf:"bytes,6,rep,name=field_indexes,json=fieldIndexes,proto3" json:"field_indexes,omitempty"`
	VirtualChannelNames  []string          `protobuf:"bytes,7,rep,name=virtual_channel_names,json=virtualChannelNames,proto3" json:"virtual_channel_names,omitempty"`
	PhysicalChannelNames []string          `protobuf:"bytes,8,rep,name=physical_channel_names,json=physicalChannelNames,proto3" json:"physical_channel_names,omitempty"`
	// deprecate
	PartitionCreatedTimestamps []uint64                  `protobuf:"varint,9,rep,packed,name=partition_created_timestamps,json=partitionCreatedTimestamps,proto3" json:"partition_created_timestamps,omitempty"`
	ShardsNum                  int32                     `protobuf:"varint,10,opt,name=shards_num,json=shardsNum,proto3" json:"shards_num,omitempty"`
	StartPositions             []*commonpb.KeyDataPair   `protobuf:"bytes,11,rep,name=start_positions,json=startPositions,proto3" json:"start_positions,omitempty"`
	ConsistencyLevel           commonpb.ConsistencyLevel `protobuf:"varint,12,opt,name=consistency_level,json=consistencyLevel,proto3,enum=milvus.proto.common.ConsistencyLevel" json:"consistency_level,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                  `json:"-"`
	XXX_unrecognized           []byte                    `json:"-"`
	XXX_sizecache              int32                     `json:"-"`
}

func (m *CollectionInfo) Reset()         { *m = CollectionInfo{} }
func (m *CollectionInfo) String() string { return proto.CompactTextString(m) }
func (*CollectionInfo) ProtoMessage()    {}
func (*CollectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{2}
}

func (m *CollectionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionInfo.Unmarshal(m, b)
}
func (m *CollectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionInfo.Marshal(b, m, deterministic)
}
func (m *CollectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionInfo.Merge(m, src)
}
func (m *CollectionInfo) XXX_Size() int {
	return xxx_messageInfo_CollectionInfo.Size(m)
}
func (m *CollectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionInfo proto.InternalMessageInfo

func (m *CollectionInfo) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *CollectionInfo) GetSchema() *schemapb.CollectionSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CollectionInfo) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CollectionInfo) GetPartitionIDs() []int64 {
	if m != nil {
		return m.PartitionIDs
	}
	return nil
}

func (m *CollectionInfo) GetPartitionNames() []string {
	if m != nil {
		return m.PartitionNames
	}
	return nil
}

func (m *CollectionInfo) GetFieldIndexes() []*FieldIndexInfo {
	if m != nil {
		return m.FieldIndexes
	}
	return nil
}

func (m *CollectionInfo) GetVirtualChannelNames() []string {
	if m != nil {
		return m.VirtualChannelNames
	}
	return nil
}

func (m *CollectionInfo) GetPhysicalChannelNames() []string {
	if m != nil {
		return m.PhysicalChannelNames
	}
	return nil
}

func (m *CollectionInfo) GetPartitionCreatedTimestamps() []uint64 {
	if m != nil {
		return m.PartitionCreatedTimestamps
	}
	return nil
}

func (m *CollectionInfo) GetShardsNum() int32 {
	if m != nil {
		return m.ShardsNum
	}
	return 0
}

func (m *CollectionInfo) GetStartPositions() []*commonpb.KeyDataPair {
	if m != nil {
		return m.StartPositions
	}
	return nil
}

func (m *CollectionInfo) GetConsistencyLevel() commonpb.ConsistencyLevel {
	if m != nil {
		return m.ConsistencyLevel
	}
	return commonpb.ConsistencyLevel_Strong
}

type PartitionInfo struct {
	PartitionID               int64    `protobuf:"varint,1,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	PartitionName             string   `protobuf:"bytes,2,opt,name=partitionName,proto3" json:"partitionName,omitempty"`
	PartitionCreatedTimestamp uint64   `protobuf:"varint,3,opt,name=partition_created_timestamp,json=partitionCreatedTimestamp,proto3" json:"partition_created_timestamp,omitempty"`
	CollectionId              int64    `protobuf:"varint,4,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *PartitionInfo) Reset()         { *m = PartitionInfo{} }
func (m *PartitionInfo) String() string { return proto.CompactTextString(m) }
func (*PartitionInfo) ProtoMessage()    {}
func (*PartitionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{3}
}

func (m *PartitionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionInfo.Unmarshal(m, b)
}
func (m *PartitionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionInfo.Marshal(b, m, deterministic)
}
func (m *PartitionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionInfo.Merge(m, src)
}
func (m *PartitionInfo) XXX_Size() int {
	return xxx_messageInfo_PartitionInfo.Size(m)
}
func (m *PartitionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionInfo proto.InternalMessageInfo

func (m *PartitionInfo) GetPartitionID() int64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *PartitionInfo) GetPartitionName() string {
	if m != nil {
		return m.PartitionName
	}
	return ""
}

func (m *PartitionInfo) GetPartitionCreatedTimestamp() uint64 {
	if m != nil {
		return m.PartitionCreatedTimestamp
	}
	return 0
}

func (m *PartitionInfo) GetCollectionId() int64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

type AliasInfo struct {
	AliasName            string   `protobuf:"bytes,1,opt,name=alias_name,json=aliasName,proto3" json:"alias_name,omitempty"`
	CollectionId         int64    `protobuf:"varint,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	CreatedTime          uint64   `protobuf:"varint,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AliasInfo) Reset()         { *m = AliasInfo{} }
func (m *AliasInfo) String() string { return proto.CompactTextString(m) }
func (*AliasInfo) ProtoMessage()    {}
func (*AliasInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{4}
}

func (m *AliasInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AliasInfo.Unmarshal(m, b)
}
func (m *AliasInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AliasInfo.Marshal(b, m, deterministic)
}
func (m *AliasInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AliasInfo.Merge(m, src)
}
func (m *AliasInfo) XXX_Size() int {
	return xxx_messageInfo_AliasInfo.Size(m)
}
func (m *AliasInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AliasInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AliasInfo proto.InternalMessageInfo

func (m *AliasInfo) GetAliasName() string {
	if m != nil {
		return m.AliasName
	}
	return ""
}

func (m *AliasInfo) GetCollectionId() int64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

func (m *AliasInfo) GetCreatedTime() uint64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

type SegmentIndexInfo struct {
	CollectionID         int64    `protobuf:"varint,1,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	PartitionID          int64    `protobuf:"varint,2,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	SegmentID            int64    `protobuf:"varint,3,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	FieldID              int64    `protobuf:"varint,4,opt,name=fieldID,proto3" json:"fieldID,omitempty"`
	IndexID              int64    `protobuf:"varint,5,opt,name=indexID,proto3" json:"indexID,omitempty"`
	BuildID              int64    `protobuf:"varint,6,opt,name=buildID,proto3" json:"buildID,omitempty"`
	EnableIndex          bool     `protobuf:"varint,7,opt,name=enable_index,json=enableIndex,proto3" json:"enable_index,omitempty"`
	CreateTime           uint64   `protobuf:"varint,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentIndexInfo) Reset()         { *m = SegmentIndexInfo{} }
func (m *SegmentIndexInfo) String() string { return proto.CompactTextString(m) }
func (*SegmentIndexInfo) ProtoMessage()    {}
func (*SegmentIndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{5}
}

func (m *SegmentIndexInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentIndexInfo.Unmarshal(m, b)
}
func (m *SegmentIndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentIndexInfo.Marshal(b, m, deterministic)
}
func (m *SegmentIndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentIndexInfo.Merge(m, src)
}
func (m *SegmentIndexInfo) XXX_Size() int {
	return xxx_messageInfo_SegmentIndexInfo.Size(m)
}
func (m *SegmentIndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentIndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentIndexInfo proto.InternalMessageInfo

func (m *SegmentIndexInfo) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

func (m *SegmentIndexInfo) GetPartitionID() int64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *SegmentIndexInfo) GetSegmentID() int64 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

func (m *SegmentIndexInfo) GetFieldID() int64 {
	if m != nil {
		return m.FieldID
	}
	return 0
}

func (m *SegmentIndexInfo) GetIndexID() int64 {
	if m != nil {
		return m.IndexID
	}
	return 0
}

func (m *SegmentIndexInfo) GetBuildID() int64 {
	if m != nil {
		return m.BuildID
	}
	return 0
}

func (m *SegmentIndexInfo) GetEnableIndex() bool {
	if m != nil {
		return m.EnableIndex
	}
	return false
}

func (m *SegmentIndexInfo) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

// TODO move to proto files of interprocess communication
type CollectionMeta struct {
	ID                   int64                      `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Schema               *schemapb.CollectionSchema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	CreateTime           uint64                     `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	SegmentIDs           []int64                    `protobuf:"varint,4,rep,packed,name=segmentIDs,proto3" json:"segmentIDs,omitempty"`
	PartitionTags        []string                   `protobuf:"bytes,5,rep,name=partition_tags,json=partitionTags,proto3" json:"partition_tags,omitempty"`
	PartitionIDs         []int64                    `protobuf:"varint,6,rep,packed,name=partitionIDs,proto3" json:"partitionIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CollectionMeta) Reset()         { *m = CollectionMeta{} }
func (m *CollectionMeta) String() string { return proto.CompactTextString(m) }
func (*CollectionMeta) ProtoMessage()    {}
func (*CollectionMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{6}
}

func (m *CollectionMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionMeta.Unmarshal(m, b)
}
func (m *CollectionMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionMeta.Marshal(b, m, deterministic)
}
func (m *CollectionMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionMeta.Merge(m, src)
}
func (m *CollectionMeta) XXX_Size() int {
	return xxx_messageInfo_CollectionMeta.Size(m)
}
func (m *CollectionMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionMeta.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionMeta proto.InternalMessageInfo

func (m *CollectionMeta) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *CollectionMeta) GetSchema() *schemapb.CollectionSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CollectionMeta) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CollectionMeta) GetSegmentIDs() []int64 {
	if m != nil {
		return m.SegmentIDs
	}
	return nil
}

func (m *CollectionMeta) GetPartitionTags() []string {
	if m != nil {
		return m.PartitionTags
	}
	return nil
}

func (m *CollectionMeta) GetPartitionIDs() []int64 {
	if m != nil {
		return m.PartitionIDs
	}
	return nil
}

type CredentialInfo struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// encrypted by bcrypt (for higher security level)
	EncryptedPassword string `protobuf:"bytes,2,opt,name=encrypted_password,json=encryptedPassword,proto3" json:"encrypted_password,omitempty"`
	Tenant            string `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	IsSuper           bool   `protobuf:"varint,4,opt,name=is_super,json=isSuper,proto3" json:"is_super,omitempty"`
	// encrypted by sha256 (for good performance in cache mapping)
	Sha256Password       string   `protobuf:"bytes,5,opt,name=sha256_password,json=sha256Password,proto3" json:"sha256_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialInfo) Reset()         { *m = CredentialInfo{} }
func (m *CredentialInfo) String() string { return proto.CompactTextString(m) }
func (*CredentialInfo) ProtoMessage()    {}
func (*CredentialInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{7}
}

func (m *CredentialInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialInfo.Unmarshal(m, b)
}
func (m *CredentialInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialInfo.Marshal(b, m, deterministic)
}
func (m *CredentialInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialInfo.Merge(m, src)
}
func (m *CredentialInfo) XXX_Size() int {
	return xxx_messageInfo_CredentialInfo.Size(m)
}
func (m *CredentialInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialInfo proto.InternalMessageInfo

func (m *CredentialInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CredentialInfo) GetEncryptedPassword() string {
	if m != nil {
		return m.EncryptedPassword
	}
	return ""
}

func (m *CredentialInfo) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *CredentialInfo) GetIsSuper() bool {
	if m != nil {
		return m.IsSuper
	}
	return false
}

func (m *CredentialInfo) GetSha256Password() string {
	if m != nil {
		return m.Sha256Password
	}
	return ""
}

func init() {
	proto.RegisterType((*IndexInfo)(nil), "milvus.proto.etcd.IndexInfo")
	proto.RegisterType((*FieldIndexInfo)(nil), "milvus.proto.etcd.FieldIndexInfo")
	proto.RegisterType((*CollectionInfo)(nil), "milvus.proto.etcd.CollectionInfo")
	proto.RegisterType((*PartitionInfo)(nil), "milvus.proto.etcd.PartitionInfo")
	proto.RegisterType((*AliasInfo)(nil), "milvus.proto.etcd.AliasInfo")
	proto.RegisterType((*SegmentIndexInfo)(nil), "milvus.proto.etcd.SegmentIndexInfo")
	proto.RegisterType((*CollectionMeta)(nil), "milvus.proto.etcd.CollectionMeta")
	proto.RegisterType((*CredentialInfo)(nil), "milvus.proto.etcd.CredentialInfo")
}

func init() { proto.RegisterFile("etcd_meta.proto", fileDescriptor_975d306d62b73e88) }

var fileDescriptor_975d306d62b73e88 = []byte{
	// 861 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0xe3, 0xfc, 0xf9, 0x24, 0x4d, 0xb7, 0x03, 0xac, 0xbc, 0x65, 0x01, 0x6f, 0x60, 0xc1,
	0x37, 0xdb, 0x8a, 0x2e, 0x70, 0x07, 0x02, 0x6a, 0xad, 0x14, 0x01, 0x55, 0x34, 0xad, 0xb8, 0xe0,
	0xc6, 0x9a, 0xd8, 0xa7, 0xcd, 0x48, 0xfe, 0x93, 0x67, 0x5c, 0xe8, 0x1b, 0xf0, 0x46, 0xdc, 0x70,
	0xcb, 0xd3, 0xf0, 0x0e, 0x08, 0xcd, 0x78, 0xec, 0xd8, 0x09, 0xe5, 0x72, 0xef, 0xf2, 0x7d, 0x33,
	0xe7, 0xf8, 0xfc, 0x7c, 0xf3, 0x05, 0x8e, 0x51, 0x46, 0x71, 0x98, 0xa2, 0x64, 0x67, 0x45, 0x99,
	0xcb, 0x9c, 0x9c, 0xa4, 0x3c, 0xb9, 0xaf, 0x44, 0x8d, 0xce, 0xd4, 0xe9, 0xe9, 0x3c, 0xca, 0xd3,
	0x34, 0xcf, 0x6a, 0xea, 0x74, 0x2e, 0xa2, 0x2d, 0xa6, 0xe6, 0xfa, 0xf2, 0x2f, 0x0b, 0x9c, 0x55,
	0x16, 0xe3, 0x6f, 0xab, 0xec, 0x36, 0x27, 0x1f, 0x00, 0x70, 0x05, 0xc2, 0x8c, 0xa5, 0xe8, 0x5a,
	0x9e, 0xe5, 0x3b, 0xd4, 0xd1, 0xcc, 0x15, 0x4b, 0x91, 0xb8, 0x30, 0xd1, 0x60, 0x15, 0xb8, 0x03,
	0xcf, 0xf2, 0x6d, 0xda, 0x40, 0x12, 0xc0, 0xbc, 0x0e, 0x2c, 0x58, 0xc9, 0x52, 0xe1, 0xda, 0x9e,
	0xed, 0xcf, 0x2e, 0x5e, 0x9c, 0xf5, 0x8a, 0x31, 0x65, 0xfc, 0x80, 0x0f, 0x3f, 0xb3, 0xa4, 0xc2,
	0x35, 0xe3, 0x25, 0x9d, 0xe9, 0xb0, 0xb5, 0x8e, 0x52, 0xf9, 0x63, 0x4c, 0x50, 0x62, 0xec, 0x0e,
	0x3d, 0xcb, 0x9f, 0xd2, 0x06, 0x92, 0x8f, 0x60, 0x16, 0x95, 0xc8, 0x24, 0x86, 0x92, 0xa7, 0xe8,
	0x8e, 0x3c, 0xcb, 0x1f, 0x52, 0xa8, 0xa9, 0x1b, 0x9e, 0xe2, 0x32, 0x80, 0xc5, 0x1b, 0x8e, 0x49,
	0xbc, 0xeb, 0xc5, 0x85, 0xc9, 0x2d, 0x4f, 0x30, 0x5e, 0x05, 0xba, 0x11, 0x9b, 0x36, 0xf0, 0xf1,
	0x36, 0x96, 0xff, 0x0c, 0x61, 0x71, 0x99, 0x27, 0x09, 0x46, 0x92, 0xe7, 0x99, 0x4e, 0xb3, 0x80,
	0x41, 0x9b, 0x61, 0xb0, 0x0a, 0xc8, 0xd7, 0x30, 0xae, 0x07, 0xa8, 0x63, 0x67, 0x17, 0x2f, 0xfb,
	0x3d, 0x9a, 0xe1, 0xee, 0x92, 0x5c, 0x6b, 0x82, 0x9a, 0xa0, 0xfd, 0x46, 0xec, 0xfd, 0x46, 0xc8,
	0x12, 0xe6, 0x05, 0x2b, 0x25, 0xd7, 0x05, 0x04, 0xc2, 0x1d, 0x7a, 0xb6, 0x6f, 0xd3, 0x1e, 0x47,
	0x3e, 0x85, 0x45, 0x8b, 0xd5, 0x62, 0x84, 0x3b, 0xf2, 0x6c, 0xdf, 0xa1, 0x7b, 0x2c, 0x79, 0x03,
	0x47, 0xb7, 0x6a, 0x28, 0xa1, 0xee, 0x0f, 0x85, 0x3b, 0xfe, 0xaf, 0xb5, 0x28, 0x8d, 0x9c, 0xf5,
	0x87, 0x47, 0xe7, 0xb7, 0x2d, 0x46, 0x41, 0x2e, 0xe0, 0xbd, 0x7b, 0x5e, 0xca, 0x8a, 0x25, 0x61,
	0xb4, 0x65, 0x59, 0x86, 0x89, 0x16, 0x88, 0x70, 0x27, 0xfa, 0xb3, 0xef, 0x98, 0xc3, 0xcb, 0xfa,
	0xac, 0xfe, 0xf6, 0x17, 0xf0, 0xb4, 0xd8, 0x3e, 0x08, 0x1e, 0x1d, 0x04, 0x4d, 0x75, 0xd0, 0xbb,
	0xcd, 0x69, 0x2f, 0xea, 0x5b, 0x78, 0xde, 0xf6, 0x10, 0xd6, 0x53, 0x89, 0xf5, 0xa4, 0x84, 0x64,
	0x69, 0x21, 0x5c, 0xc7, 0xb3, 0xfd, 0x21, 0x3d, 0x6d, 0xef, 0x5c, 0xd6, 0x57, 0x6e, 0xda, 0x1b,
	0x4a, 0xc2, 0x62, 0xcb, 0xca, 0x58, 0x84, 0x59, 0x95, 0xba, 0xe0, 0x59, 0xfe, 0x88, 0x3a, 0x35,
	0x73, 0x55, 0xa5, 0x64, 0x05, 0xc7, 0x42, 0xb2, 0x52, 0x86, 0x45, 0x2e, 0x74, 0x06, 0xe1, 0xce,
	0xf4, 0x50, 0xbc, 0xc7, 0xb4, 0x1a, 0x30, 0xc9, 0xb4, 0x54, 0x17, 0x3a, 0x70, 0xdd, 0xc4, 0x11,
	0x0a, 0x27, 0x51, 0x9e, 0x09, 0x2e, 0x24, 0x66, 0xd1, 0x43, 0x98, 0xe0, 0x3d, 0x26, 0xee, 0xdc,
	0xb3, 0xfc, 0xc5, 0xbe, 0x28, 0x4c, 0xb2, 0xcb, 0xdd, 0xed, 0x1f, 0xd5, 0x65, 0xfa, 0x24, 0xda,
	0x63, 0x96, 0x7f, 0x5a, 0x70, 0xb4, 0x6e, 0x57, 0xad, 0xf4, 0xe7, 0xc1, 0xac, 0xb3, 0x7b, 0x23,
	0xc4, 0x2e, 0x45, 0x3e, 0x81, 0xa3, 0xde, 0xde, 0xb5, 0x30, 0x1d, 0xda, 0x27, 0xc9, 0x37, 0xf0,
	0xfe, 0xff, 0x4c, 0xd6, 0x08, 0xf1, 0xd9, 0xa3, 0x83, 0x25, 0x1f, 0xc3, 0x51, 0xd4, 0x8a, 0x3a,
	0xe4, 0xf5, 0x0b, 0xb5, 0xe9, 0x7c, 0x47, 0xae, 0xe2, 0x65, 0x09, 0xce, 0x77, 0x09, 0x67, 0xa2,
	0x31, 0x13, 0xa6, 0x40, 0xcf, 0x4c, 0x34, 0xa3, 0x0b, 0x3a, 0x48, 0x38, 0x38, 0x4c, 0x48, 0x5e,
	0xc0, 0xbc, 0x5b, 0xab, 0x29, 0xd3, 0x3c, 0x21, 0x5d, 0xdd, 0xf2, 0xf7, 0x01, 0x3c, 0xb9, 0xc6,
	0xbb, 0x14, 0x33, 0xb9, 0x7b, 0xfc, 0x4b, 0xe8, 0xe6, 0x69, 0xc6, 0xd6, 0xe3, 0xf6, 0x27, 0x3b,
	0x38, 0x9c, 0xec, 0x73, 0x70, 0x84, 0xc9, 0x1c, 0xe8, 0x4f, 0xdb, 0x74, 0x47, 0xd4, 0x06, 0xa3,
	0x5e, 0x49, 0x60, 0x66, 0xd1, 0xc0, 0xae, 0xc1, 0x8c, 0xfa, 0x3e, 0xe9, 0xc2, 0x64, 0x53, 0x71,
	0x1d, 0x33, 0xae, 0x4f, 0x0c, 0x54, 0x9d, 0x62, 0xc6, 0x36, 0x09, 0xd6, 0x8f, 0xd5, 0x9d, 0x68,
	0x03, 0x9c, 0xd5, 0x9c, 0x6e, 0x6c, 0xdf, 0x3b, 0xa6, 0x07, 0x26, 0xf8, 0xb7, 0xd5, 0xb5, 0xaf,
	0x9f, 0x50, 0xb2, 0xb7, 0x6e, 0x5f, 0x1f, 0x02, 0xb4, 0x13, 0x6a, 0xcc, 0xab, 0xc3, 0x90, 0x97,
	0x1d, 0xeb, 0x0a, 0x25, 0xbb, 0x6b, 0xac, 0x6b, 0xa7, 0xd6, 0x1b, 0x76, 0x27, 0x0e, 0x5c, 0x70,
	0x7c, 0xe8, 0x82, 0xcb, 0x3f, 0x54, 0xb7, 0x25, 0xc6, 0x98, 0x49, 0xce, 0x12, 0xbd, 0xf6, 0x53,
	0x98, 0x56, 0x02, 0xcb, 0x8e, 0xe0, 0x5a, 0x4c, 0x5e, 0x01, 0xc1, 0x2c, 0x2a, 0x1f, 0x0a, 0x25,
	0xa6, 0x82, 0x09, 0xf1, 0x6b, 0x5e, 0xc6, 0xe6, 0xad, 0x9c, 0xb4, 0x27, 0x6b, 0x73, 0x40, 0x9e,
	0xc2, 0x58, 0x62, 0xc6, 0x32, 0xa9, 0x9b, 0x74, 0xa8, 0x41, 0xe4, 0x19, 0x4c, 0xb9, 0x08, 0x45,
	0x55, 0x60, 0xd9, 0xfc, 0x49, 0x71, 0x71, 0xad, 0x20, 0xf9, 0x0c, 0x8e, 0xc5, 0x96, 0x5d, 0x7c,
	0xf9, 0xd5, 0x2e, 0xfd, 0x48, 0xc7, 0x2e, 0x6a, 0xba, 0xc9, 0xfd, 0xfd, 0xeb, 0x5f, 0x3e, 0xbf,
	0xe3, 0x72, 0x5b, 0x6d, 0x94, 0x33, 0x9c, 0xd7, 0x0b, 0x78, 0xc5, 0x73, 0xf3, 0xeb, 0x9c, 0x67,
	0x52, 0xd5, 0x9c, 0x9c, 0xeb, 0x9d, 0x9c, 0x2b, 0x7f, 0x2e, 0x36, 0x9b, 0xb1, 0x46, 0xaf, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x28, 0x7e, 0xad, 0xf2, 0x07, 0x00, 0x00,
}
