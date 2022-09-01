package main

type Calculator struct {
	expr        []Item
	exprPointer int
	stack       Stack
}

func New_Calculator(expr []Item, stack *Stack) *Calculator {
	calc := new(Calculator)
	calc.exprPointer = 0
	calc.expr = expr
	calc.stack = *stack
	return calc
}

func (calc *Calculator) step() {
	item := calc.expr[calc.exprPointer]

	switch item.item_type {

	case ADD:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(x + y)

	case SUB:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(y - x)

	case MUL:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(x * y)

	case DIV:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(y / x)

	case VALUE:
		calc.stack.push(item.value)
	}

	calc.exprPointer++

}

func (c *Calculator) run() int {
	for c.exprPointer < len(c.expr) {
		c.step()
	}
	return c.stack.pop()
}
