// Find if any two values make one target value. []int{1,2,6,3}, target = 5
// return indixes of those two numbers

package main

func twoSum(nums []int, target int) []int {
	revised := map[int]int{}

	for i, num := range nums {
		compliment := target - num
		if j, exists := revised[compliment]; exists {
			return []int{i, j}
		} else {
			revised[num] = i
		}
	}

	return nil
}

// func main() {
// 	nums := []int{1, 2, 4, 5, 6}
// 	target := 5
// 	calc := twoSum(nums, target)
// 	fmt.Println("The two numbers are: ", calc)
// }
