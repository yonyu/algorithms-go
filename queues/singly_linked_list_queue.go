package queues

import (
	"errors"
	"fmt"
	"github.com/yonyu/algorithms-go/lists"
)

type LinkedQueue struct {
	count int
	front *lists.Node
	end   *lists.Node
}

func NewLinkedQueue(count int, front *lists.Node, end *lists.Node) *LinkedQueue {
	q := LinkedQueue{count, front, end}
	return &q
}

func (q LinkedQueue) Size() int {
	return q.count
}

func (q LinkedQueue) Empty() bool {
	return q.count == 0
}

func (q *LinkedQueue) Clear() {
	q.count = 0
	q.front = nil
	q.end = nil
}

func (q *LinkedQueue) Enqueue(e interface{}) {
	node := lists.NewNode(e, nil)

	// if q.end != nil {
	// 	q.end.Next = node
	// 	q.end = q.end.Next
	// } else {
	// 	q.end = node
	// }

	if q.front == nil { // no item in the queue yet
		q.front = node
		q.front.Next = nil
		q.end = nil
	} else if q.front.Next == nil{ // only one item in the queue
		q.end = node
		q.front.Next = q.end

	} else {
		q.end.Next = node
		q.end = q.end.Next
	}

	q.count++
}

func (q *LinkedQueue) Dequeue() (interface{}, error) {
	if (q.count == 0) {
		return nil, errors.New("Dequeue: the queue cannot be empty")
	}

	result := q.front
	q.count--
	q.front = q.front.Next
	return result.Item, nil
}

func (q *LinkedQueue) Front() (interface{}, error) {
	if (q.count == 0) {
		return nil, errors.New("Front: the queue cannot be empty")
	}
	return q.front.Item, nil
}

func (q LinkedQueue) ToString() string {
	items := make([]interface{}, 0)
	item := q.front
	for item != nil {
		items = append(items, item.Item)
		item = item.Next
	}
	return fmt.Sprintf("LinkedQueue:\ncount: %d\nitems: %#v\n", q.count, items)
}