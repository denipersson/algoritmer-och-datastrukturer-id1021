package main

import (
	"fmt"
	"time"
)

func test_main(c_name1 string, c_name2 string, max int) {
	new := new_map("trains.csv")
	//new.print_cities()
	t0 := time.Now()
	dist := shortest_dist(new.lookup(c_name1), new.lookup(c_name2), max)
	t_since := time.Since(t0).Milliseconds()
	fmt.Println("")
	fmt.Println(c_name1, "->", c_name2, " shortest: ", *dist, " min (", t_since, " ms)")
}
