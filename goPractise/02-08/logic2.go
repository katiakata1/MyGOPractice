// Using approach with indexes instead - more efficeint
// No map overhead: Maps in Go use hashing, which is slower than direct array indexing.
// Fixed size array: Only 26 elements, so memory allocation and access are very fast and predictable.
// Two-pass algorithm: Still O(n) time complexity, but with less constant overhead.

package main

import "fmt"

func firstUniqCh(s string) int {
	counters := [26]int{} // 26 letters from a to z

	for i := 0; i < len(s); i++ {
		counters[s[i]-'a']++
	}
	// loops through s again, that's how it determines the first ch with 1
	for i := 0; i < len(s); i++ {
		if counters[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	words := []string{"leetcode", "loveleetcode", "aabb", "abcabcde", "z"}
	for _, word := range words {
		index := firstUniqCh(word)
		fmt.Printf("The word `%s` has one unique letter at index %v\n", word, index)
	}
}
