package graph

import (
	"context"
	"testing"

	"github.com/etc-sudonters/substrate/circumstances"
	"github.com/etc-sudonters/substrate/reiterate"
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
	ctx, cc := circumstances.ContextFromTest(t)
	defer cc(circumstances.ErrTestEnded)

	expectedTrip := queue.From([]Node{A, B, D, C})
	g := FromOriginationMap(OriginationMap{
		Origination(A): {Destination(B), Destination(D)},
		Origination(B): {Destination(D), Destination(C)},
		Origination(C): {Destination(A), Destination(D)},
	})

	q := queue.Make[Node](0, expectedTrip.Len())

	err := BreadthFirst[Destination]{
		Selector: Successors,
		Visitor: VisitorFunc(func(ctx context.Context, n Node) error {
			q.Push(n)
			return nil
		}),
	}.Walk(ctx, g, A)

	if err != nil {
		t.Fatalf("error while traversing graph: %s", err)
	}

	if expectedTrip.Len() != q.Len() {
		t.Logf("visited a different number of nodes than expected: %d %+v", q.Len(), *q)
		t.Fail()
	}

	zip := reiterate.ZipTwo(*q, *expectedTrip)
	for zip.Next() {
		p := zip.Current()
		r := p.Reduce(func(a, b Node) interface{} {
			return a == b
		})
		if !r.(bool) {
			t.Fatalf("visited a different set of nodes than expected:\bexpected:\t%+v\nactual:\t%+v", *expectedTrip, *q)
		}
	}
}

func TestDFS(t *testing.T) {
	ctx, cc := circumstances.ContextFromTest(t)
	defer cc(circumstances.ErrTestEnded)

	expectedTrip := stack.From([]Node{A, B, E, D, C, F})
	g := FromOriginationMap(OriginationMap{
		Origination(A): {Destination(D), Destination(B)},
		Origination(B): {Destination(E)},
		Origination(C): {Destination(F)},
		Origination(D): {Destination(C)},
	})

	visited := stack.Make[Node](0, expectedTrip.Len())

	err := DepthFirst[Destination]{
		Selector: Successors,
		Visitor: VisitorFunc(func(_ context.Context, n Node) error {
			visited.Push(n)
			return nil
		}),
	}.Walk(ctx, g, A)

	if err != nil {
		t.Fatalf("error while traversing graph: %s", err)
	}

	zip := reiterate.ZipTwo(*visited, *expectedTrip)
	for zip.Next() {
		p := zip.Current()
		r := p.Reduce(func(a, b Node) interface{} {
			return a == b
		})
		if !r.(bool) {
			t.Fatalf("visited a different set of nodes than expected:\bexpected:\t%+v\nactual:\t%+v", *expectedTrip, *visited)
		}
	}
}
