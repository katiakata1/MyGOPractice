// Group Users by Country from JSON and Count
// 1. Unmarshal the data from JSON.
// 2. Group the users by their country.
// 3. Count how many users are from each country.
// 4. Marshal the result into JSON and write it to a file.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email`
	Country string `json:"country"`
}

func unmarshalJson(fileName string) []User {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading json file: ", err)
		return nil
	}

	var user []User
	if err := json.Unmarshal(data, &user); err != nil {
		fmt.Println("Error unmarshaling from json to struct: ", err)
		return nil
	}

	return user
}

func countCountry(users []User) map[string]int {
	newMap := map[string]int{}

	for _, user := range users {
		if num, exists := newMap[user.Country]; exists {
			newMap[user.Country] = num + 1
		} else {
			newMap[user.Country] = 1
		}
	}

	return newMap
}

func writeToFile(users map[string]int) {
	data, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		fmt.Println("Error marshaling from map to byte: ", err)
		return
	}

	if err := os.WriteFile("result.json", data, 0644); err != nil {
		fmt.Println("Error writting to file: ", err)
		return
	}

	fmt.Println("The data is written to results.json")
}

func main() {
	file := "./file.json"
	users := unmarshalJson(file)
	answer := countCountry(users)
	fmt.Println("Countries and how many people are from that country: ", answer)
	writeToFile(answer)
}
