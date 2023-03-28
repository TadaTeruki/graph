# Graph

This package provides a simple implementation of a graph data structures written in Go.

This package uses adjacency list representation for implementing graph data structures.

## features

`graph.DirectedGraph` : represents directed graph<br>
`graph.UndirectedGraph` : represents undirected graph<br>

Both structure implements `graph.Graph`.
```go
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
```

## example

```go
package main

import (
	"fmt"

	"github.com/TadaTeruki/graph"
)

func main() {
	g := graph.NewDirected(4) // Create a new graph with 4 vertices

	g.AddEdge(0, 1) // Add an edge from 0 to 1
	g.AddEdge(2, 3) // Add an edge from 0 to 2

	// Add more edges
	g.AddEdge(1, 2)
	g.AddEdge(0, 2)

	for i := 0; i < g.Order(); i++ {
		fmt.Println("Vertex", i, "has degree", g.GetDegree(i))
		g.ForEachSuccessor(i, func(j int) {
			fmt.Println("Vertex", i, "has edge for: ", j)
		})
	}
}

```