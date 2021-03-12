// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: schema.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_schema_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_schema_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3009000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3009000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
#include "common.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_schema_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_schema_2eproto {
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::FieldMetadata field_metadata[];
  static const ::PROTOBUF_NAMESPACE_ID::internal::SerializationTable serialization_table[];
  static const ::PROTOBUF_NAMESPACE_ID::uint32 offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_schema_2eproto;
namespace milvus {
namespace proto {
namespace schema {
class CollectionSchema;
class CollectionSchemaDefaultTypeInternal;
extern CollectionSchemaDefaultTypeInternal _CollectionSchema_default_instance_;
class FieldSchema;
class FieldSchemaDefaultTypeInternal;
extern FieldSchemaDefaultTypeInternal _FieldSchema_default_instance_;
}  // namespace schema
}  // namespace proto
}  // namespace milvus
PROTOBUF_NAMESPACE_OPEN
template<> ::milvus::proto::schema::CollectionSchema* Arena::CreateMaybeMessage<::milvus::proto::schema::CollectionSchema>(Arena*);
template<> ::milvus::proto::schema::FieldSchema* Arena::CreateMaybeMessage<::milvus::proto::schema::FieldSchema>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace milvus {
namespace proto {
namespace schema {

enum DataType : int {
  None = 0,
  Bool = 1,
  Int8 = 2,
  Int16 = 3,
  Int32 = 4,
  Int64 = 5,
  Float = 10,
  Double = 11,
  String = 20,
  BinaryVector = 100,
  FloatVector = 101,
  DataType_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::PROTOBUF_NAMESPACE_ID::int32>::min(),
  DataType_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::PROTOBUF_NAMESPACE_ID::int32>::max()
};
bool DataType_IsValid(int value);
constexpr DataType DataType_MIN = None;
constexpr DataType DataType_MAX = FloatVector;
constexpr int DataType_ARRAYSIZE = DataType_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* DataType_descriptor();
template<typename T>
inline const std::string& DataType_Name(T enum_t_value) {
  static_assert(::std::is_same<T, DataType>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function DataType_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    DataType_descriptor(), enum_t_value);
}
inline bool DataType_Parse(
    const std::string& name, DataType* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<DataType>(
    DataType_descriptor(), name, value);
}
// ===================================================================

class FieldSchema :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:milvus.proto.schema.FieldSchema) */ {
 public:
  FieldSchema();
  virtual ~FieldSchema();

  FieldSchema(const FieldSchema& from);
  FieldSchema(FieldSchema&& from) noexcept
    : FieldSchema() {
    *this = ::std::move(from);
  }

  inline FieldSchema& operator=(const FieldSchema& from) {
    CopyFrom(from);
    return *this;
  }
  inline FieldSchema& operator=(FieldSchema&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return GetMetadataStatic().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return GetMetadataStatic().reflection;
  }
  static const FieldSchema& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const FieldSchema* internal_default_instance() {
    return reinterpret_cast<const FieldSchema*>(
               &_FieldSchema_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(FieldSchema& a, FieldSchema& b) {
    a.Swap(&b);
  }
  inline void Swap(FieldSchema* other) {
    if (other == this) return;
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  inline FieldSchema* New() const final {
    return CreateMaybeMessage<FieldSchema>(nullptr);
  }

  FieldSchema* New(::PROTOBUF_NAMESPACE_ID::Arena* arena) const final {
    return CreateMaybeMessage<FieldSchema>(arena);
  }
  void CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void CopyFrom(const FieldSchema& from);
  void MergeFrom(const FieldSchema& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  #else
  bool MergePartialFromCodedStream(
      ::PROTOBUF_NAMESPACE_ID::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::PROTOBUF_NAMESPACE_ID::io::CodedOutputStream* output) const final;
  ::PROTOBUF_NAMESPACE_ID::uint8* InternalSerializeWithCachedSizesToArray(
      ::PROTOBUF_NAMESPACE_ID::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  inline void SharedCtor();
  inline void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(FieldSchema* other);
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "milvus.proto.schema.FieldSchema";
  }
  private:
  inline ::PROTOBUF_NAMESPACE_ID::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;
  private:
  static ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadataStatic() {
    ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&::descriptor_table_schema_2eproto);
    return ::descriptor_table_schema_2eproto.file_level_metadata[kIndexInFileMessages];
  }

  public:

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kTypeParamsFieldNumber = 6,
    kIndexParamsFieldNumber = 7,
    kNameFieldNumber = 2,
    kDescriptionFieldNumber = 4,
    kFieldIDFieldNumber = 1,
    kIsPrimaryKeyFieldNumber = 3,
    kDataTypeFieldNumber = 5,
  };
  // repeated .milvus.proto.common.KeyValuePair type_params = 6;
  int type_params_size() const;
  void clear_type_params();
  ::milvus::proto::common::KeyValuePair* mutable_type_params(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair >*
      mutable_type_params();
  const ::milvus::proto::common::KeyValuePair& type_params(int index) const;
  ::milvus::proto::common::KeyValuePair* add_type_params();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair >&
      type_params() const;

  // repeated .milvus.proto.common.KeyValuePair index_params = 7;
  int index_params_size() const;
  void clear_index_params();
  ::milvus::proto::common::KeyValuePair* mutable_index_params(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair >*
      mutable_index_params();
  const ::milvus::proto::common::KeyValuePair& index_params(int index) const;
  ::milvus::proto::common::KeyValuePair* add_index_params();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair >&
      index_params() const;

  // string name = 2;
  void clear_name();
  const std::string& name() const;
  void set_name(const std::string& value);
  void set_name(std::string&& value);
  void set_name(const char* value);
  void set_name(const char* value, size_t size);
  std::string* mutable_name();
  std::string* release_name();
  void set_allocated_name(std::string* name);

  // string description = 4;
  void clear_description();
  const std::string& description() const;
  void set_description(const std::string& value);
  void set_description(std::string&& value);
  void set_description(const char* value);
  void set_description(const char* value, size_t size);
  std::string* mutable_description();
  std::string* release_description();
  void set_allocated_description(std::string* description);

  // int64 fieldID = 1;
  void clear_fieldid();
  ::PROTOBUF_NAMESPACE_ID::int64 fieldid() const;
  void set_fieldid(::PROTOBUF_NAMESPACE_ID::int64 value);

  // bool is_primary_key = 3;
  void clear_is_primary_key();
  bool is_primary_key() const;
  void set_is_primary_key(bool value);

  // .milvus.proto.schema.DataType data_type = 5;
  void clear_data_type();
  ::milvus::proto::schema::DataType data_type() const;
  void set_data_type(::milvus::proto::schema::DataType value);

  // @@protoc_insertion_point(class_scope:milvus.proto.schema.FieldSchema)
 private:
  class _Internal;

  ::PROTOBUF_NAMESPACE_ID::internal::InternalMetadataWithArena _internal_metadata_;
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair > type_params_;
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair > index_params_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr name_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr description_;
  ::PROTOBUF_NAMESPACE_ID::int64 fieldid_;
  bool is_primary_key_;
  int data_type_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_schema_2eproto;
};
// -------------------------------------------------------------------

class CollectionSchema :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:milvus.proto.schema.CollectionSchema) */ {
 public:
  CollectionSchema();
  virtual ~CollectionSchema();

  CollectionSchema(const CollectionSchema& from);
  CollectionSchema(CollectionSchema&& from) noexcept
    : CollectionSchema() {
    *this = ::std::move(from);
  }

  inline CollectionSchema& operator=(const CollectionSchema& from) {
    CopyFrom(from);
    return *this;
  }
  inline CollectionSchema& operator=(CollectionSchema&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return GetMetadataStatic().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return GetMetadataStatic().reflection;
  }
  static const CollectionSchema& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const CollectionSchema* internal_default_instance() {
    return reinterpret_cast<const CollectionSchema*>(
               &_CollectionSchema_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(CollectionSchema& a, CollectionSchema& b) {
    a.Swap(&b);
  }
  inline void Swap(CollectionSchema* other) {
    if (other == this) return;
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  inline CollectionSchema* New() const final {
    return CreateMaybeMessage<CollectionSchema>(nullptr);
  }

  CollectionSchema* New(::PROTOBUF_NAMESPACE_ID::Arena* arena) const final {
    return CreateMaybeMessage<CollectionSchema>(arena);
  }
  void CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void CopyFrom(const CollectionSchema& from);
  void MergeFrom(const CollectionSchema& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  #else
  bool MergePartialFromCodedStream(
      ::PROTOBUF_NAMESPACE_ID::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::PROTOBUF_NAMESPACE_ID::io::CodedOutputStream* output) const final;
  ::PROTOBUF_NAMESPACE_ID::uint8* InternalSerializeWithCachedSizesToArray(
      ::PROTOBUF_NAMESPACE_ID::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  inline void SharedCtor();
  inline void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(CollectionSchema* other);
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "milvus.proto.schema.CollectionSchema";
  }
  private:
  inline ::PROTOBUF_NAMESPACE_ID::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;
  private:
  static ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadataStatic() {
    ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&::descriptor_table_schema_2eproto);
    return ::descriptor_table_schema_2eproto.file_level_metadata[kIndexInFileMessages];
  }

  public:

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kFieldsFieldNumber = 4,
    kNameFieldNumber = 1,
    kDescriptionFieldNumber = 2,
    kAutoIDFieldNumber = 3,
  };
  // repeated .milvus.proto.schema.FieldSchema fields = 4;
  int fields_size() const;
  void clear_fields();
  ::milvus::proto::schema::FieldSchema* mutable_fields(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::schema::FieldSchema >*
      mutable_fields();
  const ::milvus::proto::schema::FieldSchema& fields(int index) const;
  ::milvus::proto::schema::FieldSchema* add_fields();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::schema::FieldSchema >&
      fields() const;

  // string name = 1;
  void clear_name();
  const std::string& name() const;
  void set_name(const std::string& value);
  void set_name(std::string&& value);
  void set_name(const char* value);
  void set_name(const char* value, size_t size);
  std::string* mutable_name();
  std::string* release_name();
  void set_allocated_name(std::string* name);

  // string description = 2;
  void clear_description();
  const std::string& description() const;
  void set_description(const std::string& value);
  void set_description(std::string&& value);
  void set_description(const char* value);
  void set_description(const char* value, size_t size);
  std::string* mutable_description();
  std::string* release_description();
  void set_allocated_description(std::string* description);

  // bool autoID = 3;
  void clear_autoid();
  bool autoid() const;
  void set_autoid(bool value);

  // @@protoc_insertion_point(class_scope:milvus.proto.schema.CollectionSchema)
 private:
  class _Internal;

  ::PROTOBUF_NAMESPACE_ID::internal::InternalMetadataWithArena _internal_metadata_;
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::schema::FieldSchema > fields_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr name_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr description_;
  bool autoid_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_schema_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// FieldSchema

// int64 fieldID = 1;
inline void FieldSchema::clear_fieldid() {
  fieldid_ = PROTOBUF_LONGLONG(0);
}
inline ::PROTOBUF_NAMESPACE_ID::int64 FieldSchema::fieldid() const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.FieldSchema.fieldID)
  return fieldid_;
}
inline void FieldSchema::set_fieldid(::PROTOBUF_NAMESPACE_ID::int64 value) {
  
  fieldid_ = value;
  // @@protoc_insertion_point(field_set:milvus.proto.schema.FieldSchema.fieldID)
}

// string name = 2;
inline void FieldSchema::clear_name() {
  name_.ClearToEmptyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline const std::string& FieldSchema::name() const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.FieldSchema.name)
  return name_.GetNoArena();
}
inline void FieldSchema::set_name(const std::string& value) {
  
  name_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:milvus.proto.schema.FieldSchema.name)
}
inline void FieldSchema::set_name(std::string&& value) {
  
  name_.SetNoArena(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:milvus.proto.schema.FieldSchema.name)
}
inline void FieldSchema::set_name(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  name_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:milvus.proto.schema.FieldSchema.name)
}
inline void FieldSchema::set_name(const char* value, size_t size) {
  
  name_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:milvus.proto.schema.FieldSchema.name)
}
inline std::string* FieldSchema::mutable_name() {
  
  // @@protoc_insertion_point(field_mutable:milvus.proto.schema.FieldSchema.name)
  return name_.MutableNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline std::string* FieldSchema::release_name() {
  // @@protoc_insertion_point(field_release:milvus.proto.schema.FieldSchema.name)
  
  return name_.ReleaseNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline void FieldSchema::set_allocated_name(std::string* name) {
  if (name != nullptr) {
    
  } else {
    
  }
  name_.SetAllocatedNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), name);
  // @@protoc_insertion_point(field_set_allocated:milvus.proto.schema.FieldSchema.name)
}

// bool is_primary_key = 3;
inline void FieldSchema::clear_is_primary_key() {
  is_primary_key_ = false;
}
inline bool FieldSchema::is_primary_key() const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.FieldSchema.is_primary_key)
  return is_primary_key_;
}
inline void FieldSchema::set_is_primary_key(bool value) {
  
  is_primary_key_ = value;
  // @@protoc_insertion_point(field_set:milvus.proto.schema.FieldSchema.is_primary_key)
}

// string description = 4;
inline void FieldSchema::clear_description() {
  description_.ClearToEmptyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline const std::string& FieldSchema::description() const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.FieldSchema.description)
  return description_.GetNoArena();
}
inline void FieldSchema::set_description(const std::string& value) {
  
  description_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:milvus.proto.schema.FieldSchema.description)
}
inline void FieldSchema::set_description(std::string&& value) {
  
  description_.SetNoArena(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:milvus.proto.schema.FieldSchema.description)
}
inline void FieldSchema::set_description(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  description_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:milvus.proto.schema.FieldSchema.description)
}
inline void FieldSchema::set_description(const char* value, size_t size) {
  
  description_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:milvus.proto.schema.FieldSchema.description)
}
inline std::string* FieldSchema::mutable_description() {
  
  // @@protoc_insertion_point(field_mutable:milvus.proto.schema.FieldSchema.description)
  return description_.MutableNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline std::string* FieldSchema::release_description() {
  // @@protoc_insertion_point(field_release:milvus.proto.schema.FieldSchema.description)
  
  return description_.ReleaseNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline void FieldSchema::set_allocated_description(std::string* description) {
  if (description != nullptr) {
    
  } else {
    
  }
  description_.SetAllocatedNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), description);
  // @@protoc_insertion_point(field_set_allocated:milvus.proto.schema.FieldSchema.description)
}

