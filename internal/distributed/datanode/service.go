package grpcdatanode

import (
	"context"
	"sync"
	"time"

	"log"
	"net"
	"strconv"

	"github.com/zilliztech/milvus-distributed/internal/util/funcutil"

	dn "github.com/zilliztech/milvus-distributed/internal/datanode"
	dsc "github.com/zilliztech/milvus-distributed/internal/distributed/dataservice/client"
	msc "github.com/zilliztech/milvus-distributed/internal/distributed/masterservice/client"
	"github.com/zilliztech/milvus-distributed/internal/errors"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"

	"google.golang.org/grpc"
)

type Server struct {
	impl        *dn.DataNode
	wg          sync.WaitGroup
	grpcErrChan chan error
	grpcServer  *grpc.Server
	ctx         context.Context
	cancel      context.CancelFunc

	msFactory msgstream.Factory

	masterService *msc.GrpcClient
	dataService   *dsc.Client
}

func New(ctx context.Context, factory msgstream.Factory) (*Server, error) {
	ctx1, cancel := context.WithCancel(ctx)
	var s = &Server{
		ctx:         ctx1,
		cancel:      cancel,
		msFactory:   factory,
		grpcErrChan: make(chan error),
	}

	s.impl = dn.NewDataNode(s.ctx, s.msFactory)

	return s, nil
}

func (s *Server) startGrpcLoop(grpcPort int) {
	defer s.wg.Done()

	addr := ":" + strconv.Itoa(grpcPort)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("DataNode GrpcServer:failed to listen: %v", err)
		s.grpcErrChan <- err
		return
	}
	log.Println("DataNode:: addr:", addr)

	s.grpcServer = grpc.NewServer()
	datapb.RegisterDataNodeServer(s.grpcServer, s)

	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	go funcutil.CheckGrpcReady(ctx, s.grpcErrChan)
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Println("DataNode Start Grpc Failed!!!!")
		s.grpcErrChan <- err
	}

}

func (s *Server) SetMasterServiceInterface(ms dn.MasterServiceInterface) error {
	return s.impl.SetMasterServiceInterface(ms)
}

func (s *Server) SetDataServiceInterface(ds dn.DataServiceInterface) error {
	return s.impl.SetDataServiceInterface(ds)
}

func (s *Server) Run() error {

	if err := s.init(); err != nil {
		return err
	}
	log.Println("data node init done ...")

	if err := s.start(); err != nil {
		return err
	}
	log.Println("data node start done ...")
	return nil
}

func (s *Server) Stop() error {
	s.cancel()
	var err error
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	err = s.impl.Stop()
	if err != nil {
		return err
	}
	s.wg.Wait()
	return nil
}

func (s *Server) init() error {
	Params.Init()
	Params.Port = funcutil.GetAvailablePort()
	Params.LoadFromEnv()
	Params.LoadFromArgs()

	log.Println("DataNode, port:", Params.Port)
	s.wg.Add(1)
	go s.startGrpcLoop(Params.Port)
	// wait for grpc server loop start
	err := <-s.grpcErrChan
	if err != nil {
		return err
	}

	// --- Master Server Client ---
	log.Println("Master service address:", Params.MasterAddress)
	log.Println("Init master service client ...")
	masterClient, err := msc.NewClient(Params.MasterAddress, 20*time.Second)
	if err != nil {
		panic(err)
	}

	if err = masterClient.Init(); err != nil {
		panic(err)
	}

	if err = masterClient.Start(); err != nil {
		panic(err)
	}
	err = funcutil.WaitForComponentHealthy(masterClient, "MasterService", 100, time.Millisecond*200)

	if err != nil {
		panic(err)
	}

	if err := s.SetMasterServiceInterface(masterClient); err != nil {
		panic(err)
	}

	// --- Data Server Client ---
	log.Println("Data service address: ", Params.DataServiceAddress)
	log.Println("DataNode Init data service client ...")
	dataService := dsc.NewClient(Params.DataServiceAddress)
	if err = dataService.Init(); err != nil {
		panic(err)
	}
	if err = dataService.Start(); err != nil {
		panic(err)
	}
	err = funcutil.WaitForComponentInitOrHealthy(dataService, "DataService", 100, time.Millisecond*200)
	if err != nil {
		panic(err)
	}
	if err := s.SetDataServiceInterface(dataService); err != nil {
		panic(err)
	}

	dn.Params.Init()
	dn.Params.Port = Params.Port
	dn.Params.IP = Params.IP

	s.impl.NodeID = dn.Params.NodeID
	s.impl.UpdateStateCode(internalpb2.StateCode_INITIALIZING)

	if err := s.impl.Init(); err != nil {
		log.Println("impl init error: ", err)
		return err
	}
	return nil
}

func (s *Server) start() error {
	return s.impl.Start()
}

func (s *Server) GetComponentStates(ctx context.Context, empty *commonpb.Empty) (*internalpb2.ComponentStates, error) {
	return s.impl.GetComponentStates()
}

func (s *Server) WatchDmChannels(ctx context.Context, in *datapb.WatchDmChannelRequest) (*commonpb.Status, error) {
	return s.impl.WatchDmChannels(in)
}

func (s *Server) FlushSegments(ctx context.Context, in *datapb.FlushSegRequest) (*commonpb.Status, error) {
	if s.impl.State.Load().(internalpb2.StateCode) != internalpb2.StateCode_HEALTHY {
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    "DataNode isn't healthy.",
		}, errors.Errorf("DataNode is not ready yet")
	}
	return &commonpb.Status{
		ErrorCode: commonpb.ErrorCode_SUCCESS,
	}, s.impl.FlushSegments(in)
}
