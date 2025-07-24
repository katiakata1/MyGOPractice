// I/O errors - Input/Output — basically, anything involving reading data from or writing data to external sources like
// scanner.Scan() reads the file line by line.
// If something unexpected happens while reading the file (e.g., disk failure), it stops and stores that error internally.
// After you finish scanning, you call scanner.Err() to check if any such errors occurred.
// If yes, you can handle or report it, so your program doesn’t silently ignore it.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Message struct {
	To string `json:"to"`
}

func main() {
	file, err := os.Open("./file.jsonl")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Scan() saves each line as string, trims anything aroung {}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Now this string is concerted to bytes and unmarshalled to Message struct
		var msg Message
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			fmt.Println("Error parsing:", err)
			continue
		}

		fmt.Printf("Hello, %s!\n", msg.To)
	}

	// checking if error happened during scanning (after the loop is done)
	// scanner.Err() checks for I/O errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}

}