// .milvus.proto.schema.DataType data_type = 5;
inline void FieldSchema::clear_data_type() {
  data_type_ = 0;
}
inline ::milvus::proto::schema::DataType FieldSchema::data_type() const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.FieldSchema.data_type)
  return static_cast< ::milvus::proto::schema::DataType >(data_type_);
}
inline void FieldSchema::set_data_type(::milvus::proto::schema::DataType value) {
  
  data_type_ = value;
  // @@protoc_insertion_point(field_set:milvus.proto.schema.FieldSchema.data_type)
}

// repeated .milvus.proto.common.KeyValuePair type_params = 6;
inline int FieldSchema::type_params_size() const {
  return type_params_.size();
}
inline ::milvus::proto::common::KeyValuePair* FieldSchema::mutable_type_params(int index) {
  // @@protoc_insertion_point(field_mutable:milvus.proto.schema.FieldSchema.type_params)
  return type_params_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair >*
FieldSchema::mutable_type_params() {
  // @@protoc_insertion_point(field_mutable_list:milvus.proto.schema.FieldSchema.type_params)
  return &type_params_;
}
inline const ::milvus::proto::common::KeyValuePair& FieldSchema::type_params(int index) const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.FieldSchema.type_params)
  return type_params_.Get(index);
}
inline ::milvus::proto::common::KeyValuePair* FieldSchema::add_type_params() {
  // @@protoc_insertion_point(field_add:milvus.proto.schema.FieldSchema.type_params)
  return type_params_.Add();
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair >&
FieldSchema::type_params() const {
  // @@protoc_insertion_point(field_list:milvus.proto.schema.FieldSchema.type_params)
  return type_params_;
}

// repeated .milvus.proto.common.KeyValuePair index_params = 7;
inline int FieldSchema::index_params_size() const {
  return index_params_.size();
}
inline ::milvus::proto::common::KeyValuePair* FieldSchema::mutable_index_params(int index) {
  // @@protoc_insertion_point(field_mutable:milvus.proto.schema.FieldSchema.index_params)
  return index_params_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair >*
FieldSchema::mutable_index_params() {
  // @@protoc_insertion_point(field_mutable_list:milvus.proto.schema.FieldSchema.index_params)
  return &index_params_;
}
inline const ::milvus::proto::common::KeyValuePair& FieldSchema::index_params(int index) const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.FieldSchema.index_params)
  return index_params_.Get(index);
}
inline ::milvus::proto::common::KeyValuePair* FieldSchema::add_index_params() {
  // @@protoc_insertion_point(field_add:milvus.proto.schema.FieldSchema.index_params)
  return index_params_.Add();
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::common::KeyValuePair >&
FieldSchema::index_params() const {
  // @@protoc_insertion_point(field_list:milvus.proto.schema.FieldSchema.index_params)
  return index_params_;
}

// -------------------------------------------------------------------

// CollectionSchema

// string name = 1;
inline void CollectionSchema::clear_name() {
  name_.ClearToEmptyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline const std::string& CollectionSchema::name() const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.CollectionSchema.name)
  return name_.GetNoArena();
}
inline void CollectionSchema::set_name(const std::string& value) {
  
  name_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:milvus.proto.schema.CollectionSchema.name)
}
inline void CollectionSchema::set_name(std::string&& value) {
  
  name_.SetNoArena(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:milvus.proto.schema.CollectionSchema.name)
}
inline void CollectionSchema::set_name(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  name_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:milvus.proto.schema.CollectionSchema.name)
}
inline void CollectionSchema::set_name(const char* value, size_t size) {
  
  name_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:milvus.proto.schema.CollectionSchema.name)
}
inline std::string* CollectionSchema::mutable_name() {
  
  // @@protoc_insertion_point(field_mutable:milvus.proto.schema.CollectionSchema.name)
  return name_.MutableNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline std::string* CollectionSchema::release_name() {
  // @@protoc_insertion_point(field_release:milvus.proto.schema.CollectionSchema.name)
  
  return name_.ReleaseNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline void CollectionSchema::set_allocated_name(std::string* name) {
  if (name != nullptr) {
    
  } else {
    
  }
  name_.SetAllocatedNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), name);
  // @@protoc_insertion_point(field_set_allocated:milvus.proto.schema.CollectionSchema.name)
}

