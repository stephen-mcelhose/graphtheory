package algebra

import "github.com/stephen-mcelhose/graphtheory/graph"

func SpanningTreeCount(g graph.Graph) int {
	_, lap := LaplacianMatrix(g)
	if len(lap) <= 1 {
		return 1
	}
	cofactor := Minor(lap, 0, 0)
	return Determinant(cofactor)
}

func RootedSpanningArborescenceCount(g *graph.SimpleDigraph, root graph.Vertex) int {
	verts, lap := DirectedLaplacian(g)
	idx := -1
	for i, v := range verts {
		if v == root {
			idx = i
			break
		}
	}
	if idx < 0 {
		return 0
	}
	cofactor := Minor(lap, idx, idx)
	return Determinant(cofactor)
}
