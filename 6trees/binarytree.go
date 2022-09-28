package main

import "fmt"

type Node struct {
	key   int
	value int

	right *Node
	left  *Node
}

type Binary_Tree struct {
	root *Node
}

func (n *Node) add(key int, value int) {
	if key < n.key {
		if n.left == nil {
			n.left = &Node{key, value, nil, nil}
			return
		}
		n.left.add(key, value)
	} else if key > n.key {
		if n.right == nil {
			n.right = &Node{key, value, nil, nil}
			return
		}
		n.right.add(key, value)
	} else {
		n.value = value
	}
}
func (n *Node) lookup(key int) *int {
	if key < n.key {
		if n.left == nil {
			return nil
		}
		return n.left.lookup(key)
	} else if key > n.key {
		if n.right == nil {
			return nil
		}
		return n.right.lookup(key)
	} else {
		return &n.value
	}
}

func (n *Node) print_implicit_stack() {
	left := n.left
	right := n.right
	if left != nil {
		fmt.Println("L")
		left.print_implicit_stack()
	}
	fmt.Println(" key: ", n.key, "value: ", n.value)
	if right != nil {
		fmt.Println("R")
		right.print_implicit_stack()
	}

}

func (n *Node) print_explicit_stack() {
	iter := TreeIterator{&Stack{0, make([]Node, 10), 10, true}}
	cur := n
	for iter.hasNext() {
		if cur != nil {
			iter.stack.Push(cur)
			cur = cur.left
		} else {
			cur = iter.getNext()
			if cur == nil {
				break
			}
			fmt.Println("K: ", cur.key, " V: ", cur.value)
			cur = cur.right
		}
	}
}
