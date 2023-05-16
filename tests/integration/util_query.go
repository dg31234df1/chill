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

package integration

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/milvuspb"
	"github.com/milvus-io/milvus-proto/go-api/schemapb"
	"github.com/milvus-io/milvus/pkg/common"
)

const (
	AnnsFieldKey    = "anns_field"
	TopKKey         = "topk"
	NQKey           = "nq"
	MetricTypeKey   = common.MetricTypeKey
	SearchParamsKey = common.IndexParamsKey
	RoundDecimalKey = "round_decimal"
	OffsetKey       = "offset"
	LimitKey        = "limit"
)

func waitingForLoad(ctx context.Context, cluster *MiniCluster, collection string) {
	getLoadingProgress := func() *milvuspb.GetLoadingProgressResponse {
		loadProgress, err := cluster.proxy.GetLoadingProgress(ctx, &milvuspb.GetLoadingProgressRequest{
			CollectionName: collection,
		})
		if err != nil {
			panic("GetLoadingProgress fail")
		}
		return loadProgress
	}
	for getLoadingProgress().GetProgress() != 100 {
		select {
		case <-ctx.Done():
			panic("load timeout")
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func constructSearchRequest(
	dbName, collectionName string,
	expr string,
	vecField string,
	vectorType schemapb.DataType,
	outputFields []string,
	metricType string,
	params map[string]any,
	nq, dim int, topk, roundDecimal int,
) *milvuspb.SearchRequest {
	b, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	plg := constructPlaceholderGroup(nq, dim, vectorType)
	plgBs, err := proto.Marshal(plg)
	if err != nil {
		panic(err)
	}

	return &milvuspb.SearchRequest{
		Base:             nil,
		DbName:           dbName,
		CollectionName:   collectionName,
		PartitionNames:   nil,
		Dsl:              expr,
		PlaceholderGroup: plgBs,
		DslType:          commonpb.DslType_BoolExprV1,
		OutputFields:     outputFields,
		SearchParams: []*commonpb.KeyValuePair{
			{
				Key:   common.MetricTypeKey,
				Value: metricType,
			},
			{
				Key:   SearchParamsKey,
				Value: string(b),
			},
			{
				Key:   AnnsFieldKey,
				Value: vecField,
			},
			{
				Key:   common.TopKKey,
				Value: strconv.Itoa(topk),
			},
			{
				Key:   RoundDecimalKey,
				Value: strconv.Itoa(roundDecimal),
			},
		},
		TravelTimestamp:    0,
		GuaranteeTimestamp: 0,
	}
}

func constructPlaceholderGroup(nq, dim int, vectorType schemapb.DataType) *commonpb.PlaceholderGroup {
	values := make([][]byte, 0, nq)
	var placeholderType commonpb.PlaceholderType
	switch vectorType {
	case schemapb.DataType_FloatVector:
		placeholderType = commonpb.PlaceholderType_FloatVector
		for i := 0; i < nq; i++ {
			bs := make([]byte, 0, dim*4)
			for j := 0; j < dim; j++ {
				var buffer bytes.Buffer
				f := rand.Float32()
				err := binary.Write(&buffer, common.Endian, f)
				if err != nil {
					panic(err)
				}
				bs = append(bs, buffer.Bytes()...)
			}
			values = append(values, bs)
		}
	case schemapb.DataType_BinaryVector:
		placeholderType = commonpb.PlaceholderType_BinaryVector
		for i := 0; i < nq; i++ {
			total := dim / 8
			ret := make([]byte, total)
			_, err := rand.Read(ret)
			if err != nil {
				panic(err)
			}
			values = append(values, ret)
		}
	default:
		panic("invalid vector data type")
	}

	return &commonpb.PlaceholderGroup{
		Placeholders: []*commonpb.PlaceholderValue{
			{
				Tag:    "$0",
				Type:   placeholderType,
				Values: values,
			},
		},
	}
}
