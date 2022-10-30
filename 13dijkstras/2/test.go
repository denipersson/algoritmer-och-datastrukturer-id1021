package main

import (
	"fmt"
	"time"
)

func test_main(c_name1 string, c_name2 string, max int) {
	new := new_map("trains.csv")
	c1 := new.lookup(c_name1)
	c2 := new.lookup(c_name2)
	t0 := time.Now()
	dist := new.dijkstras_search(c1, c2)
	t_since := time.Since(t0).Nanoseconds()
	fmt.Println(c_name1, "->", c_name2, " ", dist, " min ", t_since, " ns")
	fmt.Println("")
}
