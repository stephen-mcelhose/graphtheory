package generators

import "github.com/stephen-mcelhose/graphtheory/graph"

// TournamentFromUpperTriangle builds a tournament on vertices 0..n-1.
// orient[i][j] for i<j: true means i->j, false means j->i.
func TournamentFromUpperTriangle(n int, orient [][]bool) *graph.SimpleDigraph {
	g := graph.NewSimpleDigraph()
	for i := 0; i < n; i++ {
		g.AddVertex(graph.Vertex(i))
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if orient[i][j] {
				g.AddArc(graph.Vertex(i), graph.Vertex(j))
			} else {
				g.AddArc(graph.Vertex(j), graph.Vertex(i))
			}
		}
	}
	return g
}
