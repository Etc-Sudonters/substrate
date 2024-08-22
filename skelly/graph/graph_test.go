package graph

import (
	"iter"
	"testing"

	"github.com/etc-sudonters/substrate/skelly/bitset"
	"github.com/etc-sudonters/substrate/skelly/queue"
	"github.com/etc-sudonters/substrate/skelly/stack"
)

const (
	A Node = 1 << iota
	B Node = 1 << iota
	C Node = 1 << iota
	D Node = 1 << iota
	E Node = 1 << iota
	F Node = 1 << iota
)

func TestBFS(t *testing.T) {
	expectedTrip := queue.From([]Node{A, B, D, E, C, F})
	g := FromOriginationMap(OriginationMap{
		Origination(A): bitset.CreateT(Destination(D), Destination(B)),
		Origination(B): bitset.CreateT(Destination(E)),
		Origination(C): bitset.CreateT(Destination(F)),
		Origination(D): bitset.CreateT(Destination(C)),
	})

	actualTrip := queue.Make[Node](0, expectedTrip.Len())
	iters := SuccessorsIter{g}

	for node := range iters.BFS(A) {
		actualTrip.Push(node)
	}

	if expectedTrip.Len() != actualTrip.Len() {
		t.Logf("visited a different number of nodes than expected: %d %+v", actualTrip.Len(), *actualTrip)
		t.Logf("visited a different set of nodes than expected:\nexpected:\t%+v\nactual:\t\t%+v", *expectedTrip, *actualTrip)
		t.Fatal()
	}

    for expected, actual := range ziptwo(expectedTrip.Iter, actualTrip.Iter) {
        if expected != actual {
			t.Fatalf("visited a different set of nodes than expected:\nexpected:\t%+v\nactual:\t\t%+v", *expectedTrip, *actualTrip)
        }
    }
}

func TestDFS(t *testing.T) {
	expectedTrip := stack.From([]Node{A, D, C, F, B, E})
	g := FromOriginationMap(OriginationMap{
		Origination(A): bitset.CreateT(Destination(D), Destination(B)),
		Origination(B): bitset.CreateT(Destination(E)),
		Origination(C): bitset.CreateT(Destination(F)),
		Origination(D): bitset.CreateT(Destination(C)),
	})

	actualTrip := stack.Make[Node](0, expectedTrip.Len())
	iters := SuccessorsIter{g}

	for node := range iters.DFS(A) {
		actualTrip.Push(node)
	}

	if expectedTrip.Len() != actualTrip.Len() {
		t.Logf("visited a different number of nodes than expected: %d %+v", actualTrip.Len(), *actualTrip)
		t.Logf("visited a different set of nodes than expected:\nexpected:\t%+v\nactual:\t\t%+v", *expectedTrip, *actualTrip)
		t.Fatal()
	}

	for expected, actual := range ziptwo(expectedTrip.Iter, actualTrip.Iter) {
		if expected != actual {
			t.Fatalf("visited a different set of nodes than expected:\nexpected:\t%+v\nactual:\t\t%+v", *expectedTrip, *actualTrip)
		}
	}
}

func ziptwo[A any, B any](a iter.Seq[A], b iter.Seq[B]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		nextA, stopA := iter.Pull(a)
		nextB, stopB := iter.Pull(b)

		defer stopA()
		defer stopB()

		for {
			vA, okA := nextA()
			if !okA {
				return
			}

			vB, okB := nextB()
			if !okB {
				return
			}

			if !yield(vA, vB) {
				return
			}
		}
	}
}
