package named_graphs

import (
	"github.com/stephen-mcelhose/graphtheory/algebra"
	"github.com/stephen-mcelhose/graphtheory/generators"
	"github.com/stephen-mcelhose/graphtheory/graph"
	"github.com/stephen-mcelhose/graphtheory/theory"
)

func ExampleHypercubeGraph() (order, size int, bipartite bool, treeCount int) {
	g := generators.Hypercube(3)
	return g.Order(), g.Size(), theory.IsBipartite(g), algebra.SpanningTreeCount(g)
}

func ExampleDeBruijnDigraph() (order, size int, eulerian bool, circuitLength int) {
	g := generators.DeBruijnDigraph(2, 3)
	circuit := theory.EulerianCircuitDigraph(g, 0)
	return g.Order(), g.Size(), theory.IsEulerianDigraph(g), len(circuit)
}

func ExampleTournamentGraph() []graph.Vertex {
	orient := make([][]bool, 4)
	for i := range orient {
		orient[i] = make([]bool, 4)
	}
	orient[0][1] = true
	orient[0][2] = true
	orient[0][3] = false
	orient[1][2] = true
	orient[1][3] = true
	orient[2][3] = true
	g := generators.TournamentFromUpperTriangle(4, orient)
	return theory.TournamentHamiltonianPath(g)
}
