package graph

import "testing"

func TestNeighbors(t *testing.T) {
	graph := NewGraph()
	node1 := graph.Add()
	node2 := graph.Add()
	graph.Add()
	graph.AddEdge(node1, node2)
	if !node1.HasChild(node2) {
		t.Error("node1 is supposed to have node2 as a child")
	}
}

func TestHasRoute(t *testing.T) {
	graph := NewGraph()
	node1 := graph.Add()
	node2 := graph.Add()
	node3 := graph.Add()
	node4 := graph.Add()
	node5 := graph.Add()
	graph.AddEdge(node1, node2)
	graph.AddEdge(node2, node3)
	graph.AddEdge(node3, node1)
	graph.AddEdge(node3, node4)
	graph.AddEdge(node4, node5)
	hasRoute, error := graph.HasRoute(node1, node5)
	if error != nil || hasRoute == false {
		t.Error("node1 has route to node3 but HasRoute returned false or error")
	}
}
