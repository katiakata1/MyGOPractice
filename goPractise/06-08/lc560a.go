// Practice Task: Count Subarrays That Sum to K

package main

func subarraySum(arr []int, k int) int {
	subFreq := map[int]int{0: 1}
	prefix := 0
	count := 0

	for _, num := range arr {
		prefix += num

		if freq, exist := subFreq[prefix-k]; exist {
			count += freq
		}

		subFreq[prefix]++
	}

	return count
}

// func main() {
// 	nums := []int{1, -1, 2, 3, -2, 4, 1}
// 	k := 3
// 	result := subarraySum(nums, k)
// 	fmt.Println("The sum of subarrays is:", result)
// }
