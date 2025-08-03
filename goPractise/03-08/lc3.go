// Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without duplicate characters.
// sliding windows approach

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
