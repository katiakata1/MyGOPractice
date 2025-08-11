// You also can count numbers at specific idnexes simultaneously
// because both lists if they identical whould have only 0s

package main

import "fmt"

func isAnagram1(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, num := range count {
		if num != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	answer := isAnagram1(s, t)
	fmt.Println(answer)
}
