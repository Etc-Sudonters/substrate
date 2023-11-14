package reiterate

type Iterator[E any] interface {
	MoveNext() bool
	Current() E
}
