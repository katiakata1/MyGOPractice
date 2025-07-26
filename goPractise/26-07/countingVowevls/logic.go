// Task: Count Vowels in a Name (from JSON)

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Users struct {
	Name string `json:"name"`
}

func unmarshalJson(fileName string) []Users {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return nil
	}

	var users []Users
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("Error unmarshalling from JSON to struct: ", err)
		return nil
	}

	return users
}

func countVowels(s string) int {
	s = strings.ToLower(s)
	count := 0
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
		}
	}

	return count
}

func main() {
	file := "./file.json"
	data := unmarshalJson(file)

	fmt.Println(data)

	for _, user := range data {
		number := countVowels(user.Name)
		fmt.Printf("The user %s has %d vowels in their name\n", user.Name, number)
	}
}
