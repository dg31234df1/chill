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

package paramtable

import (
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-basic/ipv4"
	"github.com/milvus-io/milvus/internal/log"
	"go.uber.org/zap"
)

const (
	// DefaultServerMaxSendSize defines the maximum size of data per grpc request can send by server side.
	DefaultServerMaxSendSize = math.MaxInt32

	// DefaultServerMaxRecvSize defines the maximum size of data per grpc request can receive by server side.
	DefaultServerMaxRecvSize = math.MaxInt32

	// DefaultClientMaxSendSize defines the maximum size of data per grpc request can send by client side.
	DefaultClientMaxSendSize = 100 * 1024 * 1024

	// DefaultClientMaxRecvSize defines the maximum size of data per grpc request can receive by client side.
	DefaultClientMaxRecvSize = 100 * 1024 * 1024

	// SuggestPulsarMaxMessageSize defines the maximum size of Pulsar message.
	SuggestPulsarMaxMessageSize = 5 * 1024 * 1024

	// DefaultRetentionDuration defines the default duration for retention which is 5 days in seconds.
	DefaultRetentionDuration = 3600 * 24 * 5
)

// GlobalParamTable is a derived struct of BaseParamTable.
// It is used to quickly and easily access global system configuration.
type GlobalParamTable struct {
	BaseParamTable
	once sync.Once

	CommonCfg     commonConfig
	KnowhereCfg   knowhereConfig
	MsgChannelCfg msgChannelConfig

	RootCoordCfg  rootCoordConfig
	ProxyCfg      proxyConfig
	QueryCoordCfg queryCoordConfig
	QueryNodeCfg  queryNodeConfig
	DataCoordCfg  dataCoordConfig
	DataNodeCfg   dataNodeConfig
	IndexCoordCfg indexCoordConfig
	IndexNodeCfg  indexNodeConfig
}

// InitOnce initialize once
func (p *GlobalParamTable) InitOnce() {
	p.once.Do(func() {
		p.Init()
	})
}

// Init initialize the global param table
func (p *GlobalParamTable) Init() {
	p.BaseParamTable.Init()

	p.CommonCfg.init(&p.BaseParamTable)
	p.KnowhereCfg.init(&p.BaseParamTable)
	p.MsgChannelCfg.init(&p.BaseParamTable)

	p.RootCoordCfg.init(&p.BaseParamTable)
	p.ProxyCfg.init(&p.BaseParamTable)
	p.QueryCoordCfg.init(&p.BaseParamTable)
	p.QueryNodeCfg.init(&p.BaseParamTable)
	p.DataCoordCfg.init(&p.BaseParamTable)
	p.DataNodeCfg.init(&p.BaseParamTable)
	p.IndexCoordCfg.init(&p.BaseParamTable)
	p.IndexNodeCfg.init(&p.BaseParamTable)
}

// SetLogConfig set log config with given role
func (p *GlobalParamTable) SetLogConfig(role string) {
	p.BaseTable.RoleName = role
	p.BaseTable.SetLogConfig()
}

///////////////////////////////////////////////////////////////////////////////
// --- common ---
type commonConfig struct {
	BaseParams *BaseParamTable

	DefaultPartitionName string
	DefaultIndexName     string
	RetentionDuration    int64
}

func (p *commonConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initDefaultPartitionName()
	p.initDefaultIndexName()
	p.initRetentionDuration()
}

func (p *commonConfig) initDefaultPartitionName() {
	p.DefaultPartitionName = p.BaseParams.LoadWithDefault("common.defaultPartitionName", "_default")
}

func (p *commonConfig) initDefaultIndexName() {
	p.DefaultIndexName = p.BaseParams.LoadWithDefault("common.defaultIndexName", "_default_idx")
}

func (p *commonConfig) initRetentionDuration() {
	p.RetentionDuration = p.BaseParams.ParseInt64WithDefault("common.retentionDuration", DefaultRetentionDuration)
}

///////////////////////////////////////////////////////////////////////////////
// --- knowhere ---
type knowhereConfig struct {
	BaseParams *BaseParamTable

	SimdType string
}

func (p *knowhereConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initSimdType()
}

func (p *knowhereConfig) initSimdType() {
	p.SimdType = p.BaseParams.LoadWithDefault("knowhere.simdType", "auto")
}

