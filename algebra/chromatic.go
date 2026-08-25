package algebra

import "github.com/stephen-mcelhose/graphtheory/graph"

// ChromaticPolynomial returns coefficients c where sum_i c[i] k^i gives the
// number of proper k-colorings. Exponential-time deletion-contraction.
func ChromaticPolynomial(g *graph.SimpleGraph) []int {
	if g.Size() == 0 {
		coeff := make([]int, g.Order()+1)
		coeff[g.Order()] = 1
		return coeff
	}
	edges := g.Edges()
	e := edges[0]
	del := g.Clone()
	del.RemoveEdge(e[0], e[1])
	con := contractEdge(g, e[0], e[1])
	return subPoly(ChromaticPolynomial(del), ChromaticPolynomial(con))
}

func contractEdge(g *graph.SimpleGraph, u, v graph.Vertex) *graph.SimpleGraph {
	out := graph.NewSimpleGraph()
	for _, x := range g.Vertices() {
		if x != v {
			out.AddVertex(x)
		}
	}
	for _, e := range g.Edges() {
		a, b := e[0], e[1]
		if a == v {
			a = u
		}
		if b == v {
			b = u
		}
		if a != b {
			out.AddEdge(a, b)
		}
	}
	return out
}

func subPoly(a, b []int) []int {
	n := len(a)
	if len(b) > n {
		n = len(b)
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		if i < len(a) {
			out[i] += a[i]
		}
		if i < len(b) {
			out[i] -= b[i]
		}
	}
	return trim(out)
}

func trim(a []int) []int {
	i := len(a) - 1
	for i > 0 && a[i] == 0 {
		i--
	}
	return a[:i+1]
}
