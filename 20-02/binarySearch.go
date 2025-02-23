package main

func binary_search(array []int, guess int) bool {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2
		if array[mid] == guess {
			return true
		}
		if array[mid] < guess {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

// func main() {
// 	arr := []int{1, 4, 7, 9, 15, 19}
// 	guess := 19

// 	if binary_search(arr, guess) {
// 		fmt.Println("Found", guess)
// 	} else {
// 		fmt.Println("Not found", guess)
// 	}
// }
