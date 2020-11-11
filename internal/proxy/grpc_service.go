package proxy

import (
	"context"
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"log"

	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/servicepb"
)

func (p *Proxy) Insert(ctx context.Context, in *servicepb.RowBatch) (*servicepb.IntegerRangeResponse, error) {
	it := &InsertTask{
		BaseInsertTask: BaseInsertTask{
			BaseMsg: msgstream.BaseMsg{
				HashValues: in.HashKeys,
			},
			InsertRequest: internalpb.InsertRequest{
				MsgType:        internalpb.MsgType_kInsert,
				CollectionName: in.CollectionName,
				PartitionTag:   in.PartitionTag,
				RowData:        in.RowData,
			},
		},
		done:                  make(chan error),
		resultChan:            make(chan *servicepb.IntegerRangeResponse),
		manipulationMsgStream: p.manipulationMsgStream,
	}
	it.ctx, it.cancel = context.WithCancel(ctx)
	// TODO: req_id, segment_id, channel_id, proxy_id, timestamps, row_ids

	defer it.cancel()

	p.taskSch.DmQueue.Enqueue(it)
	select {
	case <-ctx.Done():
		log.Print("insert timeout!")
		return &servicepb.IntegerRangeResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "insert timeout!",
			},
		}, errors.New("insert timeout!")
	case result := <-it.resultChan:
		return result, nil
	}
}

func (p *Proxy) CreateCollection(ctx context.Context, req *schemapb.CollectionSchema) (*commonpb.Status, error) {
	cct := &CreateCollectionTask{
		CreateCollectionRequest: internalpb.CreateCollectionRequest{
			MsgType: internalpb.MsgType_kCreateCollection,
			Schema:  &commonpb.Blob{},
			// TODO: req_id, timestamp, proxy_id
		},
		masterClient: p.masterClient,
		done:         make(chan error),
		resultChan:   make(chan *commonpb.Status),
	}
	schemaBytes, _ := proto.Marshal(req)
	cct.CreateCollectionRequest.Schema.Value = schemaBytes
	cct.ctx, cct.cancel = context.WithCancel(ctx)
	defer cct.cancel()

	p.taskSch.DdQueue.Enqueue(cct)
	select {
	case <-ctx.Done():
		log.Print("create collection timeout!")
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    "create collection timeout!",
		}, errors.New("create collection timeout!")
	case result := <-cct.resultChan:
		return result, nil
	}
}

func (p *Proxy) Search(ctx context.Context, req *servicepb.Query) (*servicepb.QueryResult, error) {
	qt := &QueryTask{
		SearchRequest: internalpb.SearchRequest{
			MsgType: internalpb.MsgType_kSearch,
			Query:   &commonpb.Blob{},
			// TODO: req_id, proxy_id, timestamp, result_channel_id
		},
		queryMsgStream: p.queryMsgStream,
		done:           make(chan error),
		resultBuf:      make(chan []*internalpb.SearchResult),
		resultChan:     make(chan *servicepb.QueryResult),
	}
	qt.ctx, qt.cancel = context.WithCancel(ctx)
	queryBytes, _ := proto.Marshal(req)
	qt.SearchRequest.Query.Value = queryBytes
	defer qt.cancel()

	p.taskSch.DqQueue.Enqueue(qt)
	select {
	case <-ctx.Done():
		log.Print("query timeout!")
		return &servicepb.QueryResult{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "query timeout!",
			},
		}, errors.New("query timeout!")
	case result := <-qt.resultChan:
		return result, nil
	}
}

func (p *Proxy) DropCollection(ctx context.Context, req *servicepb.CollectionName) (*commonpb.Status, error) {
	dct := &DropCollectionTask{
		DropCollectionRequest: internalpb.DropCollectionRequest{
			MsgType: internalpb.MsgType_kDropCollection,
			// TODO: req_id, timestamp, proxy_id
			CollectionName: req,
		},
		masterClient: p.masterClient,
		done:         make(chan error),
		resultChan:   make(chan *commonpb.Status),
	}
	dct.ctx, dct.cancel = context.WithCancel(ctx)
	defer dct.cancel()

	p.taskSch.DdQueue.Enqueue(dct)
	select {
	case <-ctx.Done():
		log.Print("create collection timeout!")
		return &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
			Reason:    "create collection timeout!",
		}, errors.New("create collection timeout!")
	case result := <-dct.resultChan:
		return result, nil
	}
}