///////////////////////////////////////////////////////////////////////////////
// --- msgChannel ---
type msgChannelConfig struct {
	BaseParams *BaseParamTable

	ClusterPrefix string

	RootCoordTimeTick   string
	RootCoordStatistics string
	RootCoordDml        string
	RootCoordDelta      string
	RootCoordSubName    string

	QueryCoordSearch       string
	QueryCoordSearchResult string
	QueryCoordTimeTick     string
	QueryNodeStats         string

	DataCoordStatistic   string
	DataCoordTimeTick    string
	DataCoordSegmentInfo string
	DataCoordSubName     string
}

func (p *msgChannelConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	// must init cluster prefix first
	p.initClusterPrefix()

	p.initRootCoordTimeTick()
	p.initRootCoordStatistics()
	p.initRootCoordDml()
	p.initRootCoordDelta()
	p.initRootCoordSubName()

	p.initQueryCoordSearch()
	p.initQueryCoordSearchResult()
	p.initQueryCoordTimeTick()
	p.initQueryNodeStats()

	p.initDataCoordStatistic()
	p.initDataCoordTimeTick()
	p.initDataCoordSegmentInfo()
	p.initDataCoordSubName()
}

func (p *msgChannelConfig) initClusterPrefix() {
	str, err := p.BaseParams.Load("msgChannel.chanNamePrefix.cluster")
	if err != nil {
		panic(err)
	}
	p.ClusterPrefix = str
}

func (p *msgChannelConfig) initChanNamePrefix(cfg string) string {
	value, err := p.BaseParams.Load(cfg)
	if err != nil {
		panic(err)
	}
	s := []string{p.ClusterPrefix, value}
	return strings.Join(s, "-")
}

// --- rootcoord ---
func (p *msgChannelConfig) initRootCoordTimeTick() {
	p.RootCoordTimeTick = p.initChanNamePrefix("msgChannel.chanNamePrefix.rootCoordTimeTick")
}

func (p *msgChannelConfig) initRootCoordStatistics() {
	p.RootCoordStatistics = p.initChanNamePrefix("msgChannel.chanNamePrefix.rootCoordStatistics")
}

func (p *msgChannelConfig) initRootCoordDml() {
	p.RootCoordDml = p.initChanNamePrefix("msgChannel.chanNamePrefix.rootCoordDml")
}

func (p *msgChannelConfig) initRootCoordDelta() {
	p.RootCoordDelta = p.initChanNamePrefix("msgChannel.chanNamePrefix.rootCoordDelta")
}

func (p *msgChannelConfig) initRootCoordSubName() {
	p.RootCoordSubName = p.initChanNamePrefix("msgChannel.subNamePrefix.rootCoordSubNamePrefix")
}

// --- querycoord ---
func (p *msgChannelConfig) initQueryCoordSearch() {
	p.QueryCoordSearch = p.initChanNamePrefix("msgChannel.chanNamePrefix.search")
}

func (p *msgChannelConfig) initQueryCoordSearchResult() {
	p.QueryCoordSearchResult = p.initChanNamePrefix("msgChannel.chanNamePrefix.searchResult")
}

func (p *msgChannelConfig) initQueryCoordTimeTick() {
	p.QueryCoordTimeTick = p.initChanNamePrefix("msgChannel.chanNamePrefix.queryTimeTick")
}

// --- querynode ---
func (p *msgChannelConfig) initQueryNodeStats() {
	p.QueryNodeStats = p.initChanNamePrefix("msgChannel.chanNamePrefix.queryNodeStats")
}

// --- datacoord ---
func (p *msgChannelConfig) initDataCoordStatistic() {
	p.DataCoordStatistic = p.initChanNamePrefix("msgChannel.chanNamePrefix.dataCoordStatistic")
}

func (p *msgChannelConfig) initDataCoordTimeTick() {
	p.DataCoordTimeTick = p.initChanNamePrefix("msgChannel.chanNamePrefix.dataCoordTimeTick")
}

func (p *msgChannelConfig) initDataCoordSegmentInfo() {
	p.DataCoordSegmentInfo = p.initChanNamePrefix("msgChannel.chanNamePrefix.dataCoordSegmentInfo")
}

func (p *msgChannelConfig) initDataCoordSubName() {
	p.DataCoordSubName = p.initChanNamePrefix("msgChannel.subNamePrefix.dataCoordSubNamePrefix")
}

///////////////////////////////////////////////////////////////////////////////
// --- rootcoord ---
type rootCoordConfig struct {
	BaseParams *BaseParamTable

	Address string
	Port    int

	DmlChannelNum               int64
	MaxPartitionNum             int64
	MinSegmentSizeToEnableIndex int64

	CreatedTime time.Time
	UpdatedTime time.Time
}

