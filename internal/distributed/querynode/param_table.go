package grpcquerynode

import (
	"os"
	"strconv"
	"sync"

	"github.com/zilliztech/milvus-distributed/internal/util/funcutil"
	"github.com/zilliztech/milvus-distributed/internal/util/paramtable"
)

var Params ParamTable
var once sync.Once

type ParamTable struct {
	paramtable.BaseTable

	QueryNodeIP   string
	QueryNodePort int
	QueryNodeID   UniqueID

	IndexServiceAddress string
	MasterAddress       string
	DataServiceAddress  string
	QueryServiceAddress string
}

func (pt *ParamTable) Init() {
	once.Do(func() {
		pt.BaseTable.Init()
		pt.initPort()
		pt.initMasterAddress()
		pt.initIndexServiceAddress()
		pt.initDataServiceAddress()
		pt.initQueryServiceAddress()

	})
}

func (pt *ParamTable) LoadFromArgs() {

}

func (pt *ParamTable) LoadFromEnv() {

	// todo assign by queryservice and set by read initparms
	queryNodeIDStr := os.Getenv("QUERY_NODE_ID")
	if queryNodeIDStr == "" {
		panic("Can't Get QUERY_NODE_ID")
	}

	queryID, err := strconv.Atoi(queryNodeIDStr)
	if err != nil {
		panic(err)
	}
	pt.QueryNodeID = UniqueID(queryID)

	Params.QueryNodeIP = funcutil.GetLocalIP()

}

func (pt *ParamTable) initMasterAddress() {
	ret, err := pt.Load("_MasterAddress")
	if err != nil {
		panic(err)
	}
	pt.MasterAddress = ret
}

func (pt *ParamTable) initIndexServiceAddress() {
	ret, err := pt.Load("IndexServiceAddress")
	if err != nil {
		panic(err)
	}
	pt.IndexServiceAddress = ret
}

func (pt *ParamTable) initDataServiceAddress() {
	ret, err := pt.Load("_DataServiceAddress")
	if err != nil {
		panic(err)
	}
	pt.DataServiceAddress = ret
}

func (pt *ParamTable) initQueryServiceAddress() {
	ret, err := pt.Load("_QueryServiceAddress")
	if err != nil {
		panic(err)
	}
	pt.QueryServiceAddress = ret
}

func (pt *ParamTable) initPort() {
	port := pt.ParseInt("queryNode.port")
	pt.QueryNodePort = port
}
