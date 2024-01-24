package main

import(
        "fmt"
	"container/list"
)

func HasCycle(edge []*ServiceEdge) (bool) {
	indegrees := make(map[string]int)
	direction := make(map[string][]string)
	nodes := make(map[string]struct {})
	topology := []string{}
	for _, e := range edge {
		indegrees[*e.ToNode]++
		direction[*e.FromNode] = append(direction[*e.FromNode], *e.ToNode)
		nodes[*e.FromNode] = struct{}{}
		nodes[*e.ToNode] = struct{}{}
	}
	fmt.Printf("indegrees = %v\n", indegrees)
	fmt.Printf("direction = %v\n", direction)
	fmt.Printf("nodes = %v\n", nodes)
	l := list.New()
	for e := range direction {
		if indegrees[e] == 0 {
			l.PushBack(e)
			fmt.Printf("init zero edge = %v\n", e)
		}
	}
	if l.Len() == 0 {
		return true
	}
	for l.Len() != 0 {
		e := l.Remove(l.Front()).(string)
		topology = append(topology, e)
		for _, to := range direction[e] {
			indegrees[to]--
			if indegrees[to] == 0 {
				l.PushBack(to)
				fmt.Printf("add zero edge = %v\n", to)
			}
		}
	}
	fmt.Printf("topology = %v\n", topology)
	return len(topology) != len(nodes)
}

type ServiceEdge struct {
	FromNode *string
	ToNode   *string
}

func NewEdge(from, to string) (*ServiceEdge) {
	se := new(ServiceEdge)
	se.FromNode = &from
	se.ToNode = &to
	return se
}
func main() {
	edges := make([]*ServiceEdge, 0)
	edges = append(edges, NewEdge("a", "b"))
	edges = append(edges, NewEdge("b", "a"))
	edges = append(edges, NewEdge("b", "c"))
	edges = append(edges, NewEdge("c", "d"))
	edges = append(edges, NewEdge("d", "e"))

	edges = append(edges, NewEdge("x", "y"))
	edges = append(edges, NewEdge("z", "k"))
	for _, e := range edges {
		fmt.Printf("edges = %v -> %v\n", *(e.FromNode), *(e.ToNode))
	}
	fmt.Printf("\n\nHasCycle = %t\n", HasCycle(edges))

	//fmt.Printf("\n\nHasCycle = %t\n", HasCycle(edges))
}
