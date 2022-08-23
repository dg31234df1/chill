package sessionutil

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/blang/semver/v4"
	"github.com/milvus-io/milvus/internal/common"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/util/retry"
	"go.etcd.io/etcd/api/v3/mvccpb"
	v3rpc "go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

const (
	// DefaultServiceRoot default root path used in kv by Session
	DefaultServiceRoot = "session/"
	// DefaultIDKey default id key for Session
	DefaultIDKey = "id"
	// DefaultRetryTimes default retry times when registerService or getServerByID
	DefaultRetryTimes = 30
	// DefaultTTL default ttl value when granting a lease
	DefaultTTL = 60
)

// SessionEventType session event type
type SessionEventType int

// Rewatch defines the behavior outer session watch handles ErrCompacted
// it should process the current full list of session
// and returns err if meta error or anything else goes wrong
type Rewatch func(sessions map[string]*Session) error

const (
	// SessionNoneEvent place holder for zero value
	SessionNoneEvent SessionEventType = iota
	// SessionAddEvent event type for a new Session Added
	SessionAddEvent
	// SessionDelEvent event type for a Session deleted
	SessionDelEvent
)

// Session is a struct to store service's session, including ServerID, ServerName,
// Address.
// Exclusive indicates that this server can only start one.
type Session struct {
	ctx context.Context
	// When outside context done, Session cancels its goroutines first, then uses
	// keepAliveCancel to cancel the etcd KeepAlive
	keepAliveCancel context.CancelFunc

	ServerID    int64  `json:"ServerID,omitempty"`
	ServerName  string `json:"ServerName,omitempty"`
	Address     string `json:"Address,omitempty"`
	Exclusive   bool   `json:"Exclusive,omitempty"`
	TriggerKill bool
	Version     semver.Version `json:"Version,omitempty"`

	liveCh  <-chan bool
	etcdCli *clientv3.Client
	leaseID *clientv3.LeaseID

	metaRoot string

	registered atomic.Value
}

// UnmarshalJSON unmarshal bytes to Session.
func (s *Session) UnmarshalJSON(data []byte) error {
	var raw struct {
		ServerID    int64  `json:"ServerID,omitempty"`
		ServerName  string `json:"ServerName,omitempty"`
		Address     string `json:"Address,omitempty"`
		Exclusive   bool   `json:"Exclusive,omitempty"`
		TriggerKill bool
		Version     string `json:"Version"`
	}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	if raw.Version != "" {
		s.Version, err = semver.Parse(raw.Version)
		if err != nil {
			return err
		}
	}

	s.ServerID = raw.ServerID
	s.ServerName = raw.ServerName
	s.Address = raw.Address
	s.Exclusive = raw.Exclusive
	s.TriggerKill = raw.TriggerKill
	return nil
}

// MarshalJSON marshals session to bytes.
func (s *Session) MarshalJSON() ([]byte, error) {

	verStr := s.Version.String()
	return json.Marshal(&struct {
		ServerID    int64  `json:"ServerID,omitempty"`
		ServerName  string `json:"ServerName,omitempty"`
		Address     string `json:"Address,omitempty"`
		Exclusive   bool   `json:"Exclusive,omitempty"`
		TriggerKill bool
		Version     string `json:"Version"`
	}{
		ServerID:    s.ServerID,
		ServerName:  s.ServerName,
		Address:     s.Address,
		Exclusive:   s.Exclusive,
		TriggerKill: s.TriggerKill,
		Version:     verStr,
	})

}