func (p *rootCoordConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initDmlChannelNum()
	p.initMaxPartitionNum()
	p.initMinSegmentSizeToEnableIndex()
}

func (p *rootCoordConfig) initDmlChannelNum() {
	p.DmlChannelNum = p.BaseParams.ParseInt64WithDefault("rootCoord.dmlChannelNum", 256)
}

func (p *rootCoordConfig) initMaxPartitionNum() {
	p.MaxPartitionNum = p.BaseParams.ParseInt64WithDefault("rootCoord.maxPartitionNum", 4096)
}

func (p *rootCoordConfig) initMinSegmentSizeToEnableIndex() {
	p.MinSegmentSizeToEnableIndex = p.BaseParams.ParseInt64WithDefault("rootCoord.minSegmentSizeToEnableIndex", 1024)
}

///////////////////////////////////////////////////////////////////////////////
// --- proxy ---
type proxyConfig struct {
	BaseParams *BaseParamTable

	// NetworkPort & IP are not used
	NetworkPort    int
	IP             string
	NetworkAddress string

	Alias string

	ProxyID                  UniqueID
	TimeTickInterval         time.Duration
	MsgStreamTimeTickBufSize int64
	MaxNameLength            int64
	MaxFieldNum              int64
	MaxShardNum              int32
	MaxDimension             int64
	BufFlagExpireTime        time.Duration
	BufFlagCleanupInterval   time.Duration

	// --- Channels ---
	ProxySubName string

	// required from QueryCoord
	SearchResultChannelNames   []string
	RetrieveResultChannelNames []string

	MaxTaskNum int64

	CreatedTime time.Time
	UpdatedTime time.Time
}

func (p *proxyConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initTimeTickInterval()

	p.initProxySubName()

	p.initMsgStreamTimeTickBufSize()
	p.initMaxNameLength()
	p.initMaxFieldNum()
	p.initMaxShardNum()
	p.initMaxDimension()

	p.initMaxTaskNum()
	p.initBufFlagExpireTime()
	p.initBufFlagCleanupInterval()
}

// Refresh is called after session init
func (p *proxyConfig) Refresh() {
	p.initProxySubName()
}

// InitAlias initialize Alias member.
func (p *proxyConfig) InitAlias(alias string) {
	p.Alias = alias
}

func (p *proxyConfig) initTimeTickInterval() {
	interval := p.BaseParams.ParseIntWithDefault("proxy.timeTickInterval", 200)
	p.TimeTickInterval = time.Duration(interval) * time.Millisecond
}

func (p *proxyConfig) initProxySubName() {
	cluster, err := p.BaseParams.Load("msgChannel.chanNamePrefix.cluster")
	if err != nil {
		panic(err)
	}
	subname, err := p.BaseParams.Load("msgChannel.subNamePrefix.proxySubNamePrefix")
	if err != nil {
		panic(err)
	}
	s := []string{cluster, subname, strconv.FormatInt(p.ProxyID, 10)}
	p.ProxySubName = strings.Join(s, "-")
}

func (p *proxyConfig) initMsgStreamTimeTickBufSize() {
	p.MsgStreamTimeTickBufSize = p.BaseParams.ParseInt64WithDefault("proxy.msgStream.timeTick.bufSize", 512)
}

func (p *proxyConfig) initMaxNameLength() {
	str := p.BaseParams.LoadWithDefault("proxy.maxNameLength", "255")
	maxNameLength, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	p.MaxNameLength = maxNameLength
}

func (p *proxyConfig) initMaxShardNum() {
	str := p.BaseParams.LoadWithDefault("proxy.maxShardNum", "256")
	maxShardNum, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	p.MaxShardNum = int32(maxShardNum)
}

func (p *proxyConfig) initMaxFieldNum() {
	str := p.BaseParams.LoadWithDefault("proxy.maxFieldNum", "64")
	maxFieldNum, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	p.MaxFieldNum = maxFieldNum
}

func (p *proxyConfig) initMaxDimension() {
	str := p.BaseParams.LoadWithDefault("proxy.maxDimension", "32768")
	maxDimension, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	p.MaxDimension = maxDimension
}

func (p *proxyConfig) initMaxTaskNum() {
	p.MaxTaskNum = p.BaseParams.ParseInt64WithDefault("proxy.maxTaskNum", 1024)
}

func (p *proxyConfig) initBufFlagExpireTime() {
	expireTime := p.BaseParams.ParseInt64WithDefault("proxy.bufFlagExpireTime", 3600)
	p.BufFlagExpireTime = time.Duration(expireTime) * time.Second
}

