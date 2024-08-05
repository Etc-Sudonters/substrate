package graph

import (
	"context"
	"errors"

	"github.com/etc-sudonters/substrate/skelly/bitset"
)

type (
	// used to provide human readable diagnostic output
	DebugFunc func(string, ...any)

	// calls F on current node, selected and err results from S
	DebugSelector[T Direction] struct {
		F DebugFunc
		S Selector[T]
	}

	// calls F on current node, error from V
	DebugVisitor struct {
		F DebugFunc
		V Visitor
	}

	// provides current node to every Visitor in slice
	VisitorArray []Visitor
)

func (v VisitorArray) Visit(ctx context.Context, node Node) error {
	var errs []error

	for i := range v {
		if err := v[i].Visit(ctx, node); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (d DebugVisitor) Visit(ctx context.Context, node Node) error {
	d.F("visiting node %s", node)
	err := d.V.Visit(ctx, node)
	if err != nil {
		d.F("error on visit to %s: %s", node, err)
	}
	return err
}

func (d DebugSelector[T]) Select(g Directed, n Node) (bitset.Bitset64, error) {
	d.F("selecting from %s", n)
	selected, err := d.S.Select(g, n)
	if err != nil {
		d.F("error while selecting: %s", err)
	}

	d.F("selected %+v", selected)

	return selected, err
}
