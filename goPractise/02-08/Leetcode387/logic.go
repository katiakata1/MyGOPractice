//Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
// Go knows about runes and bytes based on how you loop through them:
// In this example, it's rune because I loop through each character of the string, that can contain multiple bytes (let's say if I had emoji, it would count as a single character with many bytes)
// If I used index looping (for i := 0; i < len(s); i++) it would mean I loop through each byte. Therefore, for emoji I would get 4 bytes.

package main

func unique(s string) int {
	countLetters := map[rune]int{}

	for _, letter := range s {
		countLetters[letter]++
	}

	for i, letter := range s {
		if countLetters[letter] == 1 {
			return i
		}
	}

	return -1
}

// func main() {
// 	words := []string{"leetcode", "loveleetcode", "aabb", "abcabcde", "z"}
// 	for _, word := range words {
// 		index := unique(word)
// 		fmt.Printf("The word `%s` has one unique letter at index %v\n", word, index)
// 	}
// }
