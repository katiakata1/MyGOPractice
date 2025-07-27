// Given a list of integers, return true if any number appears more than once.

package main

func checker(nums []int) (bool, int) {
	newMap := map[int]bool{}
	for _, num := range nums {
		if _, exists := newMap[num]; exists {
			return true, num
		} else {
			newMap[num] = true
		}
	}

	return false, -1
}

// func main() {
// 	nums := []int{1, 2, 4, 1}
// 	answer, num := checker(nums)
// 	fmt.Println("Dublicate number exists:", answer, num)
// }
