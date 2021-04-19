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

package grpcproxynode

import (
	"sync"

	"github.com/zilliztech/milvus-distributed/internal/util/funcutil"
	"github.com/zilliztech/milvus-distributed/internal/util/paramtable"
)

type ParamTable struct {
	paramtable.BaseTable

	ProxyServiceAddress string
	ProxyServicePort    int

	IndexServerAddress string
	MasterAddress      string

	DataServiceAddress  string
	QueryServiceAddress string

	IP      string
	Port    int
	Address string
}

var Params ParamTable
var once sync.Once

func (pt *ParamTable) Init() {
	once.Do(func() {
		pt.BaseTable.Init()
		pt.initParams()
	})
}

func (pt *ParamTable) LoadFromArgs() {

}

func (pt *ParamTable) LoadFromEnv() {
	Params.IP = funcutil.GetLocalIP()
}

func (pt *ParamTable) initParams() {
	pt.initPoxyServicePort()
	pt.initPort()
	pt.initProxyServiceAddress()
	pt.initMasterAddress()
	pt.initIndexServerAddress()
	pt.initDataServiceAddress()
	pt.initQueryServiceAddress()
}

func (pt *ParamTable) initPoxyServicePort() {
	pt.ProxyServicePort = pt.ParseInt("proxyService.port")
}

func (pt *ParamTable) initProxyServiceAddress() {
	ret, err := pt.Load("_PROXY_SERVICE_ADDRESS")
	if err != nil {
		panic(err)
	}
	pt.ProxyServiceAddress = ret
}

// todo remove and use load from env
func (pt *ParamTable) initIndexServerAddress() {
	ret, err := pt.Load("IndexServiceAddress")
	if err != nil {
		panic(err)
	}
	pt.IndexServerAddress = ret
}

// todo remove and use load from env
func (pt *ParamTable) initMasterAddress() {
	ret, err := pt.Load("_MasterAddress")
	if err != nil {
		panic(err)
	}
	pt.MasterAddress = ret
}

// todo remove and use load from env
func (pt *ParamTable) initDataServiceAddress() {
	ret, err := pt.Load("_DataServiceAddress")
	if err != nil {
		panic(err)
	}
	pt.DataServiceAddress = ret
}

// todo remove and use load from env
func (pt *ParamTable) initQueryServiceAddress() {
	ret, err := pt.Load("_QueryServiceAddress")
	if err != nil {
		panic(err)
	}
	pt.QueryServiceAddress = ret
}

func (pt *ParamTable) initPort() {
	port := pt.ParseInt("proxyNode.port")
	pt.Port = port
}
