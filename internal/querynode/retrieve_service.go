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

package querynode

import (
	"context"
	"errors"
	"strconv"

	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/msgstream"
	"github.com/milvus-io/milvus/internal/util/trace"
)

type retrieveService struct {
	ctx    context.Context
	cancel context.CancelFunc

	historical *historical
	streaming  *streaming

	retrieveMsgStream       msgstream.MsgStream
	retrieveResultMsgStream msgstream.MsgStream

	queryNodeID         UniqueID
	retrieveCollections map[UniqueID]*retrieveCollection
}

func newRetrieveService(ctx context.Context,
	historical *historical,
	streaming *streaming,
	factory msgstream.Factory) *retrieveService {

	retrieveStream, _ := factory.NewQueryMsgStream(ctx)
	retrieveResultStream, _ := factory.NewQueryMsgStream(ctx)

	if len(Params.SearchChannelNames) > 0 && len(Params.SearchResultChannelNames) > 0 {
		consumeChannels := Params.SearchChannelNames
		consumeSubName := "RetrieveSubName"
		retrieveStream.AsConsumer(consumeChannels, consumeSubName)
		log.Debug("query node AsConsumer", zap.Any("retrieveChannels", consumeChannels), zap.Any("consumeSubName", consumeSubName))
		producerChannels := Params.SearchResultChannelNames
		retrieveResultStream.AsProducer(producerChannels)
		log.Debug("query node AsProducer", zap.Any("retrieveResultChannels", producerChannels))
	}

	retrieveServiceCtx, retrieveServiceCancel := context.WithCancel(ctx)
	return &retrieveService{
		ctx:    retrieveServiceCtx,
		cancel: retrieveServiceCancel,

		historical: historical,
		streaming:  streaming,

		retrieveMsgStream:       retrieveStream,
		retrieveResultMsgStream: retrieveResultStream,

		queryNodeID:         Params.QueryNodeID,
		retrieveCollections: make(map[UniqueID]*retrieveCollection),
	}
}

func (rs *retrieveService) start() {
	rs.retrieveMsgStream.Start()
	rs.retrieveResultMsgStream.Start()
	rs.consumeRetrieve()
}

func (rs *retrieveService) collectionCheck(collectionID UniqueID) error {
	if ok := rs.historical.replica.hasCollection(collectionID); !ok {
		err := errors.New("no collection found, collectionID = " + strconv.FormatInt(collectionID, 10))
		log.Error(err.Error())
		return err
	}

	return nil
}

func (rs *retrieveService) consumeRetrieve() {
	for {
		select {
		case <-rs.ctx.Done():
			return
		default:
			msgPack := rs.retrieveMsgStream.Consume()
			if msgPack == nil || len(msgPack.Msgs) <= 0 {
				continue
			}
			for _, msg := range msgPack.Msgs {
				rm, ok := msg.(*msgstream.RetrieveMsg)
				if !ok {
					// Not a retrieve request, discard
					continue
				}
				log.Info("RetrieveService consume retrieve message",
					zap.Int64("collectionID", rm.CollectionID),
					zap.Int64("requestID", msg.ID()),
					zap.Any("requestType", "retrieve"),
				)

				sp, ctx := trace.StartSpanFromContext(rm.TraceCtx())
				rm.SetTraceCtx(ctx)
				err := rs.collectionCheck(rm.CollectionID)
				if err != nil {
					log.Debug("Failed to check collection exist, discard.",
						zap.Int64("collectionID", rm.CollectionID),
						zap.Int64("requestID", msg.ID()),
						zap.Any("requestType", "retrieve"),
					)
					continue
				}

				_, ok = rs.retrieveCollections[rm.CollectionID]
				if !ok {
					rs.startRetrieveCollection(rm.CollectionID)
					log.Debug("Receive retrieve request on new collection, start an new retrieve collection service",
						zap.Int64("collectionID", rm.CollectionID),
						zap.Int64("requestID", msg.ID()),
						zap.Any("requestType", "retrieve"),
					)
				}

				rs.retrieveCollections[rm.CollectionID].msgBuffer <- rm
				log.Info("Put retrieve msg into msgBuffer",
					zap.Any("requestID", msg.ID),
					zap.Any("requestType", "retrieve"),
				)
				sp.Finish()
			}
		}
	}
}

func (rs *retrieveService) close() {
	if rs.retrieveMsgStream != nil {
		rs.retrieveMsgStream.Close()
	}
	if rs.retrieveResultMsgStream != nil {
		rs.retrieveResultMsgStream.Close()
	}
	for collectionID := range rs.retrieveCollections {
		rs.stopRetrieveCollection(collectionID)
	}
	rs.retrieveCollections = make(map[UniqueID]*retrieveCollection)
	rs.cancel()
}

func (rs *retrieveService) startRetrieveCollection(collectionID UniqueID) {
	ctx1, cancel := context.WithCancel(rs.ctx)
	rc := newRetrieveCollection(ctx1,
		cancel,
		collectionID,
		rs.historical,
		rs.streaming,
		rs.retrieveResultMsgStream)
	rs.retrieveCollections[collectionID] = rc
	rc.start()
}

func (rs *retrieveService) hasRetrieveCollection(collectionID UniqueID) bool {
	_, ok := rs.retrieveCollections[collectionID]
	return ok
}

func (rs *retrieveService) stopRetrieveCollection(collectionID UniqueID) {
	rc, ok := rs.retrieveCollections[collectionID]
	if !ok {
		log.Error("stopRetrieveCollection failed, collection doesn't exist", zap.Int64("collectionID", collectionID))
	}
	rc.cancel()
	delete(rs.retrieveCollections, collectionID)
}
