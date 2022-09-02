package main

import (
	"fmt"
)

func main() {
	test_static_calc()
	test_static_calc_overflow()
	test_dynamic_calc_overflow()
}

func test_static_calc() {

	fmt.Println("Test Static Calculator, expected result: 3")
	items := []Item{New_Item(VALUE, 10), New_Item(VALUE, 5), New_Item(DIV, 0)}
	//Expected result: (10 - 5 + 4)/3  = (5+4)/3 = 9/3 = 3

	stack := New_Stack(4, false)
	calc := New_Calculator(items, &stack)

	fmt.Println("", calc.Run())
	fmt.Println("")
}
func test_static_calc_overflow() {

	fmt.Println("Test Static Calculator, expected result: stack overflow")
	items := []Item{New_Item(VALUE, 10), New_Item(VALUE, 10), New_Item(VALUE, 10), New_Item(VALUE, 10), New_Item(VALUE, 10), New_Item(ADD, 0), New_Item(ADD, 0), New_Item(ADD, 0), New_Item(ADD, 0)}
	stack := New_Stack(4, false)

	calc := New_Calculator(items, &stack)
	fmt.Println("", calc.Run())
	fmt.Println("")
}

func test_dynamic_calc_overflow() {
	fmt.Println("Test Dynamic Calculator, expected result: 50")
	items := []Item{New_Item(VALUE, 10), New_Item(VALUE, 10), New_Item(VALUE, 10), New_Item(VALUE, 10), New_Item(VALUE, 10), New_Item(ADD, 0), New_Item(ADD, 0), New_Item(ADD, 0), New_Item(ADD, 0)}
	stack := New_Stack(4, true)

	calc := New_Calculator(items, &stack)
	fmt.Println("", calc.Run())
	fmt.Println("")
}

func static_large_test() {

}
