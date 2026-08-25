package graph

import "sort"

// Network is a capacitated digraph for flow algorithms.
type Network struct {
	verts map[Vertex]struct{}
	cap   map[Vertex]map[Vertex]int
}

func NewNetwork() *Network {
	return &Network{
		verts: map[Vertex]struct{}{},
		cap:   map[Vertex]map[Vertex]int{},
	}
}

func (n *Network) AddVertex(v Vertex) {
	if n.verts == nil {
		n.verts = map[Vertex]struct{}{}
		n.cap = map[Vertex]map[Vertex]int{}
	}
	n.verts[v] = struct{}{}
	if _, ok := n.cap[v]; !ok {
		n.cap[v] = map[Vertex]int{}
	}
}

func (n *Network) AddArc(u, v Vertex, capacity int) {
	if capacity < 0 {
		capacity = 0
	}
	n.AddVertex(u)
	n.AddVertex(v)
	n.cap[u][v] = capacity
}

func (n *Network) Vertices() []Vertex {
	vs := make([]Vertex, 0, len(n.verts))
	for v := range n.verts {
		vs = append(vs, v)
	}
	sort.Slice(vs, func(i, j int) bool { return vs[i] < vs[j] })
	return vs
}

func (n *Network) HasVertex(v Vertex) bool {
	_, ok := n.verts[v]
	return ok
}

func (n *Network) HasArc(u, v Vertex) bool {
	_, ok := n.cap[u][v]
	return ok
}

func (n *Network) Successors(v Vertex) []Vertex {
	res := make([]Vertex, 0, len(n.cap[v]))
	for u := range n.cap[v] {
		res = append(res, u)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func (n *Network) Predecessors(v Vertex) []Vertex {
	res := []Vertex{}
	for u := range n.verts {
		if _, ok := n.cap[u][v]; ok {
			res = append(res, u)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func (n *Network) OutDegree(v Vertex) int { return len(n.cap[v]) }
func (n *Network) InDegree(v Vertex) int  { return len(n.Predecessors(v)) }
func (n *Network) Order() int             { return len(n.verts) }

func (n *Network) Size() int {
	total := 0
	for _, v := range n.Vertices() {
		total += len(n.cap[v])
	}
	return total
}

func (n *Network) Weight(u, v Vertex) (int, bool) {
	w, ok := n.cap[u][v]
	return w, ok
}

func (n *Network) Capacity(u, v Vertex) int {
	if w, ok := n.cap[u][v]; ok {
		return w
	}
	return 0
}
