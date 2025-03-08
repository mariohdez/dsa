package queue_test

import (
	"dsa/playground/internal/queue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	t.Run("When empty, returns true", func(t *testing.T) {
		// Arrange.
		q := queue.NewRingBufferQueue(2)

		// Act.
		ok := q.Enqueue(1)

		// Assert.
		assert.True(t, ok)

	})

	t.Run("When available capacity, returns true", func(t *testing.T) {
		// Arrange.
		q := queue.NewRingBufferQueue(2)
		ok := q.Enqueue(1)

		// Act.
		ok = q.Enqueue(2)

		// Assert.
		assert.True(t, ok)
	})

	t.Run("When available capacity, returns false", func(t *testing.T) {
		// Arrange.
		q := queue.NewRingBufferQueue(2)
		ok := q.Enqueue(1)
		ok = q.Enqueue(2)

		// Act.
		ok = q.Enqueue(3)

		// Assert.
		assert.False(t, ok)
	})
}

func TestDequeue(t *testing.T) {
	t.Run("When empty, returns false", func(t *testing.T) {
		// Arrange.
		q := queue.NewRingBufferQueue(2)

		// Act.
		_, deqOk := q.Dequeue()

		// Assert.
		assert.False(t, deqOk)
	})

	t.Run("When queue is not empty, returns first enqueued element", func(t *testing.T) {
		// Arrange.
		q := queue.NewRingBufferQueue(2)
		q.Enqueue(1)
		q.Enqueue(2)

		// Act, and assert.
		r, deqOk := q.Dequeue()
		assert.True(t, deqOk)
		assert.Equal(t, rune(1), r)
		r, deqOk = q.Dequeue()
		assert.True(t, deqOk)
		assert.Equal(t, rune(2), r)
	})
}

func TestSize(t *testing.T) {
	t.Run("when no wrap around, return correct size", func(t *testing.T) {
		// Arrange.
		q := queue.NewRingBufferQueue(2)
		q.Enqueue(1)

		// Act.
		s := q.Size()

		// Assert.
		assert.Equal(t, 1, s)
	})

	t.Run("When wrap around, returns correct size", func(t *testing.T) {
		// Arrange
		q := queue.NewRingBufferQueue(2)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Dequeue()
		q.Enqueue(3)

		// Act.
		numOfElements := q.Size()

		// Assert.
		assert.Equal(t, 2, numOfElements)
	})

	t.Run("When wrap around happens multiple times, returns correct size", func(t *testing.T) {
		// Arrange
		q := queue.NewRingBufferQueue(2)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Dequeue()
		q.Dequeue()
		q.Enqueue(3)
		q.Enqueue(4)

		// Act.
		numOfElements := q.Size()

		// Assert.
		assert.Equal(t, 2, numOfElements)

	})
}
