// Package mst provides minimum spanning tree algorithms for weighted
// undirected graphs.
package mst

import (
	"sort"

	"github.com/stephen-mcelhose/graphtheory/graph"
)

// WeightedGraph is an undirected graph with integer edge weights.
type WeightedGraph struct {
	verts   map[graph.Vertex]struct{}
	weights map[graph.Vertex]map[graph.Vertex]int
}

// Edge represents a weighted undirected edge with u < v.
type Edge struct {
	U, V   graph.Vertex
	Weight int
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		verts:   map[graph.Vertex]struct{}{},
		weights: map[graph.Vertex]map[graph.Vertex]int{},
	}
}

func (g *WeightedGraph) AddVertex(v graph.Vertex) {
	g.verts[v] = struct{}{}
	if _, ok := g.weights[v]; !ok {
		g.weights[v] = map[graph.Vertex]int{}
	}
}

func (g *WeightedGraph) AddEdge(u, v graph.Vertex, w int) {
	if u == v {
		return
	}
	g.AddVertex(u)
	g.AddVertex(v)
	if u > v {
		u, v = v, u
	}
	g.weights[u][v] = w
}

func (g *WeightedGraph) Weight(u, v graph.Vertex) (int, bool) {
	if u > v {
		u, v = v, u
	}
	w, ok := g.weights[u][v]
	return w, ok
}

func (g *WeightedGraph) HasEdge(u, v graph.Vertex) bool {
	if u > v {
		u, v = v, u
	}
	_, ok := g.weights[u][v]
	return ok
}

func (g *WeightedGraph) Neighbors(v graph.Vertex) []graph.Vertex {
	res := []graph.Vertex{}
	for u := range g.weights[v] {
		res = append(res, u)
	}
	for u, m := range g.weights {
		if _, ok := m[v]; ok && u != v {
			res = append(res, u)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func (g *WeightedGraph) Vertices() []graph.Vertex {
	vs := make([]graph.Vertex, 0, len(g.verts))
	for v := range g.verts {
		vs = append(vs, v)
	}
	sort.Slice(vs, func(i, j int) bool { return vs[i] < vs[j] })
	return vs
}

func (g *WeightedGraph) HasVertex(v graph.Vertex) bool {
	_, ok := g.verts[v]
	return ok
}

func (g *WeightedGraph) Degree(v graph.Vertex) int {
	return len(g.Neighbors(v))
}

func (g *WeightedGraph) Order() int { return len(g.verts) }

func (g *WeightedGraph) Size() int {
	total := 0
	for _, m := range g.weights {
		total += len(m)
	}
	return total
}

func (g *WeightedGraph) Edges() []Edge {
	edges := []Edge{}
	for u, m := range g.weights {
		for v, w := range m {
			edges = append(edges, Edge{U: u, V: v, Weight: w})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].Weight != edges[j].Weight {
			return edges[i].Weight < edges[j].Weight
		}
		if edges[i].U != edges[j].U {
			return edges[i].U < edges[j].U
		}
		return edges[i].V < edges[j].V
	})
	return edges
}
