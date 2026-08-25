package traversal

import "github.com/stephen-mcelhose/graphtheory/graph"

func DFS(g graph.Graph, start graph.Vertex) []graph.Vertex {
	if !g.HasVertex(start) {
		return nil
	}
	seen := map[graph.Vertex]bool{}
	order := []graph.Vertex{}
	var visit func(graph.Vertex)
	visit = func(v graph.Vertex) {
		seen[v] = true
		order = append(order, v)
		for _, u := range g.Neighbors(v) {
			if !seen[u] {
				visit(u)
			}
		}
	}
	visit(start)
	return order
}
