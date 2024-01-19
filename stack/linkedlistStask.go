package stack

import (
	"sync"

	"github.com/google/uuid"
)

type Node struct {
	Value uuid.UUID
	Next *Node
}

type LinkedListStack struct {
	sync.Mutex
	Top *Node
	Size int
}

func (s *LinkedListStack) Push(value uuid.UUID) {
	s.Lock()
	defer s.Unlock()
	
	if s.Top == nil {
		node := &Node{
			Value: value,
			Next:  nil,
		}
		s.Top = node
		s.Size++
		return
	}
	node := &Node{
		Value: value,
		Next:  s.Top,
	}
	s.Top = node
}

func (s *LinkedListStack) Pop() uuid.UUID {
	if s.Size > 0 {
		value := s.Top.Value
		s.Top = s.Top.Next
		s.Size--
		return value
	}
	return uuid.Nil
}

func (s *LinkedListStack) Peek() uuid.UUID {
	if s.Size > 0 {
		return s.Top.Value
	}
	return uuid.Nil
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.Size == 0
}

func (s *LinkedListStack) Clear() {

	for s.Size > 0 {
		s.Pop()
	}
}

func (s *LinkedListStack) Contains(value uuid.UUID) bool {
	for node := s.Top; node != nil; node = node.Next {
		if node.Value == value {
			return true
		}
	}
	return false
}

func (s *LinkedListStack) ToSlice() []uuid.UUID {
	slice := make([]uuid.UUID, s.Size)
	for node, i := s.Top, 0; node != nil; node, i = node.Next, i+1 {
		slice[i] = node.Value
	}
	return slice
}

func (s *LinkedListStack) Copy() *LinkedListStack {
	stack := &LinkedListStack{}
	for node := s.Top; node != nil; node = node.Next {
		stack.Push(node.Value)
	}
	return stack
}

func (s *LinkedListStack) Reverse() {
	stack := &LinkedListStack{}
	for node := s.Top; node != nil; node = node.Next {
		stack.Push(node.Value)
	}
	s.Top = stack.Top
}

func (s *LinkedListStack) SizeOf() int {
	return s.Size
}

