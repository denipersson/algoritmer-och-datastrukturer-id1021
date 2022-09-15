package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())

	for i := 0; i <= 10000; i += 1000 {
		test_insertion_sort(i)
	}

	/*
		test_merge_sort(2)
		test_merge_sort(4)
		test_merge_sort(8)
		test_merge_sort(16)
		test_merge_sort(32)
		test_merge_sort(64)
		test_merge_sort(128)
		test_merge_sort(256)
		test_merge_sort(512)
		test_merge_sort(1024)
		test_merge_sort(2048)
		test_merge_sort(4096)
		test_merge_sort(8192)
		test_merge_sort(16384)
		test_merge_sort(16384 * 2)
		test_merge_sort(16384 * 2 * 2)
		test_merge_sort(16384 * 2 * 2 * 2)
		test_merge_sort(16384 * 2 * 2 * 2 * 2)
	*/
}
func test_selection_sort(mult int) {
	size := 10 * mult

	arr := rand_int_arr(size)
	t0 := time.Now()
	selection_sort(arr)
	t_after := time.Since(t0)
	fmt.Println(len(arr), " ", t_after.Microseconds())
}

func rand_int_arr(size int) []int {
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size)
	}
	return slice
}

func selection_sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {

		cand := i

		for j := i; j < len(arr); j++ {
			if arr[j] < arr[cand] {
				cand = j
			}
		}
		temp := arr[i]
		arr[i] = arr[cand]
		arr[cand] = temp
	}
	return arr

}

func insertion_sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			copy := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = copy
			j--
		}
	}
	return arr

}
func test_insertion_sort(mult int) {
	size := 10 * mult

	arr := rand_int_arr(size)
	t0 := time.Now()
	insertion_sort(arr)
	t_after := time.Since(t0)
	fmt.Println(len(arr), " ", t_after.Microseconds())

}

func test_merge_sort(mult int) {
	size := 10 * mult

	arr := rand_int_arr(size)
	t0 := time.Now()
	merge_sort(arr)
	t_after := time.Since(t0)
	fmt.Println(len(arr), " ", t_after.Microseconds())
}

func merge_sort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}

	arr1 := arr[:(len(arr))/2]
	arr2 := arr[(len(arr))/2:]

	arr1 = merge_sort(arr1)
	arr2 = merge_sort(arr2)

	return merge(arr1, arr2)
}
func merge(arra []int, arrb []int) []int {

	last := make([]int, len(arra)+len(arrb))
	a_index, b_index, merge_index := 0, 0, 0

	for ; a_index < len(arra) && b_index < len(arrb); merge_index++ {
		if arra[a_index] < arrb[b_index] {
			last[merge_index] = arra[a_index]
			a_index++
		} else {
			last[merge_index] = arrb[b_index]
			b_index++
		}
	}
	for len(arra) > a_index {
		last[merge_index] = arra[a_index]
		a_index++
		merge_index++
	}
	for len(arrb) > b_index {
		last[merge_index] = arrb[b_index]
		b_index++
		merge_index++
	}
	return last
}
