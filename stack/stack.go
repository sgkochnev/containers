package stack

type node[T any] struct {
	value T
	prev  *node[T]
}

type Stack[T any] struct {
	root *node[T]
	len  int
}

func New[T any](items ...T) *Stack[T] {
	s := &Stack[T]{}

	for _, item := range items {
		s.Push(item)
	}

	return s
}

func (s *Stack[T]) Push(item T) {
	n := &node[T]{
		value: item,
		prev:  s.root,
	}

	s.root = n
	s.len++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.root == nil {
		return node[T]{}.value, false
	}

	result := s.root.value
	s.root = s.root.prev
	s.len--

	return result, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.root == nil {
		return node[T]{}.value, false
	}

	return s.root.value, true
}

func (q *Stack[T]) Clear() {
	q.root = nil
	q.len = 0
}

func (s *Stack[T]) Len() int {
	return s.len
}

func (s *Stack[T]) IsEmpty() bool {
	return s.len == 0 && s.root == nil
}