// NewSession is a helper to build Session object.
// ServerID, ServerName, Address, Exclusive will be assigned after Init().
// metaRoot is a path in etcd to save session information.
// etcdEndpoints is to init etcdCli when NewSession
func NewSession(ctx context.Context, metaRoot string, client *clientv3.Client) *Session {
	session := &Session{
		ctx:      ctx,
		metaRoot: metaRoot,
		Version:  common.Version,
	}

	session.UpdateRegistered(false)

	connectEtcdFn := func() error {
		log.Debug("Session try to connect to etcd")
		ctx2, cancel2 := context.WithTimeout(session.ctx, 5*time.Second)
		defer cancel2()
		if _, err := client.Get(ctx2, "health"); err != nil {
			return err
		}
		session.etcdCli = client
		return nil
	}
	err := retry.Do(ctx, connectEtcdFn, retry.Attempts(100))
	if err != nil {
		log.Warn("failed to initialize session",
			zap.Error(err))
		return nil
	}
	log.Debug("Session connect to etcd success")
	return session
}

// Init will initialize base struct of the Session, including ServerName, ServerID,
// Address, Exclusive. ServerID is obtained in getServerID.
func (s *Session) Init(serverName, address string, exclusive bool, triggerKill bool) {
	s.ServerName = serverName
	s.Address = address
	s.Exclusive = exclusive
	s.TriggerKill = triggerKill
	s.checkIDExist()
	serverID, err := s.getServerID()
	if err != nil {
		panic(err)
	}
	s.ServerID = serverID
}

// String makes Session struct able to be logged by zap
func (s *Session) String() string {
	return fmt.Sprintf("Session:<ServerID: %d, ServerName: %s, Version: %s>", s.ServerID, s.ServerName, s.Version.String())
}

// Register will process keepAliveResponse to keep alive with etcd.
func (s *Session) Register() {
	ch, err := s.registerService()
	if err != nil {
		panic(err)
	}
	s.liveCh = s.processKeepAliveResponse(ch)
	s.UpdateRegistered(true)
}

func (s *Session) getServerID() (int64, error) {
	return s.getServerIDWithKey(DefaultIDKey)
}

func (s *Session) checkIDExist() {
	s.etcdCli.Txn(s.ctx).If(
		clientv3.Compare(
			clientv3.Version(path.Join(s.metaRoot, DefaultServiceRoot, DefaultIDKey)),
			"=",
			0)).
		Then(clientv3.OpPut(path.Join(s.metaRoot, DefaultServiceRoot, DefaultIDKey), "1")).Commit()
}

func (s *Session) getServerIDWithKey(key string) (int64, error) {
	for {
		getResp, err := s.etcdCli.Get(s.ctx, path.Join(s.metaRoot, DefaultServiceRoot, key))
		if err != nil {
			log.Warn("Session get etcd key error", zap.String("key", key), zap.Error(err))
			return -1, err
		}
		if getResp.Count <= 0 {
			log.Warn("Session there is no value", zap.String("key", key))
			continue
		}
		value := string(getResp.Kvs[0].Value)
		valueInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Warn("Session ParseInt error", zap.String("value", value), zap.Error(err))
			continue
		}
		txnResp, err := s.etcdCli.Txn(s.ctx).If(
			clientv3.Compare(
				clientv3.Value(path.Join(s.metaRoot, DefaultServiceRoot, key)),
				"=",
				value)).
			Then(clientv3.OpPut(path.Join(s.metaRoot, DefaultServiceRoot, key), strconv.FormatInt(valueInt+1, 10))).Commit()
		if err != nil {
			log.Warn("Session Txn failed", zap.String("key", key), zap.Error(err))
			return -1, err
		}

		if !txnResp.Succeeded {
			log.Warn("Session Txn unsuccessful", zap.String("key", key))
			continue
		}
		log.Debug("Session get serverID success", zap.String("key", key), zap.Int64("ServerId", valueInt))
		return valueInt, nil
	}
}

