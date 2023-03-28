// Package graph implements a simple graph data structure.
package graph

// Graph is an interface for a graph data structure.
type Graph interface {
	AddEdge(i1, i2 int)
	DeleteEdge(i1, i2 int)
	ForEachSuccessor(index int, f func(int))
	GetDegree(index int) int
	HasEdge(i1, i2 int) bool
	Order() int
	Size() int
	Clear()
}
