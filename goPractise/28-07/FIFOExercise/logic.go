// 🧠 Quick Recap on FIFO (Queue)
// Push (Enqueue): Add to the end of the queue.
// Pop (Dequeue): Remove from the front of the queue.
// Think of it like a line at a coffee shop: first person in is the first one out.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Ticket struct {
	Name  string `json:"name"`
	Issue string `json:"issue"`
}

func unmaeshallJson(fileName string) []Ticket {
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

	tickets := []Ticket{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var ticket Ticket
		if err := json.Unmarshal([]byte(line), &ticket); err != nil {
			fmt.Println("error unmarshaling form json to struct: ", err)
			return nil
		}

		tickets = append(tickets, ticket)
	}

	return tickets
}

type Queue struct {
	items []Ticket
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Push(ticket Ticket) {
	q.items = append(q.items, ticket)
}

func (q *Queue) Pop() Ticket {
	if q.IsEmpty() {
		fmt.Println("The queue is empty")
		return Ticket{}
	}

	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func (q *Queue) Peek() Ticket {
	if q.IsEmpty() {
		fmt.Println("The queue is empty")
		return Ticket{}
	}

	return q.items[0]
}

func main() {
	file := "./file.jsonl"
	tickets := unmaeshallJson(file)

	q := Queue{}

	for _, ticket := range tickets {
		q.Push(ticket)
		fmt.Printf("Adding support ticket for %s: %s\n", ticket.Name, ticket.Issue)
	}

	for !q.IsEmpty() {
		ticket := q.Pop()
		fmt.Printf("Processing support ticket for %s: %s\n", ticket.Name, ticket.Issue)
	}

}
