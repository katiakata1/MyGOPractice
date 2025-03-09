package main

type Node struct {
	value string
	next  *Node
}

type QueueList struct {
	front *Node
	rear  *Node
}

func (q *QueueList) Enqueue(value string) {
	newNode := &Node{value: value}
	if q.rear != nil {
		q.rear.next = newNode
	}
	q.rear = newNode

	if q.front == nil {
		q.front = newNode
	}
}

func (q *QueueList) Dequeue() string {
	if q.front == nil {
		return ""
	}

	value := q.front.value
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return value

}

func (q *QueueList) IsEmtpy() bool {
	return q.front == nil
}

// func main() {
// 	queue := &QueueList{}

// 	queue.Enqueue("A")
// 	queue.Enqueue("B")
// 	queue.Enqueue("C")

// 	fmt.Println(queue.Dequeue())
// 	fmt.Println(queue.IsEmtpy())

// 	fmt.Println(queue.Dequeue())
// 	fmt.Println(queue.Dequeue())
// 	fmt.Println(queue.IsEmtpy())

// }
