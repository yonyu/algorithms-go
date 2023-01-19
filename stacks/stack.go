package stacks

import (
	"github.com/yonyu/algorithms-go/containers"
)

type Stack interface {
	containers.Container
	Push(e interface{})
	Pop() (interface{}, error)
	Top() (interface{}, error)
}