package algorithms

import "github.com/stephen-mcelhose/graphtheory/mst"

func ExampleKruskal() ([]mst.Edge, int) {
	g := mst.NewWeightedGraph()
	g.AddEdge(0, 1, 3)
	g.AddEdge(1, 2, 1)
	g.AddEdge(0, 2, 2)
	g.AddEdge(1, 3, 4)
	return mst.Kruskal(g)
}
