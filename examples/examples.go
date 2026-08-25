package examples

import (
	"github.com/stephen-mcelhose/graphtheory/algebra"
	"github.com/stephen-mcelhose/graphtheory/flow"
	"github.com/stephen-mcelhose/graphtheory/generators"
	"github.com/stephen-mcelhose/graphtheory/graph"
	"github.com/stephen-mcelhose/graphtheory/theory"
	"github.com/stephen-mcelhose/graphtheory/traversal"
)

func ExampleSimpleConnectivity() (components int, centers []graph.Vertex) {
	g := generators.Path(5)
	return len(traversal.ConnectedComponents(g)), traversal.Centers(g)
}

func ExampleSpanningTreeCount() int {
	g := generators.Complete(4)
	return algebra.SpanningTreeCount(g)
}

func ExampleBipartiteMatching() int {
	g := graph.NewSimpleGraph()
	left := []graph.Vertex{0, 1, 2}
	right := []graph.Vertex{3, 4, 5}
	for _, v := range append(append([]graph.Vertex{}, left...), right...) {
		g.AddVertex(v)
	}
	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 5)
	m := flow.MaxBipartiteMatching(g, left, right)
	return len(m) / 2
}

func ExampleMaxFlow() int {
	n := graph.NewNetwork()
	n.AddArc(0, 1, 3)
	n.AddArc(0, 2, 2)
	n.AddArc(1, 2, 1)
	n.AddArc(1, 3, 2)
	n.AddArc(2, 3, 4)
	return flow.EdmondsKarp(n, 0, 3).Value
}

func ExampleDeBruijnEulerian() []graph.Vertex {
	d := generators.DeBruijnDigraph(2, 3)
	return theory.EulerianCircuitDigraph(d, 0)
}
