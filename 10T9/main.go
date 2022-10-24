package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	t9 := T9{new_node()}
	f, err := os.Open("kelly.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		t9.add(s)
	}

	fmt.Println(t9.start_search("161"))
	fmt.Println(t9.start_search("2314"))

}
