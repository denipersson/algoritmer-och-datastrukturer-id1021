package main

import "fmt"

type Array_Queue struct {
	first int
	last  int
	size  int
	queue []*Arr_Node
}

type Arr_Node struct {
	key   int
	value int
}

func new_queue(size int) Array_Queue {
	a := Array_Queue{0, 0, size, make([]*Arr_Node, size)}
	return a
}

func (arrq *Array_Queue) add(q *Arr_Node) {
	if arrq.last == arrq.size {
		fmt.Println("OVERFLOW")
	} else {
		arrq.queue[arrq.last] = q
		arrq.last++
		if arrq.last == arrq.first {
			arrq.expand_queue()
		}
		if arrq.last == arrq.size {
			arrq.last = 0 //wrap around
			arrq.add(q)
		}
	}
}
func (arrq *Array_Queue) expand_queue() {
	new_q := make([]*Arr_Node, arrq.size*2)
	j := 0
	for i := arrq.first; i < arrq.size; i++ {
		new_q[j] = arrq.queue[i]
		j++
	}
	arrq.first = 0
	for k := 0; k < arrq.last; k++ {
		new_q[j] = arrq.queue[k]
		j++
	}
	arrq.last = arrq.size
	arrq.queue = new_q
	arrq.size = len(arrq.queue)
	fmt.Println("EXPANDED QUEUE")
}

func (arrq *Array_Queue) remove() *Arr_Node {
	if arrq.first == arrq.last {
		fmt.Println("Empty")
		return nil
	} else {
		arrq.first++
		first := arrq.queue[arrq.first-1]
		arrq.queue[arrq.first-1] = nil
		return first
	}
}
func (arrq *Array_Queue) print() { //broken, needs to wrap around
	fmt.Println("CURRENT QUEUE:")
	if arrq.first == arrq.last {
		fmt.Println("Empty")
		return
	} else {
		for i := arrq.first; i < arrq.last; i++ {
			fmt.Println("KEY: ", arrq.queue[i].key, " VAL: ", arrq.queue[i].value, "i: ", arrq.first, " k: ", arrq.last, " n:", arrq.size)
		}
		for i := arrq.last; arrq.last < arrq.first; i++ {
			fmt.Println(i, "")
			fmt.Println("KEY: ", arrq.queue[i].key, " VAL: ", arrq.queue[i].value, "i: ", arrq.first, " k: ", arrq.last, " n:", arrq.size)
		}
	}

}
