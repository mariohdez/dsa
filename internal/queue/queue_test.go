package queue_test

import (
	"dsa/playground/internal/queue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue_WhenEmptyQueue_ReturnsTrue(t *testing.T) {
	// Arrange.
	q := queue.NewRingBufferQueue(2)

	// Act.
	ok := q.Enqueue(1)
	q.Print()
	ok = q.Enqueue(2)
	q.Print()

	// Assert.
	assert.True(t, ok)
}

func TestEnqueue_WhenQueueIsNonEmpty_ReturnsTrue(t *testing.T) {
	// Arrange.
	q := queue.NewLLQueue()
	q.Enqueue(1)

	// Act.
	ok := q.Enqueue(2)

	// Assert.
	assert.True(t, ok)
}

func TestDequeue_WhenEmptyQueue_ReturnsFalse(t *testing.T) {
	// Arrange.
	q := queue.NewLLQueue()

	// Act.
	_, ok := q.Dequeue()

	// Assert.
	assert.False(t, ok)
}

func TestDequeue_WhenMultipleDequeuesCreatesEmptyQueue_ReturnsFalse(t *testing.T) {
	// Arrange.
	q := queue.NewLLQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	num, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, num)

	num, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, num)

	// Act.
	_, ok = q.Dequeue()

	// Assert.
	assert.False(t, ok)
}

func TestDequeue_WhenMultipleDequeuesCreatesEmptyQueueThenEnqueueANumber_ReturnsCorrectValueAndFalse(t *testing.T) {
	// Arrange.
	q := queue.NewLLQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	num, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, num)

	num, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, num)

	q.Enqueue(10)

	// Act.
	num, ok = q.Dequeue()

	// Assert.
	assert.True(t, ok)
	assert.Equal(t, 10, num)
}
