package main

import "fmt"

func insertion(arr []int) {
	for i := 1; i < len(arr)-1; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{4, 7, 12, 2, 7, 29}
	fmt.Println("Before Sorting:", arr) // Prints original array

	insertion(arr)

	fmt.Println("Before Sorting:", arr) //Same array, but now sorted
}
