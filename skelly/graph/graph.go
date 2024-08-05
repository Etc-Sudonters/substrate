package graph

import (
	"fmt"

	"github.com/etc-sudonters/substrate/skelly/bitset"
)

// an entity in the graph
type Node uint64

// an entity in the graph that is the origination for an edge
type Origination Node

// an entity in the graph that is the destination for an edge
type Destination Node

// a connection between two entities in the graph
type Edge struct {
	O Origination
	D Destination
}

// describes a graph as the set of originating edges
type OriginationMap map[Origination]bitset.Bitset64

// describes a graph as the set of terminating edges
type destinationMap map[Destination]bitset.Bitset64

// allows specifying direction when interacting with a Directed graph
type Direction interface {
	Origination | Destination
}

func (n Node) String() string {
	return fmt.Sprintf("Node{%d}", n)
}

func (d Destination) String() string {
	return fmt.Sprintf("Destination{%d}", d)
}

func (o Origination) String() string {
	return fmt.Sprintf("Origination{%d}", o)
}

func WithCapacity(c int) Directed {
	return Directed{
		origins: make(OriginationMap, c),
		dests:   make(destinationMap, c),
	}
}

// returns a Directed with a small predetermined capacity
func New() Directed {
	return WithCapacity(16)
}

func FromOriginationMap(src OriginationMap) Directed {
	b := Builder{WithCapacity(len(src))}
	b.AddEdges(src)
	return b.G
}

// adjanceny list, insertion order is not preserved
// do not construct directly, use a provided ctor
// direct usage of Directed is readonly
type Directed struct {
	origins map[Origination]bitset.Bitset64
	dests   map[Destination]bitset.Bitset64
}

func (g Directed) NodeCount() int {
	return len(g.origins)
}

// given Node n, find all other nodes that point at it
func (g Directed) Predecessors(n Node) (bitset.Bitset64, error) {
	destinations, found := g.dests[Destination(n)]
	if !found {
		return bitset.Bitset64{}, ErrDestNotFound
	}
	return bitset.Copy(destinations), nil
}

// given Node n, find all nodes that it points at
func (g Directed) Successors(n Node) (bitset.Bitset64, error) {
	origins, found := g.origins[Origination(n)]
	if !found {
		return bitset.Bitset64{}, ErrOriginNotFound
	}
	return bitset.Copy(origins), nil
}
