package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to print the linked list
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println() // To print a newline at the end
}

// The functions that reverses the list
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // prev variables is initialised and is nil

	for head != nil {
		next := head.Next // next points to head.Next now which is &ListNode{Val: 2, Next:{Val: 3, Next:...}}
		head.Next = prev  // now head.Next &ListNode{Val: 1, Next:{Val: 2, Next:...}} points to prev which is nil, therefore head.Next is {Val: 1, Next:nil}
		prev = head       // prev was still nil, but here we make prev to point to head so now prev is {Val: 1, Next:nil}
		head = next       // head is still {Val: 1, Next:nil}, but now we make it to point to next which is &ListNode{Val: 2, Next:{Val: 3, Next:...}}
	}

	return prev
}

// the function that compares Val in original and reversed lists
func isPalindrome(head *ListNode) bool {
	reversedList := reverseList(head) // we reverse you head linked list
	currentOriginal := head           // we keep our otiginal linked list for comparison
	currentReversed := reversedList   // we keep reversed list for comparison

	for currentOriginal != nil && currentReversed != nil { // we check if these two nodes are emtpy
		if currentOriginal.Val != currentReversed.Val { // now we check Val in node
			return false
		}
		currentOriginal = currentOriginal.Next // now we need to move to Next which will have another Val
		currentReversed = currentReversed.Next
	}

	return true
}

// The function that convert strings to linked list
func createLinkedList(values []int) *ListNode {
	head := &ListNode{Val: values[0]} // we create a memory stack to hold ListNode struct with Val of values[0]. We point head to it and we will keep head here
	current := head                   // we point current to head. Now current points to {Val: values[0]}

	for _, val := range values[1:] { // we go through each value in a list (except indixes)
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

// the main function which runs
func main() {
	// Example 1: Palindrome list
	values1 := []int{1, 2, 2, 1}
	head1 := createLinkedList(values1)

	// Print original list
	fmt.Println("Original List 1:")
	printList(head1)

	// Check if palindrome
	fmt.Println("Is Palindrome?", isPalindrome(head1))

	// Example 2: Non-palindrome list
	values2 := []int{1, 5, 7, 9}
	head2 := createLinkedList(values2)

	// Print original list
	fmt.Println("Original List 2:")
	printList(head2)

	// Check if palindrome
	fmt.Println("Is Palindrome?", isPalindrome(head2))
}
