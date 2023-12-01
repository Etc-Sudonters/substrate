package bitset

import (
	"testing"

	"github.com/etc-sudonters/substrate/reiterate"
)

const (
	minU uint64 = 0
	maxU        = ^minU
)

func TestSetsBits(t *testing.T) {
	expected := []uint64{2, 2, 2}

	numbers := []int{
		1,
		65,
		129,
	}

	b := New(3)
	for i := range numbers {
		b.Set(numbers[i])
	}

	for i := range expected {
		if expected[i] != b.bytes[i] {
			t.FailNow()
		}
	}
}

func TestClearsBits(t *testing.T) {
	expected := []uint64{2, 0, 2}

	b := New(3)
	numbers := []int{1, 65, 129}
	for i := range numbers {
		b.Set(numbers[i])
	}
	b.Clear(65)

	for i := range expected {
		if expected[i] != b.bytes[i] {
			t.FailNow()
		}
	}
}

func TestTestBits(t *testing.T) {
	var b Bitset64
	b.k = 3
	b.bytes = []uint64{2, 2, 2}

	if !b.Test(1) {
		t.Log("expected 1 to be set")
		t.Fail()
	}

	if !b.Test(65) {
		t.Log("expected 65 to be set")
		t.Fail()
	}

	if !b.Test(129) {
		t.Log("expected 129 to be set")
		t.Fail()
	}
}

func TestComplement(t *testing.T) {
	b := New(3)
	b.Set(1)
	b.Set(65)
	b.Set(129)

	comp := b.Complement().bytes
	expected := maxU ^ 2

	if comp[0] != expected || comp[1] != expected || comp[2] != expected {
		t.Fail()
	}
}

func TestIntersect(t *testing.T) {
	b1 := New(3)
	b2 := New(3)

	shared := []int{1, 65, 129}
	b1.Set(144)
	b2.Set(13)

	for i := range shared {
		b1.Set(shared[i])
		b2.Set(shared[i])
	}

	I := b1.Intersect(b2).bytes

	if I[0] != 2 || I[1] != 2 || I[2] != 2 {
		t.Fail()
	}
}

func TestUnion(t *testing.T) {
	b1 := New(1)
	b2 := New(2)
	b3 := New(3)

	b1.Set(1)
	b2.Set(65)
	b3.Set(129)

	b := b1.Union(b2).Union(b3)

	if !b.Eq(FromRaw(2, 2, 2)) {
		t.Fail()
	}
}

func TestDifference(t *testing.T) {
	b1 := New(3)
	b2 := New(3)

	b1.Set(1)
	b1.Set(65)
	b2.Set(65)
	b2.Set(129)

	if !b1.Difference(b2).Eq(FromRaw(2, 0, 0)) {
		t.Log("expected only 1 to be set")
		t.Fail()
	}

	if !b2.Difference(b1).Eq(FromRaw(0, 0, 2)) {
		t.Log("expected only 129 to be set")
		t.Fail()
	}
}

func TestLength(t *testing.T) {
	b := New(3)
	b.Set(1)
	b.Set(65)
	b.Set(129)

	if l := b.Len(); l != 3 {
		t.Fatalf("expected length of 3 but got %d", l)
	}

	b = New(Buckets(10000))

	for i := 0; i < 10000; i += 2 {
		b.Set(i)
	}

	if l := b.Len(); l != 5000 {
		t.Fatalf("expected length of 5000 but got %d", l)
	}
}

func TestElems(t *testing.T) {
	b := New(3)
	b.Set(1)
	b.Set(65)
	b.Set(129)

	expected := []int{1, 65, 129}
	elems := b.Elems()

	if len(expected) != len(elems) {
		t.Fatalf("mismatched elems\nexpected:\t%+v\nactual:\t%+v", expected, elems)
	}

	pairs := reiterate.ZipTwo(expected, elems)

	for i := 0; pairs.Next(); i++ {
		p := pairs.Current()
		if p.A != p.B {
			t.Logf("expected to find %d at index %d but found %d", p.A, i, p.B)
			t.Fail()
		}
	}

	b = New(Buckets(10000))
	expected = make([]int, 0, 5000)

	for i := 0; i < 10000; i += 2 {
		b.Set(i)
		expected = append(expected, i)
	}

	elems = b.Elems()
	l := b.Len()

	if len(elems) != l {
		t.Fatalf("len(Elems()) and Len() disagree\nlen(Elems()) = %d\nLen() = %d", len(elems), l)
	}

	if l != len(expected) {
		t.Fatalf("expected length of %d but got %d", len(expected), l)
	}

	pairs = reiterate.ZipTwo(expected, elems)

	for i := 0; pairs.Next(); i++ {
		p := pairs.Current()
		if p.A != p.B {
			t.Logf("expected to find %d at index %d but found %d", p.A, i, p.B)
			t.Fail()
		}
	}
}
