package stacks

import (
	"errors"
	"github.com/yonyu/algorithms-go/lists"
)

type LinkedStack struct {
	topPtr *lists.Node
	count int
}

func (s *LinkedStack) Size() int {
	return s.count
}

func (s* LinkedStack) Empty() bool {
	return s.count == 0
}

func (s *LinkedStack) Clear() {
	s.topPtr = nil
	s.count = 0
}

func (s *LinkedStack) Push(e interface{}) {
	newNode := lists.Node{Item: e, Next: s.topPtr}
	s.topPtr = &newNode
	s.count++
}

func (s *LinkedStack) Pop() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("Pop: the stack cannot be empty")
	}

	result := s.topPtr.Next
	s.topPtr = s.topPtr.Next
	s.count--
	return result, nil
}

func (s *LinkedStack) Top() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("Top: the stack cannot be empty")
	}
	return s.topPtr.Item, nil
}