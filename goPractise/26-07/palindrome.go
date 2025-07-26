// To solve this efficiently:
// Find the middle of the list (slow/fast pointer approach).
// Reverse the second half of the list.
// Compare the first and second halves node by node.
// (Optional) Restore the list to its original form afterward.

// O(n) + O(n) + O(n) = O(n) linear time. (finding middle, reversing, comparing)
// O(1) constant space ( using only few pointers)

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Transform array to Linked List
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head

	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Find midle with slow/fast in LinkedList
func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head

	// fast checks if it's current position is nil or next one is nil
	// if the third position is nil, the loop exists and that's how we know we reached middle
	// Then we want to return slow
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow // slow this is the new middle. Returning fast so to check if it's nil it means list is odd
}

// Reverse the second half of the list starting from the middle node
// In example [1, 2, 3, 2, 1], we want to make 2, 1 after 3 to be 1, 2
// Reversing means changing the direction of Next pointers

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil // creating a variable pointing to *ListNode, which is not point to anything currently
	current := head

	for current != nil {
		next := current.Next // I want to save entire list starting from list[1:] - assume we have 1 -> 2 -> 3 -> nil
		current.Next = prev  // Now I want to break the link between 1st and 2nd node
		prev = current       // Now we want prev to point to smaller list (1->nil) rather than nil
		current = next       // and we need to move current to 2 in 2->3->nil list. the loop repeats until current reaches nil
	}

	return prev // prev is our new starting pointer
}

// Compares both halves
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// Find middle
	middle := findMiddle(head)

	// Check if list is odd
	// if fast != nil {
	// 	middle = middle.Next // Skip the exact middle node in odd-length case
	// }

	// Reverse the second half
	secondHalf := reverseList(middle)

	// Compare the first and second half
	first := head
	second := secondHalf

	for second != nil {
		if first.Val != second.Val {
			return false
		}

		first = first.Next // you need to move them now
		second = second.Next
	}

	// Optional: you can restore the list

	return true
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	nums := []int{1, 2, 3, 4, 3, 2, 1}
	linkedList := createLinkedList(nums)
	printList(linkedList)
	fmt.Println(isPalindrome(linkedList))
}
