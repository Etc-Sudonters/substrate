package rng

type Xoshiro256PP [4]uint64

func NewXoshiro256PPFromU64(seed uint64) *Xoshiro256PP {
	x := Xoshiro256PP([4]uint64{0, 0, 0, 0})
	for i := range x {
		seed = SplitMix64(seed)
		x[i] = seed
	}
	return &x
}

func (x *Xoshiro256PP) NextUint64() uint64 {
	result := rotl(x[0]+x[3], 23) + x[0]
	t := x[1] << 27
	x[2] ^= x[0]
	x[3] ^= x[1]
	x[1] ^= x[2]
	x[0] ^= x[3]
	x[2] ^= t
	x[3] = rotl(x[3], 45)
	return result
}

func (x *Xoshiro256PP) NextFloat64() float64 {
	n := x.NextUint64()
	return float64(n) * 0x1.0p-64
}

func (x *Xoshiro256PP) Uint64() uint64 {
	return x.NextUint64()
}


func SplitMix64(state uint64) uint64 {
	state += 0x9e3779b97f4a7c15
	z := (state ^ (state >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

func rotl(x uint64, k int) uint64 {
	return (x << k) | (x >> (64 - k))
}