func (p *proxyConfig) initBufFlagCleanupInterval() {
	interval := p.BaseParams.ParseInt64WithDefault("proxy.bufFlagCleanupInterval", 600)
	p.BufFlagCleanupInterval = time.Duration(interval) * time.Second
}

///////////////////////////////////////////////////////////////////////////////
// --- querycoord ---
type queryCoordConfig struct {
	BaseParams *BaseParamTable

	NodeID uint64

	Address      string
	Port         int
	QueryCoordID UniqueID

	CreatedTime time.Time
	UpdatedTime time.Time

	//---- Handoff ---
	AutoHandoff bool

	//---- Balance ---
	AutoBalance                         bool
	OverloadedMemoryThresholdPercentage float64
	BalanceIntervalSeconds              int64
	MemoryUsageMaxDifferencePercentage  float64
}

func (p *queryCoordConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	//---- Handoff ---
	p.initAutoHandoff()

	//---- Balance ---
	p.initAutoBalance()
	p.initOverloadedMemoryThresholdPercentage()
	p.initBalanceIntervalSeconds()
	p.initMemoryUsageMaxDifferencePercentage()
}

func (p *queryCoordConfig) initAutoHandoff() {
	handoff, err := p.BaseParams.Load("queryCoord.autoHandoff")
	if err != nil {
		panic(err)
	}
	p.AutoHandoff, err = strconv.ParseBool(handoff)
	if err != nil {
		panic(err)
	}
}

func (p *queryCoordConfig) initAutoBalance() {
	balanceStr := p.BaseParams.LoadWithDefault("queryCoord.autoBalance", "false")
	autoBalance, err := strconv.ParseBool(balanceStr)
	if err != nil {
		panic(err)
	}
	p.AutoBalance = autoBalance
}

func (p *queryCoordConfig) initOverloadedMemoryThresholdPercentage() {
	overloadedMemoryThresholdPercentage := p.BaseParams.LoadWithDefault("queryCoord.overloadedMemoryThresholdPercentage", "90")
	thresholdPercentage, err := strconv.ParseInt(overloadedMemoryThresholdPercentage, 10, 64)
	if err != nil {
		panic(err)
	}
	p.OverloadedMemoryThresholdPercentage = float64(thresholdPercentage) / 100
}

func (p *queryCoordConfig) initBalanceIntervalSeconds() {
	balanceInterval := p.BaseParams.LoadWithDefault("queryCoord.balanceIntervalSeconds", "60")
	interval, err := strconv.ParseInt(balanceInterval, 10, 64)
	if err != nil {
		panic(err)
	}
	p.BalanceIntervalSeconds = interval
}

func (p *queryCoordConfig) initMemoryUsageMaxDifferencePercentage() {
	maxDiff := p.BaseParams.LoadWithDefault("queryCoord.memoryUsageMaxDifferencePercentage", "30")
	diffPercentage, err := strconv.ParseInt(maxDiff, 10, 64)
	if err != nil {
		panic(err)
	}
	p.MemoryUsageMaxDifferencePercentage = float64(diffPercentage) / 100
}

///////////////////////////////////////////////////////////////////////////////
// --- querynode ---
type queryNodeConfig struct {
	BaseParams *BaseParamTable

	Alias         string
	QueryNodeIP   string
	QueryNodePort int64
	QueryNodeID   UniqueID
	// TODO: remove cacheSize
	CacheSize int64 // deprecated

	// channel prefix
	QueryNodeSubName string

	FlowGraphMaxQueueLength int32
	FlowGraphMaxParallelism int32

	// search
	SearchChannelNames         []string
	SearchResultChannelNames   []string
	SearchReceiveBufSize       int64
	SearchPulsarBufSize        int64
	SearchResultReceiveBufSize int64

	// Retrieve
	RetrieveChannelNames         []string
	RetrieveResultChannelNames   []string
	RetrieveReceiveBufSize       int64
	RetrievePulsarBufSize        int64
	RetrieveResultReceiveBufSize int64

	// stats
	StatsPublishInterval int

	GracefulTime int64
	SliceIndex   int

	// segcore
	ChunkRows int64

	CreatedTime time.Time
	UpdatedTime time.Time

	// memory limit
	OverloadedMemoryThresholdPercentage float64
}