func (p *Proxy) HasCollection(ctx context.Context, req *servicepb.CollectionName) (*servicepb.BoolResponse, error) {
	hct := &HasCollectionTask{
		HasCollectionRequest: internalpb.HasCollectionRequest{
			MsgType: internalpb.MsgType_kHasCollection,
			// TODO: req_id, timestamp, proxy_id
			CollectionName: req,
		},
		masterClient: p.masterClient,
		done:         make(chan error),
		resultChan:   make(chan *servicepb.BoolResponse),
	}
	hct.ctx, hct.cancel = context.WithCancel(ctx)
	defer hct.cancel()

	p.taskSch.DqQueue.Enqueue(hct)
	select {
	case <-ctx.Done():
		log.Print("has collection timeout!")
		return &servicepb.BoolResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "has collection timeout!",
			},
			Value: false,
		}, errors.New("has collection timeout!")
	case result := <-hct.resultChan:
		return result, nil
	}
}

func (p *Proxy) DescribeCollection(ctx context.Context, req *servicepb.CollectionName) (*servicepb.CollectionDescription, error) {
	dct := &DescribeCollectionTask{
		DescribeCollectionRequest: internalpb.DescribeCollectionRequest{
			MsgType: internalpb.MsgType_kDescribeCollection,
			// TODO: req_id, timestamp, proxy_id
			CollectionName: req,
		},
		masterClient: p.masterClient,
		done:         make(chan error),
		resultChan:   make(chan *servicepb.CollectionDescription),
	}
	dct.ctx, dct.cancel = context.WithCancel(ctx)
	defer dct.cancel()

	p.taskSch.DqQueue.Enqueue(dct)
	select {
	case <-ctx.Done():
		log.Print("has collection timeout!")
		return &servicepb.CollectionDescription{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "describe collection timeout!",
			},
		}, errors.New("describe collection timeout!")
	case result := <-dct.resultChan:
		return result, nil
	}
}

func (p *Proxy) ShowCollections(ctx context.Context, req *commonpb.Empty) (*servicepb.StringListResponse, error) {
	sct := &ShowCollectionsTask{
		ShowCollectionRequest: internalpb.ShowCollectionRequest{
			MsgType: internalpb.MsgType_kDescribeCollection,
			// TODO: req_id, timestamp, proxy_id
		},
		masterClient: p.masterClient,
		done:         make(chan error),
		resultChan:   make(chan *servicepb.StringListResponse),
	}
	sct.ctx, sct.cancel = context.WithCancel(ctx)
	defer sct.cancel()

	p.taskSch.DqQueue.Enqueue(sct)
	select {
	case <-ctx.Done():
		log.Print("show collections timeout!")
		return &servicepb.StringListResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UNEXPECTED_ERROR,
				Reason:    "show collections timeout!",
			},
		}, errors.New("show collections timeout!")
	case result := <-sct.resultChan:
		return result, nil
	}
}

func (p *Proxy) CreatePartition(ctx context.Context, in *servicepb.PartitionName) (*commonpb.Status, error) {
	return &commonpb.Status{
		ErrorCode: 0,
		Reason:    "",
	}, nil
}

func (p *Proxy) DropPartition(ctx context.Context, in *servicepb.PartitionName) (*commonpb.Status, error) {
	return &commonpb.Status{
		ErrorCode: 0,
		Reason:    "",
	}, nil
}

func (p *Proxy) HasPartition(ctx context.Context, in *servicepb.PartitionName) (*servicepb.BoolResponse, error) {
	return &servicepb.BoolResponse{
		Status: &commonpb.Status{
			ErrorCode: 0,
			Reason:    "",
		},
		Value: true,
	}, nil
}

func (p *Proxy) DescribePartition(ctx context.Context, in *servicepb.PartitionName) (*servicepb.PartitionDescription, error) {
	return &servicepb.PartitionDescription{
		Status: &commonpb.Status{
			ErrorCode: 0,
			Reason:    "",
		},
	}, nil
}

func (p *Proxy) ShowPartitions(ctx context.Context, req *servicepb.CollectionName) (*servicepb.StringListResponse, error) {
	return &servicepb.StringListResponse{
		Status: &commonpb.Status{
			ErrorCode: 0,
			Reason:    "",
		},
	}, nil
}
