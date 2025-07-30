// Find First and Last Position of Element in Sorted Array
// Given a sorted array of integers nums and a target value target, find the starting and ending position of the given target value.
// If the target is not found, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity (hint: binary search!).

package main

func searchRange(nums []int, target int) []int {
	first := binarySearch(nums, target, true)
	if first == -1 {
		return []int{-1, -1}
	}
	last := binarySearch(nums, target, false)
	return []int{first, last}
}

func binarySearch(nums []int, target int, isFirst bool) int {
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

// func main() {
// 	nums := []int{1, 2, 3, 3, 3, 3, 5, 6}
// 	target := 3
// 	answer := searchRange(nums, target)
// 	fmt.Println(answer)
// }
