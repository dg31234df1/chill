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

package importv2

import (
	"sync"

	"github.com/milvus-io/milvus/pkg/util/conc"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
)

var (
	execPool         *conc.Pool[any]
	execPoolInitOnce sync.Once
)

func initExecPool() {
	execPool = conc.NewPool[any](
		paramtable.Get().DataNodeCfg.MaxConcurrentImportTaskNum.GetAsInt(),
		conc.WithPreAlloc(true),
	)
}

func GetExecPool() *conc.Pool[any] {
	execPoolInitOnce.Do(initExecPool)
	return execPool
}
