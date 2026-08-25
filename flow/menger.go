package flow

import "github.com/stephen-mcelhose/graphtheory/graph"

// EdgeConnectivityUpperBound computes the maximum number of edge-disjoint
// s-t paths in an undirected simple graph by reducing to unit-capacity flow.
func EdgeConnectivityUpperBound(g *graph.SimpleGraph, s, t graph.Vertex) int {
	n := graph.NewNetwork()
	for _, v := range g.Vertices() {
		n.AddVertex(v)
	}
	for _, e := range g.Edges() {
		u, v := e[0], e[1]
		n.AddArc(u, v, 1)
		n.AddArc(v, u, 1)
	}
	return EdmondsKarp(n, s, t).Value
}
