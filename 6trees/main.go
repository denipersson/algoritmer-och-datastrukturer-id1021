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

func construct_random_tree(size int) (Binary_Tree, []int) {
	tree := Binary_Tree{&Node{rand.Intn(size), rand.Intn(100), nil, nil}}

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
func test_print() {

	tree, keys := construct_random_tree(10)

	keys = keys

	/*
		tree := Binary_Tree{}
		tree.root = &Node{5, rand.Intn(10), nil, nil}
		tree.root.add(1, rand.Intn(10))
		tree.root.add(7, rand.Intn(10))
		tree.root.add(8, rand.Intn(10))

			fmt.Println(tree.root.key)
			fmt.Println(tree.root.left.key)
			fmt.Println(tree.root.right.key)
			fmt.Println(tree.root.right.right.key)
	*/

	fmt.Println("IMP:")
	tree.root.print_implicit_stack()

	fmt.Println("EXP:")
	tree.root.print_explicit_stack()
}