// string description = 2;
inline void CollectionSchema::clear_description() {
  description_.ClearToEmptyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline const std::string& CollectionSchema::description() const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.CollectionSchema.description)
  return description_.GetNoArena();
}
inline void CollectionSchema::set_description(const std::string& value) {
  
  description_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:milvus.proto.schema.CollectionSchema.description)
}
inline void CollectionSchema::set_description(std::string&& value) {
  
  description_.SetNoArena(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:milvus.proto.schema.CollectionSchema.description)
}
inline void CollectionSchema::set_description(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  description_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:milvus.proto.schema.CollectionSchema.description)
}
inline void CollectionSchema::set_description(const char* value, size_t size) {
  
  description_.SetNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:milvus.proto.schema.CollectionSchema.description)
}
inline std::string* CollectionSchema::mutable_description() {
  
  // @@protoc_insertion_point(field_mutable:milvus.proto.schema.CollectionSchema.description)
  return description_.MutableNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline std::string* CollectionSchema::release_description() {
  // @@protoc_insertion_point(field_release:milvus.proto.schema.CollectionSchema.description)
  
  return description_.ReleaseNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}
inline void CollectionSchema::set_allocated_description(std::string* description) {
  if (description != nullptr) {
    
  } else {
    
  }
  description_.SetAllocatedNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), description);
  // @@protoc_insertion_point(field_set_allocated:milvus.proto.schema.CollectionSchema.description)
}

