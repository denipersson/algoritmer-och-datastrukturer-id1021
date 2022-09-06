package main

import (
	"fmt"
	"time"
)

type Calculator struct {
	expr         []Item
	expr_pointer int
	stack        Stack
}

func New_Calculator(expr []Item, stack *Stack) *Calculator {
	calc := new(Calculator)
	calc.expr_pointer = 0
	calc.expr = expr
	calc.stack = *stack
	return calc
}

func (calc *Calculator) Step() {
	item := calc.expr[calc.expr_pointer]

	switch item.item_type {

	case ADD:
		x := calc.stack.Pop()
		y := calc.stack.Pop()
		calc.stack.Push(x + y)

	case SUB:
		x := calc.stack.Pop()
		y := calc.stack.Pop()
		calc.stack.Push(y - x)

	case MUL:
		x := calc.stack.Pop()
		y := calc.stack.Pop()
		calc.stack.Push(x * y)

	case DIV:
		x := calc.stack.Pop()
		y := calc.stack.Pop()
		if x == 0 {
			x++
		}
		calc.stack.Push(y / x)

	case VALUE:
		calc.stack.Push(item.value)
	case MOD:
		x := calc.stack.Pop()
		y := calc.stack.Pop()
		calc.stack.Push(x % y)
	}

	calc.expr_pointer++

}

func (c *Calculator) Run() int {
	t0 := time.Now()
	for c.expr_pointer < len(c.expr) {
		c.Step()
	}
	t_after := time.Since(t0).Microseconds()
	if c.stack.dynamic {
		fmt.Println("D ", t_after)
	} else {
		fmt.Println("S ", t_after)
	}

	return c.stack.Pop()
}
