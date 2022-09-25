package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//test_linked_list_append()
	//test_slice_append()
	test_double_linked(10, 1)
	for i := 0; i <= 100000; i += 1000 {
		test_linked(10*i, 1*i)
	}

	/*
		for i := 0; i <= 10; i += 1 {
			test_linked(1000*i, 100*i)
		}
	*/
}
func test_linked(size int, ops int) {
	dll := Linked_List{}
	dll.head = &Node{0, nil}

	nodes := make([]Node, size)

	for i := 0; i < size; i++ {
		nodes[i] = Node{rand.Intn(size), nil}
	}

	for i := 0; i < len(nodes); i++ {
		val := nodes[i]
		n := val
		dll.add_to_front(&n)
	}

	removed := make([]Node, size)
	t0 := time.Now()
	for i := 0; i < ops; i++ {
		val := nodes[rand.Intn(size)]
		dll.remove(&val)
		removed = append(removed, val)
	}

	for i := 0; i < ops; i++ {
		dll.add_to_front(&removed[i])
	}
	t_since := time.Since(t0)

	fmt.Println(ops, " ", t_since.Nanoseconds())
}
func test_double_linked(size int, ops int) {
	dll := Double_Linked_List{}
	dll.head = &Double_Node{0, nil, nil}

	nodes := make([]Double_Node, size)

	for i := 0; i < size; i++ {
		nodes[i] = Double_Node{rand.Intn(size), nil, nil}
	}

	for i := 0; i < len(nodes); i++ {
		val := nodes[i]
		n := val
		dll.add_to_front(&n)
	}

	removed := make([]Double_Node, size)
	t0 := time.Now()
	for i := 0; i < ops; i++ {
		val := nodes[rand.Intn(size)]
		dll.remove(&val)
		removed = append(removed, val)
	}

	for i := 0; i < ops; i++ {
		dll.add_to_front(&removed[i])
	}
	t_since := time.Since(t0)

	fmt.Println(ops, " ", t_since.Nanoseconds())
}

func test_linked_list_append() {
	for i := 1; i < 100; i++ {
		size := 10000 * i
		list1 := generate_random_linked_list(1000, false)

		list2 := generate_random_linked_list(size, true)
		list1.append(list2)
	}
}
func test_slice_append() {
	for i := 1; i <= 1000; i++ {
		size := 1000 * i
		list1 := generate_random_slice(1000, false)

		list2 := generate_random_slice(size, true)

		list1 = append_slice(list1, list2)
	}
}

func generate_random_linked_list(size int, print_time bool) Linked_List {
	head := Node{rand.Intn(size), nil}
	list := Linked_List{&head}
	t0 := time.Now()
	for i := 1; i < size; i++ {
		list2 := Linked_List{&Node{rand.Intn(size), nil}}
		list.append(list2)
	}
	t_since := time.Since(t0).Nanoseconds()

	if print_time {
		fmt.Println(size, " ", t_since)
	}
	return list
}

func print_linked_list(list Linked_List) {
	cur := list.head

	for cur != nil {
		fmt.Println(cur.item)

		cur = cur.next
	}
}

func generate_random_slice(size int, printTime bool) []int {
	arr := make([]int, size)
	vals := make([]int, size)
	for i := 0; i < size; i++ {
		vals[i] = rand.Intn(size)
	}
	t0 := time.Now()
	for i := 0; i < size; i++ {
		arr[i] = vals[i]
	}
	t_since := time.Since(t0).Nanoseconds()
	if printTime {
		fmt.Println(size, " ", t_since)

	}

	return arr
}
func append_slice(slice_a []int, slice_b []int) []int {
	new_slice := make([]int, len(slice_a)+len(slice_b))

	for i := 0; i < len(slice_a); i++ {
		new_slice[i] = slice_a[i]
	}

	for i := len(slice_a); i < len(slice_b)+len(slice_a); i++ {
		new_slice[i] = slice_b[i-len(slice_a)]
	}

	return new_slice
}
