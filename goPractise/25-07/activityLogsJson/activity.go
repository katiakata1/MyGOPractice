package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type FlatEntry struct {
	Username string `json:"name"`
	Age      int    `json:"age"`
	Activity string `json:"activity"`
}

func parsingJson(fileName string) []FlatEntry {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		return nil
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error while closing the file: ", err)
			return
		}
	}()

	flatEntries := []FlatEntry{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var flatEntry FlatEntry
		if err := json.Unmarshal([]byte(line), &flatEntry); err != nil {
			fmt.Println("Error unmarhsaling:", err)
			return nil
		}

		flatEntries = append(flatEntries, flatEntry)
	}

	return flatEntries
}

// NOW GROUPING ACTIVITIES UNDER SINGLE ACTIVITY
type UserActivity struct {
	Username   string   `json:"name"`
	Age        int      `json:"age"`
	Activities []string `json:"activities"`
}

func groupActivities(entries []FlatEntry) []UserActivity {
	grouped := make(map[string]UserActivity)

	for _, entry := range entries {
		user, exists := grouped[entry.Username]
		if exists {
			// user already seen - append activity
			user.Activities = append(user.Activities, entry.Activity)
		} else {
			// not exist yet, so need to create a new user
			user = UserActivity{
				Username:   entry.Username,
				Age:        entry.Age,
				Activities: []string{entry.Activity},
			}
		}
		grouped[entry.Username] = user
	}

	// converting map to slice
	result := make([]UserActivity, 0, len(grouped))
	for _, user := range grouped {
		result = append(result, user)
	}

	return result
}

// NOW I NEED TO SORT THIS LIST BY AGE - Mergesort
func mergeSort(arr []UserActivity) []UserActivity {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []UserActivity) []UserActivity {
	result := []UserActivity{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].Age < right[j].Age {
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

func writeToFile(activities []UserActivity) {
	text, err := json.MarshalIndent(activities, "", " ")
	if err != nil {
		fmt.Println("Error marshaling from struct to json", err)
		return
	}

	file := "output_activities.json"
	err = os.WriteFile(file, text, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	file := "./file.jsonl"
	parsing := parsingJson(file)
	grouping := groupActivities(parsing)
	sorting := mergeSort(grouping)

	fmt.Println("After sorting by age:", sorting)
	writeToFile(sorting)

}
