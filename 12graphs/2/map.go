package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type maps struct {
	cities []*city
}

func new_map(file string) maps {
	new := maps{make([]*city, 541)}
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csv_reader := csv.NewReader(f)
	for true {
		line, err := csv_reader.Read()

		if err == io.EOF {
			break
		}
		c1 := new.lookup(line[0])
		c2 := new.lookup(line[1])
		dist, _ := strconv.Atoi(line[2])
		c1.add_connection(c2, dist)
	}
	return new
}
func (m *maps) lookup(name string) *city {
	i := 0
	hash_v := hash_code(name)
	for hash_v+i < len(m.cities) {
		if m.cities[hash_v+i] == nil {
			c := new_city(name)
			m.cities[hash_v+i] = &c
			break
		}
		if m.cities[hash_v+i].name == name {
			break
		}
		i++
	}
	return m.cities[hash_v+i]
}
func (m *maps) print_cities() {
	for i := 0; i < len(m.cities); i++ {
		if m.cities[i] != nil {
			fmt.Println(i, ": ", m.cities[i].name)
			m.cities[i].print_connections()
		}
	}
} /*
func shortest_dist(from *city, to *city, max int) *int {
	if max < 0 {
		return nil
	} else if from == to {
		z := 0
		return &z
	}
	var shortest *int
	shortest = &max
	for i := 0; i < len(from.connections); i++ {
		c := from.connections[i]
		dist := shortest_dist(c.connected_city, to, max-c.dist)
		if dist != nil {
			path := *dist + c.dist
			if path < *shortest {
				shortest = &path
			}
		}
	}
	return shortest
}*/
