package main

type Node struct {
	item int
	next *Node
}

func new_node(item int, next *Node) Node {
	new := Node{item, next}
	return new

}

type Double_Node struct {
	item     int
	next     *Double_Node
	previous *Double_Node
}
