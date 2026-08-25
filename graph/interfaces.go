package graph

// Vertex is the canonical vertex identifier used across the library.
type Vertex int

// Graph is a minimal interface for simple undirected graph algorithms.
type Graph interface {
	Vertices() []Vertex
	HasVertex(v Vertex) bool
	HasEdge(u, v Vertex) bool
	Neighbors(v Vertex) []Vertex
	Degree(v Vertex) int
	Order() int
	Size() int
}

// Digraph is a minimal interface for directed graph algorithms.
type Digraph interface {
	Vertices() []Vertex
	HasVertex(v Vertex) bool
	HasArc(u, v Vertex) bool
	Successors(v Vertex) []Vertex
	Predecessors(v Vertex) []Vertex
	OutDegree(v Vertex) int
	InDegree(v Vertex) int
	Order() int
	Size() int
}

// WeightedDigraph extends Digraph with capacities or weights.
type WeightedDigraph interface {
	Digraph
	Weight(u, v Vertex) (int, bool)
}
