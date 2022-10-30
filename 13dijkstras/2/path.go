package main

import (
	"container/heap"
)

type Path struct {
	to    *city
	dist  int
	index int
}
type PriorityQueue []*Path

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist > pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(p any) {
	n := len(*pq)
	item := p.(*Path)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq PriorityQueue) Peek() any {
	return pq[0]
}
func (pq PriorityQueue) Get_Item(p *city) any {
	for i := 0; i < len(pq); i++ {
		if pq[i].to == p {
			return pq[i]
		}
	}
	return nil
}

func (pq *PriorityQueue) update(path *Path, dist int) {
	path.dist = dist
	heap.Fix(pq, path.index)
}

func (m *maps) dijkstras_search(source *city, destination *city) int {

	pq := PriorityQueue{}
	pq.Push(new_path(source, 0))
	var cur *Path
	visited := make(map[*city]bool, len(m.cities))
	visited[source] = true
	for pq.Len() > 0 {
		cur = pq.Pop().(*Path)
		if cur.to == destination {
			break
		}
		adj := cur.to.connections
		for i := 0; i < len(adj); i++ {
			vis := adj[i].connected_city
			new_time := adj[i].dist + cur.dist
			if visited[vis] == false {
				visited[vis] = true
				pq.Push(new_path(vis, new_time))
			} else {
				if pq.Get_Item(vis) != nil {
					item := pq.Get_Item(vis).(*Path)
					if item.dist > new_time {
						pq.update(item, new_time)
					}
				}
			}
		}
	}

	return cur.dist
}
func new_path(to *city, dist int) *Path {
	p := Path{to, dist, 0}
	return &p
}
