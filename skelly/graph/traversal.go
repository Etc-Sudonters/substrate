package graph

import (
	"iter"

	"github.com/etc-sudonters/substrate/skelly/bitset"
	"github.com/etc-sudonters/substrate/skelly/queue"
	"github.com/etc-sudonters/substrate/skelly/stack"
)

type SuccessorsIter struct{ Directed }

func (i SuccessorsIter) BFS(from Node) iter.Seq[Node] {
	return visitSuccessors(i.Directed, &queue.Q[Node]{from})
}

func (i SuccessorsIter) DFS(from Node) iter.Seq[Node] {
	return visitSuccessors(i.Directed, &stack.S[Node]{from})
}


func visitSuccessors(g Directed, tracker tracker[Node]) iter.Seq[Node] {
	return func(yield func(v Node) bool) {
		visited := visitedset(g.NodeCount())
		var node Node
		for tracker.Len() > 0 {
			node, _ = tracker.Pop()

			if !bitset.Set(&visited, node) {
				continue
			}

			if !yield(node) {
				break
			}

			successors, err := g.Successors(node)
			if err != nil {
                panic(err)
			}

			for successor := range bitset.Iter64T[Node](successors).All {
                tracker.Push(successor)
			}
		}
	}
}

type tracker[T any] interface {
	Push(T)
	Pop() (T, error)
	Len() int
}


func visitedset(nodeCount int) bitset.Bitset64 {
	return bitset.New(bitset.Buckets(uint64(nodeCount)))
}

