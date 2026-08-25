package theory

import "github.com/stephen-mcelhose/graphtheory/graph"

func DiracCondition(g graph.Graph) bool {
	n := g.Order()
	if n < 3 {
		return false
	}
	for _, v := range g.Vertices() {
		if g.Degree(v) < n/2 {
			return false
		}
	}
	return true
}

func OreCondition(g graph.Graph) bool {
	verts := g.Vertices()
	n := len(verts)
	if n < 3 {
		return false
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			u, v := verts[i], verts[j]
			if !g.HasEdge(u, v) && g.Degree(u)+g.Degree(v) < n {
				return false
			}
		}
	}
	return true
}
