package main

type Queue_Node struct {
	referred_node *Tree_Node
	next          *Queue_Node
}

type Queue struct {
	head *Queue_Node
	tail *Queue_Node
}

func (q *Queue) add(new_node *Queue_Node) {
	if q.head == nil {
		q.head = new_node
		q.tail = q.head
	} else {
		if q.tail == nil {
			q.tail = q.head
		}
		last := q.tail
		last.next = new_node
		q.tail = new_node
	}
}

func (q *Queue) remove() *Queue_Node {
	old_head := q.head
	q.head = old_head.next
	return old_head
}
