package main

func bubble(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// func main() {
// 	arr := []int{5, 3, 7, 0, 2, 1}
// 	fmt.Println("Before Sorting:", arr) // Prints original array

// 	bubble(arr)

// 	fmt.Println("Before Sorting:", arr) // Same array, but now sorted

// }
