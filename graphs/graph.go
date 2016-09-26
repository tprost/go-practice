package graph

import "errors"

type Graph struct {
	nodes map[*GraphNode]*GraphNode
}

type GraphNode struct {
	neighbors map[*GraphNode]*GraphNode
}

func NewGraph() *Graph {
	graph := Graph{}
	graph.nodes = make(map[*GraphNode]*GraphNode)
	return &graph
}

func NewGraphNode() *GraphNode {
	node := GraphNode{}
	node.neighbors = make(map[*GraphNode]*GraphNode)
	return &node
}

func (graph *Graph) Add() *GraphNode {
	node := NewGraphNode()
	graph.nodes[node] = node
	return node
}

func (graph *Graph) AddEdge(x, y *GraphNode) error {
	_, xExists := graph.nodes[x]
	_, yExists := graph.nodes[y]
	if !xExists || !yExists {
		return errors.New("nodes are not in the graph")
	}
	x.neighbors[y] = y
	return nil
}

func (x *GraphNode) HasChild(y *GraphNode) bool {
	_, exists := x.neighbors[y]
	return exists
}

func (graph *Graph) HasRoute(x, y *GraphNode) (bool, error) {
	if x == nil || y == nil {
		return false, errors.New("nil node")
	}
	if graph.nodes[x] == nil || graph.nodes[y] == nil {
		return false, errors.New("node not in graph")
	}
	visited := map[*GraphNode]bool{}
	queue := []*GraphNode{}
	for _, neighbor := range x.neighbors {
		queue = append(queue, neighbor)
	}
	visited[x] = true
	for len(queue) > 0 {
		node := queue[0]
		if node == y {
			return true, nil
		}
		for _, neighbor := range node.neighbors {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
		visited[node] = true
		queue = queue[1:]
	}
	return false, nil
}
