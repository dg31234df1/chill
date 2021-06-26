// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package rocksmq

import (
	"os"
	"sync"

	"github.com/milvus-io/milvus/internal/allocator"

	rocksdbkv "github.com/milvus-io/milvus/internal/kv/rocksdb"
)

var Rmq *rocksmq
var once sync.Once

func InitRmq(rocksdbName string, idAllocator allocator.GIDAllocator) error {
	var err error
	Rmq, err = NewRocksMQ(rocksdbName, idAllocator)
	return err
}

func InitRocksMQ(rocksdbName string) error {
	var err error
	once.Do(func() {
		_, err := os.Stat(rocksdbName)
		if os.IsNotExist(err) {
			err = os.MkdirAll(rocksdbName, os.ModePerm)
			if err != nil {
				errMsg := "Create dir " + rocksdbName + " failed"
				panic(errMsg)
			}
		}

		kvname := rocksdbName + "_kv"
		rocksdbKV, err := rocksdbkv.NewRocksdbKV(kvname)
		if err != nil {
			panic(err)
		}
		idAllocator := allocator.NewGlobalIDAllocator("rmq_id", rocksdbKV)
		_ = idAllocator.Initialize()

		Rmq, err = NewRocksMQ(rocksdbName, idAllocator)
		if err != nil {
			panic(err)
		}
	})
	return err
}

func CloseRocksMQ() {
	if Rmq != nil && Rmq.store != nil {
		Rmq.store.Close()
	}
}
