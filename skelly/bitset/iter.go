package bitset

import (
	"iter"
	"math/bits"
)

func Iter64(b Bitset64) iter64 {
    return iter64(b)
}

func Iter64T[T ~uint64](b Bitset64) iter64T[T] {
    return iter64T[T]{ Iter64(b) }
}

type iter64 Bitset64

func (i iter64) All(yield func(v uint64) bool) {
	all(Bitset64(i))(yield)
}

func(i iter64) Buckets(yield func(v uint64) bool) {
    for _, bucket := range ToRawParts(Bitset64(i)) {
        if !yield(bucket) {
            break
        }
    }
}

type iter64T[T ~uint64] struct {
    iter64
}

func (i iter64T[T]) All(yield func(v T) bool) {
    for x := range i.iter64.All {
        if !yield(T(x)) {
            break
        }
    }
}

func all(set Bitset64) iter.Seq[uint64] {
	return func(yield func(v uint64) bool) {
	iter:
		for p, part := range ToRawParts(Bitset64(set)) {
			for part != 0 {
				tz := bits.TrailingZeros64(part)
				if !yield(uint64(tz + (p * 64))) {
					break iter
				}
				part ^= (1 << tz)
			}
		}
	}
}
