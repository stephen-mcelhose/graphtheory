package graph

import "sort"

// SimpleGraph is an undirected simple graph backed by adjacency sets.
type SimpleGraph struct {
	adj map[Vertex]map[Vertex]struct{}
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{adj: map[Vertex]map[Vertex]struct{}{}}
}

func (g *SimpleGraph) Clone() *SimpleGraph {
	out := NewSimpleGraph()
	for _, v := range g.Vertices() {
		out.AddVertex(v)
		for _, u := range g.Neighbors(v) {
			if v < u {
				out.AddEdge(v, u)
			}
		}
	}
	return out
}

func (g *SimpleGraph) AddVertex(v Vertex) {
	if g.adj == nil {
		g.adj = map[Vertex]map[Vertex]struct{}{}
	}
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = map[Vertex]struct{}{}
	}
}

func (g *SimpleGraph) AddEdge(u, v Vertex) {
	if u == v {
		return
	}
	g.AddVertex(u)
	g.AddVertex(v)
	g.adj[u][v] = struct{}{}
	g.adj[v][u] = struct{}{}
}

func (g *SimpleGraph) RemoveEdge(u, v Vertex) {
	if _, ok := g.adj[u]; ok {
		delete(g.adj[u], v)
	}
	if _, ok := g.adj[v]; ok {
		delete(g.adj[v], u)
	}
}

func (g *SimpleGraph) RemoveVertex(v Vertex) {
	for u := range g.adj[v] {
		delete(g.adj[u], v)
	}
	delete(g.adj, v)
}

func (g *SimpleGraph) HasVertex(v Vertex) bool {
	_, ok := g.adj[v]
	return ok
}

func (g *SimpleGraph) HasEdge(u, v Vertex) bool {
	_, ok := g.adj[u][v]
	return ok
}

func (g *SimpleGraph) Neighbors(v Vertex) []Vertex {
	nbrs := make([]Vertex, 0, len(g.adj[v]))
	for u := range g.adj[v] {
		nbrs = append(nbrs, u)
	}
	sort.Slice(nbrs, func(i, j int) bool { return nbrs[i] < nbrs[j] })
	return nbrs
}

func (g *SimpleGraph) Vertices() []Vertex {
	vs := make([]Vertex, 0, len(g.adj))
	for v := range g.adj {
		vs = append(vs, v)
	}
	sort.Slice(vs, func(i, j int) bool { return vs[i] < vs[j] })
	return vs
}

func (g *SimpleGraph) Degree(v Vertex) int {
	return len(g.adj[v])
}

func (g *SimpleGraph) Order() int {
	return len(g.adj)
}

func (g *SimpleGraph) Size() int {
	total := 0
	for _, v := range g.Vertices() {
		total += len(g.adj[v])
	}
	return total / 2
}

func (g *SimpleGraph) Edges() [][2]Vertex {
	edges := make([][2]Vertex, 0, g.Size())
	for _, u := range g.Vertices() {
		for _, v := range g.Neighbors(u) {
			if u < v {
				edges = append(edges, [2]Vertex{u, v})
			}
		}
	}
	return edges
}

func (g *SimpleGraph) Induced(vertices []Vertex) *SimpleGraph {
	want := map[Vertex]struct{}{}
	for _, v := range vertices {
		if g.HasVertex(v) {
			want[v] = struct{}{}
		}
	}
	out := NewSimpleGraph()
	for v := range want {
		out.AddVertex(v)
	}
	for u := range want {
		for v := range want {
			if u < v && g.HasEdge(u, v) {
				out.AddEdge(u, v)
			}
		}
	}
	return out
}
