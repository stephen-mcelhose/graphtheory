package algorithms

import (
	"github.com/stephen-mcelhose/graphtheory/generators"
	"github.com/stephen-mcelhose/graphtheory/graph"
	"github.com/stephen-mcelhose/graphtheory/theory"
)

func ExampleHamiltonianBacktracking() []graph.Vertex {
	g := generators.Cycle(5)
	return theory.HamiltonianCycle(g)
}

func ExampleHamiltonianDP() []graph.Vertex {
	g := generators.Path(5)
	return theory.HamiltonianPathDP(g)
}

func ExampleHamiltonianExactSearch() []graph.Vertex {
	g := generators.Cycle(6)
	return theory.HamiltonianCycleExactSearch(g)
}
