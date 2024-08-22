package bitset

type newtype interface {
	~uint64
}

func CreateT[T newtype](members ...T) Bitset64 {
	b := New(0)
	for _, m := range members {
		b.Set(uint64(m))
	}
	return b
}

func Set[T newtype](b *Bitset64, t T) bool {
	return b.Set(uint64(t))
}

func Unset[T newtype](b *Bitset64, t T) {
	b.Unset(uint64(t))
}

func IsSet[T newtype](b *Bitset64, t T) bool {
	return b.IsSet(uint64(t))
}
