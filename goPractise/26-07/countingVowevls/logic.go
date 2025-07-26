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

// Writing this info to a new json file
type UserVowels struct {
	Name       string `json:"name"`
	VowelCount int    `json:"vowel_count"`
}

func writeToFile(fileName string, users []UserVowels) error {
	data, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		fmt.Println("Error marhasling from struct to json")
		return err
	}

	if err := os.WriteFile(fileName, data, 0644); err != nil {
		fmt.Println("Error writing to file")
		return err
	}
	return nil
}

func main() {
	file := "./file.json"
	data := unmarshalJson(file)

	fmt.Println(data)

	var results []UserVowels
	for _, user := range data {
		number := countVowels(user.Name)
		fmt.Printf("The user %s has %d vowels in their name\n", user.Name, number)

		results = append(results, UserVowels{
			Name:       user.Name,
			VowelCount: number,
		})
	}
	if err := writeToFile("results.json", results); err != nil {
		fmt.Println("Error writting to file: ", err)
	} else {
		fmt.Println("Results are written to results.json")
	}
}
