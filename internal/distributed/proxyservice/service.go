package grpcproxyservice

import (
	"context"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	"github.com/zilliztech/milvus-distributed/internal/proto/proxypb"
	"github.com/zilliztech/milvus-distributed/internal/proxyservice"
	"github.com/zilliztech/milvus-distributed/internal/util/funcutil"
	"google.golang.org/grpc"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	grpcServer  *grpc.Server
	grpcErrChan chan error

	impl *proxyservice.ServiceImpl
}

func NewServer(ctx1 context.Context) (*Server, error) {
	ctx, cancel := context.WithCancel(ctx1)

	server := &Server{
		ctx:         ctx,
		cancel:      cancel,
		grpcErrChan: make(chan error),
	}

	var err error
	server.impl, err = proxyservice.NewServiceImpl(server.ctx)
	if err != nil {
		return nil, err
	}
	return server, err
}

func (s *Server) Run() error {

	if err := s.init(); err != nil {
		return err
	}
	log.Println("proxy service init done ...")

	if err := s.start(); err != nil {
		return err
	}
	return nil
}

func (s *Server) init() error {
	Params.Init()
	proxyservice.Params.Init()
	log.Println("init params done")

	s.wg.Add(1)
	go s.startGrpcLoop(Params.ServicePort)
	// wait for grpc server loop start
	if err := <-s.grpcErrChan; err != nil {
		return err
	}
	s.impl.UpdateStateCode(internalpb2.StateCode_INITIALIZING)
	log.Println("grpc init done ...")

	if err := s.impl.Init(); err != nil {
		return err
	}
	return nil
}

func (s *Server) startGrpcLoop(grpcPort int) {

	defer s.wg.Done()

	log.Println("network port: ", grpcPort)
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(grpcPort))
	if err != nil {
		log.Printf("GrpcServer:failed to listen: %v", err)
		s.grpcErrChan <- err
		return
	}

	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	s.grpcServer = grpc.NewServer()
	proxypb.RegisterProxyServiceServer(s.grpcServer, s)
	milvuspb.RegisterProxyServiceServer(s.grpcServer, s)

	go funcutil.CheckGrpcReady(ctx, s.grpcErrChan)
	if err := s.grpcServer.Serve(lis); err != nil {
		s.grpcErrChan <- err
	}

}

func (s *Server) start() error {
	log.Println("proxy ServiceImpl start ...")
	if err := s.impl.Start(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	err := s.impl.Stop()
	if err != nil {
		return err
	}
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	s.cancel()
	s.wg.Wait()
	return nil
}

func (s *Server) RegisterLink(ctx context.Context, empty *commonpb.Empty) (*milvuspb.RegisterLinkResponse, error) {
	return s.impl.RegisterLink()
}

func (s *Server) RegisterNode(ctx context.Context, request *proxypb.RegisterNodeRequest) (*proxypb.RegisterNodeResponse, error) {
	return s.impl.RegisterNode(request)
}

func (s *Server) InvalidateCollectionMetaCache(ctx context.Context, request *proxypb.InvalidateCollMetaCacheRequest) (*commonpb.Status, error) {
	return s.impl.InvalidateCollectionMetaCache(request)
}

func (s *Server) GetTimeTickChannel(ctx context.Context, empty *commonpb.Empty) (*milvuspb.StringResponse, error) {
	return s.impl.GetTimeTickChannel()
}

func (s *Server) GetComponentStates(ctx context.Context, empty *commonpb.Empty) (*internalpb2.ComponentStates, error) {
	return s.impl.GetComponentStates()
}
