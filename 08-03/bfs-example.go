package main

// Node represents a person in the network
type Node1 struct {
	name        string
	connections []*Node1
}

// isMangoSeller checks if a person mango seller
func isMangoSeller(person *Node1) bool {
	return len(person.name) > 0 && person.name[len(person.name)-1] == 'm'
}

// BFS to find mango seller using Nodes
func findMangoSeller(start *Node1) *Node1 {
	queue := []*Node1{start}         //queue initialised with start node
	visited := make(map[*Node1]bool) // track visited nodes

	for len(queue) > 0 {
		person := queue[0] // Dequeue first person
		queue = queue[1:]

		if isMangoSeller(person) {
			return person
		}

		if !visited[person] {
			visited[person] = true // mark as visited
			queue = append(queue, person.connections...)
		}
	}

	return nil // no mango seller found
}

// func main() {
// 	you := &Node1{name: "You"}
// 	alice := &Node1{name: "Alice"}
// 	bob := &Node1{name: "Bob"}
// 	claire := &Node1{name: "Claire"}
// 	peggy := &Node1{name: "Peggy"}
// 	anuj := &Node1{name: "Anuj"}
// 	thom := &Node1{name: "Thom"}
// 	jonny := &Node1{name: "Jonny"}

// 	//connecting nodes
// 	you.connections = []*Node1{alice, bob, claire}
// 	alice.connections = []*Node1{peggy}
// 	bob.connections = []*Node1{anuj, peggy}
// 	claire.connections = []*Node1{thom, jonny}
// 	anuj.connections = []*Node1{}
// 	peggy.connections = []*Node1{}
// 	thom.connections = []*Node1{}
// 	jonny.connections = []*Node1{}

// 	seller := findMangoSeller(you)
// 	if seller != nil {
// 		fmt.Println("Mango seller found: ", seller.name)
// 	} else {
// 		fmt.Println("No seller is found ")
// 	}

// }
