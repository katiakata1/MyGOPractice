package main

import "fmt"

// Represents each a person in social network
type Node2 struct {
	name       string
	connection []*Node2
}

// NodeFactory stores created nodes to avoid dublicates
type NodeFactory struct {
	nodes map[string]*Node2
}

// NewNodeFactory initialises the factory
func NewNodeFactory() *NodeFactory {
	return &NodeFactory{nodes: make(map[string]*Node2)}
}

// GetNode returns an existing node or creates a new one
func (nf *NodeFactory) GetNode(name string) *Node2 {
	if node, exists := nf.nodes[name]; exists {
		return node
	}
	node := &Node2{name: name}
	nf.nodes[name] = node
	return node
}

// AddConnection adds bidirectional connections
func (nf *NodeFactory) AddConnection(name1, name2 string) {
	node1 := nf.GetNode(name1)
	node2 := nf.GetNode(name2)

	node1.connection = append(node1.connection, node2)
	node2.connection = append(node2.connection, node1)
}

// BFS Search for Mango seller
func isMangoSeller1(person *Node2) bool {
	return len(person.name) > 0 && person.name[len(person.name)-1] == 'm'
}

func FindMangoSeller1(start *Node2) *Node2 {
	queue := []*Node2{start} // Creating a new slice to hold reference to nodes
	visited := make(map[*Node2]bool)

	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]

		if visited[person] {
			continue
		}

		visited[person] = true

		if isMangoSeller1(person) {
			return person
		}

		for _, contact := range person.connection {
			if !visited[contact] {
				queue = append(queue, contact)
			}
		}
	}

	return nil
}

func main() {
	factory := NewNodeFactory()

	factory.AddConnection("you", "alice")
	factory.AddConnection("you", "bob")
	factory.AddConnection("you", "claire")
	factory.AddConnection("alice", "peggy")
	factory.AddConnection("bob", "anuj")
	factory.AddConnection("claire", "thom") // Thom is the seller
	factory.AddConnection("claire", "jonny")

	// Find mando seller
	seller := FindMangoSeller1(factory.GetNode("you"))
	if seller != nil {
		fmt.Println("Found mango seller: ", seller.name)
	} else {
		fmt.Println("Did not find mango seller")
	}

}
