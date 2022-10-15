package main

import (
	"fmt"
	"math/rand"
	"time"
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

	heap := new_heap()
	node := Node{rand.Intn(20)}

	for i := 0; i < 63; i++ {
		node = Node{rand.Intn(20)}
		heap.add(&node)
	}

	for i := 0; i < 1000000; i += 1000 {
		t0 := time.Now()
		node = Node{rand.Intn(20)}
		t_since := time.Since(t0)
		fmt.Println(64+i, t_since.Nanoseconds())
		heap.add(&node)
	}

}
