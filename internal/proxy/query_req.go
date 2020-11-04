package proxy

import (
	"log"
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/golang/protobuf/proto"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/servicepb"
)

type queryReq struct {
	internalpb.SearchRequest
	result []*internalpb.SearchResult
	wg     sync.WaitGroup
	proxy  *proxyServer
}

// BaseRequest interfaces
func (req *queryReq) Type() internalpb.MsgType {
	return req.MsgType
}

func (req *queryReq) PreExecute() commonpb.Status {
	return commonpb.Status{ErrorCode: commonpb.ErrorCode_SUCCESS}
}

func (req *queryReq) Execute() commonpb.Status {
	req.proxy.reqSch.queryChan <- req
	return commonpb.Status{ErrorCode: commonpb.ErrorCode_SUCCESS}
}

func (req *queryReq) PostExecute() commonpb.Status { // send into pulsar
	req.wg.Add(1)
	return commonpb.Status{ErrorCode: commonpb.ErrorCode_SUCCESS}
}

func (req *queryReq) WaitToFinish() commonpb.Status { // wait unitl send into pulsar
	req.wg.Wait()
	return commonpb.Status{ErrorCode: commonpb.ErrorCode_SUCCESS}
}

func (s *proxyServer) restartQueryRoutine(buf_size int) error {
	s.reqSch.queryChan = make(chan *queryReq, buf_size)
	pulsarClient, err := pulsar.NewClient(pulsar.ClientOptions{URL: s.pulsarAddr})
	if err != nil {
		return nil
	}
	query, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{Topic: s.queryTopic})
	if err != nil {
		return err
	}

	result, err := pulsarClient.Subscribe(pulsar.ConsumerOptions{
		Topic:                       s.resultTopic,
		SubscriptionName:            s.resultGroup,
		Type:                        pulsar.KeyShared,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
	})
	if err != nil {
		return err
	}

	resultMap := make(map[UniqueID]*queryReq)

	go func() {
		defer result.Close()
		defer query.Close()
		defer pulsarClient.Close()
		for {
			select {
			case <-s.ctx.Done():
				return
			case qm := <-s.reqSch.queryChan:
				ts, err := s.getTimestamp(1)
				if err != nil {
					log.Printf("get time stamp failed")
					break
				}
				qm.Timestamp = uint64(ts[0])

				qb, err := proto.Marshal(qm)
				if err != nil {
					log.Printf("Marshal QueryReqMsg failed, error = %v", err)
					continue
				}
				if _, err := query.Send(s.ctx, &pulsar.ProducerMessage{Payload: qb}); err != nil {
					log.Printf("post into puslar failed, error = %v", err)
				}
				s.reqSch.qTimestampMux.Lock()
				if s.reqSch.qTimestamp <= ts[0] {
					s.reqSch.qTimestamp = ts[0]
				} else {
					log.Printf("there is some wrong with q_timestamp, it goes back, current = %d, previous = %d", ts[0], s.reqSch.qTimestamp)
				}
				s.reqSch.qTimestampMux.Unlock()
				resultMap[qm.ReqId] = qm
				//log.Printf("start search, query id = %d", qm.QueryId)
			case cm, ok := <-result.Chan():
				if !ok {
					log.Printf("consumer of result topic has closed")
					return
				}
				var rm internalpb.SearchResult
				if err := proto.Unmarshal(cm.Message.Payload(), &rm); err != nil {
					log.Printf("Unmarshal QueryReqMsg failed, error = %v", err)
					break
				}
				if rm.ProxyId != s.proxyId {
					break
				}
				qm, ok := resultMap[rm.ReqId]
				if !ok {
					log.Printf("unknown query id = %d", rm.ReqId)
					break
				}
				qm.result = append(qm.result, &rm)
				if len(qm.result) == s.numReaderNode {
					qm.wg.Done()
					delete(resultMap, rm.ReqId)
				}
				result.AckID(cm.ID())
			}

		}
	}()
	return nil
}

//func (s *proxyServer) reduceResult(query *queryReq) *servicepb.QueryResult {
//}

func (s *proxyServer) reduceResults(query *queryReq) *servicepb.QueryResult {

	var results []*internalpb.SearchResult
	var status commonpb.Status
	status.ErrorCode = commonpb.ErrorCode_UNEXPECTED_ERROR
	for _, r := range query.result {
		status = *r.Status
		if status.ErrorCode == commonpb.ErrorCode_SUCCESS {
			results = append(results, r)
		} else {
			break
		}
	}
	if len(results) != s.numReaderNode {
		status.ErrorCode = commonpb.ErrorCode_UNEXPECTED_ERROR
	}
	if status.ErrorCode != commonpb.ErrorCode_SUCCESS {
		result := servicepb.QueryResult{
			Status: &status,
		}
		return &result
	}

	if s.numReaderNode == 1 {
		result := servicepb.QueryResult{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_SUCCESS,
			},
			Hits: results[0].Hits,
		}
		return &result
	}

	return &servicepb.QueryResult{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_SUCCESS,
		},
	}
}
