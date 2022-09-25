package main

type Stack struct {
	head *Node
}

func (s *Stack) Push(n Node) {
	n.next = s.head
	s.head = &n
}

func (s *Stack) Pop() (n Node) {
	n = *s.head
	s.head = s.head.next
	n.next = nil
	return
}
