// Package graph implements a simple graph data structure.
package graph

import "container/list"

// DirectedGraph is the directed graph.
type DirectedGraph struct {
	link []list.List
}

// NewDirectedGraph creates a new directed graph with a given order (number of vertices).
func NewDirectedGraph(order int) *DirectedGraph {
	return &DirectedGraph{link: make([]list.List, order)}
}

// AddEdge adds a directed edge from i1 to i2 in the graph.
func (g *DirectedGraph) AddEdge(i1, i2 int) {
	g.link[i1].PushBack(i2)
}

// DeleteEdge removes a directed edge from i1 to i2 in the graph.
func (g *DirectedGraph) DeleteEdge(i1, i2 int) {
	for n := g.link[i1].Front(); n != nil; n = n.Next() {
		if n.Value == i2 {
			g.link[i1].Remove(n)
			break
		}
	}
}

// ForEachSuccessor applies the given function to each vertex that is a successor of the vertex with the given index.
func (g *DirectedGraph) ForEachSuccessor(index int, f func(int)) {
	for n := g.link[index].Front(); n != nil; n = n.Next() {
		f(n.Value.(int))
	}
}

// GetDegree returns the degree of the vertex with the given index.
func (g *DirectedGraph) GetDegree(index int) int {
	return g.link[index].Len()
}

// HasEdge returns true if there is an edge from i1 to i2 in the graph, and false otherwise.
func (g *DirectedGraph) HasEdge(i1, i2 int) bool {
	for n := g.link[i1].Front(); n != nil; n = n.Next() {
		if n.Value == i2 {
			return true
		}
	}
	return false
}

// Order returns the number of vertices in the graph.
func (g *DirectedGraph) Order() int {
	return len(g.link)
}

// Size returns the number of edges in the graph. Requires O(N) to compute.
func (g *DirectedGraph) Size() int {
	size := 0
	for _, neighbors := range g.link {
		size += neighbors.Len()
	}
	return size
}

// Clear removes all edges from the graph.
func (g *DirectedGraph) Clear() {
	for i := range g.link {
		g.link[i].Init()
	}
}
