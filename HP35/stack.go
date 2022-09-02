package main

import (
	"fmt"
)

type Stack struct {
	pointer int
	stack   []int
	size    int
	dynamic bool
}

func New_Stack(size int, dynamic bool) Stack {
	return Stack{0, make([]int, size), size, dynamic}
}

func (s *Stack) Push(value int) {

	if s.pointer >= len(s.stack) {
		if s.dynamic {
			s.Expand_Stack()
		} else {
			fmt.Println("Stack full(overflow)")
			return
		}
	}

	s.stack[s.pointer] = value
	s.pointer++

}
func (s *Stack) Pop() int {

	value := 0

	if s.pointer == 0 {
		fmt.Println("Stack is empty.")
		return 0
	}

	if s.dynamic {
		if s.pointer < s.size-500 {
			s.Shrink_Stack()
		}
	}
	s.pointer--
	value = s.stack[s.pointer]

	return value

}
func (s *Stack) Shrink_Stack() {
	old_stack := s.stack
	s.stack = make([]int, len(s.stack)-500)

	for i := 0; i < s.pointer; i++ {
		s.stack[i] = old_stack[i]
	}

}

func (s *Stack) Expand_Stack() {
	old_stack := s.stack
	s.stack = make([]int, len(s.stack)+500)

	for i := 0; i < s.pointer; i++ {
		s.stack[i] = old_stack[i]
	}

}
