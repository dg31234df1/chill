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

package msgdispatcher

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/schemapb"
	"github.com/milvus-io/milvus/internal/mq/msgstream"
	"github.com/milvus-io/milvus/internal/mq/msgstream/mqwrapper"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/util/paramtable"
	"github.com/milvus-io/milvus/internal/util/typeutil"
)

const (
	dim = 128
)

func newMockFactory() msgstream.Factory {
	paramtable.Init()
	return msgstream.NewRmsFactory("/tmp/milvus/rocksmq/")
}

func newMockProducer(factory msgstream.Factory, pchannel string) (msgstream.MsgStream, error) {
	stream, err := factory.NewMsgStream(context.Background())
	if err != nil {
		return nil, err
	}
	stream.AsProducer([]string{pchannel})
	stream.SetRepackFunc(defaultInsertRepackFunc)
	return stream, nil
}

func getSeekPositions(factory msgstream.Factory, pchannel string, maxNum int) ([]*msgstream.MsgPosition, error) {
	stream, err := factory.NewTtMsgStream(context.Background())
	if err != nil {
		return nil, err
	}
	defer stream.Close()
	stream.AsConsumer([]string{pchannel}, fmt.Sprintf("%d", rand.Int()), mqwrapper.SubscriptionPositionEarliest)
	positions := make([]*msgstream.MsgPosition, 0)
	for {
		select {
		case <-time.After(100 * time.Millisecond): // no message to consume
			return positions, nil
		case pack := <-stream.Chan():
			positions = append(positions, pack.EndPositions[0])
			if len(positions) >= maxNum {
				return positions, nil
			}
		}
	}
}

func genPKs(numRows int) []typeutil.IntPrimaryKey {
	ids := make([]typeutil.IntPrimaryKey, numRows)
	for i := 0; i < numRows; i++ {
		ids[i] = typeutil.IntPrimaryKey(i)
	}
	return ids
}

func genTimestamps(numRows int) []typeutil.Timestamp {
	ts := make([]typeutil.Timestamp, numRows)
	for i := 0; i < numRows; i++ {
		ts[i] = typeutil.Timestamp(i + 1)
	}
	return ts
}

func genInsertMsg(numRows int, vchannel string, msgID typeutil.UniqueID) *msgstream.InsertMsg {
	floatVec := make([]float32, numRows*dim)
	for i := 0; i < numRows*dim; i++ {
		floatVec[i] = rand.Float32()
	}
	hashValues := make([]uint32, numRows)
	for i := 0; i < numRows; i++ {
		hashValues[i] = uint32(1)
	}
	return &msgstream.InsertMsg{
		BaseMsg: msgstream.BaseMsg{HashValues: hashValues},
		InsertRequest: internalpb.InsertRequest{
			Base:       &commonpb.MsgBase{MsgType: commonpb.MsgType_Insert, MsgID: msgID},
			ShardName:  vchannel,
			Timestamps: genTimestamps(numRows),
			RowIDs:     genPKs(numRows),
			FieldsData: []*schemapb.FieldData{{
				Field: &schemapb.FieldData_Vectors{
					Vectors: &schemapb.VectorField{
						Dim:  dim,
						Data: &schemapb.VectorField_FloatVector{FloatVector: &schemapb.FloatArray{Data: floatVec}},
					},
				},
			}},
			NumRows: uint64(numRows),
			Version: internalpb.InsertDataVersion_ColumnBased,
		},
	}
}

func genDeleteMsg(numRows int, vchannel string, msgID typeutil.UniqueID) *msgstream.DeleteMsg {
	return &msgstream.DeleteMsg{
		BaseMsg: msgstream.BaseMsg{HashValues: make([]uint32, numRows)},
		DeleteRequest: internalpb.DeleteRequest{
			Base:      &commonpb.MsgBase{MsgType: commonpb.MsgType_Delete, MsgID: msgID},
			ShardName: vchannel,
			PrimaryKeys: &schemapb.IDs{
				IdField: &schemapb.IDs_IntId{
					IntId: &schemapb.LongArray{
						Data: genPKs(numRows),
					},
				},
			},
			Timestamps: genTimestamps(numRows),
			NumRows:    int64(numRows),
		},
	}
}

func genDDLMsg(msgType commonpb.MsgType) msgstream.TsMsg {
	switch msgType {
	case commonpb.MsgType_CreateCollection:
		return &msgstream.CreateCollectionMsg{
			BaseMsg: msgstream.BaseMsg{HashValues: []uint32{0}},
			CreateCollectionRequest: internalpb.CreateCollectionRequest{
				Base: &commonpb.MsgBase{MsgType: commonpb.MsgType_CreateCollection},
			},
		}
	case commonpb.MsgType_DropCollection:
		return &msgstream.DropCollectionMsg{
			BaseMsg: msgstream.BaseMsg{HashValues: []uint32{0}},
			DropCollectionRequest: internalpb.DropCollectionRequest{
				Base: &commonpb.MsgBase{MsgType: commonpb.MsgType_DropCollection},
			},
		}
	case commonpb.MsgType_CreatePartition:
		return &msgstream.CreatePartitionMsg{
			BaseMsg: msgstream.BaseMsg{HashValues: []uint32{0}},
			CreatePartitionRequest: internalpb.CreatePartitionRequest{
				Base: &commonpb.MsgBase{MsgType: commonpb.MsgType_CreatePartition},
			},
		}
	case commonpb.MsgType_DropPartition:
		return &msgstream.DropPartitionMsg{
			BaseMsg: msgstream.BaseMsg{HashValues: []uint32{0}},
			DropPartitionRequest: internalpb.DropPartitionRequest{
				Base: &commonpb.MsgBase{MsgType: commonpb.MsgType_DropPartition},
			},
		}
	}
	return nil
}

func genTimeTickMsg(ts typeutil.Timestamp) *msgstream.TimeTickMsg {
	return &msgstream.TimeTickMsg{
		BaseMsg: msgstream.BaseMsg{HashValues: []uint32{0}},
		TimeTickMsg: internalpb.TimeTickMsg{
			Base: &commonpb.MsgBase{
				MsgType:   commonpb.MsgType_TimeTick,
				Timestamp: ts,
			},
		},
	}
}

// defaultInsertRepackFunc repacks the dml messages.
func defaultInsertRepackFunc(
	tsMsgs []msgstream.TsMsg,
	hashKeys [][]int32,
) (map[int32]*msgstream.MsgPack, error) {

	if len(hashKeys) < len(tsMsgs) {
		return nil, fmt.Errorf(
			"the length of hash keys (%d) is less than the length of messages (%d)",
			len(hashKeys),
			len(tsMsgs),
		)
	}

	// after assigning segment id to msg, tsMsgs was already re-bucketed
	pack := make(map[int32]*msgstream.MsgPack)
	for idx, msg := range tsMsgs {
		if len(hashKeys[idx]) <= 0 {
			return nil, fmt.Errorf("no hash key for %dth message", idx)
		}
		key := hashKeys[idx][0]
		_, ok := pack[key]
		if !ok {
			pack[key] = &msgstream.MsgPack{}
		}
		pack[key].Msgs = append(pack[key].Msgs, msg)
	}
	return pack, nil
}
