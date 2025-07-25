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
	Age  int    `json:"age"`
}

func parseJson(fileName string) ([]User, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening the file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("warning: error closing the file: ", err)
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
			return nil, fmt.Errorf("Error unmarshaling json to struct: %v", err)
		}

		userList = append(userList, user)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Scan error: %w", err)
	}

	return userList, nil

}

// Quicksort
func quickSort(users []User, low, high int) {
	if low < high {
		pivotIndex := partition(users, low, high)
		quickSort(users, low, pivotIndex-1)
		quickSort(users, pivotIndex+1, high)
	}
}

func partition(users []User, low, high int) int {
	pivot := users[high]
	i := low - 1

	for j := low; j < high; j++ {
		if users[j].Age < pivot.Age {
			i++
			users[i], users[j] = users[j], users[i]
		}
	}

	users[i+1], users[high] = users[high], users[i+1]
	return i + 1
}

func main() {
	file := "./file.jsonl"
	output, _ := parseJson(file)
	fmt.Println("Unmarshaling Output:", output)

	quickSort(output, 0, len(output)-1)
	for _, user := range output {
		fmt.Printf("%s is %d years old\n", user.Name, user.Age)
	}
}
