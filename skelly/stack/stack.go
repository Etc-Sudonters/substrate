package stack

import "errors"

var ErrEmptyStack = errors.New("empty stack")

type S[T any] []T

func From[E any, T ~[]E](src T) *S[E] {
	l := len(src)
	dest := make(S[E], l)

	for i, e := range src {
		dest[l-1-i] = e
	}
	return &dest
}

func Make[E any](len, cap int) *S[E] {
	s := make(S[E], len, cap)
	return &s
}

func (s *S[T]) Push(t T) {
	*s = append([]T{t}, *s...)
}

func (s *S[T]) Pop() (T, error) {
	var t T
	if len(*s) == 0 {
		return t, ErrEmptyStack
	}

	items := *s
	t, *s = items[0], items[1:]

	return t, nil
}

func (s *S[T]) Len() int {
	return len(*s)
}

func (s *S[T]) Iter(yield func(T) bool) {
	arr := *s
	for _, t := range arr {
		if !yield(t) {
			break
		}
	}
}

func (s *S[T]) Top() (*T, error) {
	if len(*s) == 0 {
		return nil, ErrEmptyStack
	}

	return &(*s)[0], nil
}