// registerService registers the service to etcd so that other services
// can find that the service is online and issue subsequent operations
// RegisterService will save a key-value in etcd
// key: metaRootPath + "/services" + "/ServerName-ServerID"
// value: json format
// {
//   ServerID   int64  `json:"ServerID,omitempty"`
//	 ServerName string `json:"ServerName,omitempty"`
//	 Address    string `json:"Address,omitempty"`
//   Exclusive  bool   `json:"Exclusive,omitempty"`
// }
// Exclusive means whether this service can exist two at the same time, if so,
// it is false. Otherwise, set it to true.
func (s *Session) registerService() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	var ch <-chan *clientv3.LeaseKeepAliveResponse
	log.Debug("service begin to register to etcd", zap.String("serverName", s.ServerName), zap.Int64("ServerID", s.ServerID))
	registerFn := func() error {
		resp, err := s.etcdCli.Grant(s.ctx, DefaultTTL)
		if err != nil {
			log.Error("register service", zap.Error(err))
			return err
		}
		s.leaseID = &resp.ID

		sessionJSON, err := json.Marshal(s)
		if err != nil {
			return err
		}

		key := s.ServerName
		if !s.Exclusive {
			key = fmt.Sprintf("%s-%d", key, s.ServerID)
		}
		txnResp, err := s.etcdCli.Txn(s.ctx).If(
			clientv3.Compare(
				clientv3.Version(path.Join(s.metaRoot, DefaultServiceRoot, key)),
				"=",
				0)).
			Then(clientv3.OpPut(path.Join(s.metaRoot, DefaultServiceRoot, key), string(sessionJSON), clientv3.WithLease(resp.ID))).Commit()

		if err != nil {
			log.Warn("compare and swap error, maybe the key has already been registered", zap.Error(err))
			return err
		}

		if !txnResp.Succeeded {
			return fmt.Errorf("function CompareAndSwap error for compare is false for key: %s", key)
		}

		keepAliveCtx, keepAliveCancel := context.WithCancel(context.Background())
		s.keepAliveCancel = keepAliveCancel
		ch, err = s.etcdCli.KeepAlive(keepAliveCtx, resp.ID)
		if err != nil {
			fmt.Printf("got error during keeping alive with etcd, err: %s\n", err)
			return err
		}
		log.Info("Service registered successfully", zap.String("ServerName", s.ServerName), zap.Int64("serverID", s.ServerID))
		return nil
	}
	err := retry.Do(s.ctx, registerFn, retry.Attempts(DefaultRetryTimes))
	if err != nil {
		return nil, err
	}
	return ch, nil
}

// processKeepAliveResponse processes the response of etcd keepAlive interface
// If keepAlive fails for unexpected error, it will send a signal to the channel.
func (s *Session) processKeepAliveResponse(ch <-chan *clientv3.LeaseKeepAliveResponse) (failChannel <-chan bool) {
	failCh := make(chan bool)
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				log.Warn("keep alive", zap.Error(errors.New("context done")))
				if s.keepAliveCancel != nil {
					s.keepAliveCancel()
				}
				return
			case resp, ok := <-ch:
				if !ok {
					log.Warn("session keepalive channel closed")
					close(failCh)
					return
				}
				if resp == nil {
					log.Warn("session keepalive response failed")
					close(failCh)
					return
				}
				//failCh <- true
			}
		}
	}()
	return failCh
}

// GetSessions will get all sessions registered in etcd.
// Revision is returned for WatchServices to prevent key events from being missed.
func (s *Session) GetSessions(prefix string) (map[string]*Session, int64, error) {
	res := make(map[string]*Session)
	key := path.Join(s.metaRoot, DefaultServiceRoot, prefix)
	resp, err := s.etcdCli.Get(s.ctx, key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, 0, err
	}
	for _, kv := range resp.Kvs {
		session := &Session{}
		err = json.Unmarshal(kv.Value, session)
		if err != nil {
			return nil, 0, err
		}
		_, mapKey := path.Split(string(kv.Key))
		log.Debug("SessionUtil GetSessions ", zap.Any("prefix", prefix),
			zap.String("key", mapKey),
			zap.Any("address", session.Address))
		res[mapKey] = session
	}
	return res, resp.Header.Revision, nil
}

