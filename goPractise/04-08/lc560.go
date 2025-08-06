// Subarray Sum Equals K
// An array nums, An integer k, You need to count how many subarrays (contiguous slices) of nums add up to k.

package main

import "fmt"

func subarraySum(nums []int, k int) int {
	subFreq := map[int]int{0: 1}
	count := 0
	prefixSum := 0

	for _, num := range nums {
		prefixSum += num

		if freq, exists := subFreq[prefixSum-k]; exists {
			count += freq
		}

		subFreq[prefixSum]++
	}

	return count
}

func main() {
	nums := []int{3, 4, 7, 2, -3, 1, 4, 2}
	k := 7
	result := subarraySum(nums, k)
	fmt.Println("Number of subarrays with sum %d: %d\n", k, result)
}
