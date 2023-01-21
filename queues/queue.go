package queues

import "github.com/yonyu/algorithms-go/containers"

type Queue interface {
	containers.Container
	Enqueue(e interface{})
	Dequeue() (interface{}, error)
	Front() (interface{}, error)
}
