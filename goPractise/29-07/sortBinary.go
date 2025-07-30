// Sort numbers and then binary search
package main

import "fmt"

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func searchRange1(nums []int, target int) []int {
	first := binarySearch1(nums, target, true)
	if first == -1 {
		return []int{-1, -1}
	}

	last := binarySearch1(nums, target, false)
	return []int{first, last}
}

func binarySearch1(nums []int, target int, isFirst bool) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			if isFirst == true {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func main() {
	nums := []int{3, 4, 3, 1, 7, 4, 5, 6, 3}
	ordered := mergeSort(nums)
	fmt.Println(ordered)

	target := 4
	answer := searchRange1(ordered, target)
	fmt.Println(answer)
}
