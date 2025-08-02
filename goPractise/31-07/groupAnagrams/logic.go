// Leetcode 49
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	words := map[string][]string{}
	for _, str := range strs {
		byteString := []byte(str)

		sort.Slice(byteString, func(i, j int) bool {
			return byteString[i] < byteString[j]
		})

		sortedWord := string(byteString)

		// This can be simplified:
		// if _, exist := words[sortedWord]; exist {
		// 	words[sortedWord] = append(words[sortedWord], str)
		// } else {
		// 	words[sortedWord] = []string{str}
		// }

		words[sortedWord] = append(words[sortedWord], str)
	}

	arrayList := [][]string{}
	for _, value := range words {
		arrayList = append(arrayList, value)
	}

	return arrayList
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	words := groupAnagrams(strs)
	fmt.Println(words)
}
