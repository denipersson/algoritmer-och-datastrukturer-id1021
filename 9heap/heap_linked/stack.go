package main

import (
	"fmt"
)

type Stack struct {
	pointer int
	stack   []Node
	size    int
	dynamic bool
}

func New_Stack(size int, dynamic bool) Stack {
	return Stack{0, make([]Node, size), size, dynamic}
}

func (s *Stack) Push(node *Node) {

	if s.pointer >= len(s.stack) {
		if s.dynamic {
			s.Expand_Stack()
		} else {
			fmt.Println("Stack full(overflow)")
			return
		}
	}

	s.stack[s.pointer] = *node
	s.pointer++

}
func (s *Stack) Pop() *Node {

	if s.pointer == 0 {
		fmt.Println("Stack is empty.")
		return nil
	}

	if s.dynamic {
		if s.pointer < s.size-500 {
			s.Shrink_Stack()
		}
	}
	s.pointer--
	value := s.stack[s.pointer]

	return &value

}
func (s *Stack) Shrink_Stack() {
	old_stack := s.stack
	s.stack = make([]Node, len(s.stack)-500)

	for i := 0; i < s.pointer; i++ {
		s.stack[i] = old_stack[i]
	}

}

func (s *Stack) Expand_Stack() {
	old_stack := s.stack
	s.stack = make([]Node, len(s.stack)*2)

	for i := 0; i < s.pointer; i++ {
		s.stack[i] = old_stack[i]
	}

}
