package typeutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumHeap(t *testing.T) {
	h := []int{4, 5, 2}
	heap := NewArrayBasedMinimumHeap(h)
	assert.Equal(t, 2, heap.Peek())
	assert.Equal(t, 3, heap.Len())
	heap.Push(3)
	assert.Equal(t, 2, heap.Peek())
	assert.Equal(t, 4, heap.Len())
	heap.Push(1)
	assert.Equal(t, 1, heap.Peek())
	assert.Equal(t, 5, heap.Len())
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, heap.Peek())
		assert.Equal(t, i, heap.Pop())
	}
}

func TestMaximumHeap(t *testing.T) {
	h := []int{4, 1, 2}
	heap := NewArrayBasedMaximumHeap(h)
	assert.Equal(t, 4, heap.Peek())
	assert.Equal(t, 3, heap.Len())
	heap.Push(3)
	assert.Equal(t, 4, heap.Peek())
	assert.Equal(t, 4, heap.Len())
	heap.Push(5)
	assert.Equal(t, 5, heap.Peek())
	assert.Equal(t, 5, heap.Len())
	for i := 5; i >= 1; i-- {
		assert.Equal(t, i, heap.Peek())
		assert.Equal(t, i, heap.Pop())
	}
}
