package traversal

import "github.com/stephen-mcelhose/graphtheory/graph"

func Eccentricities(g graph.Graph) map[graph.Vertex]int {
	ecc := map[graph.Vertex]int{}
	for _, v := range g.Vertices() {
		d := Distances(g, v)
		maxd := 0
		for _, w := range g.Vertices() {
			if dist, ok := d[w]; ok {
				if dist > maxd {
					maxd = dist
				}
			} else {
				maxd = -1
				break
			}
		}
		ecc[v] = maxd
	}
	return ecc
}

func Centers(g graph.Graph) []graph.Vertex {
	ecc := Eccentricities(g)
	best := -1
	res := []graph.Vertex{}
	for _, v := range g.Vertices() {
		e := ecc[v]
		if e < 0 {
			continue
		}
		if best == -1 || e < best {
			best = e
			res = []graph.Vertex{v}
		} else if e == best {
			res = append(res, v)
		}
	}
	return res
}
