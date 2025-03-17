package main

import "fmt"

type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

// GetHeight return height of the tree
func GetHeight(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

// Helper function to calculate max
func max(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// GetBalanceFactor calculates balance factor
func GetBalanceFactor(n *Node) int {
	if n == nil {
		return 0
	}
	return GetHeight(n.left) - GetHeight(n.right)
}

// RotateRight performs rotation if left tree heavy
func RotateRight(y *Node) *Node {
	x := y.left
	T2 := x.right

	// rotate
	x.right = y
	y.left = T2

	// Recalculate the height
	y.height = max(GetHeight(y.left), GetHeight(y.right)) + 1
	x.height = max(GetHeight(x.left), GetHeight(x.right)) + 1

	return x
}

// RotateLeft performs rotation if right side is heavy
func RotateLeft(x *Node) *Node {
	y := x.right
	T2 := y.left

	// rotate
	y.left = x
	x.right = T2

	// recalculate the heights
	y.height = max(GetHeight(y.left), GetHeight(y.right)) + 1
	x.height = max(GetHeight(x.left), GetHeight(x.right)) + 1

	return y
}

// Insert keys into AVL tree
func Insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{key: key, height: 1}
	}

	if key < root.key {
		root.left = Insert(root.left, key)
	} else if key > root.key {
		root.right = Insert(root.right, key)
	} else {
		return root
	}

	root.height = 1 + max(GetHeight(root.left), GetHeight(root.right))

	balance := GetBalanceFactor(root)

	// Perform rotations if unbalanced
	// Case 1: Left Heavy (Right rotation)
	if balance > 1 && key < root.left.key {
		return RotateRight(root)
	}

	// Case 2: Right Heavy (Left Rotation)
	if balance < -1 && key > root.right.key {
		return RotateLeft(root)
	}

	// Case 3: Left-Right case (Right-left rotation)
	if balance > 1 && key > root.left.key {
		root.left = RotateLeft(root.left)
		return RotateRight(root)
	}

	// Case 4: Right-Left case (Left-Right rotation)
	if balance > 1 && key < root.right.key {
		root.right = RotateRight(root.right)
		return RotateLeft(root)
	}

	return root
}

// Search searches for a key in AVL tree
func Search(root *Node, key int) bool {
	if root == nil {
		return false
	}
	if key < root.key {
		return Search(root.left, key)
	} else if key > root.key {
		return Search(root.right, key)
	}

	return true
}

// InOrderTraversal prints the tree in sorted order
func InOrderTraversal(root *Node) {
	if root != nil {
		InOrderTraversal(root.left)
		fmt.Printf("%d", root.key)
		InOrderTraversal(root.right)
	}
}

// func main() {
// 	var root *Node // Here we don't use &Node{} because we don't want to initialise immidiately

// 	keys := []int{10, 20, 30, 40, 50, 25}
// 	for _, key := range keys {
// 		root = Insert(root, key)
// 	}

// 	fmt.Println("In-order Traversal: ")
// 	InOrderTraversal(root)
// 	fmt.Println()

// 	// Search for element
// 	fmt.Println("Searching for 20: ", Search(root, 20))
// 	fmt.Println("Searching for 15: ", Search(root, 15))
// }
