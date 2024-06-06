package io

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/milvus-io/milvus/pkg/util/conc"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
)

func TestGetOrCreateIOPool(t *testing.T) {
	paramtable.Init()
	ioConcurrency := paramtable.Get().DataNodeCfg.IOConcurrency.GetValue()
	paramtable.Get().Save(paramtable.Get().DataNodeCfg.IOConcurrency.Key, "64")
	defer func() { paramtable.Get().Save(paramtable.Get().DataNodeCfg.IOConcurrency.Key, ioConcurrency) }()
	nP := 10
	nTask := 10
	wg := sync.WaitGroup{}
	for i := 0; i < nP; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p := GetOrCreateIOPool()
			futures := make([]*conc.Future[any], 0, nTask)
			for j := 0; j < nTask; j++ {
				future := p.Submit(func() (interface{}, error) {
					return nil, nil
				})
				futures = append(futures, future)
			}
			err := conc.AwaitAll(futures...)
			assert.NoError(t, err)
		}()
	}
	wg.Wait()
}
