package stack

import (
	"sync"

	"github.com/google/uuid"
)

type ArrayStack struct {
	sync.RWMutex
	Data []uuid.UUID
}

func (s *ArrayStack) Push(value uuid.UUID) {
	s.Lock()
	defer s.Unlock()

	s.Data = append(s.Data, value)
}

func (s *ArrayStack) Pop() uuid.UUID {
	s.Lock()
	defer s.Unlock()

	if len(s.Data) > 0 {
		value := s.Data[len(s.Data)-1]
		s.Data = s.Data[:len(s.Data)-1]
		return value
	}

	return uuid.Nil
}

func (s *ArrayStack) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.Data)
}

func (s *ArrayStack) Peek() uuid.UUID {
	s.RLock()
	defer s.RUnlock()

	if len(s.Data) > 0 {
		value := s.Data[len(s.Data)-1]
		return value
	}

	return uuid.Nil
}

func (s *ArrayStack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.Data) == 0
}

func (s *ArrayStack) Clear() {
	s.Lock()
	defer s.Unlock()
	s.Data = nil
}

func (s *ArrayStack) Contains(value uuid.UUID) bool {
	s.RLock()
	defer s.RUnlock()

	for _, v := range s.Data {
		if v == value {
			return true
		}
	}
	return false
}

func (s *ArrayStack) Copy() *ArrayStack {
	s.RLock()
	defer s.RUnlock()

	stack := &ArrayStack{}
	stack.Data = make([]uuid.UUID, len(s.Data))
	copy(stack.Data, s.Data)
	return stack
}

func (s *ArrayStack) Reverse() {
	s.Lock()
	defer s.Unlock()

	for i, j := 0, len(s.Data)-1; i < j; i, j = i+1, j-1 {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
}
