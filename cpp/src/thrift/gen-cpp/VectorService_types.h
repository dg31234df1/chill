/**
 * Autogenerated by Thrift Compiler (0.12.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef VectorService_TYPES_H
#define VectorService_TYPES_H

#include <iosfwd>

#include <thrift/Thrift.h>
#include <thrift/TApplicationException.h>
#include <thrift/TBase.h>
#include <thrift/protocol/TProtocol.h>
#include <thrift/transport/TTransport.h>

#include <thrift/stdcxx.h>




struct VecErrCode {
  enum type {
    SUCCESS = 0,
    ILLEGAL_ARGUMENT = 1,
    GROUP_NOT_EXISTS = 2,
    ILLEGAL_TIME_RANGE = 3,
    ILLEGAL_VECTOR_DIMENSION = 4,
    OUT_OF_MEMORY = 5
  };
};

extern const std::map<int, const char*> _VecErrCode_VALUES_TO_NAMES;

std::ostream& operator<<(std::ostream& out, const VecErrCode::type& val);

class VecException;

class VecGroup;

class VecTensor;

class VecTensorList;

class VecSearchResult;

class VecSearchResultList;

class VecDateTime;

class VecTimeRange;

class VecTimeRangeList;

typedef struct _VecException__isset {
  _VecException__isset() : code(false), reason(false) {}
  bool code :1;
  bool reason :1;
} _VecException__isset;

class VecException : public ::apache::thrift::TException {
 public:

  VecException(const VecException&);
  VecException& operator=(const VecException&);
  VecException() : code((VecErrCode::type)0), reason() {
  }

  virtual ~VecException() throw();
  VecErrCode::type code;
  std::string reason;

  _VecException__isset __isset;

  void __set_code(const VecErrCode::type val);

  void __set_reason(const std::string& val);

  bool operator == (const VecException & rhs) const
  {
    if (!(code == rhs.code))
      return false;
    if (!(reason == rhs.reason))
      return false;
    return true;
  }
  bool operator != (const VecException &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecException & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
  mutable std::string thriftTExceptionMessageHolder_;
  const char* what() const throw();
};

void swap(VecException &a, VecException &b);

std::ostream& operator<<(std::ostream& out, const VecException& obj);

typedef struct _VecGroup__isset {
  _VecGroup__isset() : id(false), dimension(false), index_type(false) {}
  bool id :1;
  bool dimension :1;
  bool index_type :1;
} _VecGroup__isset;

class VecGroup : public virtual ::apache::thrift::TBase {
 public:

  VecGroup(const VecGroup&);
  VecGroup& operator=(const VecGroup&);
  VecGroup() : id(), dimension(0), index_type(0) {
  }

  virtual ~VecGroup() throw();
  std::string id;
  int32_t dimension;
  int32_t index_type;

  _VecGroup__isset __isset;

  void __set_id(const std::string& val);

  void __set_dimension(const int32_t val);

  void __set_index_type(const int32_t val);

  bool operator == (const VecGroup & rhs) const
  {
    if (!(id == rhs.id))
      return false;
    if (!(dimension == rhs.dimension))
      return false;
    if (!(index_type == rhs.index_type))
      return false;
    return true;
  }
  bool operator != (const VecGroup &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecGroup & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VecGroup &a, VecGroup &b);

std::ostream& operator<<(std::ostream& out, const VecGroup& obj);

typedef struct _VecTensor__isset {
  _VecTensor__isset() : uid(false), tensor(false) {}
  bool uid :1;
  bool tensor :1;
} _VecTensor__isset;

class VecTensor : public virtual ::apache::thrift::TBase {
 public:

  VecTensor(const VecTensor&);
  VecTensor& operator=(const VecTensor&);
  VecTensor() : uid() {
  }

  virtual ~VecTensor() throw();
  std::string uid;
  std::vector<double>  tensor;

  _VecTensor__isset __isset;

  void __set_uid(const std::string& val);

  void __set_tensor(const std::vector<double> & val);

  bool operator == (const VecTensor & rhs) const
  {
    if (!(uid == rhs.uid))
      return false;
    if (!(tensor == rhs.tensor))
      return false;
    return true;
  }
  bool operator != (const VecTensor &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecTensor & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VecTensor &a, VecTensor &b);

std::ostream& operator<<(std::ostream& out, const VecTensor& obj);

typedef struct _VecTensorList__isset {
  _VecTensorList__isset() : tensor_list(false) {}
  bool tensor_list :1;
} _VecTensorList__isset;

class VecTensorList : public virtual ::apache::thrift::TBase {
 public:

  VecTensorList(const VecTensorList&);
  VecTensorList& operator=(const VecTensorList&);
  VecTensorList() {
  }

  virtual ~VecTensorList() throw();
  std::vector<VecTensor>  tensor_list;

  _VecTensorList__isset __isset;

  void __set_tensor_list(const std::vector<VecTensor> & val);

  bool operator == (const VecTensorList & rhs) const
  {
    if (!(tensor_list == rhs.tensor_list))
      return false;
    return true;
  }
  bool operator != (const VecTensorList &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecTensorList & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VecTensorList &a, VecTensorList &b);

std::ostream& operator<<(std::ostream& out, const VecTensorList& obj);

typedef struct _VecSearchResult__isset {
  _VecSearchResult__isset() : id_list(false) {}
  bool id_list :1;
} _VecSearchResult__isset;

class VecSearchResult : public virtual ::apache::thrift::TBase {
 public:

  VecSearchResult(const VecSearchResult&);
  VecSearchResult& operator=(const VecSearchResult&);
  VecSearchResult() {
  }

  virtual ~VecSearchResult() throw();
  std::vector<std::string>  id_list;

  _VecSearchResult__isset __isset;

  void __set_id_list(const std::vector<std::string> & val);

  bool operator == (const VecSearchResult & rhs) const
  {
    if (!(id_list == rhs.id_list))
      return false;
    return true;
  }
  bool operator != (const VecSearchResult &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecSearchResult & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VecSearchResult &a, VecSearchResult &b);

std::ostream& operator<<(std::ostream& out, const VecSearchResult& obj);

typedef struct _VecSearchResultList__isset {
  _VecSearchResultList__isset() : result_list(false) {}
  bool result_list :1;
} _VecSearchResultList__isset;

class VecSearchResultList : public virtual ::apache::thrift::TBase {
 public:

  VecSearchResultList(const VecSearchResultList&);
  VecSearchResultList& operator=(const VecSearchResultList&);
  VecSearchResultList() {
  }

  virtual ~VecSearchResultList() throw();
  std::vector<VecSearchResult>  result_list;

  _VecSearchResultList__isset __isset;

  void __set_result_list(const std::vector<VecSearchResult> & val);

  bool operator == (const VecSearchResultList & rhs) const
  {
    if (!(result_list == rhs.result_list))
      return false;
    return true;
  }
  bool operator != (const VecSearchResultList &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecSearchResultList & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VecSearchResultList &a, VecSearchResultList &b);

std::ostream& operator<<(std::ostream& out, const VecSearchResultList& obj);

typedef struct _VecDateTime__isset {
  _VecDateTime__isset() : year(false), month(false), day(false), hour(false), minute(false), second(false) {}
  bool year :1;
  bool month :1;
  bool day :1;
  bool hour :1;
  bool minute :1;
  bool second :1;
} _VecDateTime__isset;

class VecDateTime : public virtual ::apache::thrift::TBase {
 public:

  VecDateTime(const VecDateTime&);
  VecDateTime& operator=(const VecDateTime&);
  VecDateTime() : year(0), month(0), day(0), hour(0), minute(0), second(0) {
  }

  virtual ~VecDateTime() throw();
  int32_t year;
  int32_t month;
  int32_t day;
  int32_t hour;
  int32_t minute;
  int32_t second;

  _VecDateTime__isset __isset;

  void __set_year(const int32_t val);

  void __set_month(const int32_t val);

  void __set_day(const int32_t val);

  void __set_hour(const int32_t val);

  void __set_minute(const int32_t val);

  void __set_second(const int32_t val);

  bool operator == (const VecDateTime & rhs) const
  {
    if (!(year == rhs.year))
      return false;
    if (!(month == rhs.month))
      return false;
    if (!(day == rhs.day))
      return false;
    if (!(hour == rhs.hour))
      return false;
    if (!(minute == rhs.minute))
      return false;
    if (!(second == rhs.second))
      return false;
    return true;
  }
  bool operator != (const VecDateTime &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecDateTime & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VecDateTime &a, VecDateTime &b);

std::ostream& operator<<(std::ostream& out, const VecDateTime& obj);

typedef struct _VecTimeRange__isset {
  _VecTimeRange__isset() : time_begin(false), begine_closed(false), time_end(false), end_closed(false) {}
  bool time_begin :1;
  bool begine_closed :1;
  bool time_end :1;
  bool end_closed :1;
} _VecTimeRange__isset;

class VecTimeRange : public virtual ::apache::thrift::TBase {
 public:

  VecTimeRange(const VecTimeRange&);
  VecTimeRange& operator=(const VecTimeRange&);
  VecTimeRange() : begine_closed(0), end_closed(0) {
  }

  virtual ~VecTimeRange() throw();
  VecDateTime time_begin;
  bool begine_closed;
  VecDateTime time_end;
  bool end_closed;

  _VecTimeRange__isset __isset;

  void __set_time_begin(const VecDateTime& val);

  void __set_begine_closed(const bool val);

  void __set_time_end(const VecDateTime& val);

  void __set_end_closed(const bool val);

  bool operator == (const VecTimeRange & rhs) const
  {
    if (!(time_begin == rhs.time_begin))
      return false;
    if (!(begine_closed == rhs.begine_closed))
      return false;
    if (!(time_end == rhs.time_end))
      return false;
    if (!(end_closed == rhs.end_closed))
      return false;
    return true;
  }
  bool operator != (const VecTimeRange &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecTimeRange & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VecTimeRange &a, VecTimeRange &b);

std::ostream& operator<<(std::ostream& out, const VecTimeRange& obj);

typedef struct _VecTimeRangeList__isset {
  _VecTimeRangeList__isset() : range_list(false) {}
  bool range_list :1;
} _VecTimeRangeList__isset;

class VecTimeRangeList : public virtual ::apache::thrift::TBase {
 public:

  VecTimeRangeList(const VecTimeRangeList&);
  VecTimeRangeList& operator=(const VecTimeRangeList&);
  VecTimeRangeList() {
  }

  virtual ~VecTimeRangeList() throw();
  std::vector<VecTimeRange>  range_list;

  _VecTimeRangeList__isset __isset;

  void __set_range_list(const std::vector<VecTimeRange> & val);

  bool operator == (const VecTimeRangeList & rhs) const
  {
    if (!(range_list == rhs.range_list))
      return false;
    return true;
  }
  bool operator != (const VecTimeRangeList &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const VecTimeRangeList & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(VecTimeRangeList &a, VecTimeRangeList &b);

std::ostream& operator<<(std::ostream& out, const VecTimeRangeList& obj);



#endif
