package main

/*
type Tree_Iterator_Arr struct {
	queue *Array_Queue
}

func (t *Tree_Iterator_Arr) hasNext() bool {
	if t.queue.first != t.queue.last {
		return true
	}
	return false

}
func (t *Tree_Iterator_Arr) next() *Queue_Node {
	if t.hasNext() {
		next := t.queue.queue[t.queue.first]
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
*/
