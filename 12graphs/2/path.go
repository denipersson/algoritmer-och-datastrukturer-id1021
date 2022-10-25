package main

type paths struct {
	path []*city
	sp   int
}

func new_path() paths {
	p := paths{make([]*city, 54), 0}
	return p
}
func (p *paths) shortest_dist(from *city, to *city, max int) *int {
	if max < 0 {
		return nil
	} else if from == to {
		z := 0
		return &z
	}
	for i := 0; i < p.sp; i++ {
		if p.path[i] == from {
			return nil
		}

	}
	var shortest *int
	shortest = &max

	p.path[p.sp] = from
	p.sp++
	for i := 0; i < len(from.connections); i++ {
		c := from.connections[i]
		dist := p.shortest_dist(c.connected_city, to, max-c.dist)
		if dist != nil {
			path := *dist + c.dist
			if path < *shortest {
				shortest = &path
				max = *shortest + c.dist
			}
		}
	}
	p.sp--
	p.path[p.sp] = nil
	return shortest
}
