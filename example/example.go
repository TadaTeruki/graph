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
