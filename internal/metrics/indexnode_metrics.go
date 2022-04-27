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

package metrics

import (
	"github.com/milvus-io/milvus/internal/util/typeutil"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	IndexNodeBuildIndexTaskCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.IndexNodeRole,
			Name:      "index_task_count",
			Help:      "number of tasks that index node received",
		}, []string{nodeIDLabelName, statusLabelName})

	IndexNodeLoadFieldLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.IndexNodeRole,
			Name:      "load_field_latency",
			Help:      "latency of loading the field data",
			Buckets:   buckets,
		}, []string{nodeIDLabelName})

	IndexNodeDecodeFieldLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.IndexNodeRole,
			Name:      "decode_field_latency",
			Help:      "latency of decode field data",
			Buckets:   buckets,
		}, []string{nodeIDLabelName})

	IndexNodeKnowhereBuildIndexLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.IndexNodeRole,
			Name:      "build_index_latency",
			Help:      "latency of building the index by knowhere",
			Buckets:   buckets,
		}, []string{nodeIDLabelName})

	IndexNodeEncodeIndexFileLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.IndexNodeRole,
			Name:      "encode_index_latency",
			Help:      "latency of encoding the index file",
			Buckets:   buckets,
		}, []string{nodeIDLabelName})

	IndexNodeSaveIndexFileLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: milvusNamespace,
			Subsystem: typeutil.IndexNodeRole,
			Name:      "save_index_latency",
			Help:      "latency of saving the index file",
			Buckets:   buckets,
		}, []string{nodeIDLabelName})
)

//RegisterIndexNode registers IndexNode metrics
func RegisterIndexNode(registry *prometheus.Registry) {
	registry.MustRegister(IndexNodeBuildIndexTaskCounter)
	registry.MustRegister(IndexNodeLoadFieldLatency)
	registry.MustRegister(IndexNodeDecodeFieldLatency)
	registry.MustRegister(IndexNodeKnowhereBuildIndexLatency)
	registry.MustRegister(IndexNodeEncodeIndexFileLatency)
	registry.MustRegister(IndexNodeSaveIndexFileLatency)
}
