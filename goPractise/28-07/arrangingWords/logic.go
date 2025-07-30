// Arrange words by their lenght. If two words have same lefth, keep them in a original order

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func unmaeshallJson(fileName string) []string {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("error reading the file", err)
		return nil
	}

	var fruits []string
	if err := json.Unmarshal(data, &fruits); err != nil {
		fmt.Println("error unmarshalling from json to array")
		return nil
	}

	return fruits
}

func mergeSort(arr []string) []string {
	for len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []string) []string {
	result := []string{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if len(left[i]) <= len(right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}

func main() {
	file := "./file.json"
	fruits := unmaeshallJson(file)
	fmt.Println("Fruits are: ", fruits)

	sorted := mergeSort(fruits)
	fmt.Println("Fruits after sorting:", sorted)
}
