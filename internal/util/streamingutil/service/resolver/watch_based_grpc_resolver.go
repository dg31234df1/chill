package resolver

import (
	"github.com/cockroachdb/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"

	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/lifetime"
)

var _ resolver.Resolver = (*watchBasedGRPCResolver)(nil)

// newWatchBasedGRPCResolver creates a new watch based grpc resolver.
func newWatchBasedGRPCResolver(cc resolver.ClientConn, logger *log.MLogger) *watchBasedGRPCResolver {
	return &watchBasedGRPCResolver{
		lifetime: lifetime.NewLifetime(lifetime.Working),
		cc:       cc,
		logger:   logger,
	}
}

// watchBasedGRPCResolver is a watch based grpc resolver.
type watchBasedGRPCResolver struct {
	lifetime lifetime.Lifetime[lifetime.State]

	cc     resolver.ClientConn
	logger *log.MLogger
}

// ResolveNow will be called by gRPC to try to resolve the target name
// again. It's just a hint, resolver can ignore this if it's not necessary.
//
// It could be called multiple times concurrently.
func (r *watchBasedGRPCResolver) ResolveNow(_ resolver.ResolveNowOptions) {
}

// Close closes the resolver.
// Do nothing.
func (r *watchBasedGRPCResolver) Close() {
	r.lifetime.SetState(lifetime.Stopped)
	r.lifetime.Wait()
	r.lifetime.Close()
}

func (r *watchBasedGRPCResolver) Update(state VersionedState) error {
	if r.lifetime.Add(lifetime.IsWorking) != nil {
		return errors.New("resolver is closed")
	}
	defer r.lifetime.Done()

	if err := r.cc.UpdateState(state.State); err != nil {
		// watch based resolver could ignore the error.
		r.logger.Warn("fail to update resolver state", zap.Error(err))
	}
	r.logger.Info("update resolver state success", zap.Any("state", state.State))
	return nil
}

// State returns the state of the resolver.
func (r *watchBasedGRPCResolver) State() lifetime.State {
	return r.lifetime.GetState()
}
