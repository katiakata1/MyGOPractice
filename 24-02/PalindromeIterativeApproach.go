package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	//check base case: if list is empty or has 1 value it is consideted to be palindrome
	if head == nil || head.Next == nil {
		return true
	}

	// now set the slow and fast pointers
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Now you need to reverse second half and set first half to point to head
	secondHalf := reverseList(slow)
	firstHalf := head

	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		secondHalf = secondHalf.Next
		firstHalf = firstHalf.Next
	}

	return true
}

// First check if array is emtpy, empty are not needed to be converted
// then set head to point newly created node (space in memory), head will stay pointing here
// move current based on lenght of arr
func convertToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head

	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// Set current point to head, don't change head because it will alter original list
// then go through each Val until current is nil
func printList(head *ListNode) {
	current := head

	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}

	fmt.Println()
}

func main() {
	arr := []int{1, 2, 3, 3, 2, 1}
	converted := convertToList(arr)
	printList(converted)
	fmt.Println("Is this list Palindrome: ", isPalindrome(converted))

	arr2 := []int{3, 6, 8, 9, 3}
	converted2 := convertToList(arr2)
	printList(converted2)
	fmt.Println("Is this list Palindrome: ", isPalindrome(converted2))
}
