package reiterate

func Flatten[T any, U any](src Iterator[T], f func(T) Iterator[U]) Iterator[U] {
	return &flattener[T, U]{src, f, nil}
}

type flattener[T any, U any] struct {
	src Iterator[T]
	f   func(T) Iterator[U]
	sub Iterator[U]
}

func (f *flattener[T, U]) MoveNext() bool {
	if f.sub != nil && f.sub.MoveNext() {
		return true
	}

	for f.src.MoveNext() {
		f.sub = f.f(f.src.Current())
		if f.sub.MoveNext() {
			return true
		}
	}

	return false
}

func (f *flattener[T, U]) Current() U {
	return f.sub.Current()
}
