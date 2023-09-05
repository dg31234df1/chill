package lock

import (
	"sync"
	"testing"
	"time"

	"github.com/milvus-io/milvus/pkg/util/paramtable"
	"github.com/stretchr/testify/assert"
)

func TestMetricsLockLock(t *testing.T) {
	params := paramtable.Get()
	lManager := &MetricsLockManager{
		rwLocks: make(map[string]*MetricsRWMutex, 0),
	}
	params.Init(paramtable.NewBaseTable(paramtable.SkipRemote(true)))
	params.Save(params.CommonCfg.EnableLockMetrics.Key, "true")
	params.Save(params.CommonCfg.LockSlowLogInfoThreshold.Key, "10")
	lName := "testLock"
	lockDuration := 10 * time.Millisecond

	testRWLock := lManager.applyRWLock(lName)
	wg := sync.WaitGroup{}
	testRWLock.Lock("main_thread")
	go func() {
		wg.Add(1)
		before := time.Now()
		testRWLock.Lock("sub_thread")
		lkDuration := time.Since(before)
		assert.True(t, lkDuration >= lockDuration)
		testRWLock.UnLock("sub_threadXX")
		testRWLock.UnLock("sub_thread")
		wg.Done()
	}()
	time.Sleep(lockDuration)
	testRWLock.UnLock("main_thread")
	wg.Wait()
}

func TestMetricsLockRLock(t *testing.T) {
	params := paramtable.Get()
	lManager := &MetricsLockManager{
		rwLocks: make(map[string]*MetricsRWMutex, 0),
	}
	params.Init(paramtable.NewBaseTable(paramtable.SkipRemote(true)))
	params.Save(params.CommonCfg.EnableLockMetrics.Key, "true")
	params.Save(params.CommonCfg.LockSlowLogWarnThreshold.Key, "10")
	lName := "testLock"
	lockDuration := 10 * time.Millisecond

	testRWLock := lManager.applyRWLock(lName)
	wg := sync.WaitGroup{}
	testRWLock.RLock("main_thread")
	go func() {
		wg.Add(1)
		before := time.Now()
		testRWLock.Lock("sub_thread")
		lkDuration := time.Since(before)
		assert.True(t, lkDuration >= lockDuration)
		testRWLock.UnLock("sub_thread")
		wg.Done()
	}()
	time.Sleep(lockDuration)
	assert.Equal(t, 1, len(testRWLock.acquireTimeMap))
	testRWLock.RUnLock("main_threadXXX")
	assert.Equal(t, 1, len(testRWLock.acquireTimeMap))
	testRWLock.RUnLock("main_thread")
	wg.Wait()
	assert.Equal(t, 0, len(testRWLock.acquireTimeMap))
}
