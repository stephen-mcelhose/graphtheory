package graph

import "sort"

// MultiDigraph is a lightweight directed multigraph with arc multiplicities.
type MultiDigraph struct {
	verts map[Vertex]struct{}
	mult  map[Vertex]map[Vertex]int
}

func NewMultiDigraph() *MultiDigraph {
	return &MultiDigraph{
		verts: map[Vertex]struct{}{},
		mult:  map[Vertex]map[Vertex]int{},
	}
}

func (g *MultiDigraph) AddVertex(v Vertex) {
	if g.verts == nil {
		g.verts = map[Vertex]struct{}{}
		g.mult = map[Vertex]map[Vertex]int{}
	}
	g.verts[v] = struct{}{}
	if _, ok := g.mult[v]; !ok {
		g.mult[v] = map[Vertex]int{}
	}
}

func (g *MultiDigraph) AddArc(u, v Vertex) {
	g.AddVertex(u)
	g.AddVertex(v)
	g.mult[u][v]++
}

func (g *MultiDigraph) Multiplicity(u, v Vertex) int {
	return g.mult[u][v]
}

func (g *MultiDigraph) Vertices() []Vertex {
	vs := make([]Vertex, 0, len(g.verts))
	for v := range g.verts {
		vs = append(vs, v)
	}
	sort.Slice(vs, func(i, j int) bool { return vs[i] < vs[j] })
	return vs
}

func (g *MultiDigraph) OutDegree(v Vertex) int {
	total := 0
	for _, m := range g.mult[v] {
		total += m
	}
	return total
}

func (g *MultiDigraph) InDegree(v Vertex) int {
	total := 0
	for _, u := range g.Vertices() {
		total += g.mult[u][v]
	}
	return total
}
