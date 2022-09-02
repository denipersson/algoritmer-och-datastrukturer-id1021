package main

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
		calc.stack.Push(y / x)

	case VALUE:
		calc.stack.Push(item.value)
	}

	calc.expr_pointer++

}

func (c *Calculator) Run() int {
	for c.expr_pointer < len(c.expr) {
		c.Step()
	}
	return c.stack.Pop()
}
