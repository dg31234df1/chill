package proxynode

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/util/paramtable"
)

type ParamTable struct {
	paramtable.BaseTable
}

var Params ParamTable

func (pt *ParamTable) Init() {
	pt.BaseTable.Init()

	err := pt.LoadYaml("advanced/proxy_node.yaml")
	if err != nil {
		panic(err)
	}

	proxyIDStr := os.Getenv("PROXY_ID")
	if proxyIDStr == "" {
		proxyIDList := pt.ProxyIDList()
		if len(proxyIDList) <= 0 {
			proxyIDStr = "0"
		} else {
			proxyIDStr = strconv.Itoa(int(proxyIDList[0]))
		}
	}
	pt.Save("_proxyID", proxyIDStr)
}

func (pt *ParamTable) NetworkPort() int {
	return pt.ParseInt("proxyNode.port")
}

func (pt *ParamTable) NetworkAddress() string {
	addr, err := pt.Load("proxyNode.address")
	if err != nil {
		panic(err)
	}

	hostName, _ := net.LookupHost(addr)
	if len(hostName) <= 0 {
		if ip := net.ParseIP(addr); ip == nil {
			panic("invalid ip proxyNode.address")
		}
	}

	port, err := pt.Load("proxyNode.port")
	if err != nil {
		panic(err)
	}
	_, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	return addr + ":" + port
}

func (pt *ParamTable) ProxyServiceAddress() string {
	addressFromEnv := os.Getenv("PROXY_SERVICE_ADDRESS")
	if len(addressFromEnv) > 0 {
		// TODO: or write to param table?
		return addressFromEnv
	}

	addr, err := pt.Load("proxyService.address")
	if err != nil {
		panic(err)
	}

	hostName, _ := net.LookupHost(addr)
	if len(hostName) <= 0 {
		if ip := net.ParseIP(addr); ip == nil {
			panic("invalid ip proxyService.address")
		}
	}

	port, err := pt.Load("proxyService.port")
	if err != nil {
		panic(err)
	}
	_, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	return addr + ":" + port
}

func (pt *ParamTable) MasterAddress() string {
	ret, err := pt.Load("_MasterAddress")
	if err != nil {
		panic(err)
	}
	return ret
}

func (pt *ParamTable) PulsarAddress() string {
	ret, err := pt.Load("_PulsarAddress")
	if err != nil {
		panic(err)
	}
	return ret
}

func (pt *ParamTable) ProxyNum() int {
	ret := pt.ProxyIDList()
	return len(ret)
}

func (pt *ParamTable) queryNodeNum() int {
	return len(pt.queryNodeIDList())
}

func (pt *ParamTable) queryNodeIDList() []UniqueID {
	queryNodeIDStr, err := pt.Load("nodeID.queryNodeIDList")
	if err != nil {
		panic(err)
	}
	var ret []UniqueID
	queryNodeIDs := strings.Split(queryNodeIDStr, ",")
	for _, i := range queryNodeIDs {
		v, err := strconv.Atoi(i)
		if err != nil {
			log.Panicf("load proxynode id list error, %s", err.Error())
		}
		ret = append(ret, UniqueID(v))
	}
	return ret
}

func (pt *ParamTable) ProxyID() UniqueID {
	proxyID, err := pt.Load("_proxyID")
	if err != nil {
		panic(err)
	}
	ID, err := strconv.Atoi(proxyID)
	if err != nil {
		panic(err)
	}
	return UniqueID(ID)
}

func (pt *ParamTable) TimeTickInterval() time.Duration {
	internalStr, err := pt.Load("proxyNode.timeTickInterval")
	if err != nil {
		panic(err)
	}
	interval, err := strconv.Atoi(internalStr)
	if err != nil {
		panic(err)
	}
	return time.Duration(interval) * time.Millisecond
}

func (pt *ParamTable) sliceIndex() int {
	proxyID := pt.ProxyID()
	proxyIDList := pt.ProxyIDList()
	for i := 0; i < len(proxyIDList); i++ {
		if proxyID == proxyIDList[i] {
			return i
		}
	}
	return -1
}

func (pt *ParamTable) InsertChannelNames() []string {
	prefix, err := pt.Load("msgChannel.chanNamePrefix.insert")
	if err != nil {
		panic(err)
	}
	prefix += "-"
	iRangeStr, err := pt.Load("msgChannel.channelRange.insert")
	if err != nil {
		panic(err)
	}
	channelIDs := paramtable.ConvertRangeToIntSlice(iRangeStr, ",")
	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}

	proxyNum := pt.ProxyNum()
	sep := len(channelIDs) / proxyNum
	index := pt.sliceIndex()
	if index == -1 {
		panic("ProxyID not Match with Config")
	}
	start := index * sep
	return ret[start : start+sep]
}