func (p *queryNodeConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initCacheSize()
	p.initGracefulTime()

	p.initFlowGraphMaxQueueLength()
	p.initFlowGraphMaxParallelism()

	p.initSearchReceiveBufSize()
	p.initSearchPulsarBufSize()
	p.initSearchResultReceiveBufSize()

	p.initQueryNodeSubName()

	p.initStatsPublishInterval()

	p.initSegcoreChunkRows()

	p.initOverloadedMemoryThresholdPercentage()
}

// InitAlias initializes an alias for the QueryNode role.
func (p *queryNodeConfig) InitAlias(alias string) {
	p.Alias = alias
}

// Refresh is called after session init
func (p *queryNodeConfig) Refresh() {
	p.initQueryNodeSubName()
}

func (p *queryNodeConfig) initCacheSize() {
	defer log.Debug("init cacheSize", zap.Any("cacheSize (GB)", p.CacheSize))

	const defaultCacheSize = 32 // GB
	p.CacheSize = defaultCacheSize

	var err error
	cacheSize := os.Getenv("CACHE_SIZE")
	if cacheSize == "" {
		cacheSize, err = p.BaseParams.Load("queryNode.cacheSize")
		if err != nil {
			return
		}
	}
	value, err := strconv.ParseInt(cacheSize, 10, 64)
	if err != nil {
		return
	}
	p.CacheSize = value
}

// advanced params
// stats
func (p *queryNodeConfig) initStatsPublishInterval() {
	p.StatsPublishInterval = p.BaseParams.ParseIntWithDefault("queryNode.stats.publishInterval", 1000)
}

// dataSync:
func (p *queryNodeConfig) initFlowGraphMaxQueueLength() {
	p.FlowGraphMaxQueueLength = p.BaseParams.ParseInt32WithDefault("queryNode.dataSync.flowGraph.maxQueueLength", 1024)
}

func (p *queryNodeConfig) initFlowGraphMaxParallelism() {
	p.FlowGraphMaxParallelism = p.BaseParams.ParseInt32WithDefault("queryNode.dataSync.flowGraph.maxParallelism", 1024)
}

// msgStream
func (p *queryNodeConfig) initSearchReceiveBufSize() {
	p.SearchReceiveBufSize = p.BaseParams.ParseInt64WithDefault("queryNode.msgStream.search.recvBufSize", 512)
}

func (p *queryNodeConfig) initSearchPulsarBufSize() {
	p.SearchPulsarBufSize = p.BaseParams.ParseInt64WithDefault("queryNode.msgStream.search.pulsarBufSize", 512)
}

func (p *queryNodeConfig) initSearchResultReceiveBufSize() {
	p.SearchResultReceiveBufSize = p.BaseParams.ParseInt64WithDefault("queryNode.msgStream.searchResult.recvBufSize", 64)
}

// ------------------------  channel names
func (p *queryNodeConfig) initQueryNodeSubName() {
	cluster, err := p.BaseParams.Load("msgChannel.chanNamePrefix.cluster")
	if err != nil {
		panic(err)
	}
	subname, err := p.BaseParams.Load("msgChannel.subNamePrefix.queryNodeSubNamePrefix")
	if err != nil {
		log.Warn(err.Error())
	}

	s := []string{cluster, subname}
	p.QueryNodeSubName = strings.Join(s, "-")
}

func (p *queryNodeConfig) initGracefulTime() {
	p.GracefulTime = p.BaseParams.ParseInt64("queryNode.gracefulTime")
	log.Debug("query node init gracefulTime", zap.Any("gracefulTime", p.GracefulTime))
}

func (p *queryNodeConfig) initSegcoreChunkRows() {
	p.ChunkRows = p.BaseParams.ParseInt64WithDefault("queryNode.segcore.chunkRows", 32768)
}

func (p *queryNodeConfig) initOverloadedMemoryThresholdPercentage() {
	overloadedMemoryThresholdPercentage := p.BaseParams.LoadWithDefault("queryCoord.overloadedMemoryThresholdPercentage", "90")
	thresholdPercentage, err := strconv.ParseInt(overloadedMemoryThresholdPercentage, 10, 64)
	if err != nil {
		panic(err)
	}
	p.OverloadedMemoryThresholdPercentage = float64(thresholdPercentage) / 100
}

///////////////////////////////////////////////////////////////////////////////
// --- datacoord ---
type dataCoordConfig struct {
	BaseParams *BaseParamTable

	NodeID int64

	IP      string
	Port    int
	Address string

	// --- ETCD ---
	ChannelWatchSubPath string

	// --- SEGMENTS ---
	SegmentMaxSize          float64
	SegmentSealProportion   float64
	SegAssignmentExpiration int64

	CreatedTime time.Time
	UpdatedTime time.Time

	EnableCompaction        bool
	EnableAutoCompaction    bool
	EnableGarbageCollection bool

	RetentionDuration          int64
	CompactionEntityExpiration int64

	// Garbage Collection
	GCInterval         time.Duration
	GCMissingTolerance time.Duration
	GCDropTolerance    time.Duration
}

