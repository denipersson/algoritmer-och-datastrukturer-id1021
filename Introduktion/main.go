package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 1; i <= 300; i++ {
		test_search(i)
	}

	/*
		total_duplicates := 0
		//time_1_min := time.Now()

		for i := 1; i <= 300; i++ {
			total_duplicates += test_find_duplicates(i)

			//if time.Since(time_1_min).Seconds() > 60 {
			//	break
			//}
		}
		fmt.Println("Total duplicates found: ", total_duplicates)
	*/
}
func test_search(mult int) {
	search_time := search_item(11*mult, 10*mult, 100)
	fmt.Printf("%d %f\n", mult*10, search_time)
}
func test_find_duplicates(mult int) int {
	find_time, duplicates := search_duplicates(10*mult, 100)

	fmt.Printf("%d %f\n", mult*10, find_time)
	return duplicates
}

func test_time_to_access(n int) {
	time_to_access := rand_access(n)
	fmt.Println(n, " Resolution ", time_to_access, " ns")
}

func rand_access(n int) float64 {

	k := 1_000_000

	indx := make([]int, n)

	for i := 0; i < n; i++ {
		indx[i] = rand.Intn(n)
	}

	slice := make([]int, n)

	for i := 0; i < n; i++ {
		slice[i] = 1
	}

	sum := 0

	t0 := time.Now()
	for j := 0; j < k; j++ {
		for i := 0; i < n; i++ {
			sum += slice[indx[i]]
		}
	}
	t_access := time.Now().Sub(t0)

	t0 = time.Now()

	for j := 0; j < k; j++ {
		for i := 0; i < n; i++ {
			sum += 1
		}
	}
	t_dummy := time.Now().Sub(t0)

	return (float64(t_access-t_dummy) / (float64)(k*n))

}

func search_item(key_count int, values_count int, k int) float64 {

	keys := make([]int, key_count)
	values := make([]int, values_count)
	var t_total time.Duration
	sum := 0

	for j := 0; j < k; j++ {
		for ki := 0; ki < key_count; ki++ {
			keys[ki] = rand.Intn(10 * values_count)
		}
		for v := 0; v < values_count; v++ {
			values[v] = rand.Intn(10 * values_count)
		}

		t0 := time.Now()

		for ki := 0; ki < key_count; ki++ {
			key := keys[ki]
			for v := 0; v < values_count; v++ {
				if values[v] == key {
					sum++
					break
				}
			}
		}
		t1 := time.Now().Sub(t0)
		t_total += t1
	}
	return float64(t_total) / float64(k)

}

func search_duplicates(size int, k int) (float64, int) {

	keys := make([]int, size)
	values := make([]int, size)
	var t_total time.Duration
	duplicates := 0
	rand.Seed(time.Now().UTC().UnixNano())

	for j := 0; j < k; j++ {
		for ki := 0; ki < size; ki++ {
			keys[ki] = rand.Intn(10 * size)
			values[ki] = rand.Intn(10 * size)
		}

		t0 := time.Now()

		for ki := 0; ki < size; ki++ {
			key := keys[ki]
			for v := 0; v < size; v++ {
				if values[v] == key {
					duplicates++
				}
			}
		}
		t1 := time.Now().Sub(t0)
		t_total += t1
	}

	return float64(t_total) / float64(k), duplicates

}
