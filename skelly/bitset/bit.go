package bitset

import (
	"fmt"
	"math/bits"

	"github.com/etc-sudonters/substrate/reiterate"
)

func Set[E ~int](b Bitset64, e E) {
	b.Set(int(e))
}

func Clear[E ~int](b Bitset64, e E) {
	b.Clear(int(e))
}

func Test[E ~int](b Bitset64, e E) bool {
	return b.Test(int(e))
}

type Bitset64 struct {
	k     int
	bytes []uint64
}

func (b Bitset64) String() string {
	return fmt.Sprintf("Bitset64 { k: %d }", b.k)
}

func IsEmpty(b Bitset64) bool {
	for i := range b.bytes {
		if b.bytes[i] != 0 {
			return false
		}
	}
	return true
}

func Buckets(max int) int {
	return max/bs64Size + 1
}

func New(k int) Bitset64 {
	bytes := make([]uint64, k)
	return Bitset64{k, bytes}
}

func FromRaw(parts ...uint64) Bitset64 {
	var b Bitset64
	b.k = len(parts)
	b.bytes = parts
	return b
}

func ToRawParts(b Bitset64) []uint64 {
	ret := make([]uint64, len(b.bytes))
	copy(ret, b.bytes)
	return ret
}

func Copy(b Bitset64) Bitset64 {
	n := New(b.k)
	copy(n.bytes, b.bytes)
	return n
}

func (b Bitset64) Set(i int) {
	idx := bs64idx(i)
	bit := bs64bit(i)
	b.bytes[idx] |= bit
}

func (b Bitset64) Clear(i int) {
	b.bytes[bs64idx(i)] &= ^bs64bit(i)
}

func (b Bitset64) Test(i int) bool {
	bit := bs64bit(i)
	return bit == (bit & b.bytes[bs64idx(i)])
}

func (b Bitset64) Complement() Bitset64 {
	n := Copy(b)
	for i, bits := range n.bytes {
		n.bytes[i] = ^bits
	}
	return n
}

func (b Bitset64) Intersect(n Bitset64) Bitset64 {
	if n.k > b.k {
		b, n = n, b
	}

	ret := Copy(b)
	for i, d := range n.bytes {
		ret.bytes[i] &= d
	}

	return ret
}

func (b Bitset64) Union(n Bitset64) Bitset64 {
	if n.k > b.k {
		b, n = n, b
	}

	ret := Copy(b)
	for i, bits := range n.bytes {
		ret.bytes[i] |= bits
	}

	return ret
}

func (b Bitset64) Difference(n Bitset64) Bitset64 {
	return b.Intersect(n.Complement())
}

func (b Bitset64) Eq(n Bitset64) bool {
	if b.k != n.k {
		return false
	}

	pairs := reiterate.ZipTwo(b.bytes, n.bytes)

	for pairs.Next() {
		p := pairs.Current()
		if p == nil {
			break
		}
		if p.A != p.B {
			return false
		}
	}

	return true
}

func (b Bitset64) Len() int {
	var count int

	for _, bucket := range b.bytes {
		for ; bucket != 0; bucket &= bucket - 1 {
			count++
		}
	}

	return count
}

func (b Bitset64) Elems() []int {
	var elems []int

	for k, bucket := range b.bytes {
		for bucket != 0 {
			tz := bits.TrailingZeros64(bucket)
			elems = append(elems, k*64+tz)
			bucket ^= (1 << tz)
		}
	}

	return elems
}

const bs64Size = 64

func bs64idx(i int) int {
	return i / bs64Size
}

func bs64bit(i int) uint64 {
	return 1 << (i % bs64Size)
}