func (p *dataCoordConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initChannelWatchPrefix()

	p.initSegmentMaxSize()
	p.initSegmentSealProportion()
	p.initSegAssignmentExpiration()

	p.initEnableCompaction()
	p.initEnableAutoCompaction()
	p.initCompactionEntityExpiration()

	p.initEnableGarbageCollection()
	p.initGCInterval()
	p.initGCMissingTolerance()
	p.initGCDropTolerance()
}

func (p *dataCoordConfig) initSegmentMaxSize() {
	p.SegmentMaxSize = p.BaseParams.ParseFloatWithDefault("dataCoord.segment.maxSize", 512.0)
}

func (p *dataCoordConfig) initSegmentSealProportion() {
	p.SegmentSealProportion = p.BaseParams.ParseFloatWithDefault("dataCoord.segment.sealProportion", 0.75)
}

func (p *dataCoordConfig) initSegAssignmentExpiration() {
	p.SegAssignmentExpiration = p.BaseParams.ParseInt64WithDefault("dataCoord.segment.assignmentExpiration", 2000)
}

func (p *dataCoordConfig) initChannelWatchPrefix() {
	// WARN: this value should not be put to milvus.yaml. It's a default value for channel watch path.
	// This will be removed after we reconstruct our config module.
	p.ChannelWatchSubPath = "channelwatch"
}

func (p *dataCoordConfig) initEnableCompaction() {
	p.EnableCompaction = p.BaseParams.ParseBool("dataCoord.enableCompaction", false)
}

// -- GC --
func (p *dataCoordConfig) initEnableGarbageCollection() {
	p.EnableGarbageCollection = p.BaseParams.ParseBool("dataCoord.enableGarbageCollection", false)
}

func (p *dataCoordConfig) initGCInterval() {
	p.GCInterval = time.Duration(p.BaseParams.ParseInt64WithDefault("dataCoord.gc.interval", 60*60)) * time.Second
}

func (p *dataCoordConfig) initGCMissingTolerance() {
	p.GCMissingTolerance = time.Duration(p.BaseParams.ParseInt64WithDefault("dataCoord.gc.missingTolerance", 24*60*60)) * time.Second
}

func (p *dataCoordConfig) initGCDropTolerance() {
	p.GCDropTolerance = time.Duration(p.BaseParams.ParseInt64WithDefault("dataCoord.gc.dropTolerance", 24*60*60)) * time.Second
}

func (p *dataCoordConfig) initEnableAutoCompaction() {
	p.EnableAutoCompaction = p.BaseParams.ParseBool("dataCoord.compaction.enableAutoCompaction", false)
}

func (p *dataCoordConfig) initCompactionEntityExpiration() {
	p.CompactionEntityExpiration = p.BaseParams.ParseInt64WithDefault("dataCoord.compaction.entityExpiration", math.MaxInt64)
	p.CompactionEntityExpiration = func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	}(p.CompactionEntityExpiration, p.RetentionDuration)
}

///////////////////////////////////////////////////////////////////////////////
// --- datanode ---
type dataNodeConfig struct {
	BaseParams *BaseParamTable

	// ID of the current DataNode
	NodeID UniqueID

	// IP of the current DataNode
	IP string

	// Port of the current DataNode
	Port                    int
	FlowGraphMaxQueueLength int32
	FlowGraphMaxParallelism int32
	FlushInsertBufferSize   int64
	InsertBinlogRootPath    string
	StatsBinlogRootPath     string
	DeleteBinlogRootPath    string
	Alias                   string // Different datanode in one machine

	// Channel subscribition name -
	DataNodeSubName string

	// etcd
	ChannelWatchSubPath string

	CreatedTime time.Time
	UpdatedTime time.Time
}

func (p *dataNodeConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initFlowGraphMaxQueueLength()
	p.initFlowGraphMaxParallelism()
	p.initFlushInsertBufferSize()
	p.initInsertBinlogRootPath()
	p.initStatsBinlogRootPath()
	p.initDeleteBinlogRootPath()

	p.initDataNodeSubName()
	p.initChannelWatchPath()
}

// Refresh is called after session init
func (p *dataNodeConfig) Refresh() {
	p.initDataNodeSubName()
}

