package main

import (
	"dsa/playground/internal/queue"
	"fmt"
)

func main() {
	q := queue.NewRingBufferQueue(2)
	if !q.Enqueue('A') {
		panic("Queue Enqueue A")
	}
	if q.Size() != 1 {
		panic("the number of elements should be 1")
	}
	if !q.Enqueue('B') {
		panic("Queue Enqueue B")
	}
	if q.Size() != 2 {
		panic("the number of elements should be 2")
	}
	if q.Enqueue('C') {
		panic("Was able to enqueue, when I shouldn't have")
	}
	if q.Size() != 2 {
		panic("the number of elements should be 2")
	}
	dequeuedRune, ok := q.Dequeue()
	if !ok {
		panic("Queue Dequeue A")
	}
	fmt.Println("dequeued", string(dequeuedRune))

	if q.Size() != 1 {
		panic("the number of elements should be 1")
	}
	if !q.Enqueue('D') {
		panic("Queue Enqueue D")
	}
	if q.Size() != 2 {
		panic("the number of elements should be 2")
	}
	dequeuedRune, ok = q.Dequeue()
	if !ok {
		panic("Queue Dequeue B")
	}
	fmt.Println("dequeued", string(dequeuedRune))
	if q.Size() != 1 {
		panic("the number of elements should be 1")
	}
	dequeuedRune, ok = q.Dequeue()
	if !ok {
		panic("Queue Dequeue D")
	}
	fmt.Println("dequeued", string(dequeuedRune))
	if q.Size() != 0 {
		panic("the number of elements should be 0")
	}
	dequeuedRune, ok = q.Dequeue()
	if ok {
		panic("Was able to dequeue but queue should be empty...")
	}
	if q.Size() != 0 {
		panic("the number of elements should be 0")
	}
}
