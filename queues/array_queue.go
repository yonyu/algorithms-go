package queues

import (
	"errors"
	"fmt"
)

// Queue by using circular array as the underlying data store
type ArrayQueue struct {
	front 	int // the front index
	count 	int
	items 	[] interface{}
}

func NewArrayQueue() *ArrayQueue {
	q := ArrayQueue{0, 0, make([]interface{}, 4)}
	return &q
}

func (q ArrayQueue) Size() int {
	return q.count
}

func (q ArrayQueue) Empty() bool {
	return q.count == 0
}

func (q *ArrayQueue) Clear() {
	if q.count > 0 {
		q.items = make([]interface{}, 0)
	}
	q.count = 0
	q.front = 0
}

func (q *ArrayQueue) Enqueue(e interface{}) {
	if q.count == len(q.items) {
		q.items = append(q.items, q.items...)
	}
	q.items[(q.front + q.count ) % len(q.items)] = e

	q.count++
}

func (q *ArrayQueue) Dequeue() (interface{}, error) {
	if (q.count == 0) {
		return nil, errors.New("Dequeue: the queue cannot be empty")
	}

	result := q.items[q.front]
	q.count--
	q.front++
	return result, nil
}

func (q *ArrayQueue) Front() (interface{}, error) {
	if (q.count == 0) {
		return nil, errors.New("Front: the queue cannot be empty")
	}
	return q.items[q.front], nil
}

func (q ArrayQueue) ToString() string {
	cur := q.items[q.front:len(q.items)]
	cur = append(cur, q.items[0:q.front]...)
	cur = cur[0:q.count]
	return fmt.Sprintf("ArrayQueue:\ncount: %d\nfront: %d\nlen: %d\ncap: %d\nitems: %v\n", q.count, q.front, len(q.items), cap(q.items), cur)
}
