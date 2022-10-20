package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Zip struct {
	data []Node
	max  int
}

func make_zip(file_name string) Zip {
	data := make([]Node, 10000)

	f, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csv_reader := csv.NewReader(f)
	i := 0
	max := 0
	for true {
		line, err := csv_reader.Read()

		if err == io.EOF {
			break
		}
		pop, _ := strconv.Atoi(line[2])
		zip := zipcode_string_to_int(line[0])
		data[i] = Node{zip, line[1], pop}
		i++
		max = i // - 1 ?
	}

	return Zip{data, max}
}
func zipcode_string_to_int(str string) int {
	new_str := ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			continue
		}
		new_str += string(str[i])
	}
	integer, _ := strconv.Atoi(new_str)
	return integer
}
func (z *Zip) linear_search(code int) bool {
	for i := 0; i < z.max; i++ {
		if z.data[i].code == code {
			return true
		}
	}
	return false
}

func (z *Zip) binary_search(code int) bool {
	first := 0
	last := z.max
	index := (last - first) / 2
	for true {
		if z.data[index].code == code {
			return true
		}
		if z.data[index].code < code && index < last {
			first = index
			index = last - (last-first)/2
			continue
		}
		if z.data[index].code > code && index > first {
			last = index
			index = first + (last-first)/2
			continue
		}
		return false //not found
	}
	return false
}
