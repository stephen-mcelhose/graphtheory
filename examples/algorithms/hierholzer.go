package algorithms

import (
	"github.com/stephen-mcelhose/graphtheory/graph"
	"github.com/stephen-mcelhose/graphtheory/theory"
)

// ExampleHierholzerUndirected builds a small Eulerian multigraph and returns
// the Eulerian circuit found by Hierholzer's algorithm.
func ExampleHierholzerUndirected() []graph.Vertex {
	g := graph.NewMultiGraph()
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 0)
	return theory.EulerianCircuit(g, 0)
}

// ExampleHierholzerDigraph builds a small Eulerian digraph and returns
// the directed Eulerian circuit found by Hierholzer's algorithm.
func ExampleHierholzerDigraph() []graph.Vertex {
	g := graph.NewSimpleDigraph()
	g.AddArc(0, 1)
	g.AddArc(1, 2)
	g.AddArc(2, 0)
	return theory.EulerianCircuitDigraph(g, 0)
}
