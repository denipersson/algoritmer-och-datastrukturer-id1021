package main

type TreeIterator struct {
	stack *Stack
}

func (t *TreeIterator) hasNext() bool {
	if len(t.stack.stack) > 0 {
		return true
	}
	return false

}
func (t *TreeIterator) getNext() *Node {
	if t.hasNext() {
		return t.stack.Pop()
	}
	return nil
}
