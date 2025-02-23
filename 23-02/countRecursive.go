package main

func count(arr []string) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + count(arr[1:])
}

// func main() {
// 	array := []string{"m", "g", "d", "w", "e", "m", "g", "d", "w", "e"}

// 	fmt.Println(count(array))

// }
