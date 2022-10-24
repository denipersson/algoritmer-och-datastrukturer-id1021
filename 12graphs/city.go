package main

import "fmt"

type city struct {
	name        string
	connections []connection
}

func new_city(name string) city {
	n := city{name, make([]connection, 0)}
	return n
}

func (c *city) add_connection(c_to_connect *city, dist int) {
	con1 := connection{c_to_connect, dist}
	c.connections = append(c.connections, con1)
	con2 := connection{c, dist}
	c_to_connect.connections = append(c_to_connect.connections, con2)
}

func hash_code(name string) int {
	hash := 7
	mod := 541
	for i := 0; i < len(name); i++ {
		hash = (hash * 31 % mod) + int(name[i])
	}
	return hash % mod
}
func (c *city) print_connections() {
	for i := 0; i < len(c.connections); i++ {
		fmt.Println(c.name, " -> ", c.connections[i].connected_city.name, "| ", c.connections[i].dist, " minutes")
	}
}
