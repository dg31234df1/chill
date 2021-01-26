package master

import (
	"context"
	"log"
	"math"
	"sync/atomic"

	"github.com/zilliztech/milvus-distributed/internal/errors"
	ms "github.com/zilliztech/milvus-distributed/internal/msgstream"
)

type (
	TimeTickBarrier interface {
		GetTimeTick() (Timestamp, error)
		Start() error
		Close()
	}

	softTimeTickBarrier struct {
		peer2LastTt   map[UniqueID]Timestamp
		minTtInterval Timestamp
		lastTt        int64
		outTt         chan Timestamp
		ttStream      ms.MsgStream
		ctx           context.Context
		cancel        context.CancelFunc
	}

	hardTimeTickBarrier struct {
		peer2Tt  map[UniqueID]Timestamp
		outTt    chan Timestamp
		ttStream ms.MsgStream
		ctx      context.Context
		cancel   context.CancelFunc
	}

	proxyServiceTimeTickBarrier struct {
		ttStream ms.MsgStream
		ctx      context.Context
		cancel   context.CancelFunc
	}
)

func (ttBarrier *proxyServiceTimeTickBarrier) GetTimeTick() (Timestamp, error) {
	select {
	case <-ttBarrier.ctx.Done():
		return 0, errors.Errorf("[GetTimeTick] closed.")
	case ttmsgs := <-ttBarrier.ttStream.Chan():
		log.Println("ttmsgs: ", ttmsgs)
		tempMin := Timestamp(math.MaxUint64)
		for _, ttmsg := range ttmsgs.Msgs {
			timeTickMsg, ok := ttmsg.(*ms.TimeTickMsg)
			if !ok {
				log.Println("something wrong in time tick message!")
			}
			if timeTickMsg.Base.Timestamp < tempMin {
				tempMin = timeTickMsg.Base.Timestamp
			}
		}
		log.Println("[GetTimeTick]: ", tempMin)
		return tempMin, nil
	}
}

func (ttBarrier *proxyServiceTimeTickBarrier) Start() error {
	ttBarrier.ttStream.Start()
	return nil
}

func (ttBarrier *proxyServiceTimeTickBarrier) Close() {
	ttBarrier.ttStream.Close()
	ttBarrier.cancel()
}

func newProxyServiceTimeTickBarrier(ctx context.Context, stream ms.MsgStream) TimeTickBarrier {
	ctx1, cancel := context.WithCancel(ctx)
	return &proxyServiceTimeTickBarrier{
		ttStream: stream,
		ctx:      ctx1,
		cancel:   cancel,
	}
}

func (ttBarrier *softTimeTickBarrier) GetTimeTick() (Timestamp, error) {
	select {
	case <-ttBarrier.ctx.Done():
		return 0, errors.Errorf("[GetTimeTick] closed.")
	case ts, ok := <-ttBarrier.outTt:
		if !ok {
			return 0, errors.Errorf("[GetTimeTick] closed.")
		}
		num := len(ttBarrier.outTt)
		for i := 0; i < num; i++ {
			ts, ok = <-ttBarrier.outTt
			if !ok {
				return 0, errors.Errorf("[GetTimeTick] closed.")
			}
		}
		atomic.StoreInt64(&(ttBarrier.lastTt), int64(ts))
		return ts, ttBarrier.ctx.Err()
	}
}

func (ttBarrier *softTimeTickBarrier) Start() error {
	go func() {
		for {
			select {
			case <-ttBarrier.ctx.Done():
				log.Printf("[TtBarrierStart] %s\n", ttBarrier.ctx.Err())
				return

			case ttmsgs := <-ttBarrier.ttStream.Chan():
				if len(ttmsgs.Msgs) > 0 {
					for _, timetickmsg := range ttmsgs.Msgs {
						ttmsg := timetickmsg.(*ms.TimeTickMsg)
						oldT, ok := ttBarrier.peer2LastTt[ttmsg.Base.SourceID]
						// log.Printf("[softTimeTickBarrier] peer(%d)=%d\n", ttmsg.PeerID, ttmsg.Timestamp)

						if !ok {
							log.Printf("[softTimeTickBarrier] Warning: peerID %d not exist\n", ttmsg.Base.SourceID)
							continue
						}
						if ttmsg.Base.Timestamp > oldT {
							ttBarrier.peer2LastTt[ttmsg.Base.SourceID] = ttmsg.Base.Timestamp

							// get a legal Timestamp
							ts := ttBarrier.minTimestamp()
							lastTt := atomic.LoadInt64(&(ttBarrier.lastTt))
							if lastTt != 0 && ttBarrier.minTtInterval > ts-Timestamp(lastTt) {
								continue
							}
							ttBarrier.outTt <- ts
						}
					}
				}
			}
		}
	}()
	return nil
}

