package main

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sum(arr[1:])
}

// func main() {
// 	array := []int{3, 5, 7, 9}

// 	fmt.Println(sum(array))
// }
