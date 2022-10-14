package main

import "fmt"

type Node struct {
	priority int
	right    *Node
	left     *Node
	size     int
}

type Heap struct {
	root *Node
}

func (heap *Heap) add(priority int) int {
	new_node := Node{priority, nil, nil, 0}
	if heap.root == nil {
		heap.root = &new_node
		return 1
	}
	if heap.root.priority > priority {
		temp := heap.root.priority
		heap.root.priority = priority
		new_node.priority = temp
	}
	steps := 0
	heap.root.add_down(&new_node, &steps)
	return steps
	//fmt.Println(new_node.priority, " added in steps: ", steps)
}
func (node *Node) add_down(new_node *Node, steps *int) *int {
	node.size += new_node.size + 1
	*steps += 1
	if node.left == nil {
		node.left = new_node
		return steps
	} else if node.right == nil {
		node.right = new_node
		return steps
	}
	var target *Node
	if node.left.size <= node.right.size { // go down the branch with the smallest size
		target = node.left
	} else {
		target = node.right
	}
	if target.priority > new_node.priority {
		temp := target.priority
		target.priority = new_node.priority
		new_node.priority = temp
	}
	target.add_down(new_node, steps)
	return steps
}
func (heap *Heap) push(increment int) int {
	heap.root.priority += increment
	steps := 0
	var target *Node
	if heap.root.left == nil {
		target = heap.root.right
	} else if heap.root.right == nil {
		target = heap.root.left
	} else {
		if heap.root.left.priority <= heap.root.right.priority {
			target = heap.root.left
		} else {
			target = heap.root.right
		}
	}
	if target == nil {
		steps++
		return steps
	}
	if target.priority < heap.root.priority {
		temp := target.priority
		target.priority = heap.root.priority
		heap.root.priority = temp
	}
	target.push_down(&steps)
	return steps
}
func (node *Node) push_down(steps *int) *int {
	*steps++
	var target *Node
	if node.left == nil {
		target = node.right
	} else if node.right == nil {
		target = node.left
	} else {
		if node.left.priority <= node.right.priority {
			target = node.left
		} else {
			target = node.right
		}
	}
	if target == nil {
		return steps
	}
	if target.priority < node.priority {
		temp := target.priority
		target.priority = node.priority
		node.priority = temp
	}
	return target.push_down(steps)
}

func (heap *Heap) remove() *Node {
	return_node := heap.root
	if return_node == nil {
		return nil
	}
	var branch_to_fix *Node
	if heap.root.left == nil {
		heap.root.right = heap.root.left
		heap.root.right = nil
		return return_node
	} else if heap.root.right == nil {
		heap.root = heap.root.left
		return return_node
	}
	var new_root *Node
	if heap.root.left.priority < heap.root.right.priority {
		new_root = heap.root.left
		branch_to_fix = heap.root.right
	} else {
		new_root = heap.root.right
		branch_to_fix = heap.root.left
	}
	heap.root = new_root
	steps := 0
	heap.root.add_down(branch_to_fix, &steps)
	return return_node
}

func (n *Node) print_implicit_stack() {
	left := n.left
	right := n.right
	if left != nil {
		fmt.Println("L")
		left.print_implicit_stack()
	}
	fmt.Println(" priority: ", n.priority)
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
			fmt.Println("P: ", cur.priority, " [", cur.size, "]")
			cur = cur.right
		}
	}
}
