package graph

import "sort"

// MultiGraph is a lightweight undirected multigraph with edge multiplicities.
type MultiGraph struct {
	verts map[Vertex]struct{}
	mult  map[Vertex]map[Vertex]int
}

func NewMultiGraph() *MultiGraph {
	return &MultiGraph{
		verts: map[Vertex]struct{}{},
		mult:  map[Vertex]map[Vertex]int{},
	}
}

func (g *MultiGraph) AddVertex(v Vertex) {
	if g.verts == nil {
		g.verts = map[Vertex]struct{}{}
		g.mult = map[Vertex]map[Vertex]int{}
	}
	g.verts[v] = struct{}{}
	if _, ok := g.mult[v]; !ok {
		g.mult[v] = map[Vertex]int{}
	}
}

func (g *MultiGraph) AddEdge(u, v Vertex) {
	if u == v {
		return
	}
	g.AddVertex(u)
	g.AddVertex(v)
	g.mult[u][v]++
	g.mult[v][u]++
}

func (g *MultiGraph) Multiplicity(u, v Vertex) int {
	return g.mult[u][v]
}

func (g *MultiGraph) Vertices() []Vertex {
	vs := make([]Vertex, 0, len(g.verts))
	for v := range g.verts {
		vs = append(vs, v)
	}
	sort.Slice(vs, func(i, j int) bool { return vs[i] < vs[j] })
	return vs
}

func (g *MultiGraph) Degree(v Vertex) int {
	total := 0
	for _, m := range g.mult[v] {
		total += m
	}
	return total
}

func (g *MultiGraph) HasVertex(v Vertex) bool {
	_, ok := g.verts[v]
	return ok
}

func (g *MultiGraph) Order() int { return len(g.verts) }

func (g *MultiGraph) Size() int {
	total := 0
	for _, u := range g.Vertices() {
		for _, v := range g.Vertices() {
			if u < v {
				total += g.mult[u][v]
			}
		}
	}
	return total
}

func (g *MultiGraph) Neighbors(v Vertex) []Vertex {
	res := make([]Vertex, 0, len(g.mult[v]))
	for u, m := range g.mult[v] {
		if m > 0 {
			res = append(res, u)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func (g *MultiGraph) RemoveEdge(u, v Vertex) {
	if g.mult[u][v] > 0 {
		g.mult[u][v]--
		g.mult[v][u]--
	}
}

func (g *MultiGraph) Clone() *MultiGraph {
	c := NewMultiGraph()
	for _, v := range g.Vertices() {
		c.AddVertex(v)
	}
	for _, u := range g.Vertices() {
		for _, v := range g.Vertices() {
			if u < v {
				for k := 0; k < g.Multiplicity(u, v); k++ {
					c.AddEdge(u, v)
				}
			}
		}
	}
	return c
}
