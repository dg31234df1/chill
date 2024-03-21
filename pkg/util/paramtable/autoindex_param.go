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

package paramtable

import (
	"fmt"

	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/config"
	"github.com/milvus-io/milvus/pkg/util/funcutil"
	"github.com/milvus-io/milvus/pkg/util/indexparamcheck"
)

// /////////////////////////////////////////////////////////////////////////////
// --- common ---
type autoIndexConfig struct {
	Enable         ParamItem `refreshable:"true"`
	EnableOptimize ParamItem `refreshable:"true"`

	IndexParams           ParamItem  `refreshable:"true"`
	PrepareParams         ParamItem  `refreshable:"true"`
	ExtraParams           ParamItem  `refreshable:"true"`
	IndexType             ParamItem  `refreshable:"true"`
	AutoIndexTypeName     ParamItem  `refreshable:"true"`
	AutoIndexSearchConfig ParamItem  `refreshable:"true"`
	AutoIndexTuningConfig ParamGroup `refreshable:"true"`

	ScalarAutoIndexEnable  ParamItem `refreshable:"true"`
	ScalarAutoIndexParams  ParamItem `refreshable:"true"`
	ScalarNumericIndexType ParamItem `refreshable:"true"`
	ScalarVarcharIndexType ParamItem `refreshable:"true"`
	ScalarBoolIndexType    ParamItem `refreshable:"true"`
}

func (p *autoIndexConfig) init(base *BaseTable) {
	p.Enable = ParamItem{
		Key:          "autoIndex.enable",
		Version:      "2.2.0",
		DefaultValue: "false",
		PanicIfEmpty: true,
	}
	p.Enable.Init(base.mgr)

	p.EnableOptimize = ParamItem{
		Key:          "autoIndex.optimize",
		Version:      "2.4.0",
		DefaultValue: "true",
		PanicIfEmpty: true,
	}
	p.EnableOptimize.Init(base.mgr)

	p.IndexParams = ParamItem{
		Key:          "autoIndex.params.build",
		Version:      "2.2.0",
		DefaultValue: `{"M": 18,"efConstruction": 240,"index_type": "HNSW", "metric_type": "IP"}`,
	}
	p.IndexParams.Init(base.mgr)

	p.PrepareParams = ParamItem{
		Key:     "autoIndex.params.prepare",
		Version: "2.3.2",
	}
	p.PrepareParams.Init(base.mgr)

	p.ExtraParams = ParamItem{
		Key:     "autoIndex.params.extra",
		Version: "2.2.0",
	}
	p.ExtraParams.Init(base.mgr)

	p.IndexType = ParamItem{
		Version: "2.2.0",
		Formatter: func(v string) string {
			m := p.IndexParams.GetAsJSONMap()
			if m == nil {
				return ""
			}
			return m[common.IndexTypeKey]
		},
	}
	p.IndexType.Init(base.mgr)

	p.AutoIndexTypeName = ParamItem{
		Key:     "autoIndex.type",
		Version: "2.2.0",
	}
	p.AutoIndexTypeName.Init(base.mgr)

	p.AutoIndexSearchConfig = ParamItem{
		Key:     "autoindex.params.search",
		Version: "2.2.0",
	}
	p.AutoIndexSearchConfig.Init(base.mgr)

	p.AutoIndexTuningConfig = ParamGroup{
		KeyPrefix: "autoindex.params.tuning.",
		Version:   "2.3.0",
	}
	p.AutoIndexTuningConfig.Init(base.mgr)

	p.panicIfNotValidAndSetDefaultMetricType(base.mgr)

	p.ScalarAutoIndexEnable = ParamItem{
		Key:          "scalarAutoIndex.enable",
		Version:      "2.3.4",
		DefaultValue: "false",
		PanicIfEmpty: true,
	}
	p.ScalarAutoIndexEnable.Init(base.mgr)

	p.ScalarAutoIndexParams = ParamItem{
		Key:          "scalarAutoIndex.params.build",
		Version:      "2.3.4",
		DefaultValue: `{"numeric": "INVERTED","varchar": "INVERTED","bool": "INVERTED"}`,
	}
	p.ScalarAutoIndexParams.Init(base.mgr)

	p.ScalarNumericIndexType = ParamItem{
		Version: "2.4.0",
		Formatter: func(v string) string {
			m := p.ScalarAutoIndexParams.GetAsJSONMap()
			if m == nil {
				return ""
			}
			return m["numeric"]
		},
	}
	p.ScalarNumericIndexType.Init(base.mgr)

	p.ScalarVarcharIndexType = ParamItem{
		Version: "2.4.0",
		Formatter: func(v string) string {
			m := p.ScalarAutoIndexParams.GetAsJSONMap()
			if m == nil {
				return ""
			}
			return m["varchar"]
		},
	}
	p.ScalarVarcharIndexType.Init(base.mgr)

	p.ScalarBoolIndexType = ParamItem{
		Version: "2.4.0",
		Formatter: func(v string) string {
			m := p.ScalarAutoIndexParams.GetAsJSONMap()
			if m == nil {
				return ""
			}
			return m["bool"]
		},
	}
	p.ScalarBoolIndexType.Init(base.mgr)
}

func (p *autoIndexConfig) panicIfNotValidAndSetDefaultMetricType(mgr *config.Manager) {
	m := p.IndexParams.GetAsJSONMap()
	if m == nil {
		panic("autoIndex.build not invalid, should be json format")
	}

	indexType, ok := m[common.IndexTypeKey]
	if !ok {
		panic("autoIndex.build not invalid, index type not found")
	}

	checker, err := indexparamcheck.GetIndexCheckerMgrInstance().GetChecker(indexType)
	if err != nil {
		panic(fmt.Sprintf("autoIndex.build not invalid, unsupported index type: %s", indexType))
	}

	checker.SetDefaultMetricTypeIfNotExist(m)

	if err := checker.StaticCheck(m); err != nil {
		panic(fmt.Sprintf("autoIndex.build not invalid, parameters not invalid, error: %s", err.Error()))
	}

	p.reset(m, mgr)
}

func (p *autoIndexConfig) reset(m map[string]string, mgr *config.Manager) {
	j := funcutil.MapToJSON(m)
	mgr.SetConfig("autoIndex.params.build", string(j))
}
