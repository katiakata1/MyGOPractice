// JSON unmarshalling with palindrome checking for names.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name string `json:"name"`
}

func parsingJson(fileName string) []User {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error while closing file: ", err)
			return
		}
	}()

	userList := []User{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var user User
		if err := json.Unmarshal([]byte(line), &user); err != nil {
			fmt.Println("Error unmarshaling json to struct:", err)
			return nil
		}

		userList = append(userList, user)

	}

	return userList
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	file := "./file.jsonl"
	user := parsingJson(file)

	fmt.Println("Unmarshalled output: ", user)

	for _, username := range user {
		answer := isPalindrome(username.Name)
		fmt.Printf("The user's name %s is palindrome - %t\n", username.Name, answer)
	}
}
