package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func toLinkedList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	current := head

	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func isPalindrome(head *ListNode) bool {
	return isPalindromeRecursive(&head, head)
}

func isPalindromeRecursive(head **ListNode, current *ListNode) bool {
	if current == nil {
		return true
	}

	if !isPalindromeRecursive(head, current.Next) {
		return false
	}

	if (*head).Val != current.Val {
		return false
	}
	*head = (*head).Next

	return true
}

// func main() {
// 	arr := []int{1, 2, 2, 1}
// 	linkedList := toLinkedList(arr)
// 	printList(linkedList)
// 	fmt.Println("Is it Palindrome: ", isPalindrome(linkedList))

// 	arr1 := []int{3, 9, 3, 6}
// 	linkedList1 := toLinkedList(arr1)
// 	printList(linkedList1)
// 	fmt.Println("Is it Palindrome: ", isPalindrome(linkedList1))
// }