// GetSessionsWithVersionRange will get all sessions with provided prefix and version range in etcd.
// Revision is returned for WatchServices to prevent missing events.
func (s *Session) GetSessionsWithVersionRange(prefix string, r semver.Range) (map[string]*Session, int64, error) {
	res := make(map[string]*Session)
	key := path.Join(s.metaRoot, DefaultServiceRoot, prefix)
	resp, err := s.etcdCli.Get(s.ctx, key, clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, 0, err
	}
	for _, kv := range resp.Kvs {
		session := &Session{}
		err = json.Unmarshal(kv.Value, session)
		if err != nil {
			return nil, 0, err
		}
		if !r(session.Version) {
			log.Debug("Session version out of range", zap.String("version", session.Version.String()), zap.Int64("serverID", session.ServerID))
			continue
		}
		_, mapKey := path.Split(string(kv.Key))
		log.Debug("SessionUtil GetSessions ", zap.String("prefix", prefix),
			zap.String("key", mapKey),
			zap.String("address", session.Address))
		res[mapKey] = session
	}
	return res, resp.Header.Revision, nil
}

// SessionEvent indicates the changes of other servers.
// if a server is up, EventType is SessAddEvent.
// if a server is down, EventType is SessDelEvent.
// Session Saves the changed server's information.
type SessionEvent struct {
	EventType SessionEventType
	Session   *Session
}

type sessionWatcher struct {
	s        *Session
	rch      clientv3.WatchChan
	eventCh  chan *SessionEvent
	prefix   string
	rewatch  Rewatch
	validate func(*Session) bool
}

func (w *sessionWatcher) start() {
	go func() {
		for {
			select {
			case <-w.s.ctx.Done():
				return
			case wresp, ok := <-w.rch:
				if !ok {
					log.Warn("session watch channel closed")
					return
				}
				w.handleWatchResponse(wresp)
			}
		}
	}()
}

// WatchServices watches the service's up and down in etcd, and sends event to
// eventChannel.
// prefix is a parameter to know which service to watch and can be obtained in
// typeutil.type.go.
// revision is a etcd reversion to prevent missing key events and can be obtained
// in GetSessions.
// If a server up, an event will be add to channel with eventType SessionAddType.
// If a server down, an event will be add to channel with eventType SessionDelType.
func (s *Session) WatchServices(prefix string, revision int64, rewatch Rewatch) (eventChannel <-chan *SessionEvent) {
	w := &sessionWatcher{
		s:        s,
		eventCh:  make(chan *SessionEvent, 100),
		rch:      s.etcdCli.Watch(s.ctx, path.Join(s.metaRoot, DefaultServiceRoot, prefix), clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(revision)),
		prefix:   prefix,
		rewatch:  rewatch,
		validate: func(s *Session) bool { return true },
	}
	w.start()
	return w.eventCh
}

// WatchServicesWithVersionRange watches the service's up and down in etcd, and sends event toeventChannel.
// Acts like WatchServices but with extra version range check.
// prefix is a parameter to know which service to watch and can be obtained intypeutil.type.go.
// revision is a etcd reversion to prevent missing key events and can be obtained in GetSessions.
// If a server up, an event will be add to channel with eventType SessionAddType.
// If a server down, an event will be add to channel with eventType SessionDelType.
func (s *Session) WatchServicesWithVersionRange(prefix string, r semver.Range, revision int64, rewatch Rewatch) (eventChannel <-chan *SessionEvent) {
	w := &sessionWatcher{
		s:        s,
		eventCh:  make(chan *SessionEvent, 100),
		rch:      s.etcdCli.Watch(s.ctx, path.Join(s.metaRoot, DefaultServiceRoot, prefix), clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(revision)),
		prefix:   prefix,
		rewatch:  rewatch,
		validate: func(s *Session) bool { return r(s.Version) },
	}
	w.start()
	return w.eventCh
}

