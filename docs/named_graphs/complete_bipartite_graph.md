# Complete Bipartite Graph

## Definition
The complete bipartite graph #$#K_{m,n}#$# splits vertices into two parts and joins every cross-part pair.

## Go example
```go
g := generators.CompleteBipartite(3, 4)
left := []graph.Vertex{0, 1, 2}
right := []graph.Vertex{3, 4, 5, 6}
m := flow.MaxBipartiteMatching(g, left, right)
```

## Why it matters
It is central for Hall's theorem, bipartite matching, and flow reductions.

## Expected properties
- size: #$#mn#$#
- bipartite by definition
- maximum matching size is #$#\min(m,n)#$#
- spanning-tree count is #$#m^{n-1} n^{m-1}#$#
