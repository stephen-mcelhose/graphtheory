# Kruskal's Algorithm

## Purpose
Kruskal's algorithm constructs a **minimum spanning tree** of a weighted connected graph.

## Problem family
This is a **greedy optimization** algorithm for MST problems, not an Eulerian or Hamiltonian traversal algorithm.

## Intuition
Sort all edges by weight and keep taking the cheapest edge that does not create a cycle.

## In this project
- Code: `mst/mst.go` (`Kruskal`)
- Example: `examples/algorithms/kruskal.go`

## Tiny example
Edges: `(1,2):1`, `(0,2):2`, `(0,1):3`, `(1,3):4`

Chosen by Kruskal:
- `(1,2)`
- `(0,2)`
- skip `(0,1)` because it closes a cycle
- `(1,3)`

Total weight: `7`

## Related terms
- minimum spanning tree
- greedy algorithm
- weighted graph
- Prim's algorithm
