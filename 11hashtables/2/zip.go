package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type Zip struct {
	data []*Node
	max  int
}

func make_zip(file_name string) Zip {
	data := make([]*Node, 12427)

	f, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csv_reader := csv.NewReader(f)
	i := 0
	max := 0
	//collisions := 0
	for true {
		line, err := csv_reader.Read()

		if err == io.EOF {
			break
		}
		pop, _ := strconv.Atoi(line[2])
		zip := zipcode_string_to_int(line[0])

		if data[hash(zip)] != nil { //append to bucket
			k := 0
			for true {
				if data[hash(zip)+k] == nil {
					data[hash(zip)+k] = &Node{zip, line[1], pop}
					break
				}
				k++
				if len(data) <= hash(zip)+k {
					fmt.Println("FULL", zip)
				}
			}
			//collisions++
		} else {
			data[hash(zip)] = &Node{zip, line[1], pop}
		}
		i++
		max = i
	}
	//fmt.Println("collisions: ", collisions)

	return Zip{data, max}
}
func hash(zip_code int) int {
	return int(math.Mod(float64(zip_code), float64(12427)))
}
func (zip *Zip) lookup(zipcode int) bool {
	i := 0
	hash_v := hash(zipcode)
	for hash_v+i < len(zip.data) {
		if zip.data[hash_v+i].code == zipcode {
			fmt.Println(i)
			return true
		}
		i++
		if zip.data[hash_v+i] == nil {
			break
		}
	}
	return false
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

/*
func (z *Zip) linear_search(code int) bool {
	for i := 0; i < z.max; i++ {
		if z.data[i].code == code {
			return true
		}
	}
	return false
}
*/
/*
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
*/
