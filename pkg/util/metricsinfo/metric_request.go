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

package metricsinfo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/commonpbutil"
)

const (
	// MetricTypeKey are the key of metric type in GetMetrics request.
	MetricTypeKey = common.MetricTypeKey

	// SystemInfoMetrics means users request for system information metrics.
	SystemInfoMetrics = "system_info"

	// CollectionStorageMetrics means users request for collection storage metrics.
	CollectionStorageMetrics = "collection_storage"

	// MetricRequestTypeKey is a key for identify request type.
	MetricRequestTypeKey = "req_type"

	// MetricRequestParamsSeparator is a separator that parameter value will be joined be separator
	MetricRequestParamsSeparator = ","

	// QuerySegmentDist request for segment distribution on the query node
	QuerySegmentDist = "query_segment_dist"

	// QueryChannelDist request for channel distribution on the query node
	QueryChannelDist = "query_channel_dist"

	// QueryTasks request for get tasks on the querycoord
	QueryTasks = "query_tasks"

	// QueryReplicas request for get replica on the querycoord
	QueryReplicas = "query_replica"

	// QueryResourceGroups request for get resource groups on the querycoord
	QueryResourceGroups = "query_resource_group"

	// MetricRequestParamVerboseKey as a request parameter decide to whether return verbose value
	MetricRequestParamVerboseKey = "verbose"
)

type MetricsRequestAction func(ctx context.Context, req *milvuspb.GetMetricsRequest, jsonReq gjson.Result) (string, error)

var metricsReqType2Action = make(map[string]MetricsRequestAction)

func RegisterMetricsRequest(reqType string, action MetricsRequestAction) {
	_, ok := metricsReqType2Action[reqType]
	if ok {
		log.Info("metrics request type already exists", zap.String("reqType", reqType))
		return
	}

	metricsReqType2Action[reqType] = action
}

func ExecuteMetricsRequest(ctx context.Context, req *milvuspb.GetMetricsRequest) (string, error) {
	jsonReq := gjson.Parse(req.Request)
	reqType, err := ParseMetricRequestType(jsonReq)
	if err != nil {
		log.Warn("failed to parse metric type", zap.Error(err))
		return "", err
	}

	action, ok := metricsReqType2Action[reqType]
	if !ok {
		log.Warn("unimplemented metric request type", zap.String("req_type", reqType))
		return "", errors.New(MsgUnimplementedMetric)
	}

	actionRet, err := action(ctx, req, jsonReq)
	if err != nil {
		msg := fmt.Sprintf("failed to execute %s", reqType)
		log.Warn(msg, zap.Error(err))
		return "", err
	}
	return actionRet, nil
}

func RequestWithVerbose(jsonReq gjson.Result) bool {
	v := jsonReq.Get(MetricRequestParamVerboseKey)
	if !v.Exists() {
		return false
	}
	return v.Bool()
}

// ParseMetricRequestType returns the metric type of req
func ParseMetricRequestType(jsonRet gjson.Result) (string, error) {
	v := jsonRet.Get(MetricRequestTypeKey)
	if v.Exists() {
		return v.String(), nil
	}

	v = jsonRet.Get(MetricTypeKey)
	if v.Exists() {
		return v.String(), nil
	}

	return "", fmt.Errorf("%s or %s not found in request", MetricTypeKey, MetricRequestTypeKey)
}

// ConstructRequestByMetricType constructs a request according to the metric type
func ConstructRequestByMetricType(metricType string) (*milvuspb.GetMetricsRequest, error) {
	m := make(map[string]interface{})
	m[MetricTypeKey] = metricType
	binary, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to construct request by metric type %s: %s", metricType, err.Error())
	}
	// TODO:: switch metricType to different msgType and return err when metricType is not supported
	return &milvuspb.GetMetricsRequest{
		Base: commonpbutil.NewMsgBase(
			commonpbutil.WithMsgType(commonpb.MsgType_SystemInfo),
		),
		Request: string(binary),
	}, nil
}

func ConstructGetMetricsRequest(m map[string]interface{}) (*milvuspb.GetMetricsRequest, error) {
	binary, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to construct request: %s", err.Error())
	}

	return &milvuspb.GetMetricsRequest{
		Base: commonpbutil.NewMsgBase(
			commonpbutil.WithMsgType(commonpb.MsgType_SystemInfo),
		),
		Request: string(binary),
	}, nil
}
