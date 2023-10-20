package reiterate

func Map[A any, AT ~[]A, B any, BT ~[]B](as AT, f func(A) B) BT {
	bs := make(BT, len(as))

	for i, a := range as {
		bs[i] = f(a)
	}

	return bs
}

func Filter[A any, AT ~[]A, F func(A) bool](as AT, f F) AT {
	var na AT

	for _, a := range as {
		if f(a) {
			na = append(na, a)
		}
	}

	return na
}

func IndexOf[A comparable, AS ~[]A](a A, as AS) int {
	for i, b := range as {
		if a == b {
			return i
		}
	}

	return -1
}

func Contains[A comparable, AS ~[]A](a A, as AS) bool {
	return -1 != IndexOf(a, as)
}

func InPlaceReverse[A any, AS ~[]A](a AS) {
	middle := len(a) / 2

	for i := 0; i < middle; i++ {
		n := len(a) - i - 1
		b := a[i]
		c := a[n]
		a[i], a[n] = c, b
	}
}

func MakeReversed[A any, AS ~[]A](a AS) AS {
	rev := make(AS, len(a))

	middle := len(a) / 2

	for i := 0; i < middle; i++ {
		n := len(a) - i - 1
		rev[i], rev[n] = a[n], a[i]
	}

	if len(a)&1 == 1 {
		// this little piggy too
		rev[middle] = a[middle]
	}

	return rev
}
