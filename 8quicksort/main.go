package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test_quick_sort()
}

func test_quick_sort() {
	arr := create_random_array(20)
	fmt.Println("FIRST:\n", arr)
	quick_sort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func create_random_array(size int) []int {
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(size)
	return arr
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
		fmt.Println(*array, ", pivot: ", (*array)[end])
		for (*array)[left_i] < pivot && left_i < end {
			left_i++
		}
		for (*array)[right_i] >= pivot && right_i > 0 {
			right_i--
		}
		if left_i < right_i {
			fmt.Println((*array)[left_i], "<- left  right ->", (*array)[right_i])
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
