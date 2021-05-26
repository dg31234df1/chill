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

package masterservice

import (
	"context"
	"fmt"
	"sync"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/proxypb"
	"github.com/milvus-io/milvus/internal/types"
	"github.com/milvus-io/milvus/internal/util/sessionutil"
	"go.uber.org/zap"
)

type proxyClientManager struct {
	core        *Core
	lock        sync.Mutex
	proxyClient map[int64]types.ProxyNode
}

func newProxyClientManager(c *Core) *proxyClientManager {
	return &proxyClientManager{
		core:        c,
		lock:        sync.Mutex{},
		proxyClient: make(map[int64]types.ProxyNode),
	}
}

func (p *proxyClientManager) GetProxyClients(sess []*sessionutil.Session) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for _, s := range sess {
		if _, ok := p.proxyClient[s.ServerID]; ok {
			continue
		}
		pc, err := p.core.NewProxyClient(s)
		if err != nil {
			log.Debug("create proxy client failed", zap.String("proxy address", s.Address), zap.Int64("proxy id", s.ServerID), zap.Error(err))
			continue
		}
		p.proxyClient[s.ServerID] = pc
		log.Debug("create proxy client", zap.String("proxy address", s.Address), zap.Int64("proxy id", s.ServerID))
	}
}

func (p *proxyClientManager) AddProxyClient(s *sessionutil.Session) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if _, ok := p.proxyClient[s.ServerID]; ok {
		return
	}
	pc, err := p.core.NewProxyClient(s)
	if err != nil {
		log.Debug("create proxy client", zap.String("proxy address", s.Address), zap.Int64("proxy id", s.ServerID), zap.Error(err))
		return
	}
	p.proxyClient[s.ServerID] = pc
	log.Debug("create proxy client", zap.String("proxy address", s.Address), zap.Int64("proxy id", s.ServerID))
}

func (p *proxyClientManager) DelProxyClient(s *sessionutil.Session) {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.proxyClient, s.ServerID)
	log.Debug("remove proxy client", zap.String("proxy address", s.Address), zap.Int64("proxy id", s.ServerID))
}

func (p *proxyClientManager) InvalidateCollectionMetaCache(ctx context.Context, request *proxypb.InvalidateCollMetaCacheRequest) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.proxyClient) == 0 {
		log.Debug("proxy client is empty,InvalidateCollectionMetaCache will not send to any client")
		return
	}

	for k, f := range p.proxyClient {
		err := func() error {
			defer func() {
				if err := recover(); err != nil {
					log.Debug("call InvalidateCollectionMetaCache panic", zap.Int64("proxy id", k), zap.Any("msg", err))
				}

			}()
			sta, err := f.InvalidateCollectionMetaCache(ctx, request)
			if err != nil {
				return fmt.Errorf("grpc fail,error=%w", err)
			}
			if sta.ErrorCode != commonpb.ErrorCode_Success {
				return fmt.Errorf("message = %s", sta.Reason)
			}
			return nil
		}()
		if err != nil {
			log.Error("call invalidate collection meta failed", zap.Int64("proxy id", k), zap.Error(err))
		} else {
			log.Debug("send invalidate collection meta cache to proxy node", zap.Int64("node id", k))
		}

	}
}
