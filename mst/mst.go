package mst

import (
	"sort"

	"github.com/stephen-mcelhose/graphtheory/graph"
)

type disjointSet struct {
	parent map[graph.Vertex]graph.Vertex
	rank   map[graph.Vertex]int
}

func newDisjointSet(vertices []graph.Vertex) *disjointSet {
	d := &disjointSet{parent: map[graph.Vertex]graph.Vertex{}, rank: map[graph.Vertex]int{}}
	for _, v := range vertices {
		d.parent[v] = v
		d.rank[v] = 0
	}
	return d
}

func (d *disjointSet) find(v graph.Vertex) graph.Vertex {
	if d.parent[v] != v {
		d.parent[v] = d.find(d.parent[v])
	}
	return d.parent[v]
}

func (d *disjointSet) union(a, b graph.Vertex) bool {
	ra, rb := d.find(a), d.find(b)
	if ra == rb {
		return false
	}
	if d.rank[ra] < d.rank[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	if d.rank[ra] == d.rank[rb] {
		d.rank[ra]++
	}
	return true
}

// Kruskal returns the MST edges and total weight using Kruskal's algorithm.
func Kruskal(g *WeightedGraph) ([]Edge, int) {
	edges := g.Edges()
	dsu := newDisjointSet(g.Vertices())
	mst := []Edge{}
	total := 0
	for _, e := range edges {
		if dsu.union(e.U, e.V) {
			mst = append(mst, e)
			total += e.Weight
		}
	}
	return mst, total
}

// Prim returns the MST edges and total weight using Prim's algorithm.
func Prim(g *WeightedGraph, start graph.Vertex) ([]Edge, int) {
	if !g.HasVertex(start) {
		return nil, 0
	}
	inTree := map[graph.Vertex]bool{start: true}
	mst := []Edge{}
	total := 0
	for len(inTree) < g.Order() {
		best := Edge{Weight: 1 << 30}
		found := false
		for u := range inTree {
			for _, v := range g.Neighbors(u) {
				if inTree[v] {
					continue
				}
				w, _ := g.Weight(u, v)
				e := Edge{U: u, V: v, Weight: w}
				if e.U > e.V {
					e.U, e.V = e.V, e.U
				}
				if !found || e.Weight < best.Weight || (e.Weight == best.Weight && (e.U < best.U || (e.U == best.U && e.V < best.V))) {
					best = e
					found = true
				}
			}
		}
		if !found {
			return nil, 0
		}
		inTree[best.U] = true
		inTree[best.V] = true
		mst = append(mst, best)
		total += best.Weight
	}
	sort.Slice(mst, func(i, j int) bool {
		if mst[i].U != mst[j].U {
			return mst[i].U < mst[j].U
		}
		return mst[i].V < mst[j].V
	})
	return mst, total
}
