package queue

import "errors"

type Q[T any] []T

func Make[E any](len, cap int) *Q[E] {
	q := make(Q[E], len, cap)
	return &q
}

func From[E any, T ~[]E](src T) *Q[E] {
	q := make(Q[E], len(src))
	copy(q, src)
	return &q
}

func (q *Q[T]) Push(t T) {
	*q = append(*q, t)
}

func (q *Q[T]) Pop() (T, error) {
	var t T
	if len(*q) == 0 {
		return t, errors.New("empty queue")
	}

	items := *q
	t, *q = items[0], items[1:]

	return t, nil
}

func (q *Q[T]) Len() int {
	return len(*q)
}

func(q *Q[T]) Iter(yield func(T) bool) {
    for _, t := range *q {
        if !yield(t) {
            break
        }
    }
}
