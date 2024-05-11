package queues

import (
	"errors"
	"fmt"
)

// Queue by using circular array as the underlying data store
type ArrayQueue[T any] struct {
	front 	int // the front index
	count 	int
	items 	[]T
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	q := ArrayQueue[T]{0, 0, make([]T, 4)}
	return &q
}

func (q ArrayQueue[T]) Size() int {
	return q.count
}

func (q ArrayQueue[T]) Empty() bool {
	return q.count == 0
}

func (q *ArrayQueue[T]) Clear() {
	if q.count > 0 {
		q.items = make([]T, 0)
	}
	q.count = 0
	q.front = 0
}

func (q *ArrayQueue[T]) Enqueue(e T) {
	if q.count == len(q.items) {
		q.items = append(q.items, q.items...)
	}
	q.items[(q.front + q.count ) % len(q.items)] = e

	q.count++
}

func (q *ArrayQueue[T]) Dequeue() (T, error) {
	if (q.count == 0) {
		var result T
		return result, errors.New("Dequeue: the queue cannot be empty")
	}

	result := q.items[q.front]
	q.count--
	q.front++
	return result, nil
}

func (q *ArrayQueue[T]) Front() (T, error) {
	if (q.count == 0) {
		var result T
		return result, errors.New("Front: the queue cannot be empty")
	}
	return q.items[q.front], nil
}

func (q ArrayQueue[T]) ToString() string {
	cur := q.items[q.front:len(q.items)]
	cur = append(cur, q.items[0:q.front]...)
	cur = cur[0:q.count]
	return fmt.Sprintf("ArrayQueue:\ncount: %d\nfront: %d\nlen: %d\ncap: %d\nitems: %v\n", q.count, q.front, len(q.items), cap(q.items), cur)
}