// InitAlias init this DataNode alias
func (p *dataNodeConfig) InitAlias(alias string) {
	p.Alias = alias
}

func (p *dataNodeConfig) initFlowGraphMaxQueueLength() {
	p.FlowGraphMaxQueueLength = p.BaseParams.ParseInt32WithDefault("dataNode.dataSync.flowGraph.maxQueueLength", 1024)
}

func (p *dataNodeConfig) initFlowGraphMaxParallelism() {
	p.FlowGraphMaxParallelism = p.BaseParams.ParseInt32WithDefault("dataNode.dataSync.flowGraph.maxParallelism", 1024)
}

func (p *dataNodeConfig) initFlushInsertBufferSize() {
	p.FlushInsertBufferSize = p.BaseParams.ParseInt64("_DATANODE_INSERTBUFSIZE")
}

func (p *dataNodeConfig) initInsertBinlogRootPath() {
	// GOOSE TODO: rootPath change to TenentID
	rootPath, err := p.BaseParams.Load("minio.rootPath")
	if err != nil {
		panic(err)
	}
	p.InsertBinlogRootPath = path.Join(rootPath, "insert_log")
}

func (p *dataNodeConfig) initStatsBinlogRootPath() {
	rootPath, err := p.BaseParams.Load("minio.rootPath")
	if err != nil {
		panic(err)
	}
	p.StatsBinlogRootPath = path.Join(rootPath, "stats_log")
}

func (p *dataNodeConfig) initDeleteBinlogRootPath() {
	rootPath, err := p.BaseParams.Load("minio.rootPath")
	if err != nil {
		panic(err)
	}
	p.DeleteBinlogRootPath = path.Join(rootPath, "delta_log")
}

func (p *dataNodeConfig) initDataNodeSubName() {
	cluster, err := p.BaseParams.Load("msgChannel.chanNamePrefix.cluster")
	if err != nil {
		panic(err)
	}
	subname, err := p.BaseParams.Load("msgChannel.subNamePrefix.dataNodeSubNamePrefix")
	if err != nil {
		panic(err)
	}
	s := []string{cluster, subname, strconv.FormatInt(p.NodeID, 10)}
	p.DataNodeSubName = strings.Join(s, "-")
}

func (p *dataNodeConfig) initChannelWatchPath() {
	p.ChannelWatchSubPath = "channelwatch"
}

///////////////////////////////////////////////////////////////////////////////
// --- indexcoord ---
type indexCoordConfig struct {
	BaseParams *BaseParamTable

	Address string
	Port    int

	IndexStorageRootPath string

	CreatedTime time.Time
	UpdatedTime time.Time
}

func (p *indexCoordConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initIndexStorageRootPath()
}

// initIndexStorageRootPath initializes the root path of index files.
func (p *indexCoordConfig) initIndexStorageRootPath() {
	rootPath, err := p.BaseParams.Load("minio.rootPath")
	if err != nil {
		panic(err)
	}
	p.IndexStorageRootPath = path.Join(rootPath, "index_files")
}

///////////////////////////////////////////////////////////////////////////////
// --- indexnode ---
type indexNodeConfig struct {
	BaseParams *BaseParamTable

	IP      string
	Address string
	Port    int

	NodeID int64
	Alias  string

	IndexStorageRootPath string

	CreatedTime time.Time
	UpdatedTime time.Time
}

func (p *indexNodeConfig) init(bp *BaseParamTable) {
	p.BaseParams = bp

	p.initIndexStorageRootPath()
}

// InitAlias initializes an alias for the IndexNode role.
func (p *indexNodeConfig) InitAlias(alias string) {
	p.Alias = alias
}

func (p *indexNodeConfig) initIndexStorageRootPath() {
	rootPath, err := p.BaseParams.Load("minio.rootPath")
	if err != nil {
		panic(err)
	}
	p.IndexStorageRootPath = path.Join(rootPath, "index_files")
}

///////////////////////////////////////////////////////////////////////////////
// --- grpc ---
type grpcConfig struct {
	BaseParamTable

	once   sync.Once
	Domain string
	IP     string
	Port   int
}

func (p *grpcConfig) init(domain string) {
	p.BaseParamTable.Init()
	p.Domain = domain

	p.LoadFromEnv()
	p.LoadFromArgs()
	p.initPort()
}

// LoadFromEnv is used to initialize configuration items from env.
func (p *grpcConfig) LoadFromEnv() {
	p.IP = ipv4.LocalIP()
}

