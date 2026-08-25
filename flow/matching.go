package flow

import "github.com/stephen-mcelhose/graphtheory/graph"

type Matching map[graph.Vertex]graph.Vertex

func MaxBipartiteMatching(g graph.Graph, left []graph.Vertex, right []graph.Vertex) Matching {
	isRight := map[graph.Vertex]bool{}
	for _, v := range right {
		isRight[v] = true
	}
	matchR := map[graph.Vertex]graph.Vertex{}
	var aug func(graph.Vertex, map[graph.Vertex]bool) bool
	aug = func(u graph.Vertex, seen map[graph.Vertex]bool) bool {
		for _, v := range g.Neighbors(u) {
			if !isRight[v] || seen[v] {
				continue
			}
			seen[v] = true
			mu, ok := matchR[v]
			if !ok || aug(mu, seen) {
				matchR[v] = u
				return true
			}
		}
		return false
	}
	for _, u := range left {
		seen := map[graph.Vertex]bool{}
		aug(u, seen)
	}
	out := Matching{}
	for v, u := range matchR {
		out[u] = v
	}
	return out
}
