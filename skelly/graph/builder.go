package graph

import (
	"errors"

	"github.com/etc-sudonters/substrate/skelly/bitset"
)

var ErrNodeNotFound = errors.New("node not found")
var ErrOriginNotFound = errors.New("origin not found")
var ErrDestNotFound = errors.New("destination not found")

// adds nodes and edges to a Directed
type Builder struct {
	G Directed
}

// adds a node to the graph if it doesn't exist already
func (b *Builder) AddNode(n Node) {
	if _, exists := b.G.dests[Destination(n)]; !exists {
		b.G.dests[Destination(n)] = bitset.Bitset64{}
		b.G.origins[Origination(n)] = bitset.Bitset64{}
	}
}

func (b *Builder) AddNodes(ns []Node) {
	for _, n := range ns {
		b.AddNode(n)
	}
}

// connects o -> d, if either node doesn't exist they are created
//
//	duplicate edges are not added, order isn't guaranteed
func (b *Builder) AddEdge(o Origination, d Destination) error {
	b.AddNode(Node(o))
	b.AddNode(Node(d))

	origins := b.G.origins[o]
	bitset.Set(&origins, d)
	b.G.origins[o] = origins

	dests := b.G.dests[d]
	bitset.Set(&dests, o)
	b.G.dests[d] = dests
	return nil
}

func (b *Builder) AddEdges(e OriginationMap) error {
	for o, neighbors := range e {
		biter := bitset.Iter64T[Destination](neighbors)
        for dest := range biter.All {
			if err := b.AddEdge(o, dest); err != nil {
				return err
			}
		}
	}
	return nil
}
