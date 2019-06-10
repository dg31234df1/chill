/**
 * Autogenerated by Thrift Compiler (0.12.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef megasearch_TYPES_H
#define megasearch_TYPES_H

#include <iosfwd>

#include <thrift/Thrift.h>
#include <thrift/TApplicationException.h>
#include <thrift/TBase.h>
#include <thrift/protocol/TProtocol.h>
#include <thrift/transport/TTransport.h>

#include <thrift/stdcxx.h>


namespace megasearch { namespace thrift {

struct ErrorCode {
  enum type {
    SUCCESS = 0,
    CONNECT_FAILED = 1,
    PERMISSION_DENIED = 2,
    TABLE_NOT_EXISTS = 3,
    ILLEGAL_ARGUMENT = 4,
    ILLEGAL_RANGE = 5,
    ILLEGAL_DIMENSION = 6
  };
};

extern const std::map<int, const char*> _ErrorCode_VALUES_TO_NAMES;

std::ostream& operator<<(std::ostream& out, const ErrorCode::type& val);

class Exception;

class TableSchema;

class Range;

class RowRecord;

class QueryResult;

class TopKQueryResult;

typedef struct _Exception__isset {
  _Exception__isset() : code(false), reason(false) {}
  bool code :1;
  bool reason :1;
} _Exception__isset;

class Exception : public ::apache::thrift::TException {
 public:

  Exception(const Exception&);
  Exception& operator=(const Exception&);
  Exception() : code((ErrorCode::type)0), reason() {
  }

  virtual ~Exception() throw();
  ErrorCode::type code;
  std::string reason;

  _Exception__isset __isset;

  void __set_code(const ErrorCode::type val);

  void __set_reason(const std::string& val);

  bool operator == (const Exception & rhs) const
  {
    if (!(code == rhs.code))
      return false;
    if (!(reason == rhs.reason))
      return false;
    return true;
  }
  bool operator != (const Exception &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Exception & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
  mutable std::string thriftTExceptionMessageHolder_;
  const char* what() const throw();
};

void swap(Exception &a, Exception &b);

std::ostream& operator<<(std::ostream& out, const Exception& obj);

typedef struct _TableSchema__isset {
  _TableSchema__isset() : index_type(true), dimension(true), store_raw_vector(true) {}
  bool index_type :1;
  bool dimension :1;
  bool store_raw_vector :1;
} _TableSchema__isset;

class TableSchema : public virtual ::apache::thrift::TBase {
 public:

  TableSchema(const TableSchema&);
  TableSchema& operator=(const TableSchema&);
  TableSchema() : table_name(), index_type(0), dimension(0LL), store_raw_vector(false) {
  }

  virtual ~TableSchema() throw();
  std::string table_name;
  int32_t index_type;
  int64_t dimension;
  bool store_raw_vector;

  _TableSchema__isset __isset;

  void __set_table_name(const std::string& val);

  void __set_index_type(const int32_t val);

  void __set_dimension(const int64_t val);

  void __set_store_raw_vector(const bool val);

  bool operator == (const TableSchema & rhs) const
  {
    if (!(table_name == rhs.table_name))
      return false;
    if (!(index_type == rhs.index_type))
      return false;
    if (!(dimension == rhs.dimension))
      return false;
    if (!(store_raw_vector == rhs.store_raw_vector))
      return false;
    return true;
  }
  bool operator != (const TableSchema &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TableSchema & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TableSchema &a, TableSchema &b);

std::ostream& operator<<(std::ostream& out, const TableSchema& obj);

typedef struct _Range__isset {
  _Range__isset() : start_value(false), end_value(false) {}
  bool start_value :1;
  bool end_value :1;
} _Range__isset;

class Range : public virtual ::apache::thrift::TBase {
 public:

  Range(const Range&);
  Range& operator=(const Range&);
  Range() : start_value(), end_value() {
  }

  virtual ~Range() throw();
  std::string start_value;
  std::string end_value;

  _Range__isset __isset;

  void __set_start_value(const std::string& val);

  void __set_end_value(const std::string& val);

  bool operator == (const Range & rhs) const
  {
    if (!(start_value == rhs.start_value))
      return false;
    if (!(end_value == rhs.end_value))
      return false;
    return true;
  }
  bool operator != (const Range &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const Range & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(Range &a, Range &b);

std::ostream& operator<<(std::ostream& out, const Range& obj);


class RowRecord : public virtual ::apache::thrift::TBase {
 public:

  RowRecord(const RowRecord&);
  RowRecord& operator=(const RowRecord&);
  RowRecord() : vector_data() {
  }

  virtual ~RowRecord() throw();
  std::string vector_data;

  void __set_vector_data(const std::string& val);

  bool operator == (const RowRecord & rhs) const
  {
    if (!(vector_data == rhs.vector_data))
      return false;
    return true;
  }
  bool operator != (const RowRecord &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const RowRecord & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(RowRecord &a, RowRecord &b);

std::ostream& operator<<(std::ostream& out, const RowRecord& obj);

typedef struct _QueryResult__isset {
  _QueryResult__isset() : id(false), score(false) {}
  bool id :1;
  bool score :1;
} _QueryResult__isset;

class QueryResult : public virtual ::apache::thrift::TBase {
 public:

  QueryResult(const QueryResult&);
  QueryResult& operator=(const QueryResult&);
  QueryResult() : id(0), score(0) {
  }

  virtual ~QueryResult() throw();
  int64_t id;
  double score;

  _QueryResult__isset __isset;

  void __set_id(const int64_t val);

  void __set_score(const double val);

  bool operator == (const QueryResult & rhs) const
  {
    if (!(id == rhs.id))
      return false;
    if (!(score == rhs.score))
      return false;
    return true;
  }
  bool operator != (const QueryResult &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const QueryResult & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(QueryResult &a, QueryResult &b);

std::ostream& operator<<(std::ostream& out, const QueryResult& obj);

typedef struct _TopKQueryResult__isset {
  _TopKQueryResult__isset() : query_result_arrays(false) {}
  bool query_result_arrays :1;
} _TopKQueryResult__isset;

class TopKQueryResult : public virtual ::apache::thrift::TBase {
 public:

  TopKQueryResult(const TopKQueryResult&);
  TopKQueryResult& operator=(const TopKQueryResult&);
  TopKQueryResult() {
  }

  virtual ~TopKQueryResult() throw();
  std::vector<QueryResult>  query_result_arrays;

  _TopKQueryResult__isset __isset;

  void __set_query_result_arrays(const std::vector<QueryResult> & val);

  bool operator == (const TopKQueryResult & rhs) const
  {
    if (!(query_result_arrays == rhs.query_result_arrays))
      return false;
    return true;
  }
  bool operator != (const TopKQueryResult &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TopKQueryResult & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TopKQueryResult &a, TopKQueryResult &b);

std::ostream& operator<<(std::ostream& out, const TopKQueryResult& obj);

}} // namespace

#endif