// LoadFromArgs is used to initialize configuration items from args.
func (p *grpcConfig) LoadFromArgs() {

}

func (p *grpcConfig) initPort() {
	p.Port = p.ParseInt(p.Domain + ".port")
}

// GetAddress return grpc address
func (p *grpcConfig) GetAddress() string {
	return p.IP + ":" + strconv.Itoa(p.Port)
}

// GrpcServerConfig is configuration for grpc server.
type GrpcServerConfig struct {
	grpcConfig

	ServerMaxSendSize int
	ServerMaxRecvSize int
}

// InitOnce initialize grpc server config once
func (p *GrpcServerConfig) InitOnce(domain string) {
	p.once.Do(func() {
		p.init(domain)
	})
}

func (p *GrpcServerConfig) init(domain string) {
	p.grpcConfig.init(domain)

	p.initServerMaxSendSize()
	p.initServerMaxRecvSize()
}

func (p *GrpcServerConfig) initServerMaxSendSize() {
	var err error

	valueStr, err := p.Load(p.Domain + ".grpc.serverMaxSendSize")
	if err != nil {
		p.ServerMaxSendSize = DefaultServerMaxSendSize
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Warn("Failed to parse grpc.serverMaxSendSize, set to default",
			zap.String("rol", p.Domain), zap.String("grpc.serverMaxSendSize", valueStr),
			zap.Error(err))

		p.ServerMaxSendSize = DefaultServerMaxSendSize
	} else {
		p.ServerMaxSendSize = value
	}

	log.Debug("initServerMaxSendSize",
		zap.String("role", p.Domain), zap.Int("grpc.serverMaxSendSize", p.ServerMaxSendSize))
}

func (p *GrpcServerConfig) initServerMaxRecvSize() {
	var err error

	valueStr, err := p.Load(p.Domain + ".grpc.serverMaxRecvSize")
	if err != nil {
		p.ServerMaxRecvSize = DefaultServerMaxRecvSize
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Warn("Failed to parse grpc.serverMaxRecvSize, set to default",
			zap.String("role", p.Domain), zap.String("grpc.serverMaxRecvSize", valueStr),
			zap.Error(err))

		p.ServerMaxRecvSize = DefaultServerMaxRecvSize
	} else {
		p.ServerMaxRecvSize = value
	}

	log.Debug("initServerMaxRecvSize",
		zap.String("role", p.Domain), zap.Int("grpc.serverMaxRecvSize", p.ServerMaxRecvSize))
}

// GrpcClientConfig is configuration for grpc client.
type GrpcClientConfig struct {
	grpcConfig

	ClientMaxSendSize int
	ClientMaxRecvSize int
}

// InitOnce initialize grpc client config once
func (p *GrpcClientConfig) InitOnce(domain string) {
	p.once.Do(func() {
		p.init(domain)
	})
}

func (p *GrpcClientConfig) init(domain string) {
	p.grpcConfig.init(domain)

	p.initClientMaxSendSize()
	p.initClientMaxRecvSize()
}

func (p *GrpcClientConfig) initClientMaxSendSize() {
	var err error

	valueStr, err := p.Load(p.Domain + ".grpc.clientMaxSendSize")
	if err != nil {
		p.ClientMaxSendSize = DefaultClientMaxSendSize
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Warn("Failed to parse grpc.clientMaxSendSize, set to default",
			zap.String("role", p.Domain), zap.String("grpc.clientMaxSendSize", valueStr),
			zap.Error(err))

		p.ClientMaxSendSize = DefaultClientMaxSendSize
	} else {
		p.ClientMaxSendSize = value
	}

	log.Debug("initClientMaxSendSize",
		zap.String("role", p.Domain), zap.Int("grpc.clientMaxSendSize", p.ClientMaxSendSize))
}

func (p *GrpcClientConfig) initClientMaxRecvSize() {
	var err error

	valueStr, err := p.Load(p.Domain + ".grpc.clientMaxRecvSize")
	if err != nil {
		p.ClientMaxRecvSize = DefaultClientMaxRecvSize
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Warn("Failed to parse grpc.clientMaxRecvSize, set to default",
			zap.String("role", p.Domain), zap.String("grpc.clientMaxRecvSize", valueStr),
			zap.Error(err))

		p.ClientMaxRecvSize = DefaultClientMaxRecvSize
	} else {
		p.ClientMaxRecvSize = value
	}

	log.Debug("initClientMaxRecvSize",
		zap.String("role", p.Domain), zap.Int("grpc.clientMaxRecvSize", p.ClientMaxRecvSize))
}
