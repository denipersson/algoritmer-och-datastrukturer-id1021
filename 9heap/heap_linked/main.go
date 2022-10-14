package main

import (
	"fmt"
	"math/rand"
)

func main() {

	test_print()
	/*
	   test_look_up(2)

	   	for i := 1; i <= 10000; i += 100 {
	   		test_look_up(i * 10)
	   	}
	*/
}

func test_print() {

	heap := Heap{&Node{5, nil, nil, 0}}

	for i := 0; i < 63; i++ {
		heap.add(i + rand.Intn(20) + 1)
	}

	for i := 0; i < 1000; i++ {
		incr := rand.Intn(10) + 10
		fmt.Println(64+i, " ", heap.add(incr))
	}
	heap.root.print_explicit_stack()
	/*
	   heap.add(3)
	   heap.root.print_explicit_stack()
	   //fmt.Println("")
	   //fmt.Println(heap.root.priority)
	   //fmt.Println(heap.root.left.priority)
	   //fmt.Println(heap.root.right.priority)
	   //fmt.Println(heap.root.left.left.priority)
	   heap.add(10)
	   heap.root.print_explicit_stack()
	   fmt.Println("Removed:", heap.remove().priority)
	   heap.root.print_explicit_stack()
	   heap.add(7)
	   heap.root.print_explicit_stack()
	   //fmt.Println("Removed: ", heap.remove().priority)
	*/
}
