package querynode

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/zilliztech/milvus-distributed/internal/util/paramtable"
)

type ParamTable struct {
	paramtable.BaseTable

	PulsarAddress string
	ETCDAddress   string
	MetaRootPath  string

	QueryNodeIP                 string
	QueryNodePort               int64
	QueryNodeID                 UniqueID
	QueryNodeNum                int
	QueryTimeTickChannelName    string
	QueryTimeTickReceiveBufSize int64

	FlowGraphMaxQueueLength int32
	FlowGraphMaxParallelism int32

	// minio
	MinioEndPoint        string
	MinioAccessKeyID     string
	MinioSecretAccessKey string
	MinioUseSSLStr       bool
	MinioBucketName      string

	// dm
	InsertChannelNames   []string
	InsertChannelRange   []int
	InsertReceiveBufSize int64
	InsertPulsarBufSize  int64

	// dd
	DDChannelNames   []string
	DDReceiveBufSize int64
	DDPulsarBufSize  int64

	// search
	SearchChannelNames         []string
	SearchResultChannelNames   []string
	SearchReceiveBufSize       int64
	SearchPulsarBufSize        int64
	SearchResultReceiveBufSize int64

	// stats
	StatsPublishInterval int
	StatsChannelName     string
	StatsReceiveBufSize  int64

	// load index
	LoadIndexChannelNames   []string
	LoadIndexReceiveBufSize int64
	LoadIndexPulsarBufSize  int64

	GracefulTime        int64
	MsgChannelSubName   string
	DefaultPartitionTag string
	SliceIndex          int
}

var Params ParamTable

func (p *ParamTable) Init() {
	p.BaseTable.Init()
	err := p.LoadYaml("advanced/query_node.yaml")
	if err != nil {
		panic(err)
	}

	err = p.LoadYaml("milvus.yaml")
	if err != nil {
		panic(err)
	}

	queryNodeIDStr := os.Getenv("QUERY_NODE_ID")
	if queryNodeIDStr == "" {
		queryNodeIDList := p.QueryNodeIDList()
		if len(queryNodeIDList) <= 0 {
			queryNodeIDStr = "0"
		} else {
			queryNodeIDStr = strconv.Itoa(int(queryNodeIDList[0]))
		}
	}

	queryNodeAddress := os.Getenv("QUERY_NODE_ADDRESS")
	if queryNodeAddress == "" {
		p.QueryNodeIP = "localhost"
		p.QueryNodePort = 20010
	} else {
		ipAndPort := strings.Split(queryNodeAddress, ":")
		if len(ipAndPort) != 2 {
			panic("illegal query node address")
		}
		p.QueryNodeIP = ipAndPort[0]
		port, err := strconv.ParseInt(ipAndPort[1], 10, 64)
		if err != nil {
			panic(err)
		}
		p.QueryNodePort = port
	}

	err = p.LoadYaml("advanced/common.yaml")
	if err != nil {
		panic(err)
	}
	err = p.Save("_queryNodeID", queryNodeIDStr)
	if err != nil {
		panic(err)
	}

	p.initQueryNodeID()
	p.initQueryNodeNum()
	p.initQueryTimeTickChannelName()
	p.initQueryTimeTickReceiveBufSize()

	p.initMinioEndPoint()
	p.initMinioAccessKeyID()
	p.initMinioSecretAccessKey()
	p.initMinioUseSSLStr()
	p.initMinioBucketName()

	p.initPulsarAddress()
	p.initETCDAddress()
	p.initMetaRootPath()

	p.initGracefulTime()
	p.initMsgChannelSubName()
	p.initDefaultPartitionTag()
	p.initSliceIndex()

	p.initFlowGraphMaxQueueLength()
	p.initFlowGraphMaxParallelism()

	p.initInsertChannelNames()
	p.initInsertChannelRange()
	p.initInsertReceiveBufSize()
	p.initInsertPulsarBufSize()

	p.initDDChannelNames()
	p.initDDReceiveBufSize()
	p.initDDPulsarBufSize()

	p.initSearchChannelNames()
	p.initSearchResultChannelNames()
	p.initSearchReceiveBufSize()
	p.initSearchPulsarBufSize()
	p.initSearchResultReceiveBufSize()

	p.initStatsPublishInterval()
	p.initStatsChannelName()
	p.initStatsReceiveBufSize()

	p.initLoadIndexChannelNames()
	p.initLoadIndexReceiveBufSize()
	p.initLoadIndexPulsarBufSize()
}

// ---------------------------------------------------------- query node
func (p *ParamTable) initQueryNodeID() {
	queryNodeID, err := p.Load("_queryNodeID")
	if err != nil {
		panic(err)
	}
	id, err := strconv.Atoi(queryNodeID)
	if err != nil {
		panic(err)
	}
	p.QueryNodeID = UniqueID(id)
}

