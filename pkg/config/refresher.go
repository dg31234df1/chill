// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package config

import (
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/milvus-io/milvus/pkg/log"
)

type refresher struct {
	refreshInterval  time.Duration
	intervalDone     chan struct{}
	intervalInitOnce sync.Once
	eh               EventHandler

	fetchFunc func() error
	stopOnce  sync.Once
	wg        sync.WaitGroup
}

func newRefresher(interval time.Duration, fetchFunc func() error) *refresher {
	return &refresher{
		refreshInterval: interval,
		intervalDone:    make(chan struct{}),
		fetchFunc:       fetchFunc,
	}
}

func (r *refresher) start(name string) {
	if r.refreshInterval > 0 {
		r.intervalInitOnce.Do(func() {
			r.wg.Add(1)
			go r.refreshPeriodically(name)
		})
	}
}

func (r *refresher) stop() {
	r.stopOnce.Do(func() {
		close(r.intervalDone)
		r.wg.Wait()
	})
}

func (r *refresher) refreshPeriodically(name string) {
	defer r.wg.Done()
	ticker := time.NewTicker(r.refreshInterval)
	defer ticker.Stop()
	log.Debug("start refreshing configurations", zap.String("source", name))
	for {
		select {
		case <-ticker.C:
			err := r.fetchFunc()
			if err != nil {
				log.Error("can not pull configs", zap.Error(err))
				r.stop()
			}
		case <-r.intervalDone:
			log.Info("stop refreshing configurations", zap.String("source", name))
			return
		}
	}
}

func (r *refresher) fireEvents(name string, source, target map[string]string) error {
	events, err := PopulateEvents(name, source, target)
	if err != nil {
		log.Warn("generating event error", zap.Error(err))
		return err
	}
	// Generate OnEvent Callback based on the events created
	if r.eh != nil {
		for _, e := range events {
			r.eh.OnEvent(e)
		}
	}
	return nil
}
