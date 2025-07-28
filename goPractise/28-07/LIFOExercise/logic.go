// Recent Messages Tracker using JSONL + Stack (LIFO)

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Message struct {
	Name    string `json:"sender"`
	Content string `json:"content"`
}

func unmaeshallJson(fileName string) []Message {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening the file:", err)
		return nil
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("error closing the file: ", err)
			return
		}
	}()

	messages := []Message{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var message Message
		if err := json.Unmarshal([]byte(line), &message); err != nil {
			fmt.Println("error unmarshalling from json to struct: ", err)
			return nil
		}

		messages = append(messages, message)
	}

	return messages
}

type Stack struct {
	items []Message
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(val Message) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() Message {
	if s.IsEmpty() {
		fmt.Println("The stack is empty")
		return Message{}
	}

	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Peek() Message {
	if s.IsEmpty() {
		fmt.Println("The stack is empty")
		return Message{}
	}

	return s.items[len(s.items)-1]
}

func main() {
	file := "./file.jsonl"
	unmarshalled := unmaeshallJson(file)
	s := Stack{}

	for _, message := range unmarshalled {
		fmt.Printf("The name: %s and the message: %s\n", message.Name, message.Content)

		s.Push(message)
	}

	peek := s.Peek()
	fmt.Println("The last item in the stack is: ", peek)

	popped := s.Pop()
	fmt.Println("The popped item is: ", popped)

	// Pushing item manually
	s.Push(unmarshalled[0])
	peek = s.Peek()
	fmt.Println("After adding manually, The last item in the stack is: ", peek)

}
