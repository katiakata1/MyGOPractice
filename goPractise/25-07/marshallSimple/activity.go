package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Activity struct {
	Username string `json:"username"`
	Action   string `json:"action"`
	Time     string `json:"time"`
}

func main() {
	logs := []Activity{
		{"alice", "login", "2025-07-25T08:00:00Z"},
		{"tom", "upload_file", "2025-07-25T08:15:00Z"},
		{"charlie", "logout", "2025-07-25T08:30:00Z"},
	}

	data, err := json.MarshalIndent(logs, "", " ")
	if err != nil {
		fmt.Println("Error marhsalling: ", err)
		return
	}

	// Save the file
	err = os.WriteFile("activity_logs.json", data, 0644)
	if err != nil {
		fmt.Println("Error writting file: ", err)
		return
	}

	fmt.Println("Successfully wrote to activity_logs.json")
}
