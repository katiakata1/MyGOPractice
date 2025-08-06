package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	sorted := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)
	return sorted
}

func binarySearch(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]+nums[right] == target {
			return []int{left, right}
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}

func main() {
	numbers := []int{8, 3, 5, 1, 7, 4}
	sortedNums := mergeSort(numbers)
	target := 9
	result := binarySearch(sortedNums, target)
	fmt.Println("The two indixes are:", result)
}
