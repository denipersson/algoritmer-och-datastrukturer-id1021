package main

type Tree_Iterator struct {
	queue *Queue
}

func (t *Tree_Iterator) hasNext() bool {
	if t.queue.head != nil {
		return true
	}
	return false

}
func (t *Tree_Iterator) next() *Queue_Node {
	if t.hasNext() {
		next := t.queue.head
		if next.referred_node.left != nil {
			t.queue.add(&Queue_Node{next.referred_node.left, nil})
		}
		if next.referred_node.right != nil {
			t.queue.add(&Queue_Node{next.referred_node.right, nil})
		}
		return t.queue.remove()
	}
	return nil
}
