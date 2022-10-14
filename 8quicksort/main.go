package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i <= 10000001; i += 100000 {
		vals := rand.Perm(i*int(math.Log(float64(i))) + 1)

		test_quick_sort(1*i, vals)

	}

	/*
	   t := create_random_linked_list(10)
	   fmt.Println(t.String())
	   t.swap_nodes(t.start, t.end)
	   fmt.Println(t.String())
	   t.swap_nodes(t.start.next, t.end.prev)
	   fmt.Println(t.String())
	*/
}

func test_quick_sort(size int, vals []int) {
	arr := create_random_array(size, vals)
	t0 := time.Now()
	quick_sort(&arr, 0, len(arr)-1)
	t_since := time.Since(t0).Nanoseconds()

	fmt.Println(size, t_since)

	//fmt.Println(size*10, " ", t_since.Nanoseconds()/10)
}
func test_quick_sort_LL(size int, vals []int) {
	ll := create_random_linked_list(size, vals)
	t0 := time.Now()
	ll.quick_sort_LL(ll.start, ll.end)
	t_since := time.Since(t0).Nanoseconds()

	fmt.Println(size, t_since)
	//fmt.Println(size*10, " ", t_since.Nanoseconds()/10)
}

func create_random_array(size int, vals []int) []int {
	rand.Seed(time.Now().Unix())
	arr := vals
	return arr
}
func create_random_linked_list(size int, vals []int) LinkedList {

	list := new_doubly_linked_list(vals[0])

	for i := 1; i < size; i++ {
		list.append(&Node{vals[i], nil, nil})
	}

	return list

}

func quick_sort(array *[]int, start int, end int) {
	if start < end {
		p := partition(array, start, end)
		quick_sort(array, start, p-1)
		quick_sort(array, p+1, end)
	}
}
func partition(array *[]int, start int, end int) int {

	left_i := start
	right_i := end
	pivot := (*array)[end]

	for left_i < right_i {
		for (*array)[left_i] < pivot && left_i < end {
			left_i++
		}
		for (*array)[right_i] >= pivot && right_i > 0 {
			right_i--
		}
		if left_i < right_i {
			swap(array, left_i, right_i)
		} else {
			swap(array, left_i, end)
		}
	}

	return left_i
}

func swap(array *[]int, i int, j int) {
	if i < 0 || j < 0 || i >= len(*array) || j >= len(*array) {
		fmt.Println("Invalid input")
		return
	}
	left_val := (*array)[i]
	right_val := (*array)[j]
	(*array)[i] = right_val
	(*array)[j] = left_val
}
