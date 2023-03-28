package graph_test

import (
	"testing"

	"github.com/TadaTeruki/graph"
)

func TestDirectedGraph(t *testing.T) {
	// Create a new graph with 4 vertices
	g := graph.NewDirectedGraph(4)

	// Test the Order method
	if g.Order() != 4 {
		t.Errorf("Expected order 4, got %d", g.Order())
	}

	// Test the Size method (should be 0 since the graph is empty)
	if g.Size() != 0 {
		t.Errorf("Expected size 0, got %d", g.Size())
	}

	// Test the AddEdge method
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	// Test the Size method (should be 4 since we added 4 edges)
	if g.Size() != 4 {
		t.Errorf("Expected size 4, got %d", g.Size())
	}

	// Test the HasEdge method
	if !g.HasEdge(0, 1) {
		t.Error("Expected edge (0,1) to exist")
	}
	if g.HasEdge(1, 0) {
		t.Error("Expected edge (1,0) to not exist")
	}
	if !g.HasEdge(2, 3) {
		t.Error("Expected edge (2,3) to exist")
	}
	if g.HasEdge(3, 2) {
		t.Error("Expected edge (3,2) to not exist")
	}

	// Test the DeleteEdge method
	g.DeleteEdge(0, 1)
	if g.HasEdge(0, 1) {
		t.Error("Expected edge (0,1) to be deleted")
	}
	if g.Size() != 3 {
		t.Errorf("Expected size 3, got %d", g.Size())
	}

	// Test the ForEachSuccessor method
	var visited []int
	g.ForEachSuccessor(2, func(vertex int) {
		visited = append(visited, vertex)
	})
	if len(visited) != 1 {
		t.Error("Expected 1 successor for vertex 2")
	}
}

func TestUndirectedGraph(t *testing.T) {
	// Create a new graph with 4 vertices
	g := graph.NewUndirectedGraph(4)

	// Test the Order method
	if g.Order() != 4 {
		t.Errorf("Expected order 4, got %d", g.Order())
	}

	// Test the Size method (should be 0 since the graph is empty)
	if g.Size() != 0 {
		t.Errorf("Expected size 0, got %d", g.Size())
	}

	// Test the AddEdge method
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	// Test the Size method (should be 4 since we added 4 edges)
	if g.Size() != 4 {
		t.Errorf("Expected size 4, got %d", g.Size())
	}

	// Test the HasEdge method
	if !g.HasEdge(0, 1) {
		t.Error("Expected edge (0,1) to exist")
	}
	if !g.HasEdge(1, 0) {
		t.Error("Expected edge (1,0) to exist")
	}
	if g.HasEdge(0, 3) {
		t.Error("Expected edge (0,3) to not exist")
	}
	if g.HasEdge(3, 0) {
		t.Error("Expected edge (3,0) to not exist")
	}

	// Test the DeleteEdge method
	g.DeleteEdge(0, 1)
	if g.HasEdge(0, 1) {
		t.Error("Expected edge (0,1) to be deleted")
	}
	if g.HasEdge(1, 0) {
		t.Error("Expected edge (1,0) to be deleted")
	}
	if g.Size() != 3 {
		t.Errorf("Expected size 3, got %d", g.Size())
	}

	// Test the ForEachSuccessor method
	var visited []int
	g.ForEachSuccessor(2, func(vertex int) {
		visited = append(visited, vertex)
	})
	if len(visited) != 3 {
		t.Error("Expected 3 successors for vertex 2")
	}
}
