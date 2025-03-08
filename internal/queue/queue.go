package queue

type RingBufferQueue struct {
	buff     []rune
	capacity int
	head     int
	tail     int
}

func NewRingBufferQueue(bufferSize int) *RingBufferQueue {
	return &RingBufferQueue{
		buff:     make([]rune, bufferSize+1),
		capacity: bufferSize + 1,
	}
}

func (q *RingBufferQueue) Capacity() int {
	return q.capacity - 1
}

func (q *RingBufferQueue) Enqueue(r rune) bool {
	if q.isFull() {
		return false
	}

	q.buff[q.tail] = r
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *RingBufferQueue) Dequeue() (rune, bool) {
	if q.isEmpty() {
		return 0, false
	}

	r := q.buff[q.head]
	q.head = (q.head + 1) % q.capacity
	return r, true
}

func (q *RingBufferQueue) isFull() bool {
	return (q.tail+1)%q.capacity == q.head
}

func (q *RingBufferQueue) isEmpty() bool {
	return q.head == q.tail
}

func (q *RingBufferQueue) NumberOfElementsInQueue() int {
	return (q.tail - q.head + q.capacity) % q.capacity
}
