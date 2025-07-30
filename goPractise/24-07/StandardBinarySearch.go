package main

import "fmt"

func binarySearch(nums []int, target int) int {
	low := nums[0]
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	numbers := []int{2, 3, 7, 9, 13, 17}
	target := 17
	answer := binarySearch(numbers, target)

	fmt.Println("the answer is:", answer)
}
