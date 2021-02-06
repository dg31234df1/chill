package timesync

import (
	"context"
	"log"
	"math"
	"sync"
	"sync/atomic"

	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"

	"github.com/zilliztech/milvus-distributed/internal/errors"
	ms "github.com/zilliztech/milvus-distributed/internal/msgstream"
)

type (
	Timestamp = typeutil.Timestamp
	UniqueID  = typeutil.UniqueID

	TimeTickBarrier interface {
		GetTimeTick() (Timestamp, error)
		Start()
		Close()
	}

	softTimeTickBarrier struct {
		peer2LastTt   map[UniqueID]Timestamp
		minTtInterval Timestamp
		lastTt        int64
		outTt         chan Timestamp
		ttStream      ms.MsgStream
		ctx           context.Context
	}

	hardTimeTickBarrier struct {
		peer2Tt    map[UniqueID]Timestamp
		outTt      chan Timestamp
		ttStream   ms.MsgStream
		ctx        context.Context
		wg         sync.WaitGroup
		loopCtx    context.Context
		loopCancel context.CancelFunc
	}
)

func NewSoftTimeTickBarrier(ctx context.Context, ttStream ms.MsgStream, peerIds []UniqueID, minTtInterval Timestamp) *softTimeTickBarrier {
	if len(peerIds) <= 0 {
		log.Printf("[newSoftTimeTickBarrier] Error: peerIds is empty!\n")
		return nil
	}

	sttbarrier := softTimeTickBarrier{}
	sttbarrier.minTtInterval = minTtInterval
	sttbarrier.ttStream = ttStream
	sttbarrier.outTt = make(chan Timestamp, 1024)
	sttbarrier.peer2LastTt = make(map[UniqueID]Timestamp)
	sttbarrier.ctx = ctx
	for _, id := range peerIds {
		sttbarrier.peer2LastTt[id] = Timestamp(0)
	}
	if len(peerIds) != len(sttbarrier.peer2LastTt) {
		log.Printf("[newSoftTimeTickBarrier] Warning: there are duplicate peerIds!\n")
	}

	return &sttbarrier
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

func (ttBarrier *softTimeTickBarrier) Start() {
	for {
		select {
		case <-ttBarrier.ctx.Done():
			log.Printf("[TtBarrierStart] %s\n", ttBarrier.ctx.Err())
			return
		default:
		}
		ttmsgs := ttBarrier.ttStream.Consume()
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

func (ttBarrier *hardTimeTickBarrier) Start() {
	// Last timestamp synchronized
	ttBarrier.wg.Add(1)
	ttBarrier.loopCtx, ttBarrier.loopCancel = context.WithCancel(ttBarrier.ctx)
	state := Timestamp(0)
	go func(ctx context.Context) {
		defer ttBarrier.wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Printf("[TtBarrierStart] %s\n", ttBarrier.ctx.Err())
				return
			default:
			}
			ttmsgs := ttBarrier.ttStream.Consume()
			if len(ttmsgs.Msgs) > 0 {
				log.Printf("receive tt msg")
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
					log.Printf("new state %d", newState)
					if newState > state {
						ttBarrier.outTt <- newState
						log.Printf("outtttt")
						state = newState
					}
				}
			}
		}
	}(ttBarrier.loopCtx)
}

func (ttBarrier *hardTimeTickBarrier) Close() {
	ttBarrier.loopCancel()
	ttBarrier.wg.Wait()
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

func NewHardTimeTickBarrier(ctx context.Context, ttStream ms.MsgStream, peerIds []UniqueID) *hardTimeTickBarrier {
	if len(peerIds) <= 0 {
		log.Printf("[newSoftTimeTickBarrier] Error: peerIds is empty!")
		return nil
	}

	sttbarrier := hardTimeTickBarrier{}
	sttbarrier.ttStream = ttStream
	sttbarrier.outTt = make(chan Timestamp, 1024)

	sttbarrier.peer2Tt = make(map[UniqueID]Timestamp)
	sttbarrier.ctx = ctx
	for _, id := range peerIds {
		sttbarrier.peer2Tt[id] = Timestamp(0)
	}
	if len(peerIds) != len(sttbarrier.peer2Tt) {
		log.Printf("[newSoftTimeTickBarrier] Warning: there are duplicate peerIds!")
	}

	return &sttbarrier
}
