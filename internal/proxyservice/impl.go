package proxyservice

import (
	"context"
	"fmt"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/msgstream/util"

	"github.com/zilliztech/milvus-distributed/internal/msgstream/pulsarms"

	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"

	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"

	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/proto/proxypb"
)

const (
	timeoutInterval = time.Second * 10
)

func (s *ServiceImpl) fillNodeInitParams() error {
	s.nodeStartParams = make([]*commonpb.KeyValuePair, 0)
	nodeParams := &ParamTable{}
	nodeParams.Init()
	err := nodeParams.LoadYaml("advanced/proxy_node.yaml")
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) Init() error {
	err := s.fillNodeInitParams()
	if err != nil {
		return err
	}

	serviceTimeTickMsgStream := pulsarms.NewPulsarTtMsgStream(s.ctx, 1024)
	serviceTimeTickMsgStream.SetPulsarClient(Params.PulsarAddress())
	serviceTimeTickMsgStream.CreatePulsarProducers([]string{Params.ServiceTimeTickChannel()})

	nodeTimeTickMsgStream := pulsarms.NewPulsarMsgStream(s.ctx, 1024)
	nodeTimeTickMsgStream.SetPulsarClient(Params.PulsarAddress())
	nodeTimeTickMsgStream.CreatePulsarConsumers(Params.NodeTimeTickChannel(),
		"proxyservicesub", // TODO: add config
		util.NewUnmarshalDispatcher(),
		1024)

	ttBarrier := newSoftTimeTickBarrier(s.ctx, nodeTimeTickMsgStream, []UniqueID{0}, 10)
	s.tick = newTimeTick(s.ctx, ttBarrier, serviceTimeTickMsgStream)

	// dataServiceAddr := Params.DataServiceAddress()
	// s.dataServiceClient = dataservice.NewClient(dataServiceAddr)

	// insertChannelsRequest := &datapb.InsertChannelRequest{}
	// insertChannelNames, err := s.dataServiceClient.GetInsertChannels(insertChannelsRequest)
	// if err != nil {
	// 	return err
	// }

	// if len(insertChannelNames.Values) > 0 {
	// 	namesStr := strings.Join(insertChannelNames.Values, ",")
	// 	s.nodeStartParams = append(s.nodeStartParams, &commonpb.KeyValuePair{Key: KInsertChannelNames, Value: namesStr})
	// }

	s.state.State.StateCode = internalpb2.StateCode_HEALTHY

	return nil
}

func (s *ServiceImpl) Start() error {
	s.sched.Start()
	return s.tick.Start()
}

func (s *ServiceImpl) Stop() error {
	s.sched.Close()
	s.tick.Close()
	return nil
}

func (s *ServiceImpl) GetComponentStates() (*internalpb2.ComponentStates, error) {
	return s.state, nil
}

func (s *ServiceImpl) GetTimeTickChannel() (string, error) {
	return Params.ServiceTimeTickChannel(), nil
}

func (s *ServiceImpl) GetStatisticsChannel() (string, error) {
	panic("implement me")
}

func (s *ServiceImpl) RegisterLink() (*milvuspb.RegisterLinkResponse, error) {
	fmt.Println("register link")
	ctx, cancel := context.WithTimeout(s.ctx, timeoutInterval)
	defer cancel()

	t := &RegisterLinkTask{
		Condition: NewTaskCondition(ctx),
		nodeInfos: s.nodeInfos,
	}

	var err error

	err = s.sched.RegisterLinkTaskQueue.Enqueue(t)
	if err != nil {
		return &milvuspb.RegisterLinkResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
			Address: nil,
		}, nil
	}

	err = t.WaitToFinish()
	if err != nil {
		return &milvuspb.RegisterLinkResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
			Address: nil,
		}, nil
	}

	return t.response, nil
}

func (s *ServiceImpl) RegisterNode(request *proxypb.RegisterNodeRequest) (*proxypb.RegisterNodeResponse, error) {
	fmt.Println("RegisterNode: ", request)
	ctx, cancel := context.WithTimeout(s.ctx, timeoutInterval)
	defer cancel()

	t := &RegisterNodeTask{
		request:     request,
		startParams: s.nodeStartParams,
		Condition:   NewTaskCondition(ctx),
		allocator:   s.allocator,
		nodeInfos:   s.nodeInfos,
	}

	var err error

	err = s.sched.RegisterNodeTaskQueue.Enqueue(t)
	if err != nil {
		return &proxypb.RegisterNodeResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
			InitParams: nil,
		}, nil
	}

	err = t.WaitToFinish()
	if err != nil {
		return &proxypb.RegisterNodeResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    err.Error(),
			},
			InitParams: nil,
		}, nil
	}

	return t.response, nil
}

func (s *ServiceImpl) InvalidateCollectionMetaCache(request *proxypb.InvalidateCollMetaCacheRequest) error {
	fmt.Println("InvalidateCollectionMetaCache")
	ctx, cancel := context.WithTimeout(s.ctx, timeoutInterval)
	defer cancel()

	t := &InvalidateCollectionMetaCacheTask{
		request:   request,
		Condition: NewTaskCondition(ctx),
	}

	var err error

	err = s.sched.RegisterNodeTaskQueue.Enqueue(t)
	if err != nil {
		return err
	}

	err = t.WaitToFinish()
	if err != nil {
		return err
	}

	return nil
}
