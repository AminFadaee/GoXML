package utility

import (
	"errors"
	"sync"
)

type stack[T any] struct {
	lock sync.Mutex
	s    []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{sync.Mutex{}, make([]T, 0)}
}

func (s *stack[T]) Push(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		var nullResult T
		return nullResult, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
