// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

package main

func isAnagram(s string, t string) bool {
	arrayS := [26]int{}
	arrayT := [26]int{}

	for i, _ := range s {
		arrayS[s[i]-'a']++
	}

	for i, _ := range t {
		arrayT[t[i]-'a']++
	}

	if arrayS == arrayT {
		return true
	}

	return false
}

// func main() {
// 	s := "anagram"
// 	t := "nagaram"
// 	answer := isAnagram(s, t)
// 	fmt.Println(answer)
// }
