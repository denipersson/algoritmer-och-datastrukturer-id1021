package main

import "fmt"

type Heap_List struct {
	head *Node
}
type Node struct {
	val  int
	next *Node
	prev *Node
}

func new_heap_list(first_val int) Heap_List {
	n := Node{first_val, nil, nil}
	return Heap_List{&n}

}

func (hl *Heap_List) add(val int) {
	new_node := Node{val, nil, nil}
	curr := hl.head
	found := false
	if curr == nil {
		hl.head = &new_node
		return
	}
	for curr.next != nil {
		if new_node.val < curr.val {
			found = true
			break
		}
		curr = curr.next
	}
	if found {
		if curr == hl.head {
			hl.head = &new_node
		}
		old_prev := curr.prev

		if old_prev != nil {
			old_prev.next = &new_node
		}
		curr.prev = &new_node
		new_node.prev = old_prev
		new_node.next = curr
	} else {
		curr.next = &new_node
		new_node.prev = curr
	}
}
func (hl *Heap_List) remove() {
	if hl.head != nil {
		hl.head = hl.head.next
	}
}
func (hl *Heap_List) print_list() {
	curr := hl.head
	for curr != nil {
		fmt.Print(curr.val, " ")

		curr = curr.next
	}
	fmt.Println("")
}
