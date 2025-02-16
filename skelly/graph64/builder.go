package graph64

import (
	"github.com/etc-sudonters/substrate/skelly/bitset64"
)

type Builder struct {
	Graph *Directed
}

func (this *Builder) AddNode(n Node) {
	if _, exists := this.Graph.origins[n]; !exists {
		this.Graph.origins[n] = bitset64.Bitset{}
	}
}

func (this *Builder) AddNodes(ns []Node) {
	for _, n := range ns {
		this.AddNode(n)
	}
}

func (this *Builder) AddEdge(origin, dest Node) {
	this.AddNode(origin)
	this.AddNode(dest)
	origins := this.Graph.origins[origin]
	bitset64.Set(&origins, dest)
	this.Graph.origins[origin] = origins
}

func (this *Builder) AddRoot(root Node) {
	bitset64.Set(&this.Graph.roots, root)
}
