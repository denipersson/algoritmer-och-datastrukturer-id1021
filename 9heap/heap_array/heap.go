package main

import "fmt"

type Heap struct {
	heap []*Node
}
type Node struct {
	priority int
}

func new_heap() *Heap {

	heap := Heap{make([]*Node, 0)}
	return &heap
}
func (h *Heap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}
func (h *Heap) peek() *int {
	if h.heap[0] == nil {
		return nil
	}
	return &h.heap[0].priority
}
func (h *Heap) size() int {
	return len(h.heap)
}
func (h *Heap) add(n *Node) {
	h.heap = append(h.heap, n)
	h.heapify_up((len(h.heap) - 1))
}
func (h *Heap) heapify_up(i int) {
	for h.heap[get_parent_index(i)].priority < h.heap[i].priority {
		h.swap(i, get_parent_index(i))
		i = get_parent_index(i)
	}
}
func (h *Heap) remove() *Node {
	if h.size() == 0 {
		return nil
	}
	root := h.get_root()
	h.swap(0, h.size()-1)
	h.heap = h.heap[:h.size()-1]
	h.heapify_down(0)
	return root
}
func (h *Heap) heapify_down(i int) {
	l, r := get_left_index(i), get_right_index(i)
	largest := i
	if l < h.size() && h.heap[i].priority < h.heap[l].priority {
		largest = l
	}
	if r < h.size() && h.heap[largest].priority < h.heap[r].priority {
		largest = r
	}
	if largest != i {
		h.swap(i, largest)
		h.heapify_down(largest)
	}
}
func (heap *Heap) get_root() *Node {
	return heap.heap[0]
}
func get_left_index(index int) int {
	left := index*2 + 1
	return left
}
func get_right_index(index int) int {
	right := index*2 + 2
	return right
}
func get_parent_index(i int) int {
	return (i - 1) / 2
}
func (heap *Heap) print() {
	for i := 0; i < len(heap.heap); i++ {
		if heap.heap[i] == nil {
			break
		}
		fmt.Println(heap.heap[i].priority)
	}
}

/*
func (heap *Heap) push(increment int) int {
	heap.root.priority += increment
	steps := 0
	return steps
}

func (heap *Heap) remove() *Node {
	return_node := heap.root
}
*/
