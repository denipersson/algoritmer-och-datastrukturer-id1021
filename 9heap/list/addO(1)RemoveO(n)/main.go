package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	max := 100000
	for i := 1; i <= max; i += max/100 + 1 {
		test_heap_list(i)
	}

	//test_heap_list(10)

}

func test_heap_list(size int) {
	hl := create_random_heap_list(size)
	//hl.print_list()

	t0 := time.Now()
	hl.remove()
	//hl.add(size / 2)
	t_since := time.Since(t0).Nanoseconds()
	fmt.Println(size, " ", t_since)
	//hl.print_list()
}

func create_random_heap_list(size int) Heap_List {
	hl := new_heap_list(rand.Intn(size))

	for i := 1; i < size; i++ {
		hl.add(rand.Intn(size))
	}
	return hl
}