// bool autoID = 3;
inline void CollectionSchema::clear_autoid() {
  autoid_ = false;
}
inline bool CollectionSchema::autoid() const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.CollectionSchema.autoID)
  return autoid_;
}
inline void CollectionSchema::set_autoid(bool value) {
  
  autoid_ = value;
  // @@protoc_insertion_point(field_set:milvus.proto.schema.CollectionSchema.autoID)
}

// repeated .milvus.proto.schema.FieldSchema fields = 4;
inline int CollectionSchema::fields_size() const {
  return fields_.size();
}
inline void CollectionSchema::clear_fields() {
  fields_.Clear();
}
inline ::milvus::proto::schema::FieldSchema* CollectionSchema::mutable_fields(int index) {
  // @@protoc_insertion_point(field_mutable:milvus.proto.schema.CollectionSchema.fields)
  return fields_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::schema::FieldSchema >*
CollectionSchema::mutable_fields() {
  // @@protoc_insertion_point(field_mutable_list:milvus.proto.schema.CollectionSchema.fields)
  return &fields_;
}
inline const ::milvus::proto::schema::FieldSchema& CollectionSchema::fields(int index) const {
  // @@protoc_insertion_point(field_get:milvus.proto.schema.CollectionSchema.fields)
  return fields_.Get(index);
}
inline ::milvus::proto::schema::FieldSchema* CollectionSchema::add_fields() {
  // @@protoc_insertion_point(field_add:milvus.proto.schema.CollectionSchema.fields)
  return fields_.Add();
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::milvus::proto::schema::FieldSchema >&
CollectionSchema::fields() const {
  // @@protoc_insertion_point(field_list:milvus.proto.schema.CollectionSchema.fields)
  return fields_;
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace schema
}  // namespace proto
}  // namespace milvus

PROTOBUF_NAMESPACE_OPEN

template <> struct is_proto_enum< ::milvus::proto::schema::DataType> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::milvus::proto::schema::DataType>() {
  return ::milvus::proto::schema::DataType_descriptor();
}

PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_schema_2eproto
