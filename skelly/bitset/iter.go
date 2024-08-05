package bitset

import (
	"math/bits"

	"github.com/etc-sudonters/substrate/reiterate"
)

func Iter(b Bitset64) reiterate.Iterator[int] {
	return &iter{ToRawParts(b), 0, -1}
}

type iter struct {
	parts   []uint64
	current uint64
	partIdx int
}

func (b *iter) MoveNext() bool {
	if b.current == 0 {
		b.partIdx++
		if b.partIdx >= len(b.parts) {
			return false
		}

		b.current = b.parts[b.partIdx]
		return true
	}

	b.current ^= (1 << bits.TrailingZeros64(b.current))
	return true
}

func (b *iter) Current() int {
	return b.partIdx*64 + bits.TrailingZeros64(b.current)
}
