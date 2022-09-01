package main

import "fmt"

func main() {

}

func test_static_calc() {
	item1 := New_Item(VALUE, 5)
	item2 := New_Item(VALUE, 10)
	item3 := New_Item(DIV, 0)
	item4 := New_Item(VALUE, 4)
	item5 := New_Item(ADD, 5)

	items := []Item{item1, item2, item3, item4, item5}

	stack := New_Stack(4, false)

	calc := New_Calculator(items, &stack)

	fmt.Printf("%d", calc.run())
}
