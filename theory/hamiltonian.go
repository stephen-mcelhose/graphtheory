package theory

import (
	"math/bits"

	"github.com/stephen-mcelhose/graphtheory/graph"
)

func HamiltonianPath(g graph.Graph) []graph.Vertex {
	verts := g.Vertices()
	n := len(verts)
	used := map[graph.Vertex]bool{}
	path := []graph.Vertex{}
	var dfs func(graph.Vertex) bool
	dfs = func(v graph.Vertex) bool {
		path = append(path, v)
		used[v] = true
		if len(path) == n {
			return true
		}
		for _, u := range g.Neighbors(v) {
			if !used[u] && dfs(u) {
				return true
			}
		}
		used[v] = false
		path = path[:len(path)-1]
		return false
	}
	for _, v := range verts {
		for k := range used {
			delete(used, k)
		}
		path = path[:0]
		if dfs(v) {
			return append([]graph.Vertex{}, path...)
		}
	}
	return nil
}

func HamiltonianCycle(g graph.Graph) []graph.Vertex {
	path := HamiltonianPath(g)
	if len(path) == 0 {
		return nil
	}
	if g.HasEdge(path[0], path[len(path)-1]) {
		return append(path, path[0])
	}
	return nil
}

func TournamentHamiltonianPath(g *graph.SimpleDigraph) []graph.Vertex {
	path := []graph.Vertex{}
	for _, v := range g.Vertices() {
		inserted := false
		for i := 0; i < len(path); i++ {
			if g.HasArc(v, path[i]) {
				path = append(path[:i], append([]graph.Vertex{v}, path[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			path = append(path, v)
		}
	}
	return path
}

// HamiltonianPathDP uses dynamic programming on subsets to decide and recover
// a Hamiltonian path. It is practical only for small graphs.
func HamiltonianPathDP(g graph.Graph) []graph.Vertex {
	verts := g.Vertices()
	n := len(verts)
	if n == 0 {
		return nil
	}
	if n > 20 {
		return nil
	}
	index := map[graph.Vertex]int{}
	for i, v := range verts {
		index[v] = i
	}
	parent := map[[2]int]int{}
	reachable := map[[2]int]bool{}
	for i := 0; i < n; i++ {
		reachable[[2]int{1 << i, i}] = true
		parent[[2]int{1 << i, i}] = -1
	}
	for mask := 1; mask < (1 << n); mask++ {
		for end := 0; end < n; end++ {
			state := [2]int{mask, end}
			if !reachable[state] {
				continue
			}
			for _, nbr := range g.Neighbors(verts[end]) {
				next := index[nbr]
				if mask&(1<<next) != 0 {
					continue
				}
				nextState := [2]int{mask | (1 << next), next}
				if !reachable[nextState] {
					reachable[nextState] = true
					parent[nextState] = end
				}
			}
		}
	}
	full := (1 << n) - 1
	end := -1
	for i := 0; i < n; i++ {
		if reachable[[2]int{full, i}] {
			end = i
			break
		}
	}
	if end == -1 {
		return nil
	}
	path := make([]graph.Vertex, bits.OnesCount(uint(full)))
	mask := full
	cur := end
	for i := len(path) - 1; i >= 0; i-- {
		path[i] = verts[cur]
		prev := parent[[2]int{mask, cur}]
		mask ^= 1 << cur
		cur = prev
		if cur == -1 && i != 0 {
			break
		}
	}
	return path
}

// HamiltonianCycleExactSearch uses exact backtracking search for a cycle.
func HamiltonianCycleExactSearch(g graph.Graph) []graph.Vertex {
	verts := g.Vertices()
	n := len(verts)
	if n == 0 {
		return nil
	}
	start := verts[0]
	used := map[graph.Vertex]bool{start: true}
	path := []graph.Vertex{start}
	var dfs func(graph.Vertex) bool
	dfs = func(v graph.Vertex) bool {
		if len(path) == n {
			return g.HasEdge(v, start)
		}
		for _, u := range g.Neighbors(v) {
			if used[u] {
				continue
			}
			used[u] = true
			path = append(path, u)
			if dfs(u) {
				return true
			}
			path = path[:len(path)-1]
			used[u] = false
		}
		return false
	}
	if dfs(start) {
		return append(path, start)
	}
	return nil
}
