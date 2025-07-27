// Given a slice of strings, group them by their lengths and count how many strings there are for each length.
package main

import "fmt"

func countLen(words []string) map[int]int {
	newMap := map[int]int{}

	for _, word := range words {
		lenght := len(word)
		if l, exists := newMap[lenght]; exists {
			newMap[lenght] = l + 1
		} else {
			newMap[lenght] = 1
		}
	}

	return newMap
}

func main() {
	words := []string{"go", "golang", "code", "chat", "gopher", "ai", "openai", "open"}
	result := countLen(words)
	fmt.Println("Kay is a word len, and value is number of how many time it repeats: ", result)
}
