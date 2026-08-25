package traversal

import "github.com/stephen-mcelhose/graphtheory/graph"

func BFS(g graph.Graph, start graph.Vertex) []graph.Vertex {
	if !g.HasVertex(start) {
		return nil
	}
	seen := map[graph.Vertex]bool{start: true}
	queue := []graph.Vertex{start}
	order := []graph.Vertex{}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, u := range g.Neighbors(v) {
			if !seen[u] {
				seen[u] = true
				queue = append(queue, u)
			}
		}
	}
	return order
}

func Distances(g graph.Graph, start graph.Vertex) map[graph.Vertex]int {
	dist := map[graph.Vertex]int{}
	if !g.HasVertex(start) {
		return dist
	}
	queue := []graph.Vertex{start}
	dist[start] = 0
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, u := range g.Neighbors(v) {
			if _, ok := dist[u]; !ok {
				dist[u] = dist[v] + 1
				queue = append(queue, u)
			}
		}
	}
	return dist
}
