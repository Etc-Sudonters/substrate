package reiterate

func FilterIter[E any](i Iterator[E], f func(E) bool) Iterator[E] {
	return &filter[E]{i, f}
}

type filter[E any] struct {
	i Iterator[E]
	f func(E) bool
}

func (f *filter[E]) MoveNext() bool {
	for {
		if !f.i.MoveNext() {
			return false
		}

		if f.f(f.i.Current()) {
			return true
		}
	}
}

func (f filter[E]) Current() E {
	return f.i.Current()
}
