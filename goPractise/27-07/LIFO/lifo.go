// Last In - First Out

package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}

	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

// Returns what's the last item in the array
func (s *Stack) Peep() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	return s.items[len(s.items)-1]
}

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(2)
	stack.Push(48)

	fmt.Println("The stack after pushing is:", stack.items)

	top := stack.Peep()

	fmt.Println("Top of stack (peek): ", top)

	popped := stack.Pop()

	fmt.Println("Popped:", popped)

	fmt.Println("The stack after pushing is:", stack.items)

}
