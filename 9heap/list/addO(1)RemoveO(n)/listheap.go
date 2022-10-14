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
	old_head := hl.head
	hl.head = &new_node
	new_node.next = old_head
	old_head.prev = &new_node
}
func (hl *Heap_List) remove() {
	lowest := hl.head
	curr := hl.head.next
	if curr == nil {
		return
	}
	for curr.next != nil {

		if curr.val < lowest.val {
			lowest = curr
		}
		curr = curr.next
	}
	if lowest == hl.head {
		hl.head = lowest.next
	}
	if lowest.next != nil {
		lowest.next.prev = lowest.prev
	}
	if lowest.prev != nil {
		lowest.prev.next = lowest.next
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
