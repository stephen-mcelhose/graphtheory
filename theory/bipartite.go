package theory

import "github.com/stephen-mcelhose/graphtheory/graph"

func TwoColoring(g graph.Graph) (map[graph.Vertex]int, bool) {
	color := map[graph.Vertex]int{}
	for _, s := range g.Vertices() {
		if _, ok := color[s]; ok {
			continue
		}
		queue := []graph.Vertex{s}
		color[s] = 0
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			for _, u := range g.Neighbors(v) {
				if c, ok := color[u]; ok {
					if c == color[v] {
						return nil, false
					}
				} else {
					color[u] = 1 - color[v]
					queue = append(queue, u)
				}
			}
		}
	}
	return color, true
}

func IsBipartite(g graph.Graph) bool {
	_, ok := TwoColoring(g)
	return ok
}
