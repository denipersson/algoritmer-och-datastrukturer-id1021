package main

type Node struct {
	next []*Node
	word bool
}

func new_node() *Node {
	n := Node{make([]*Node, 27), false}
	return &n
}
