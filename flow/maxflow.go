package flow

import "github.com/stephen-mcelhose/graphtheory/graph"

type MaxFlowResult struct {
	Value int
	Flow  map[graph.Vertex]map[graph.Vertex]int
}

func EdmondsKarp(n *graph.Network, s, t graph.Vertex) MaxFlowResult {
	flow := map[graph.Vertex]map[graph.Vertex]int{}
	for _, u := range n.Vertices() {
		flow[u] = map[graph.Vertex]int{}
	}

	type step struct {
		parent  graph.Vertex
		forward bool
	}

	bfs := func() (map[graph.Vertex]step, map[graph.Vertex]bool) {
		parent := map[graph.Vertex]step{}
		seen := map[graph.Vertex]bool{s: true}
		queue := []graph.Vertex{s}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range n.Vertices() {
				if !seen[v] && n.Capacity(u, v)-flow[u][v] > 0 {
					seen[v] = true
					parent[v] = step{parent: u, forward: true}
					queue = append(queue, v)
				}
				if !seen[v] && flow[v][u] > 0 {
					seen[v] = true
					parent[v] = step{parent: u, forward: false}
					queue = append(queue, v)
				}
			}
		}
		return parent, seen
	}

	value := 0
	for {
		parent, seen := bfs()
		if !seen[t] {
			break
		}
		aug := int(^uint(0) >> 1)
		v := t
		for v != s {
			st := parent[v]
			u := st.parent
			residual := 0
			if st.forward {
				residual = n.Capacity(u, v) - flow[u][v]
			} else {
				residual = flow[v][u]
			}
			if residual < aug {
				aug = residual
			}
			v = u
		}
		v = t
		for v != s {
			st := parent[v]
			u := st.parent
			if st.forward {
				flow[u][v] += aug
			} else {
				flow[v][u] -= aug
			}
			v = u
		}
		value += aug
	}
	return MaxFlowResult{Value: value, Flow: flow}
}
