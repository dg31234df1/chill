package syncutil

// Future is a future value that can be set and retrieved.
type Future[T any] struct {
	ch    chan struct{}
	value T
}

// NewFuture creates a new future.
func NewFuture[T any]() *Future[T] {
	return &Future[T]{
		ch: make(chan struct{}),
	}
}

// Set sets the value of the future.
func (f *Future[T]) Set(value T) {
	f.value = value
	close(f.ch)
}

// Get retrieves the value of the future if set, otherwise block until set.
func (f *Future[T]) Get() T {
	<-f.ch
	return f.value
}

// Done returns a channel that is closed when the future is set.
func (f *Future[T]) Done() <-chan struct{} {
	return f.ch
}

// Ready returns true if the future is set.
func (f *Future[T]) Ready() bool {
	select {
	case <-f.ch:
		return true
	default:
		return false
	}
}
