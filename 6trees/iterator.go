package main

type Iterator interface {
	hasNext() bool
	getNext() *Node
}
