package generators

import "github.com/stephen-mcelhose/graphtheory/graph"

func Hypercube(dim int) *graph.SimpleGraph {
	n := 1 << dim
	g := graph.NewSimpleGraph()
	for i := 0; i < n; i++ {
		g.AddVertex(graph.Vertex(i))
	}
	for i := 0; i < n; i++ {
		for b := 0; b < dim; b++ {
			j := i ^ (1 << b)
			if i < j {
				g.AddEdge(graph.Vertex(i), graph.Vertex(j))
			}
		}
	}
	return g
}
