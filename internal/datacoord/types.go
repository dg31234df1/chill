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

package datacoord

import (
	"context"

	"github.com/milvus-io/milvus/internal/proto/indexpb"
	"github.com/milvus-io/milvus/internal/types"
)

type Task interface {
	GetTaskID() int64
	GetNodeID() int64
	ResetNodeID()
	CheckTaskHealthy(mt *meta) bool
	SetState(state indexpb.JobState, failReason string)
	GetState() indexpb.JobState
	GetFailReason() string
	UpdateVersion(ctx context.Context, meta *meta) error
	UpdateMetaBuildingState(nodeID int64, meta *meta) error
	AssignTask(ctx context.Context, client types.IndexNodeClient, dependency *taskScheduler) (bool, bool)
	QueryResult(ctx context.Context, client types.IndexNodeClient)
	DropTaskOnWorker(ctx context.Context, client types.IndexNodeClient) bool
	SetJobInfo(meta *meta) error
}
