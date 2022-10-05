package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//benchmark_breadth_first()
	//test_breadth_first()

	test_array_queue()
}
func benchmark_breadth_first() {
	for i := 1; i <= 1000; i++ {
		size := 10 * i
		q := generate_random_queue(size)
		iter := Tree_Iterator{&q}
		t0 := time.Now()
		for iter.hasNext() {
			iter.next()
		}
		t_since := time.Since(t0).Nanoseconds()
		fmt.Println(size, " ", t_since)
	}

}
func test_breadth_first() {
	tree := Binary_Tree{&Tree_Node{5, rand.Intn(100), nil, nil}}
	tree.root.add(3, rand.Intn(100))
	tree.root.add(4, rand.Intn(100))
	tree.root.add(7, rand.Intn(100))
	tree.root.add(2, rand.Intn(100))
	tree.root.add(8, rand.Intn(100))
	tree.root.add(6, rand.Intn(100))
	q := Queue{&Queue_Node{tree.root, nil}, nil}

	iter := Tree_Iterator{&q}

	for iter.hasNext() {
		n := iter.next()
		println("K: ", n.referred_node.key, " V:", n.referred_node.value)
	}
}
func test_array_queue() {

	q := new_queue(7)

	q.add(&Arr_Node{3, 5})
	q.add(&Arr_Node{6, 3})
	q.add(&Arr_Node{4, 2})
	q.add(&Arr_Node{5, 1})
	q.add(&Arr_Node{7, 8})
	fmt.Println("ADDED 5 NODES")
	q.print()

	q.remove()
	q.remove()
	fmt.Println("REMOVED 2 NODES")
	q.print()
	q.add(&Arr_Node{1, 5})
	q.add(&Arr_Node{18, 7})
	fmt.Println("ADDED 2 NODES")
	q.print()
	q.add(&Arr_Node{10, 0})
	q.add(&Arr_Node{11, 2})
	fmt.Println("ADDED 2 NODES")
	q.print()
}

func generate_random_queue(size int) Queue {

	tree, keys := construct_random_tree(size)

	q := Queue{&Queue_Node{tree.root, nil}, nil}

	for i := 0; i < size; i++ {
		new_queue_node := Queue_Node{&Tree_Node{keys[i], rand.Intn(100), nil, nil}, nil}
		q.add(&new_queue_node)
	}

	return q
}

func construct_random_tree(size int) (Binary_Tree, []int) {
	tree := Binary_Tree{&Tree_Node{rand.Intn(size), rand.Intn(100), nil, nil}}

	keys := make([]int, size)

	for i := 0; i < size; i++ {
		keys[i] = rand.Intn(size)
	}

	for i := 0; i < size; i++ {
		tree.root.add(keys[i], rand.Intn(100))
	}

	return tree, keys
}

func test_look_up(size int) {
	tree, keys := construct_random_tree(size)

	precision := len(keys)

	var t_since float64
	t0 := time.Now()
	for i := 0; i < precision; i++ {
		tree.root.lookup(keys[i])
		t_since += float64(time.Since(t0).Microseconds())
	}
	t_since = t_since / float64(precision)

	fmt.Println(size, " ", t_since)
}
