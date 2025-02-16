package graph64

import (
	"errors"
	"fmt"
	"github.com/etc-sudonters/substrate/skelly/bitset64"
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
		origins: make(map[Node]bitset64.Bitset, c),
		roots:   bitset64.Bitset{},
	}
}

func New() Directed {
	return WithCapacity(16)
}

type Directed struct {
	origins map[Node]bitset64.Bitset
	roots   bitset64.Bitset
}

func (g Directed) Roots() bitset64.Bitset {
	return bitset64.Copy(g.roots)
}

func (g Directed) NodeCount() int {
	return len(g.origins)
}

// given Node n, find all nodes that it points at
func (g Directed) Successors(n Node) (bitset64.Bitset, error) {
	origins, found := g.origins[n]
	if !found {
		return bitset64.Bitset{}, ErrOriginNotFound
	}
	return bitset64.Copy(origins), nil
}

var ErrOriginNotFound = errors.New("origin not found")
