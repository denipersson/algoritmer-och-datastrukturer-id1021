package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	prev *Node
	next *Node
}
type LinkedList struct {
	start *Node
	end   *Node
}

func (dll *LinkedList) quick_sort_LL(start *Node, end *Node) {
	if start != end && start != nil &&
		end != nil && end.next != start {
		node := dll.partition_LL(start, end)
		if node != nil {
			dll.quick_sort_LL(node.next, end)
			dll.quick_sort_LL(start, node.prev)
		}
	}

}

func (dll *LinkedList) partition_LL(start *Node, end *Node) *Node {

	current := start
	search := start.prev
	temp := 0

	for current != nil && current != end {
		if current.data <= end.data {
			if search == nil {
				search = start
			} else {
				search = search.next
			}
			temp = current.data
			current.data = search.data
			search.data = temp
		}
		current = current.next
	}
	if search != nil {
		search = search.next
	}
	if search == nil {
		search = start
	}
	temp = end.data
	end.data = search.data
	search.data = temp
	return search
}

func (dll *LinkedList) swap_nodes(a *Node, b *Node) {

	if a == dll.start {
		dll.start = b
	} else if b == dll.start {
		dll.start = a
	}
	if a == dll.end {
		dll.end = b
	} else if b == dll.end {
		dll.end = a
	}
	temp := a.next
	a.next = b.next
	b.next = temp

	if a.next != nil {
		a.next.prev = a
	}
	if b.next != nil {
		b.next.prev = b
	}

	temp = a.prev
	a.prev = b.prev
	b.prev = temp

	if a.prev != nil {
		a.prev.next = a

	}
	if b.prev != nil {
		b.prev.next = b
	}
	/*

		temp := a.data
		a.data = b.data
		b.data = temp*/

}

func (dll *LinkedList) String() string {
	sb := strings.Builder{}

	for iter := dll.start; iter != nil; iter = iter.next {
		sb.WriteString(fmt.Sprintf("%d ", iter.data))
	}
	sb.WriteString(fmt.Sprintf("\n"))
	return sb.String()
}
func (dll *LinkedList) append(n *Node) {
	dll.end.next = n
	n.prev = dll.end
	dll.end = n
}
func (dll *LinkedList) remove(n *Node) {
	prev := n.prev
	next := n.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	n.next = nil
	n.prev = nil
}

func new_doubly_linked_list(value int) LinkedList {
	first_node := Node{value, nil, nil}
	d := LinkedList{&first_node, &first_node}
	return d
}
