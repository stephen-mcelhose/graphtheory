# Lesson: Eulerian Graphs

## Definition
A graph is Eulerian when its edges can be traversed exactly once in a closed walk.

## Intuition
This is an "use every edge once" problem, not an "visit every vertex once" problem. That distinction is the main contrast with Hamiltonian problems.

## Tiny example
Edge list:
```text
V = {0,1,2}
E = {(0,1),(1,2),(2,0)}
```
This graph is Eulerian because each vertex has even degree.

## In code
- `theory.IsEulerianDigraph`
- `theory.EulerianCircuitDigraph`

## Worked code sketch
```go
g := generators.DeBruijnDigraph(2, 3)
ok := theory.IsEulerianDigraph(g)
circuit := theory.EulerianCircuitDigraph(g, 0)
```

## Common confusion
A Hamiltonian cycle uses every vertex once. An Eulerian circuit uses every edge once. These are different goals.

## Status
Implemented constructively.

## Exercise prompt
Explain why a cycle graph #$#C_n#$# is Eulerian, and compare that to whether it is bipartite.
