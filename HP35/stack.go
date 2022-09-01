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

func (s *Stack) push(value int) {

	if s.pointer < len(s.stack)-1 {
		s.pointer++
		s.stack[s.pointer] = value
	} else {
		fmt.Println("Stack full(overflow).")
	}
}
func (s *Stack) pop() int {

	value := 0

	if s.pointer < 0 {
		fmt.Println("Stack is empty.")
	} else {
		value = s.stack[s.pointer]
		s.pointer--
	}

	return value

}
