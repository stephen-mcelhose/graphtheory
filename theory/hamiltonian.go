package theory

import "github.com/stephen-mcelhose/graphtheory/graph"

func HamiltonianPath(g graph.Graph) []graph.Vertex {
	verts := g.Vertices()
	n := len(verts)
	used := map[graph.Vertex]bool{}
	path := []graph.Vertex{}
	var dfs func(graph.Vertex) bool
	dfs = func(v graph.Vertex) bool {
		path = append(path, v)
		used[v] = true
		if len(path) == n {
			return true
		}
		for _, u := range g.Neighbors(v) {
			if !used[u] && dfs(u) {
				return true
			}
		}
		used[v] = false
		path = path[:len(path)-1]
		return false
	}
	for _, v := range verts {
		for k := range used {
			delete(used, k)
		}
		path = path[:0]
		if dfs(v) {
			return append([]graph.Vertex{}, path...)
		}
	}
	return nil
}

func HamiltonianCycle(g graph.Graph) []graph.Vertex {
	path := HamiltonianPath(g)
	if len(path) == 0 {
		return nil
	}
	if g.HasEdge(path[0], path[len(path)-1]) {
		return append(path, path[0])
	}
	return nil
}

func TournamentHamiltonianPath(g *graph.SimpleDigraph) []graph.Vertex {
	path := []graph.Vertex{}
	for _, v := range g.Vertices() {
		inserted := false
		for i := 0; i < len(path); i++ {
			if g.HasArc(v, path[i]) {
				path = append(path[:i], append([]graph.Vertex{v}, path[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			path = append(path, v)
		}
	}
	return path
}
