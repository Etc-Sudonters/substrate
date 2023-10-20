package bag

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

func Max[A constraints.Ordered](a, b A) A {
	if a > b {
		return a
	}
	return b
}

func Min[A constraints.Ordered](a, b A) A {
	if a < b {
		return a
	}
	return b
}

func Shuffle[T any, E ~[]T](elms E) {
	rand.Shuffle(len(elms), func(i, j int) {
		elms[i], elms[j] = elms[j], elms[i]
	})
}
