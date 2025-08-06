// Leetcode 167: Two Sum II – Input Array Is Sorted

package main

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left, right}
		} else if numbers[left]+numbers[right] < target {
			left += 1
		} else {
			right -= 1
		}
	}

	return []int{-1, -1}
}

// func main() {
// 	numbers := []int{1, 3, 4, 5, 7, 10, 12}
// 	target := 15
// 	result := twoSum(numbers, target)
// 	fmt.Println("The numbers are: ", result)
// }
