package reiterate

func MapIter[T any, U any](src Iterator[T], f func(T) U) Iterator[U] {
	return &mapiter[T, U]{src, f}
}

type mapiter[T any, U any] struct {
	src Iterator[T]
	f   func(T) U
}

func (m *mapiter[T, U]) MoveNext() bool {
	return m.src.MoveNext()
}

func (m mapiter[T, U]) Current() U {
	return m.f(m.src.Current())
}
