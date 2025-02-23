package main

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, val := range arr[1:] {
		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)
}

// func main() {
// 	array := []int{6, 2, 4, 1, 12, 45}
// 	fmt.Println(quicksort(array))
// }