func (p *ParamTable) initQueryNodeNum() {
	p.QueryNodeNum = len(p.QueryNodeIDList())
}

func (p *ParamTable) initQueryTimeTickChannelName() {
	ch, err := p.Load("msgChannel.chanNamePrefix.queryTimeTick")
	if err != nil {
		log.Fatal(err)
	}
	p.QueryTimeTickChannelName = ch
}

func (p *ParamTable) initQueryTimeTickReceiveBufSize() {
	p.QueryTimeTickReceiveBufSize = p.ParseInt64("queryNode.msgStream.timeTick.recvBufSize")
}

// ---------------------------------------------------------- minio
func (p *ParamTable) initMinioEndPoint() {
	url, err := p.Load("_MinioAddress")
	if err != nil {
		panic(err)
	}
	p.MinioEndPoint = url
}

func (p *ParamTable) initMinioAccessKeyID() {
	id, err := p.Load("minio.accessKeyID")
	if err != nil {
		panic(err)
	}
	p.MinioAccessKeyID = id
}

func (p *ParamTable) initMinioSecretAccessKey() {
	key, err := p.Load("minio.secretAccessKey")
	if err != nil {
		panic(err)
	}
	p.MinioSecretAccessKey = key
}

func (p *ParamTable) initMinioUseSSLStr() {
	ssl, err := p.Load("minio.useSSL")
	if err != nil {
		panic(err)
	}
	sslBoolean, err := strconv.ParseBool(ssl)
	if err != nil {
		panic(err)
	}
	p.MinioUseSSLStr = sslBoolean
}

func (p *ParamTable) initMinioBucketName() {
	bucketName, err := p.Load("minio.bucketName")
	if err != nil {
		panic(err)
	}
	p.MinioBucketName = bucketName
}

func (p *ParamTable) initPulsarAddress() {
	url, err := p.Load("_PulsarAddress")
	if err != nil {
		panic(err)
	}
	p.PulsarAddress = url
}

func (p *ParamTable) initInsertChannelRange() {
	insertChannelRange, err := p.Load("msgChannel.channelRange.insert")
	if err != nil {
		panic(err)
	}
	p.InsertChannelRange = paramtable.ConvertRangeToIntRange(insertChannelRange, ",")
}

// advanced params
// stats
func (p *ParamTable) initStatsPublishInterval() {
	p.StatsPublishInterval = p.ParseInt("queryNode.stats.publishInterval")
}

// dataSync:
func (p *ParamTable) initFlowGraphMaxQueueLength() {
	p.FlowGraphMaxQueueLength = p.ParseInt32("queryNode.dataSync.flowGraph.maxQueueLength")
}

func (p *ParamTable) initFlowGraphMaxParallelism() {
	p.FlowGraphMaxParallelism = p.ParseInt32("queryNode.dataSync.flowGraph.maxParallelism")
}

// msgStream
func (p *ParamTable) initInsertReceiveBufSize() {
	p.InsertReceiveBufSize = p.ParseInt64("queryNode.msgStream.insert.recvBufSize")
}

func (p *ParamTable) initInsertPulsarBufSize() {
	p.InsertPulsarBufSize = p.ParseInt64("queryNode.msgStream.insert.pulsarBufSize")
}

func (p *ParamTable) initDDReceiveBufSize() {
	revBufSize, err := p.Load("queryNode.msgStream.dataDefinition.recvBufSize")
	if err != nil {
		panic(err)
	}
	bufSize, err := strconv.Atoi(revBufSize)
	if err != nil {
		panic(err)
	}
	p.DDReceiveBufSize = int64(bufSize)
}

func (p *ParamTable) initDDPulsarBufSize() {
	pulsarBufSize, err := p.Load("queryNode.msgStream.dataDefinition.pulsarBufSize")
	if err != nil {
		panic(err)
	}
	bufSize, err := strconv.Atoi(pulsarBufSize)
	if err != nil {
		panic(err)
	}
	p.DDPulsarBufSize = int64(bufSize)
}

func (p *ParamTable) initSearchReceiveBufSize() {
	p.SearchReceiveBufSize = p.ParseInt64("queryNode.msgStream.search.recvBufSize")
}

func (p *ParamTable) initSearchPulsarBufSize() {
	p.SearchPulsarBufSize = p.ParseInt64("queryNode.msgStream.search.pulsarBufSize")
}

func (p *ParamTable) initSearchResultReceiveBufSize() {
	p.SearchResultReceiveBufSize = p.ParseInt64("queryNode.msgStream.searchResult.recvBufSize")
}

func (p *ParamTable) initStatsReceiveBufSize() {
	p.StatsReceiveBufSize = p.ParseInt64("queryNode.msgStream.stats.recvBufSize")
}

