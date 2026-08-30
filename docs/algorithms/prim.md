# Prim's Algorithm

## Purpose
Prim's algorithm constructs a **minimum spanning tree** of a weighted connected graph.

## Problem family
This is a **greedy optimization** algorithm for MST problems, not a Hamiltonian traversal algorithm.

## Intuition
Start from one vertex and repeatedly add the cheapest edge that connects the current tree to a new vertex.

## In this project
- Code: `mst/mst.go` (`Prim`)
- Example: `examples/algorithms/prim.go`

## Related terms
- minimum spanning tree
- greedy algorithm
- weighted graph
- Kruskal's algorithm
