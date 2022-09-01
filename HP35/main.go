package main

import (
	"fmt"
)

func main() {
	test_static_calc()
}

func test_static_calc() {

	items := []Item{New_Item(VALUE, 10), New_Item(VALUE, 5), New_Item(SUB, 0),
		New_Item(VALUE, 4), New_Item(ADD, 0), New_Item(VALUE, 3), New_Item(DIV, 0)}
	//expected result: (10 - 5 + 4)/3  = (5+4)/3 = 9/3 = 3

	stack := New_Stack(4, false)

	calc := New_Calculator(items, &stack)

	fmt.Println("", calc.run())
}
