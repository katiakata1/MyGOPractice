// HashMap is very effient for TwoSum problem - map lookup
// Time Complexity: O(1) on average because it uses hast to jump to the right "bucket"
// Normal array search (loop throug it) would have O(n) time complexity because it needs to go through all elements in the array

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return []int{}
}

func main() {
	numbers := []int{3, 5, 7, 6}
	target := 11
	answer := twoSum(numbers, target)
	fmt.Println("Here it is:", answer)
}
