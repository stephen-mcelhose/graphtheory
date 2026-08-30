package theory

import "github.com/stephen-mcelhose/graphtheory/graph"

func IsEulerianUndirected(g *graph.MultiGraph) bool {
	for _, v := range g.Vertices() {
		if g.Degree(v)%2 != 0 {
			return false
		}
	}
	return true
}

// EulerianCircuit constructs an Eulerian circuit of an undirected multigraph
// using Hierholzer's algorithm. Returns nil if no circuit exists.
func EulerianCircuit(g *graph.MultiGraph, start graph.Vertex) []graph.Vertex {
	if !IsEulerianUndirected(g) || !g.HasVertex(start) {
		return nil
	}
	work := g.Clone()
	stack := []graph.Vertex{start}
	circuit := []graph.Vertex{}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		nb := work.Neighbors(v)
		if len(nb) == 0 {
			circuit = append(circuit, v)
			stack = stack[:len(stack)-1]
		} else {
			u := nb[0]
			work.RemoveEdge(v, u)
			stack = append(stack, u)
		}
	}
	for i, j := 0, len(circuit)-1; i < j; i, j = i+1, j-1 {
		circuit[i], circuit[j] = circuit[j], circuit[i]
	}
	return circuit
}

func IsEulerianDigraph(g *graph.SimpleDigraph) bool {
	for _, v := range g.Vertices() {
		if g.InDegree(v) != g.OutDegree(v) {
			return false
		}
	}
	return true
}

func EulerianCircuitDigraph(g *graph.SimpleDigraph, start graph.Vertex) []graph.Vertex {
	if !IsEulerianDigraph(g) || !g.HasVertex(start) {
		return nil
	}
	work := g.Clone()
	stack := []graph.Vertex{start}
	circuit := []graph.Vertex{}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		succ := work.Successors(v)
		if len(succ) == 0 {
			circuit = append(circuit, v)
			stack = stack[:len(stack)-1]
		} else {
			u := succ[0]
			work.RemoveArc(v, u)
			stack = append(stack, u)
		}
	}
	for i, j := 0, len(circuit)-1; i < j; i, j = i+1, j-1 {
		circuit[i], circuit[j] = circuit[j], circuit[i]
	}
	return circuit
}
