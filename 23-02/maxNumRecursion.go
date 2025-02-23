package main

func max(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}
	sub_max := max(arr[1:])

	if arr[0] > sub_max {
		return arr[0]
	} else {
		return sub_max
	}
}

// func main() {
// 	array := []int{4, 5, 8, 2, 4, 9}
// 	fmt.Println(max(array))
// }