func newSoftTimeTickBarrier(ctx context.Context,
	ttStream *ms.MsgStream,
	peerIds []UniqueID,
	minTtInterval Timestamp) *softTimeTickBarrier {

	if len(peerIds) <= 0 {
		log.Printf("[newSoftTimeTickBarrier] Error: peerIds is empty!\n")
		return nil
	}

	sttbarrier := softTimeTickBarrier{}
	sttbarrier.minTtInterval = minTtInterval
	sttbarrier.ttStream = *ttStream
	sttbarrier.outTt = make(chan Timestamp, 1024)
	sttbarrier.ctx, sttbarrier.cancel = context.WithCancel(ctx)
	sttbarrier.peer2LastTt = make(map[UniqueID]Timestamp)
	for _, id := range peerIds {
		sttbarrier.peer2LastTt[id] = Timestamp(0)
	}
	if len(peerIds) != len(sttbarrier.peer2LastTt) {
		log.Printf("[newSoftTimeTickBarrier] Warning: there are duplicate peerIds!\n")
	}

	return &sttbarrier
}

func (ttBarrier *softTimeTickBarrier) Close() {
	ttBarrier.cancel()
}

func (ttBarrier *softTimeTickBarrier) minTimestamp() Timestamp {
	tempMin := Timestamp(math.MaxUint64)
	for _, tt := range ttBarrier.peer2LastTt {
		if tt < tempMin {
			tempMin = tt
		}
	}
	return tempMin
}

func (ttBarrier *hardTimeTickBarrier) GetTimeTick() (Timestamp, error) {
	select {
	case <-ttBarrier.ctx.Done():
		return 0, errors.Errorf("[GetTimeTick] closed.")
	case ts, ok := <-ttBarrier.outTt:
		if !ok {
			return 0, errors.Errorf("[GetTimeTick] closed.")
		}
		return ts, ttBarrier.ctx.Err()
	}
}

func (ttBarrier *hardTimeTickBarrier) Start() error {
	go func() {
		// Last timestamp synchronized
		state := Timestamp(0)
		for {
			select {
			case <-ttBarrier.ctx.Done():
				log.Printf("[TtBarrierStart] %s\n", ttBarrier.ctx.Err())
				return

			case ttmsgs := <-ttBarrier.ttStream.Chan():
				if len(ttmsgs.Msgs) > 0 {
					for _, timetickmsg := range ttmsgs.Msgs {

						// Suppose ttmsg.Timestamp from stream is always larger than the previous one,
						// that `ttmsg.Timestamp > oldT`
						ttmsg := timetickmsg.(*ms.TimeTickMsg)

						oldT, ok := ttBarrier.peer2Tt[ttmsg.Base.SourceID]
						if !ok {
							log.Printf("[hardTimeTickBarrier] Warning: peerID %d not exist\n", ttmsg.Base.SourceID)
							continue
						}

						if oldT > state {
							log.Printf("[hardTimeTickBarrier] Warning: peer(%d) timestamp(%d) ahead\n",
								ttmsg.Base.SourceID, ttmsg.Base.Timestamp)
						}

						ttBarrier.peer2Tt[ttmsg.Base.SourceID] = ttmsg.Base.Timestamp

						newState := ttBarrier.minTimestamp()
						if newState > state {
							ttBarrier.outTt <- newState
							state = newState
						}
					}
				}
			}
		}
	}()
	return nil
}

func (ttBarrier *hardTimeTickBarrier) minTimestamp() Timestamp {
	tempMin := Timestamp(math.MaxUint64)
	for _, tt := range ttBarrier.peer2Tt {
		if tt < tempMin {
			tempMin = tt
		}
	}
	return tempMin
}

func newHardTimeTickBarrier(ctx context.Context,
	ttStream *ms.MsgStream,
	peerIds []UniqueID) *hardTimeTickBarrier {

	if len(peerIds) <= 0 {
		log.Printf("[newSoftTimeTickBarrier] Error: peerIds is empty!")
		return nil
	}

	sttbarrier := hardTimeTickBarrier{}
	sttbarrier.ttStream = *ttStream
	sttbarrier.outTt = make(chan Timestamp, 1024)
	sttbarrier.ctx, sttbarrier.cancel = context.WithCancel(ctx)

	sttbarrier.peer2Tt = make(map[UniqueID]Timestamp)
	for _, id := range peerIds {
		sttbarrier.peer2Tt[id] = Timestamp(0)
	}
	if len(peerIds) != len(sttbarrier.peer2Tt) {
		log.Printf("[newSoftTimeTickBarrier] Warning: there are duplicate peerIds!")
	}

	return &sttbarrier
}

func (ttBarrier *hardTimeTickBarrier) Close() {
	ttBarrier.cancel()
}
