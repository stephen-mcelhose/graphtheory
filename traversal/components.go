package traversal

import "github.com/stephen-mcelhose/graphtheory/graph"

func ConnectedComponents(g graph.Graph) [][]graph.Vertex {
	seen := map[graph.Vertex]bool{}
	comps := [][]graph.Vertex{}
	for _, v := range g.Vertices() {
		if seen[v] {
			continue
		}
		comp := []graph.Vertex{}
		queue := []graph.Vertex{v}
		seen[v] = true
		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]
			comp = append(comp, x)
			for _, y := range g.Neighbors(x) {
				if !seen[y] {
					seen[y] = true
					queue = append(queue, y)
				}
			}
		}
		comps = append(comps, comp)
	}
	return comps
}

func IsConnected(g graph.Graph) bool {
	if g.Order() == 0 {
		return true
	}
	return len(ConnectedComponents(g)) == 1
}
