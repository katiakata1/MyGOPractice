// First In - First out

package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("The queue is empty")
	}

	first := q.items[0]
	q.items = q.items[1:] // removing first item
	return first
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		panic("The queue is empty")
	}

	return q.items[0]
}

func main() {
	q := &Queue{}

	q.Enqueue(45)
	q.Enqueue(3)
	q.Enqueue(1)

	fmt.Println("The queue a the start: ", q.items)

	first := q.Peek()
	fmt.Println("The first item in the queue: ", first)

	dequed := q.Dequeue()
	fmt.Println("The item dequied ", dequed)
	fmt.Println("The queue now: ", q.items)

}
