package graph

import "sort"

// SimpleDigraph is a directed simple graph backed by successor and predecessor sets.
type SimpleDigraph struct {
	out map[Vertex]map[Vertex]struct{}
	in  map[Vertex]map[Vertex]struct{}
}

func NewSimpleDigraph() *SimpleDigraph {
	return &SimpleDigraph{
		out: map[Vertex]map[Vertex]struct{}{},
		in:  map[Vertex]map[Vertex]struct{}{},
	}
}

func (g *SimpleDigraph) Clone() *SimpleDigraph {
	out := NewSimpleDigraph()
	for _, v := range g.Vertices() {
		out.AddVertex(v)
		for _, u := range g.Successors(v) {
			out.AddArc(v, u)
		}
	}
	return out
}

func (g *SimpleDigraph) AddVertex(v Vertex) {
	if g.out == nil {
		g.out = map[Vertex]map[Vertex]struct{}{}
		g.in = map[Vertex]map[Vertex]struct{}{}
	}
	if _, ok := g.out[v]; !ok {
		g.out[v] = map[Vertex]struct{}{}
	}
	if _, ok := g.in[v]; !ok {
		g.in[v] = map[Vertex]struct{}{}
	}
}

func (g *SimpleDigraph) AddArc(u, v Vertex) {
	g.AddVertex(u)
	g.AddVertex(v)
	g.out[u][v] = struct{}{}
	g.in[v][u] = struct{}{}
}

func (g *SimpleDigraph) RemoveArc(u, v Vertex) {
	if _, ok := g.out[u]; ok {
		delete(g.out[u], v)
	}
	if _, ok := g.in[v]; ok {
		delete(g.in[v], u)
	}
}

func (g *SimpleDigraph) HasVertex(v Vertex) bool {
	_, ok := g.out[v]
	return ok
}

func (g *SimpleDigraph) HasArc(u, v Vertex) bool {
	_, ok := g.out[u][v]
	return ok
}

func (g *SimpleDigraph) Successors(v Vertex) []Vertex {
	res := make([]Vertex, 0, len(g.out[v]))
	for u := range g.out[v] {
		res = append(res, u)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func (g *SimpleDigraph) Predecessors(v Vertex) []Vertex {
	res := make([]Vertex, 0, len(g.in[v]))
	for u := range g.in[v] {
		res = append(res, u)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func (g *SimpleDigraph) Vertices() []Vertex {
	vs := make([]Vertex, 0, len(g.out))
	for v := range g.out {
		vs = append(vs, v)
	}
	sort.Slice(vs, func(i, j int) bool { return vs[i] < vs[j] })
	return vs
}

func (g *SimpleDigraph) OutDegree(v Vertex) int {
	return len(g.out[v])
}

func (g *SimpleDigraph) InDegree(v Vertex) int {
	return len(g.in[v])
}

func (g *SimpleDigraph) Order() int {
	return len(g.out)
}

func (g *SimpleDigraph) Size() int {
	total := 0
	for _, v := range g.Vertices() {
		total += len(g.out[v])
	}
	return total
}

func (g *SimpleDigraph) Reverse() *SimpleDigraph {
	r := NewSimpleDigraph()
	for _, v := range g.Vertices() {
		r.AddVertex(v)
	}
	for _, u := range g.Vertices() {
		for _, v := range g.Successors(u) {
			r.AddArc(v, u)
		}
	}
	return r
}

func (g *SimpleDigraph) UnderlyingSimpleGraph() *SimpleGraph {
	out := NewSimpleGraph()
	for _, v := range g.Vertices() {
		out.AddVertex(v)
	}
	for _, u := range g.Vertices() {
		for _, v := range g.Successors(u) {
			out.AddEdge(u, v)
		}
	}
	return out
}