func (p *ParamTable) initETCDAddress() {
	ETCDAddress, err := p.Load("_EtcdAddress")
	if err != nil {
		panic(err)
	}
	p.ETCDAddress = ETCDAddress
}

func (p *ParamTable) initMetaRootPath() {
	rootPath, err := p.Load("etcd.rootPath")
	if err != nil {
		panic(err)
	}
	subPath, err := p.Load("etcd.metaSubPath")
	if err != nil {
		panic(err)
	}
	p.MetaRootPath = rootPath + "/" + subPath
}

func (p *ParamTable) initGracefulTime() {
	p.GracefulTime = p.ParseInt64("queryNode.gracefulTime")
}

func (p *ParamTable) initInsertChannelNames() {

	prefix, err := p.Load("msgChannel.chanNamePrefix.insert")
	if err != nil {
		log.Fatal(err)
	}
	prefix += "-"
	channelRange, err := p.Load("msgChannel.channelRange.insert")
	if err != nil {
		panic(err)
	}
	channelIDs := paramtable.ConvertRangeToIntSlice(channelRange, ",")

	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}
	sep := len(channelIDs) / p.QueryNodeNum
	index := p.SliceIndex
	if index == -1 {
		panic("queryNodeID not Match with Config")
	}
	start := index * sep
	p.InsertChannelNames = ret[start : start+sep]
}

func (p *ParamTable) initSearchChannelNames() {
	prefix, err := p.Load("msgChannel.chanNamePrefix.search")
	if err != nil {
		log.Fatal(err)
	}
	prefix += "-"
	channelRange, err := p.Load("msgChannel.channelRange.search")
	if err != nil {
		panic(err)
	}

	channelIDs := paramtable.ConvertRangeToIntSlice(channelRange, ",")

	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}
	p.SearchChannelNames = ret
}

func (p *ParamTable) initSearchResultChannelNames() {
	prefix, err := p.Load("msgChannel.chanNamePrefix.searchResult")
	if err != nil {
		log.Fatal(err)
	}
	prefix += "-"
	channelRange, err := p.Load("msgChannel.channelRange.searchResult")
	if err != nil {
		panic(err)
	}

	channelIDs := paramtable.ConvertRangeToIntSlice(channelRange, ",")

	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}
	p.SearchResultChannelNames = ret
}

func (p *ParamTable) initMsgChannelSubName() {
	// TODO: subName = namePrefix + "-" + queryNodeID, queryNodeID is assigned by master
	name, err := p.Load("msgChannel.subNamePrefix.queryNodeSubNamePrefix")
	if err != nil {
		log.Panic(err)
	}
	queryNodeIDStr, err := p.Load("_QueryNodeID")
	if err != nil {
		panic(err)
	}
	p.MsgChannelSubName = name + "-" + queryNodeIDStr
}

func (p *ParamTable) initStatsChannelName() {
	channels, err := p.Load("msgChannel.chanNamePrefix.queryNodeStats")
	if err != nil {
		panic(err)
	}
	p.StatsChannelName = channels
}

func (p *ParamTable) initDDChannelNames() {
	prefix, err := p.Load("msgChannel.chanNamePrefix.dataDefinition")
	if err != nil {
		panic(err)
	}
	prefix += "-"
	iRangeStr, err := p.Load("msgChannel.channelRange.dataDefinition")
	if err != nil {
		panic(err)
	}
	channelIDs := paramtable.ConvertRangeToIntSlice(iRangeStr, ",")
	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}
	p.DDChannelNames = ret
}

func (p *ParamTable) initDefaultPartitionTag() {
	defaultTag, err := p.Load("common.defaultPartitionTag")
	if err != nil {
		panic(err)
	}

	p.DefaultPartitionTag = defaultTag
}

func (p *ParamTable) initSliceIndex() {
	queryNodeID := p.QueryNodeID
	queryNodeIDList := p.QueryNodeIDList()
	for i := 0; i < len(queryNodeIDList); i++ {
		if queryNodeID == queryNodeIDList[i] {
			p.SliceIndex = i
			return
		}
	}
	p.SliceIndex = -1
}

func (p *ParamTable) initLoadIndexChannelNames() {
	loadIndexChannelName, err := p.Load("msgChannel.chanNamePrefix.cmd")
	if err != nil {
		panic(err)
	}
	p.LoadIndexChannelNames = []string{loadIndexChannelName}
}

func (p *ParamTable) initLoadIndexReceiveBufSize() {
	p.LoadIndexReceiveBufSize = p.ParseInt64("queryNode.msgStream.loadIndex.recvBufSize")
}

func (p *ParamTable) initLoadIndexPulsarBufSize() {
	p.LoadIndexPulsarBufSize = p.ParseInt64("queryNode.msgStream.loadIndex.pulsarBufSize")
}