func (pt *ParamTable) DeleteChannelNames() []string {
	prefix, err := pt.Load("msgChannel.chanNamePrefix.delete")
	if err != nil {
		panic(err)
	}
	prefix += "-"
	dRangeStr, err := pt.Load("msgChannel.channelRange.delete")
	if err != nil {
		panic(err)
	}
	channelIDs := paramtable.ConvertRangeToIntSlice(dRangeStr, ",")
	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}
	return ret
}

func (pt *ParamTable) K2SChannelNames() []string {
	prefix, err := pt.Load("msgChannel.chanNamePrefix.k2s")
	if err != nil {
		panic(err)
	}
	prefix += "-"
	k2sRangeStr, err := pt.Load("msgChannel.channelRange.k2s")
	if err != nil {
		panic(err)
	}
	channelIDs := paramtable.ConvertRangeToIntSlice(k2sRangeStr, ",")
	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}
	return ret
}

func (pt *ParamTable) SearchChannelNames() []string {
	prefix, err := pt.Load("msgChannel.chanNamePrefix.search")
	if err != nil {
		panic(err)
	}
	prefix += "-"
	sRangeStr, err := pt.Load("msgChannel.channelRange.search")
	if err != nil {
		panic(err)
	}
	channelIDs := paramtable.ConvertRangeToIntSlice(sRangeStr, ",")
	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}
	return ret
}

func (pt *ParamTable) SearchResultChannelNames() []string {
	prefix, err := pt.Load("msgChannel.chanNamePrefix.searchResult")
	if err != nil {
		panic(err)
	}
	prefix += "-"
	sRangeStr, err := pt.Load("msgChannel.channelRange.searchResult")
	if err != nil {
		panic(err)
	}
	channelIDs := paramtable.ConvertRangeToIntSlice(sRangeStr, ",")
	var ret []string
	for _, ID := range channelIDs {
		ret = append(ret, prefix+strconv.Itoa(ID))
	}
	proxyNum := pt.ProxyNum()
	sep := len(channelIDs) / proxyNum
	index := pt.sliceIndex()
	if index == -1 {
		panic("ProxyID not Match with Config")
	}
	start := index * sep

	return ret[start : start+sep]
}

func (pt *ParamTable) ProxySubName() string {
	prefix, err := pt.Load("msgChannel.subNamePrefix.proxySubNamePrefix")
	if err != nil {
		panic(err)
	}
	proxyIDStr, err := pt.Load("_proxyID")
	if err != nil {
		panic(err)
	}
	return prefix + "-" + proxyIDStr
}

func (pt *ParamTable) ProxyTimeTickChannelNames() []string {
	prefix, err := pt.Load("msgChannel.chanNamePrefix.proxyTimeTick")
	if err != nil {
		panic(err)
	}
	prefix += "-0"
	return []string{prefix}
}

func (pt *ParamTable) DataDefinitionChannelNames() []string {
	prefix, err := pt.Load("msgChannel.chanNamePrefix.dataDefinition")
	if err != nil {
		panic(err)
	}
	prefix += "-0"
	return []string{prefix}
}

func (pt *ParamTable) MsgStreamInsertBufSize() int64 {
	return pt.ParseInt64("proxyNode.msgStream.insert.bufSize")
}

func (pt *ParamTable) MsgStreamSearchBufSize() int64 {
	return pt.ParseInt64("proxyNode.msgStream.search.bufSize")
}

func (pt *ParamTable) MsgStreamSearchResultBufSize() int64 {
	return pt.ParseInt64("proxyNode.msgStream.searchResult.recvBufSize")
}

func (pt *ParamTable) MsgStreamSearchResultPulsarBufSize() int64 {
	return pt.ParseInt64("proxyNode.msgStream.searchResult.pulsarBufSize")
}

func (pt *ParamTable) MsgStreamTimeTickBufSize() int64 {
	return pt.ParseInt64("proxyNode.msgStream.timeTick.bufSize")
}

func (pt *ParamTable) MaxNameLength() int64 {
	str, err := pt.Load("proxyNode.maxNameLength")
	if err != nil {
		panic(err)
	}
	maxNameLength, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return maxNameLength
}

func (pt *ParamTable) MaxFieldNum() int64 {
	str, err := pt.Load("proxyNode.maxFieldNum")
	if err != nil {
		panic(err)
	}
	maxFieldNum, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return maxFieldNum
}

func (pt *ParamTable) MaxDimension() int64 {
	str, err := pt.Load("proxyNode.maxDimension")
	if err != nil {
		panic(err)
	}
	maxDimension, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return maxDimension
}

func (pt *ParamTable) defaultPartitionTag() string {
	tag, err := pt.Load("common.defaultPartitionTag")
	if err != nil {
		panic(err)
	}
	return tag
}
