package main

type Tree_Node struct {
	key   int
	value int

	right *Tree_Node
	left  *Tree_Node
}

type Binary_Tree struct {
	root *Tree_Node
}

func (n *Tree_Node) add(key int, value int) {
	if key < n.key {
		if n.left == nil {
			n.left = &Tree_Node{key, value, nil, nil}
			return
		}
		n.left.add(key, value)
	} else if key > n.key {
		if n.right == nil {
			n.right = &Tree_Node{key, value, nil, nil}
			return
		}
		n.right.add(key, value)
	} else {
		n.value = value
	}
}
func (n *Tree_Node) lookup(key int) *int {
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
