// Package graph implements a simple graph data structure.
package graph

import "container/list"

// UndirectedGraph is the undirected graph.
type UndirectedGraph struct {
	link []list.List
}

// NewUndirected creates a new undirected graph with a given order (number of vertices).
func NewUndirected(order int) *UndirectedGraph {
	return &UndirectedGraph{link: make([]list.List, order)}
}

// AddEdge adds an undirected edge between i1 and i2 in the graph.
func (g *UndirectedGraph) AddEdge(i1, i2 int) {
	g.link[i1].PushBack(i2)
	g.link[i2].PushBack(i1)
}

// DeleteEdge removes an undirected edge between i1 and i2 in the graph.
func (g *UndirectedGraph) DeleteEdge(i1, i2 int) {
	for n := g.link[i1].Front(); n != nil; n = n.Next() {
		if n.Value == i2 {
			g.link[i1].Remove(n)
			break
		}
	}
	for n := g.link[i2].Front(); n != nil; n = n.Next() {
		if n.Value == i1 {
			g.link[i2].Remove(n)
			break
		}
	}
}

// ForEachSuccessor applies the given function to each vertex that is a successor of the vertex with the given index.
func (g *UndirectedGraph) ForEachSuccessor(index int, f func(int)) {
	for n := g.link[index].Front(); n != nil; n = n.Next() {
		f(n.Value.(int))
	}
}

// GetDegree returns the degree of the vertex with the given index.
func (g *UndirectedGraph) GetDegree(index int) int {
	return g.link[index].Len()
}

// HasEdge returns true if there is an edge between i1 and i2 in the graph, and false otherwise.
func (g *UndirectedGraph) HasEdge(i1, i2 int) bool {
	for n := g.link[i1].Front(); n != nil; n = n.Next() {
		if n.Value == i2 {
			return true
		}
	}
	return false
}

// Order returns the number of vertices in the graph.
func (g *UndirectedGraph) Order() int {
	return len(g.link)
}

// Size returns the number of edges in the graph. Requires O(N) to compute.
func (g *UndirectedGraph) Size() int {
	size := 0
	for _, neighbors := range g.link {
		size += neighbors.Len()
	}
	return size / 2
}

// Clear removes all edges from the graph.
func (g *UndirectedGraph) Clear() {
	for i := range g.link {
		g.link[i].Init()
	}
}
