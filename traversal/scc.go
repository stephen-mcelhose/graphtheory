package traversal

import "github.com/stephen-mcelhose/graphtheory/graph"

func StronglyConnectedComponents(g *graph.SimpleDigraph) [][]graph.Vertex {
	seen := map[graph.Vertex]bool{}
	order := []graph.Vertex{}
	var dfs1 func(graph.Vertex)
	dfs1 = func(v graph.Vertex) {
		seen[v] = true
		for _, u := range g.Successors(v) {
			if !seen[u] {
				dfs1(u)
			}
		}
		order = append(order, v)
	}
	for _, v := range g.Vertices() {
		if !seen[v] {
			dfs1(v)
		}
	}
	rev := g.Reverse()
	seen = map[graph.Vertex]bool{}
	res := [][]graph.Vertex{}
	var dfs2 func(graph.Vertex, *[]graph.Vertex)
	dfs2 = func(v graph.Vertex, comp *[]graph.Vertex) {
		seen[v] = true
		*comp = append(*comp, v)
		for _, u := range rev.Successors(v) {
			if !seen[u] {
				dfs2(u, comp)
			}
		}
	}
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		if seen[v] {
			continue
		}
		comp := []graph.Vertex{}
		dfs2(v, &comp)
		res = append(res, comp)
	}
	return res
}

func WeaklyConnectedComponents(g *graph.SimpleDigraph) [][]graph.Vertex {
	return ConnectedComponents(g.UnderlyingSimpleGraph())
}
