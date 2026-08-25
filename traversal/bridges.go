package traversal

import "github.com/stephen-mcelhose/graphtheory/graph"

func Bridges(g graph.Graph) [][2]graph.Vertex {
	var time int
	disc := map[graph.Vertex]int{}
	low := map[graph.Vertex]int{}
	parent := map[graph.Vertex]graph.Vertex{}
	bridges := [][2]graph.Vertex{}

	var dfs func(graph.Vertex)
	dfs = func(v graph.Vertex) {
		time++
		disc[v] = time
		low[v] = time
		for _, u := range g.Neighbors(v) {
			if disc[u] == 0 {
				parent[u] = v
				dfs(u)
				if low[u] < low[v] {
					low[v] = low[u]
				}
				if low[u] > disc[v] {
					if v < u {
						bridges = append(bridges, [2]graph.Vertex{v, u})
					} else {
						bridges = append(bridges, [2]graph.Vertex{u, v})
					}
				}
			} else if u != parent[v] {
				if disc[u] < low[v] {
					low[v] = disc[u]
				}
			}
		}
	}

	for _, v := range g.Vertices() {
		if disc[v] == 0 {
			dfs(v)
		}
	}
	return bridges
}
