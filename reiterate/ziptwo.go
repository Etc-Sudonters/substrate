package reiterate

import "github.com/etc-sudonters/substrate/bag"

type ziptwo[A any, B any] struct {
	a   []A
	b   []B
	cur int
	k   int
}

type zippedtwo[A any, B any] struct {
	A A
	B B
}

func (z zippedtwo[A, B]) Reduce(f func(A, B) interface{}) interface{} {
	return f(z.A, z.B)
}

func ZipTwo[A any, B any, AS ~[]A, BS ~[]B](a AS, b BS) *ziptwo[A, B] {
	return &ziptwo[A, B]{
		a:   []A(a),
		b:   []B(b),
		cur: -1,
		k:   bag.Min(len(a), len(b)),
	}
}

func (z *ziptwo[A, B]) Next() bool {
	if z.cur+1 >= z.k {
		return false
	}
	z.cur += 1
	return true
}

func (z *ziptwo[A, B]) Current() *zippedtwo[A, B] {
	if z.cur > z.k {
		return nil
	}
	return &zippedtwo[A, B]{A: z.a[z.cur], B: z.b[z.cur]}
}
