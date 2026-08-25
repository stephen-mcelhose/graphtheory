package algebra

import "github.com/stephen-mcelhose/graphtheory/graph"

func AdjacencyMatrix(g graph.Graph) (verts []graph.Vertex, m [][]int) {
	verts = g.Vertices()
	index := map[graph.Vertex]int{}
	for i, v := range verts {
		index[v] = i
	}
	n := len(verts)
	m = make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	for _, u := range verts {
		for _, v := range g.Neighbors(u) {
			m[index[u]][index[v]] = 1
		}
	}
	return verts, m
}

func DigraphAdjacencyMatrix(g *graph.SimpleDigraph) (verts []graph.Vertex, m [][]int) {
	verts = g.Vertices()
	index := map[graph.Vertex]int{}
	for i, v := range verts {
		index[v] = i
	}
	n := len(verts)
	m = make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	for _, u := range verts {
		for _, v := range g.Successors(u) {
			m[index[u]][index[v]] = 1
		}
	}
	return verts, m
}

func LaplacianMatrix(g graph.Graph) (verts []graph.Vertex, l [][]int) {
	verts, a := AdjacencyMatrix(g)
	n := len(verts)
	l = make([][]int, n)
	for i := 0; i < n; i++ {
		l[i] = make([]int, n)
		deg := 0
		for j := 0; j < n; j++ {
			if i != j {
				l[i][j] = -a[i][j]
			}
			deg += a[i][j]
		}
		l[i][i] = deg
	}
	return verts, l
}

func DirectedLaplacian(g *graph.SimpleDigraph) (verts []graph.Vertex, l [][]int) {
	verts = g.Vertices()
	index := map[graph.Vertex]int{}
	for i, v := range verts {
		index[v] = i
	}
	n := len(verts)
	l = make([][]int, n)
	for i := range l {
		l[i] = make([]int, n)
	}
	for _, u := range verts {
		i := index[u]
		l[i][i] = g.OutDegree(u)
		for _, v := range g.Successors(u) {
			j := index[v]
			l[i][j]--
		}
	}
	return verts, l
}
