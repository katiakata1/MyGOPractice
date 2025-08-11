// Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without duplicate characters.
// sliding windows approach

// The point here is to save a character in the map if it doesn't exist as key and its current index as value
// Then you ittirate through word, chech characters and if they exist in the map
// if the do, then you check their index in the map and if it's greater or equal to index of the start, you move start to i+1
// meaning that you one step ahead of the letter which was last seen already

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxLenght := 0
	lastSeen := map[byte]int{}

	for end := 0; end < len(s); end++ {
		char := s[end]

		if i, exists := lastSeen[char]; exists && i >= start {
			start = i + 1
		}

		lastSeen[char] = end

		if end-start+1 > maxLenght {
			maxLenght = end - start + 1
		}
	}

	return maxLenght
}

func main() {
	string := "abcabcbb"
	num := lengthOfLongestSubstring(string)
	fmt.Println("The longest Substring Without Repeating Characters is: ", num)
}
