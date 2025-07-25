// QUICKSORT - Sorting algorithm
// swap - It avoids creating extra arrays, so it’s space-efficient (O(1) extra space)
// we are always operating on the original array
// Avg: O(n log n), Worst: O(n²); O(log n), not stable, Fast, in-place, divide & conquer

package main

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			// swapping arr[j] and arr[i]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	//swapping pivot next to the numbers less than pivot
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

// quickSort sorts the array in-place indices low and high
func quickSort(arr []int, low, high int) {
	if low < high {
		// partition the array to get a pivot index
		pivotIndex := partition(arr, low, high)

		// Now devide this into left and right - recursion
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

// func main() {
// 	nums := []int{4, 2, 8, 4, 1, 5, 3, 9, 13, 56, 3}
// 	fmt.Println("Before:", nums)
// 	quickSort(nums, 0, len(nums)-1)
// 	fmt.Println("After:", nums)
// }
