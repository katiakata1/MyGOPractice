package main

type Queue struct {
	items []string
}

func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() string {
	if len(q.items) == 0 {
		return ""
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func (q *Queue) IsEmpty() bool {
	if len(q.items) == 0 {
		return true
	}

	return false
}

// func main() {
// 	queue := Queue{}

// 	queue.Enqueue("A")
// 	queue.Enqueue("B")
// 	queue.Enqueue("C")

// 	fmt.Println(queue.Dequeue())
// 	fmt.Println(queue.IsEmpty())

// 	fmt.Println(queue.Dequeue())
// 	fmt.Println(queue.Dequeue())

// 	fmt.Println(queue.IsEmpty())
// }
