// parser data To and Times, and print it out as many times as Times say

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Message struct {
	To    string `json:"to"`
	Times int    `json:"times"`
}

func main() {
	file, err := os.Open("./file.jsonl")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var msg Message
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		for i := 0; i < msg.Times; i++ {
			fmt.Printf("Hello, %s!\n", msg.To)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}
}
