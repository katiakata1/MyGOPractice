// Dijkstar's algorithm is best to finding fastest path

package main

import (
	"fmt"
	"math"
)

// Edge is connection between two nodes
type Edge struct {
	Target *Node
	Weight float64
}

type Node struct {
	Name     string
	Edges    []*Edge
	Cost     float64
	Previous *Node
	Visited  bool
}

type Graph struct {
	Nodes []*Node
}

// AddNodes adds all nodes into the graph
func (g *Graph) AddNodes(name string) *Node {
	node := &Node{Name: name, Cost: math.Inf(1)}
	g.Nodes = append(g.Nodes, node)
	return node
}

// AddEdge connects two nodes together
func (n *Node) AddEdge(target *Node, weight float64) {
	n.Edges = append(n.Edges, &Edge{Target: target, Weight: weight})
}

// Find the lowestCostNode
func (g *Graph) lowestCostNode() *Node {
	var lowestCostNode *Node
	lowestCost := math.Inf(1)

	for _, node := range g.Nodes {
		if !node.Visited && node.Cost < lowestCost {
			lowestCost = node.Cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

// Dij alg
func (g *Graph) Dij(start *Node) {
	start.Cost = 0

	for {
		current := g.lowestCostNode()
		if current == nil {
			break
		}

		current.Visited = true

		for _, edge := range current.Edges {
			newCost := current.Cost + edge.Weight
			if newCost < edge.Target.Cost {
				edge.Target.Cost = newCost
				edge.Target.Previous = current
			}
		}
	}
}

// Find the shortest path
func (g *Graph) GetShortestPath(end *Node) ([]string, float64) {
	path := []string{}
	current := end

	for current != nil {
		path = append([]string{current.Name}, path...)
		current = current.Previous
	}

	return path, end.Cost
}

func main() {
	graph := &Graph{}

	start := graph.AddNodes("start")
	a := graph.AddNodes("A")
	b := graph.AddNodes("B")
	fin := graph.AddNodes("fin")

	start.AddEdge(a, 6)
	start.AddEdge(b, 2)
	a.AddEdge(fin, 1)
	b.AddEdge(a, 3)
	b.AddEdge(fin, 5)

	graph.Dij(start)
	path, cost := graph.GetShortestPath(fin)

	fmt.Printf("Shortest path from Start to Fin: %v with a cost of %.2f\n", path, cost)

}
