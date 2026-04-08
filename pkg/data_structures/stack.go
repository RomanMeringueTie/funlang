package data_structures

import "log"

type Stack[T any] struct {
	s []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (s *Stack[T]) Push(v T) {
	s.s = append(s.s, v)
}

func (s *Stack[T]) Pop() T {
	l := len(s.s)
	if l == 0 {
		log.Fatal("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func (s *Stack[T]) Size() int {
	return len(s.s)
}
