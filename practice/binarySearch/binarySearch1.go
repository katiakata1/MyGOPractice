package main

func binarySearch(list []int, target int) int {
	left := 0
	rigth := len(list) - 1

	for left <= rigth {
		mid := left + (rigth-left)/2
		if list[mid] == target {
			return mid
		} else if list[mid] < target {
			left = mid + 1
		} else {
			rigth = mid - 1
		}
	}

	return -1
}

func main() {
	list := []int{1, 2, 3, 6, 9, 13, 16, 18, 25, 46, 35}

	target := 25
	result := binarySearch(list, target)
	if result != -1 {
		println("Element found at index:", result)
	} else {
		println("Element not found in the list")
	}
	target = 100
	result = binarySearch(list, target)
	if result != -1 {
		println("Element found at index:", result)
	} else {
		println("Element not found in the list")
	}
	target = 1
	result = binarySearch(list, target)
	if result != -1 {
		println("Element found at index:", result)
	} else {
		println("Element not found in the list")
	}
	target = 46
	result = binarySearch(list, target)
	if result != -1 {
		println("Element found at index:", result)
	} else {
		println("Element not found in the list")
	}
}