func (w *sessionWatcher) handleWatchResponse(wresp clientv3.WatchResponse) {
	if wresp.Err() != nil {
		err := w.handleWatchErr(wresp.Err())
		if err != nil {
			log.Error("failed to handle watch session response", zap.Error(err))
			panic(err)
		}
		return
	}
	for _, ev := range wresp.Events {
		session := &Session{}
		var eventType SessionEventType
		switch ev.Type {
		case mvccpb.PUT:
			log.Debug("watch services",
				zap.Any("add kv", ev.Kv))
			err := json.Unmarshal([]byte(ev.Kv.Value), session)
			if err != nil {
				log.Error("watch services", zap.Error(err))
				continue
			}
			if !w.validate(session) {
				continue
			}
			eventType = SessionAddEvent
		case mvccpb.DELETE:
			log.Debug("watch services",
				zap.Any("delete kv", ev.PrevKv))
			err := json.Unmarshal([]byte(ev.PrevKv.Value), session)
			if err != nil {
				log.Error("watch services", zap.Error(err))
				continue
			}
			if !w.validate(session) {
				continue
			}
			eventType = SessionDelEvent
		}
		log.Debug("WatchService", zap.Any("event type", eventType))
		w.eventCh <- &SessionEvent{
			EventType: eventType,
			Session:   session,
		}
	}
}

func (w *sessionWatcher) handleWatchErr(err error) error {
	// if not ErrCompacted, just close the channel
	if err != v3rpc.ErrCompacted {
		//close event channel
		log.Warn("Watch service found error", zap.Error(err))
		close(w.eventCh)
		return err
	}

	sessions, revision, err := w.s.GetSessions(w.prefix)
	if err != nil {
		log.Warn("GetSession before rewatch failed", zap.String("prefix", w.prefix), zap.Error(err))
		close(w.eventCh)
		return err
	}
	// rewatch is nil, no logic to handle
	if w.rewatch == nil {
		log.Warn("Watch service with ErrCompacted but no rewatch logic provided")
	} else {
		err = w.rewatch(sessions)
	}
	if err != nil {
		log.Warn("WatchServices rewatch failed", zap.String("prefix", w.prefix), zap.Error(err))
		close(w.eventCh)
		return err
	}

	w.rch = w.s.etcdCli.Watch(w.s.ctx, path.Join(w.s.metaRoot, DefaultServiceRoot, w.prefix), clientv3.WithPrefix(), clientv3.WithPrevKV(), clientv3.WithRev(revision))
	return nil
}

// LivenessCheck performs liveness check with provided context and channel
// ctx controls the liveness check loop
// ch is the liveness signal channel, ch is closed only when the session is expired
// callback is the function to call when ch is closed, note that callback will not be invoked when loop exits due to context
func (s *Session) LivenessCheck(ctx context.Context, callback func()) {
	for {
		select {
		case _, ok := <-s.liveCh:
			// ok, still alive
			if ok {
				continue
			}
			// not ok, connection lost
			log.Warn("connection lost detected, shuting down")
			if callback != nil {
				go callback()
			}
			return
		case <-ctx.Done():
			log.Debug("liveness exits due to context done")
			// cancel the etcd keepAlive context
			if s.keepAliveCancel != nil {
				s.keepAliveCancel()
			}
			return
		}
	}
}

// Revoke revokes the internal leaseID for the session key
func (s *Session) Revoke(timeout time.Duration) {
	if s == nil {
		return
	}
	if s.etcdCli == nil || s.leaseID == nil {
		return
	}
	// can NOT use s.ctx, it may be Done here
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// ignores resp & error, just do best effort to revoke
	_, _ = s.etcdCli.Revoke(ctx, *s.leaseID)
}

// UpdateRegistered update the state of registered.
func (s *Session) UpdateRegistered(b bool) {
	s.registered.Store(b)
}

// Registered check if session was registered into etcd.
func (s *Session) Registered() bool {
	b, ok := s.registered.Load().(bool)
	if !ok {
		return false
	}
	return b
}
