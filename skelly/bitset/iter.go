package bitset

import (
	"math/bits"

	"github.com/etc-sudonters/substrate/reiterate"
)

func Iter(b Bitset64) reiterate.Iterator[uint64] {
	return &iter{ToRawParts(b), 0, -1}
}

type iter struct {
	parts   []uint64
	current uint64
	partIdx int
}

func (b *iter) MoveNext() bool {
	if b.partIdx >= len(b.parts) {
		return false
	}

	for b.current == 0 {
		b.partIdx++
		if b.partIdx >= len(b.parts) {
			return false
		}

		candidate := b.parts[b.partIdx]
		if candidate != 0 {
			b.current = candidate
			break
		}
	}

	b.current ^= (1 << bits.TrailingZeros64(b.current))
	return b.current != 0
}

func (b *iter) Current() uint64 {
	return uint64(b.partIdx)*64 + uint64(bits.TrailingZeros64(b.current))
}
