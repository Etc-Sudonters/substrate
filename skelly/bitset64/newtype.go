package bitset64

type u64 interface {
	~uint64
}

func CreateT[T u64](members ...T) Bitset {
	b := New(0)
	for _, m := range members {
		b.Set(uint64(m))
	}
	return b
}

func Set[T u64](b *Bitset, t T) bool {
	return b.Set(uint64(t))
}

func Unset[T u64](b *Bitset, t T) {
	b.Unset(uint64(t))
}

func IsSet[T u64](b *Bitset, t T) bool {
	return b.IsSet(uint64(t))
}

func Intersects(this, other Bitset) bool {
	return !this.Intersect(other).IsEmpty()
}
