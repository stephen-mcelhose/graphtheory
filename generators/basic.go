package generators

import "github.com/stephen-mcelhose/graphtheory/graph"

func Empty(n int) *graph.SimpleGraph {
	g := graph.NewSimpleGraph()
	for i := 0; i < n; i++ {
		g.AddVertex(graph.Vertex(i))
	}
	return g
}

func Complete(n int) *graph.SimpleGraph {
	g := Empty(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			g.AddEdge(graph.Vertex(i), graph.Vertex(j))
		}
	}
	return g
}

func Path(n int) *graph.SimpleGraph {
	g := Empty(n)
	for i := 0; i+1 < n; i++ {
		g.AddEdge(graph.Vertex(i), graph.Vertex(i+1))
	}
	return g
}

func Cycle(n int) *graph.SimpleGraph {
	g := Path(n)
	if n > 2 {
		g.AddEdge(graph.Vertex(0), graph.Vertex(n-1))
	}
	return g
}

func CompleteBipartite(m, n int) *graph.SimpleGraph {
	g := graph.NewSimpleGraph()
	for i := 0; i < m+n; i++ {
		g.AddVertex(graph.Vertex(i))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			g.AddEdge(graph.Vertex(i), graph.Vertex(m+j))
		}
	}
	return g
}
