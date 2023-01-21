package queues

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	front 	int // the frond index
	count 	int // this simplifies the calculation comparing to keeping the end index
	items 	[] interface{}
}

func NewArrayQueue() *ArrayQueue {
	q := ArrayQueue{0, 0, make([]interface{}, 0, 3)}
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
	capacity := cap(q.items)

	//when the queue is full, doubling the capacity of the underlying slice
	if q.count == capacity {
		// a new empty slice with capacity doubled
		newItems := make([]interface{}, 0, cap(q.items) *2) 

		// copy over the data from the current q.items : [q.front:len(q.items)), [0:q.front)
		newItems = append(newItems, q.items[q.front:len(q.items)]...)
		newItems = append(newItems, q.items[0:q.front]...)

		q.items = newItems //[:] // value copy
		q.front = 0
		capacity = cap(q.items)
	}

	// using a circular array to store items
	if (q.front + q.count >= capacity) {
		q.items[(q.front + q.count) % capacity] = e // the slot is allocated already
	} else {
		q.items = append(q.items, e) 
	}

	// if q.count == len(q.items) {
	// 	q.items = append(q.items, q.items...)
	// }
	// q.items[(q.front + q.count ) % len(q.items)] = e

	q.count++
}

func (q *ArrayQueue) Dequeue() (interface{}, error) {
	if (q.count == 0) {
		return nil, errors.New("Dequeue: the queue cannot be empty")
	}

	result := q.items[q.front]
	q.count--
	q.front = (q.front + 1) % cap(q.items)
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
