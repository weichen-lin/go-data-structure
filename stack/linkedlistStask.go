package stack

import (
	"sync"

	"github.com/google/uuid"
)

type Node struct {
	Value uuid.UUID
	Next  *Node
}

type LinkedListStack struct {
	sync.RWMutex
	Top  *Node
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
	s.Size++
}

func (s *LinkedListStack) Pop() uuid.UUID {
	s.Lock()
	defer s.Unlock()
	if s.Size > 0 {
		value := s.Top.Value
		s.Top = s.Top.Next
		s.Size--
		return value
	}
	return uuid.Nil
}

func (s *LinkedListStack) Peek() uuid.UUID {
	s.RLock()
	defer s.RUnlock()
	if s.Size > 0 {
		return s.Top.Value
	}
	return uuid.Nil
}

func (s *LinkedListStack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.Size == 0
}

func (s *LinkedListStack) Clear() {
	// 這邊不能加鎖，會遇到重入鎖的問題
	// s.Lock()
	// defer s.Unlock()
	for s.Size > 0 {
		s.Pop()
	}
}

func (s *LinkedListStack) Contains(value uuid.UUID) bool {
	s.RLock()
	defer s.RUnlock()

	for node := s.Top; node != nil; node = node.Next {
		if node.Value == value {
			return true
		}
	}
	return false
}

func (s *LinkedListStack) ToSlice() []uuid.UUID {
	s.RLock()
	defer s.RUnlock()
	slice := make([]uuid.UUID, s.Size)
	for node, i := s.Top, 0; node != nil; node, i = node.Next, i+1 {
		slice[i] = node.Value
	}
	return slice
}

func (s *LinkedListStack) Copy() *LinkedListStack {
	s.RLock()
	defer s.RUnlock()
	stack := &LinkedListStack{}
	for node := s.Top; node != nil; node = node.Next {
		stack.Push(node.Value)
	}
	stack.Reverse()
	return stack
}

func (s *LinkedListStack) Reverse() {
	s.Lock()
	defer s.Unlock()

	var prev *Node

	for current := s.Top; current != nil; {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	s.Top = prev
}
