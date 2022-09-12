package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//rand.Seed(time.Now().UnixMicro())
	//test()

	//total_duplicates := 0
	for i := 1; i <= 1000; i++ {
		//total_duplicates += test_find_rewritten_duplicates(i)
		//total_duplicates += test_find_duplicates(i)
		//if time.Since(time_1_min).Seconds() > 60 {
		//	break
		//}
		var t_after int64
		t0 := time.Now()

		for k := 0; k < 30; k++ {
			search_unsorted(sorted_slice(i * 10))
			t_after += int64(time.Since(t0).Nanoseconds())
		}
		fmt.Println(i*10, t_after/30)

	}

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

func sorted_slice_no_key(size int) []int {

	slice := make([]int, size)
	nxt := 0
	for i := 0; i < size; i++ {
		nxt += rand.Intn(10) + 1
		slice[i] = nxt
	}
	return slice
}

func search_sorted(array []int, key int) bool {
	for i := 0; array[i] <= key; i++ {
		if array[i] == key {
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

func test_find_duplicates(mult int) int {
	//t0 := time.Now()
	duplicates := search_duplicates(10*mult, 100)
	//t_after := time.Since(t0).Nanoseconds() / 100

	//fmt.Printf("%d %d\n", mult*10, t_after)
	return duplicates
}
func test_find_rewritten_duplicates(mult int) int {
	t0 := time.Now()
	duplicates := rewritten_search_duplicates(10*mult, 1000)
	t_after := time.Since(t0).Nanoseconds() / 1000

	fmt.Printf("%d %d\n", mult*10, t_after)

	return duplicates
}

func search_duplicates(size int, k int) int {

	keys := sorted_slice_no_key(size)
	values := make([]int, size)
	copy(keys, values)
	duplicates := 0
	rand.Seed(time.Now().UTC().UnixNano())
	t0 := time.Now()
	var t_after int64
	for j := 0; j < k; j++ {

		for ki := 0; ki < size; ki++ {
			key := keys[ki]

			if binary_search(values, key) {
				duplicates++

				t_after += int64(time.Since(t0).Milliseconds())
			}
			t_after += int64(time.Since(t0).Milliseconds())
		}
	}
	fmt.Println(size, " ", t_after/int64(k))
	return duplicates

}

func rewritten_search_duplicates(size int, k int) int {

	keys := sorted_slice_no_key(size)
	values := sorted_slice_no_key(size)
	rand.Seed(time.Now().UTC().UnixNano())
	var t_after int64

	ki := 0
	vi := 0
	duplicates := 0

	t0 := time.Now()

	for ki < len(keys) && vi < len(values) {
		if keys[ki] == values[vi] {
			ki++
			duplicates++
			t_after += int64(time.Since(t0).Milliseconds())

			continue
		} else if keys[ki] < values[vi] {
			ki++
			t_after += int64(time.Since(t0).Milliseconds())

			continue
		} else {
			t_after += int64(time.Since(t0).Milliseconds())

			vi++
		}
	}

	//fmt.Println(size, " ", t_after/int64(k))
	return duplicates / k

}

/*

'int array1_index = 0;
int array2_index = 0;
int antal_hittade = 0;
while(array1_index < array1.length){
    if(array1[array1_index] == array2[array2_index]){
        array1_index++;
        antal_hittade++;
        continue;
        }

    else if(array1[array1_index] < array2[array2_index]){
        array1_index++;
        continue;
        }

    array2_index++
    }'

*/
