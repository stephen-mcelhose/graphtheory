package named_graphs

import (
	"github.com/stephen-mcelhose/graphtheory/flow"
	"github.com/stephen-mcelhose/graphtheory/generators"
	"github.com/stephen-mcelhose/graphtheory/graph"
	"github.com/stephen-mcelhose/graphtheory/theory"
	"github.com/stephen-mcelhose/graphtheory/traversal"
)

func ExampleEmptyGraph() (order, size int, connected bool) {
	g := generators.Empty(5)
	return g.Order(), g.Size(), traversal.IsConnected(g)
}

func ExampleCompleteGraph() (order, size int, dirac bool, ore bool) {
	g := generators.Complete(5)
	return g.Order(), g.Size(), theory.DiracCondition(g), theory.OreCondition(g)
}

func ExamplePathGraph() (order, size int, centers []graph.Vertex, isTree bool) {
	g := generators.Path(6)
	return g.Order(), g.Size(), traversal.Centers(g), theory.IsTree(g)
}

func ExampleCycleGraph() (order, size int, bipartite bool) {
	g := generators.Cycle(6)
	return g.Order(), g.Size(), theory.IsBipartite(g)
}

func ExampleCompleteBipartiteGraph() (order, size int, bipartite bool, matchingSize int) {
	g := generators.CompleteBipartite(3, 4)
	left := []graph.Vertex{0, 1, 2}
	right := []graph.Vertex{3, 4, 5, 6}
	m := flow.MaxBipartiteMatching(g, left, right)
	return g.Order(), g.Size(), theory.IsBipartite(g), len(m)
}
