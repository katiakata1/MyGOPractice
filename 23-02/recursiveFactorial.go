package main

func recursive(num int) int {
	if num == 0 {
		return 1
	}
	return num * (recursive(num - 1))
}

// func main() {
// 	fact := recursive(7)

// 	fmt.Println(fact)
// }
