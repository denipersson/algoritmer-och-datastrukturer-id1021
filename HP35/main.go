package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		test_static_calc()
		test_static_calc_overflow()
		test_dynamic_calc_overflow()
	*/
	rand.Seed(time.Now().UTC().UnixNano())
	/*
		static_large_test()
		println("______________________")
		for i := 0; i < 5; i++ {
			static_large_test()
		}
		for i := 0; i < 5; i++ {
			dynamic_large_test()
		}
	*/
	personal_number_last_digit()

}

func personal_number_last_digit() {
	items := []Item{
		New_Item(VALUE, 9), New_Item(VALUE, 2), New_Item(MUL, 0), New_Item(ADD, 0),
		New_Item(VALUE, 6), New_Item(VALUE, 1), New_Item(MUL, 0), New_Item(ADD, 0),
		New_Item(VALUE, 0), New_Item(VALUE, 2), New_Item(MUL, 0), New_Item(ADD, 0),
		New_Item(VALUE, 9), New_Item(VALUE, 1), New_Item(MUL, 0), New_Item(ADD, 0),
		New_Item(VALUE, 2), New_Item(VALUE, 2), New_Item(MUL, 0), New_Item(ADD, 0),
		New_Item(VALUE, 9), New_Item(VALUE, 1), New_Item(MUL, 0), New_Item(ADD, 0),
		New_Item(VALUE, 5), New_Item(VALUE, 2), New_Item(MUL, 0), New_Item(ADD, 0),
		New_Item(VALUE, 5), New_Item(VALUE, 1), New_Item(MUL, 0), New_Item(ADD, 0),
		New_Item(VALUE, 1), New_Item(VALUE, 2), New_Item(MUL, 0), New_Item(ADD, 0),
	}

	stack := New_Stack(4, true)
	calc := New_Calculator(items, &stack)
	fmt.Println("", calc.Run())

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
	//fmt.Println("Static test of 999999 items: ")
	items := make([]Item, 999999)

	for i := 0; i < 999999; i++ {
		var t Item_type
		if i < 500000 {
			t = VALUE
		} else {
			t = Item_type(rand.Intn(4))
		}
		items[i] = New_Item(t, rand.Intn(10)+1)
	}
	stack := New_Stack(999999, false)

	calc := New_Calculator(items, &stack)
	calc.Run()
	//fmt.Println("Result: ", calc.Run())

}
func dynamic_large_test() {
	//fmt.Println("Dynamic test of 999999 items: ")
	items := make([]Item, 999999)

	for i := 0; i < 999999; i++ {
		var t Item_type
		if i < 500000 {
			t = VALUE
		} else {
			t = Item_type(rand.Intn(4))
		}
		items[i] = New_Item(t, rand.Intn(10)+1)
	}
	stack := New_Stack(4, true)

	calc := New_Calculator(items, &stack)
	calc.Run()
	//fmt.Println("Result: ", calc.Run())

}
