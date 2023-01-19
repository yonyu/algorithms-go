package stacks

import (
	"errors"
)

// We use a dynamic array called store to hold stack elements.
type ArrayStack struct {
	store []interface{} // top is always store[len(store)-1]
}

func (s *ArrayStack) Size() int {
	return len(s.store)
}

func (s *ArrayStack) Empty() bool {
	return len(s.store) == 0
}

func (s *ArrayStack) Clear() {
	s.store = make([]interface{}, 0, 10)
}

func (s *ArrayStack) Push(e interface{}) {
	s.store = append(s.store, e)
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, errors.New("Pop: the stack cannot be empty")
	}
	result := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1] // do so as resizing slices is inexpensive in Go

	return result, nil
}

func (s *ArrayStack) Top() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, errors.New("Top: the stack cannot be empty")
	}
	return s.store[len(s.store) - 1], nil
}