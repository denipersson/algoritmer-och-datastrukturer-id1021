package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//rand.Seed(time.Now().UnixMicro())
	test()
}

func test() {

	var (
		//bin        float64 = 0
		//sort       float64 = 0
		//unsort     float64 = 0
		slice_size int = 100
		key        int
		slice      []int
		t0         time.Time
		t1         float64
		bin        float64
		loops      int = 10
	)
	fmt.Println("UNSORTED:")
	for mult := 0; slice_size < 64000000; mult++ {

		slice_size = int(math.Pow(2, float64(mult)))
		bin = 0
		fmt.Printf("%d ", slice_size)
		for i := 0; i < loops; i++ {

			slice, key = sorted_slice(slice_size)
			t0 = time.Now()
			if binary_search(slice, key) {
				t1 = float64(time.Since(t0))
				bin += t1

			}
		}
		fmt.Printf("%f\n", bin/float64(loops)/1000)
	}
	/*
		fmt.Println("SORTED:")
		for i := 0; i < loop_size; i++ {
			t0 := time.Now()

			if search_sorted(slice, key) {
				sort += float64(time.Since(t0))
			}

		}

		fmt.Println(slice_size, " ", sort/float64(loop_size))

		fmt.Println("BINARYSORT:")
		for i := 0; i < loop_size; i++ {
			t0 := time.Now()
			if binary_search(slice, key) {
				bin += float64(time.Since(t0))
			}

		}

		fmt.Println(slice_size, " ", bin/float64(loop_size))
	*/

}

func search_unsorted(array []int, key int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			return true
		}
	}
	return false
}
func unsorted_slice(size int) ([]int, int) {

	slice := make([]int, size)

	for i := 1; i < size; i++ {
		slice[i] = rand.Intn(i)
	}

	key := slice[rand.Intn(len(slice))]

	return slice, key

}

func sorted_slice(size int) ([]int, int) {

	slice := make([]int, size)
	nxt := 0
	for i := 0; i < size; i++ {
		nxt += rand.Intn(10) + 1
		slice[i] = nxt
	}

	key := slice[rand.Intn(size)]

	return slice, key
}

func search_sorted(slice []int, key int) bool {

	increment := len(slice) / 10

	if increment < 1 {
		increment = 1
	}

	for i := 0; i < len(slice); i += increment {
		if slice[i] > key {
			i -= increment
			increment = 1
		}

		if slice[i] == key {
			return true
		}
	}
	return false
}

func binary_search(slice []int, key int) bool {
	first := 0
	last := len(slice) - 1
	index := (last - first) / 2
	for true {
		if slice[index] == key {
			return true
		}
		if slice[index] < key && index < last {

			first = index
			index = last - (last-first)/2
			continue
		}
		if slice[index] > key && index > first {

			last = index
			index = first + (last-first)/2
			continue
		}

	}
	return false

}
