package main

import (
	"fmt"
	"time"
)

func main() {
	//test_linear_zip(98499)
	//test_binary_zip(98499)
	test_look_up(98499)

}

/*
	func test_linear_zip(post_num int) {
		zip := make_zip("postnr.csv")
		t0 := time.Now()
		var t_since int64
		f := zip.linear_search(post_num)
		t_since += time.Since(t0).Nanoseconds()
		fmt.Println("linear: ", post_num, " found: ", f, " ", t_since)
	}

	func test_binary_zip(post_num int) {
		zip := make_zip("postnr.csv")
		t0 := time.Now()
		f := zip.binary_search(post_num)
		t_since := time.Since(t0).Nanoseconds()
		fmt.Println("binary: ", post_num, " found: ", f, " ", t_since)
	}
*/
func test_look_up(post_num int) {
	zip := make_zip("postnr.csv")
	t0 := time.Now()
	f := zip.lookup(post_num)
	t_since := time.Since(t0).Nanoseconds()
	fmt.Println("lookup: ", post_num, " found: ", f, " ", t_since)
}
