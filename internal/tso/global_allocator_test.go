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

package tso

import (
	"math"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/milvus-io/milvus/internal/util/etcd"
	"github.com/milvus-io/milvus/internal/util/tsoutil"
	"github.com/stretchr/testify/assert"
)

var gTestTsoAllocator *GlobalTSOAllocator

func TestGlobalTSOAllocator_Initialize(t *testing.T) {
	endpoints := os.Getenv("ETCD_ENDPOINTS")
	if endpoints == "" {
		endpoints = "localhost:2379"
	}
	etcdEndpoints := strings.Split(endpoints, ",")
	etcdCli, err := etcd.GetRemoteEtcdClient(etcdEndpoints)
	assert.Nil(t, err)
	defer etcdCli.Close()

	etcdKV := tsoutil.NewTSOKVBase(etcdCli, "/test/root/kv", "tsoTest")
	gTestTsoAllocator = NewGlobalTSOAllocator("timestamp", etcdKV)
	err = gTestTsoAllocator.Initialize()
	assert.Nil(t, err)

	err = gTestTsoAllocator.UpdateTSO()
	assert.Nil(t, err)

	time.Sleep(3 * time.Second)
	err = gTestTsoAllocator.Initialize()
	assert.Nil(t, err)

	t.Run("GenerateTSO", func(t *testing.T) {
		count := 1000
		perCount := uint32(100)
		startTs, err := gTestTsoAllocator.GenerateTSO(perCount)
		assert.Nil(t, err)
		lastPhysical, lastLogical := tsoutil.ParseTS(startTs)
		for i := 0; i < count; i++ {
			ts, _ := gTestTsoAllocator.GenerateTSO(perCount)
			physical, logical := tsoutil.ParseTS(ts)
			if lastPhysical.Equal(physical) {
				diff := logical - lastLogical
				assert.Equal(t, uint64(perCount), diff)
			}
			lastPhysical, lastLogical = physical, logical
		}
	})
}

func TestGlobalTSOAllocator_All(t *testing.T) {
	endpoints := os.Getenv("ETCD_ENDPOINTS")
	if endpoints == "" {
		endpoints = "localhost:2379"
	}
	etcdEndpoints := strings.Split(endpoints, ",")
	etcdCli, err := etcd.GetRemoteEtcdClient(etcdEndpoints)
	assert.NoError(t, err)
	defer etcdCli.Close()
	etcdKV := tsoutil.NewTSOKVBase(etcdCli, "/test/root/kv", "tsoTest")

	gTestTsoAllocator = NewGlobalTSOAllocator("timestamp", etcdKV)
	t.Run("Initialize", func(t *testing.T) {
		err := gTestTsoAllocator.Initialize()
		assert.Nil(t, err)

		err = gTestTsoAllocator.UpdateTSO()
		assert.Nil(t, err)
	})

	t.Run("GenerateTSO", func(t *testing.T) {
		count := 1000
		perCount := uint32(100)
		startTs, err := gTestTsoAllocator.GenerateTSO(perCount)
		assert.Nil(t, err)
		lastPhysical, lastLogical := tsoutil.ParseTS(startTs)
		for i := 0; i < count; i++ {
			ts, err2 := gTestTsoAllocator.GenerateTSO(perCount)
			assert.Nil(t, err2)
			physical, logical := tsoutil.ParseTS(ts)
			if lastPhysical.Equal(physical) {
				diff := logical - lastLogical
				assert.Equal(t, uint64(perCount), diff)
			}
			lastPhysical, lastLogical = physical, logical
		}
	})

	gTestTsoAllocator.SetLimitMaxLogic(false)
	t.Run("GenerateTSO2", func(t *testing.T) {
		count := 1000
		maxL := 1 << 18
		startTs, err := gTestTsoAllocator.GenerateTSO(uint32(maxL))
		step := 10
		perCount := uint32(step) << 18 // 10 ms
		assert.Nil(t, err)
		lastPhysical, lastLogical := tsoutil.ParseTS(startTs)
		for i := 0; i < count; i++ {
			ts, _ := gTestTsoAllocator.GenerateTSO(perCount)
			physical, logical := tsoutil.ParseTS(ts)
			assert.Equal(t, logical, lastLogical)
			assert.Equal(t, physical, lastPhysical.Add(time.Millisecond*time.Duration(step)))
			lastPhysical = physical
		}
		err = gTestTsoAllocator.UpdateTSO()
		assert.Nil(t, err)
	})

	gTestTsoAllocator.SetLimitMaxLogic(true)
	t.Run("SetTSO", func(t *testing.T) {
		ts, err2 := gTestTsoAllocator.GenerateTSO(1)
		assert.Nil(t, err2)
		curTime, logical := tsoutil.ParseTS(ts)
		nextTime := curTime.Add(2 * time.Second)
		physical := nextTime.UnixNano() / int64(time.Millisecond)
		err := gTestTsoAllocator.SetTSO(tsoutil.ComposeTS(physical, int64(logical)))
		assert.Nil(t, err)
	})

	t.Run("UpdateTSO", func(t *testing.T) {
		err := gTestTsoAllocator.UpdateTSO()
		assert.Nil(t, err)
	})

	t.Run("Alloc", func(t *testing.T) {
		_, err := gTestTsoAllocator.Alloc(100)
		assert.Nil(t, err)
	})

	t.Run("AllocOne", func(t *testing.T) {
		_, err := gTestTsoAllocator.AllocOne()
		assert.Nil(t, err)
	})

	t.Run("Reset", func(t *testing.T) {
		gTestTsoAllocator.Reset()
	})
}

