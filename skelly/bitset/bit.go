package bitset

import (
	"math/bits"
)

func Buckets(i uint64) int {
	return int(i / 64)
}

func BitIndex(i uint64) uint64 {
	return 1 << (i % 64)
}

func New(i int) Bitset64 {
	var b Bitset64
	b.buckets = make([]uint64, i)
	return b
}

func Create(members ...uint64) Bitset64 {
	b := New(0)
	for _, m := range members {
		b.Set(m)
	}
	return b
}

func WithBucketsFor(i uint64) Bitset64 {
	return New(Buckets(i))
}

func FromRaw(parts ...uint64) Bitset64 {
	var b Bitset64
	b.buckets = parts
	return b
}

func ToRawParts(b Bitset64) []uint64 {
	ret := make([]uint64, len(b.buckets))
	copy(ret, b.buckets)
	return ret
}

type Bitset64 struct {
	buckets []uint64
}

func IsEmpty(b Bitset64) bool {
	for i := range b.buckets {
		if b.buckets[i] != 0 {
			return false
		}
	}
	return true
}

func Copy(b Bitset64) Bitset64 {
	var n Bitset64
	n.buckets = make([]uint64, len(b.buckets))
	copy(n.buckets, b.buckets)
	return n
}

func (b *Bitset64) resize(bucket int) {
	if bucket < len(b.buckets) {
		return
	}

	buckets := make([]uint64, bucket+1)
	copy(buckets, b.buckets)
	b.buckets = buckets
}

func (b *Bitset64) Set(i uint64) bool {
	idx := Buckets(i)
	bit := BitIndex(i)
	b.resize(idx)
    bucket := b.buckets[idx]
	b.buckets[idx] = bucket|bit
    return bucket&bit == 0
}

func (b Bitset64) Unset(i uint64) {
	bucket := Buckets(i)

	if bucket >= len(b.buckets) {
		return
	}

	b.buckets[bucket] &= ^BitIndex(i)
}

func (b Bitset64) IsSet(i uint64) bool {
	bucket := Buckets(i)
	if bucket >= len(b.buckets) {
		return false
	}

	bit := BitIndex(i)
	return bit == (bit & b.buckets[bucket])
}

func (b Bitset64) Complement() Bitset64 {
	n := Copy(b)
	for i, bits := range n.buckets {
		n.buckets[i] = ^bits
	}
	return n
}

func (b Bitset64) Intersect(n Bitset64) Bitset64 {
	buckets := min(len(b.buckets), len(n.buckets))
	r := Bitset64{}
	r.buckets = make([]uint64, buckets)

	for i := range r.buckets {
		r.buckets[i] = b.buckets[i] & n.buckets[i]
	}

	return r
}

func (b Bitset64) Union(n Bitset64) Bitset64 {
	if len(n.buckets) > len(b.buckets) {
		b, n = n, b
	}

	ret := Copy(b)
	for i, bits := range n.buckets {
		ret.buckets[i] |= bits
	}

	return ret
}

func (b Bitset64) Difference(n Bitset64) Bitset64 {
	buckets := make([]uint64, max(len(b.buckets), len(n.buckets)))
	copy(buckets, n.buckets)
	n.buckets = buckets
	return b.Intersect(n.Complement())
}

func (b Bitset64) Eq(n Bitset64) bool {
	hi, lo := b.buckets, n.buckets

	if len(lo) > len(hi) {
		hi, lo = lo, hi
	}

	if len(lo) < len(hi) {
		for _, bucket := range hi[len(lo):] {
			if bucket != 0 {
				return false
			}
		}
	}

	for i := range lo {
		if lo[i] != hi[i] {
			return false
		}
	}

	return true
}

func (b Bitset64) Len() int {
	var count int

	for _, bucket := range b.buckets {
		for ; bucket != 0; bucket &= bucket - 1 {
			count++
		}
	}

	return count
}

func (b Bitset64) Elems() []uint64 {
	var elems []uint64

	for k, bucket := range b.buckets {
		k := uint64(k)
		for bucket != 0 {
			tz := uint64(bits.TrailingZeros64(bucket))
			elems = append(elems, k*64+tz)
			bucket ^= (1 << tz)
		}
	}

	return elems
}
