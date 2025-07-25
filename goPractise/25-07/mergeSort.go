// Time: O(n log n); Space: O(n); stable, Stable, good for linked lists

package main

import "fmt"

func mergeSort(arr []int) []int {
	// Base case: single element is already sorted
	for len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	fmt.Printf("Merging left: %v and right: %v\n", left, right)
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	fmt.Printf("→ Merged result: %v\n\n", result)
	return result
}

func main() {
	arr := []int{5, 7, 4, 3, 10, 11, 2}
	fmt.Println("Original array:", arr)
	sorted := mergeSort(arr)
	fmt.Println("Sorted array:", sorted)
}
