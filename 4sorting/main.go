package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())

	for i := 0; i < 1000; i++ {
		//test_section_sort(i)
		test_insertion_sort(i)
	}
}
func test_selection_sort(mult int) {
	size := 10 * mult

	arr := rand_int_arr(size)
	//fmt.Println("Unsorted: ", arr)
	/*sorted_arr := */
	t0 := time.Now()
	selection_sort(arr)
	t_after := time.Since(t0)
	fmt.Println(len(arr), " ", t_after.Microseconds())
	//fmt.Println("Sorted(section sort): ", sorted_arr)
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

		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
	return arr

}
func test_insertion_sort(mult int) {
	size := 10 * mult

	arr := rand_int_arr(size)
	//fmt.Println("Unsorted: ", arr)
	t0 := time.Now()
	/*sorted_arr := */ insertion_sort(arr)
	t_after := time.Since(t0)
	fmt.Println(len(arr), " ", t_after.Microseconds())
	//fmt.Println("Sorted(insertion sort): ", sorted_arr)
}

/*

	for (int i = 0; i < array.length -1; i++) {
		// let's set the first candidate to the index itself int cand = i;
		for (int j = i; j < array.length ; j++) {
			// If the element at position j is smaller than the value
			// at the candidate position - then you have a new candidate
			// posistion.
		}
		// Swap the item at position i with the item at the candidate position.
	}
*/