func TestGlobalTSOAllocator_Fail(t *testing.T) {
	endpoints := os.Getenv("ETCD_ENDPOINTS")
	if endpoints == "" {
		endpoints = "localhost:2379"
	}
	etcdEndpoints := strings.Split(endpoints, ",")
	etcdCli, err := etcd.GetRemoteEtcdClient(etcdEndpoints)
	assert.NoError(t, err)
	defer etcdCli.Close()
	etcdKV := tsoutil.NewTSOKVBase(etcdCli, "/test/root/kv", "tsoTest")
	assert.NoError(t, err)
	gTestTsoAllocator = NewGlobalTSOAllocator("timestamp", etcdKV)
	t.Run("Initialize", func(t *testing.T) {
		err := gTestTsoAllocator.Initialize()
		assert.Nil(t, err)
	})

	t.Run("GenerateTSO_invalid", func(t *testing.T) {
		_, err := gTestTsoAllocator.GenerateTSO(0)
		assert.NotNil(t, err)
	})

	gTestTsoAllocator.SetLimitMaxLogic(true)
	t.Run("SetTSO_invalid", func(t *testing.T) {
		err := gTestTsoAllocator.SetTSO(0)
		assert.NotNil(t, err)

		err = gTestTsoAllocator.SetTSO(math.MaxUint64)
		assert.NotNil(t, err)
	})

	t.Run("Alloc_invalid", func(t *testing.T) {
		_, err := gTestTsoAllocator.Alloc(0)
		assert.NotNil(t, err)

		_, err = gTestTsoAllocator.Alloc(math.MaxUint32)
		assert.NotNil(t, err)
	})

	t.Run("Reset", func(t *testing.T) {
		gTestTsoAllocator.Reset()
	})
}

func TestGlobalTSOAllocator_Update(t *testing.T) {
	endpoints := os.Getenv("ETCD_ENDPOINTS")
	if endpoints == "" {
		endpoints = "localhost:2379"
	}
	etcdEndpoints := strings.Split(endpoints, ",")
	etcdCli, err := etcd.GetRemoteEtcdClient(etcdEndpoints)
	assert.NoError(t, err)
	defer etcdCli.Close()
	etcdKV := tsoutil.NewTSOKVBase(etcdCli, "/test/root/kv", "tsoTest")
	assert.NoError(t, err)
	gTestTsoAllocator = NewGlobalTSOAllocator("timestamp", etcdKV)
	err = gTestTsoAllocator.Initialize()
	assert.Nil(t, err)

	err = gTestTsoAllocator.UpdateTSO()
	assert.Nil(t, err)
	time.Sleep(160 * time.Millisecond)
	err = gTestTsoAllocator.UpdateTSO()
	assert.Nil(t, err)
	time.Sleep(500 * time.Millisecond)
	err = gTestTsoAllocator.UpdateTSO()
	assert.Nil(t, err)
}

func TestGlobalTSOAllocator_load(t *testing.T) {
	endpoints := os.Getenv("ETCD_ENDPOINTS")
	if endpoints == "" {
		endpoints = "localhost:2379"
	}
	etcdEndpoints := strings.Split(endpoints, ",")
	etcdCli, err := etcd.GetRemoteEtcdClient(etcdEndpoints)
	assert.NoError(t, err)
	defer etcdCli.Close()
	etcdKV := tsoutil.NewTSOKVBase(etcdCli, "/test/root/kv", "tsoTest")
	assert.NoError(t, err)
	gTestTsoAllocator = NewGlobalTSOAllocator("timestamp", etcdKV)
	err = gTestTsoAllocator.Initialize()
	assert.Nil(t, err)

	err = gTestTsoAllocator.UpdateTSO()
	assert.Nil(t, err)

	ts, _ := gTestTsoAllocator.GenerateTSO(1)
	curTime, logical := tsoutil.ParseTS(ts)
	nextTime := curTime.Add(2 * time.Second)
	physical := nextTime.UnixNano() / int64(time.Millisecond)
	target := tsoutil.ComposeTS(physical, int64(logical))
	err = gTestTsoAllocator.SetTSO(target)
	assert.Nil(t, err)

	gTestTsoAllocator = NewGlobalTSOAllocator("timestamp", etcdKV)
	err = gTestTsoAllocator.Initialize()
	assert.Nil(t, err)

	ts2, err2 := gTestTsoAllocator.GenerateTSO(1)
	assert.Nil(t, err2)
	curTime2, _ := tsoutil.ParseTS(ts2)
	assert.True(t, ts2 >= target)
	assert.True(t, curTime2.UnixNano() >= nextTime.UnixNano())
}
