package bitset64

import (
	"iter"
	"math/bits"
)

type IterOf[T ~uint64] interface {
	All(func(T) bool)
	Buckets(func(uint64) bool)
	UntilEmpty(func(T) bool)
}

func Iter(b *Bitset) iter64 {
	return iter64{b}
}

func IterT[T ~uint64](b *Bitset) iter64T[T] {
	return iter64T[T]{Iter(b)}
}

type iter64 struct {
	set *Bitset
}

func (i iter64) All(yield func(v uint64) bool) {
	all(i.set)(yield)
}

func (i iter64) Buckets(yield func(v uint64) bool) {
	parts := ToRawParts(*i.set)
	for _, bucket := range parts {
		if !yield(bucket) {
			break
		}
	}
}

func (i iter64) UntilEmpty(yield func(uint64) bool) {
	for !i.set.IsEmpty() {
		if !yield(i.set.Pop()) {
			return
		}
	}
}

type iter64T[T ~uint64] struct {
	iter64
}

func (i iter64T[T]) UntilEmpty(yield func(T) bool) {
	for x := range i.iter64.UntilEmpty {
		if !yield(T(x)) {
			return
		}
	}
}

func (i iter64T[T]) All(yield func(v T) bool) {
	for x := range i.iter64.All {
		if !yield(T(x)) {
			break
		}
	}
}

func all(set *Bitset) iter.Seq[uint64] {
	return func(yield func(v uint64) bool) {
		parts := ToRawParts(*set)
	iter:
		for p, part := range parts {
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
