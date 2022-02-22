// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kv

import (
	"github.com/milvus-io/milvus/internal/util/typeutil"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Value is interface for kv-value, needed to support string and byte slice
type Value interface {
	Serialize() []byte
	String() string
}

// StringValue type alias for string to implement Value
type StringValue string

// Serialize serialize the StringValue to byte array.
func (s StringValue) Serialize() []byte {
	return []byte(s)
}

// String return the value of StringValue.
func (s StringValue) String() string {
	return string(s)
}

// BytesValue type alias for byte slice to implement value
type BytesValue []byte

// Serialize return the byte array.
func (s BytesValue) Serialize() []byte {
	return s
}

// String return the string of byte array.
func (s BytesValue) String() string {
	return string(s)
}

// ValueKV example usage to change KV interface to
type ValueKV interface {
	Save(key string, value Value) error
	Load(key string) (Value, error)
}

// BaseKV contains basic kv operations, including save, load and remove.
type BaseKV interface {
	Load(key string) (string, error)
	MultiLoad(keys []string) ([]string, error)
	LoadWithPrefix(key string) ([]string, []string, error)
	Save(key, value string) error
	MultiSave(kvs map[string]string) error
	Remove(key string) error
	MultiRemove(keys []string) error
	RemoveWithPrefix(key string) error

	Close()
}

// DataKV persists the data.
type DataKV interface {
	BaseKV
	LoadPartial(key string, start, end int64) ([]byte, error)
	GetSize(key string) (int64, error)
}

// TxnKV contains extra txn operations of kv. The extra operations is transactional.
type TxnKV interface {
	BaseKV
	MultiSaveAndRemove(saves map[string]string, removals []string) error
	MultiRemoveWithPrefix(keys []string) error
	MultiSaveAndRemoveWithPrefix(saves map[string]string, removals []string) error
}

// MetaKv is TxnKV for metadata. It should save data with lease.
type MetaKv interface {
	TxnKV
	GetPath(key string) string
	LoadWithPrefix(key string) ([]string, []string, error)
	LoadWithPrefix2(key string) ([]string, []string, []int64, error)
	LoadWithRevision(key string) ([]string, []string, int64, error)
	Watch(key string) clientv3.WatchChan
	WatchWithPrefix(key string) clientv3.WatchChan
	WatchWithRevision(key string, revision int64) clientv3.WatchChan
	SaveWithLease(key, value string, id clientv3.LeaseID) error
	Grant(ttl int64) (id clientv3.LeaseID, err error)
	KeepAlive(id clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error)
	CompareValueAndSwap(key, value, target string, opts ...clientv3.OpOption) error
	CompareVersionAndSwap(key string, version int64, target string, opts ...clientv3.OpOption) error
}

// SnapShotKV is TxnKV for snapshot data. It must save timestamp.
type SnapShotKV interface {
	Save(key string, value string, ts typeutil.Timestamp) error
	Load(key string, ts typeutil.Timestamp) (string, error)
	MultiSave(kvs map[string]string, ts typeutil.Timestamp) error
	LoadWithPrefix(key string, ts typeutil.Timestamp) ([]string, []string, error)
	MultiSaveAndRemoveWithPrefix(saves map[string]string, removals []string, ts typeutil.Timestamp) error
}
