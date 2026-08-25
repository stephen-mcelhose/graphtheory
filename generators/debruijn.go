package generators

import "github.com/stephen-mcelhose/graphtheory/graph"

func powInt(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}

// DeBruijnDigraph constructs the k-ary order-n de Bruijn digraph.
func DeBruijnDigraph(k, n int) *graph.SimpleDigraph {
	g := graph.NewSimpleDigraph()
	if n <= 0 {
		return g
	}
	verts := powInt(k, n-1)
	for i := 0; i < verts; i++ {
		g.AddVertex(graph.Vertex(i))
	}
	mod := 1
	if n > 1 {
		mod = powInt(k, n-2)
	}
	for i := 0; i < verts; i++ {
		suffix := 0
		if n > 1 {
			suffix = i % mod
		}
		for d := 0; d < k; d++ {
			j := suffix*k + d
			g.AddArc(graph.Vertex(i), graph.Vertex(j))
		}
	}
	return g
}
