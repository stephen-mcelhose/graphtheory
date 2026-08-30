# Hierholzer's Algorithm

## Purpose
Hierholzer's algorithm constructs an **Eulerian circuit** when one exists.

## Problem family
This is an **edge-traversal** algorithm.
It belongs to the Eulerian family of problems, not the minimum spanning tree family.

## What it solves
Given a graph that satisfies the Eulerian conditions, Hierholzer's algorithm builds a circuit that uses every edge exactly once.

For a connected undirected graph, an Eulerian circuit exists when every vertex has even degree.
For a directed graph, an Eulerian circuit exists when every vertex has equal in-degree and out-degree, together with the appropriate connectivity condition.

## Intuition
The algorithm:
1. starts walking along unused edges
2. forms a cycle
3. splices in additional cycles if unused edges remain
4. returns one full Eulerian circuit

## In this project
Code locations:
- `theory/eulerian.go`
  - `EulerianCircuit` for undirected multigraphs
  - `EulerianCircuitDigraph` for digraphs
  - `IsEulerianUndirected`
  - `IsEulerianDigraph`

Examples:
- `examples/algorithms/hierholzer.go`

## Why it matters
Hierholzer's algorithm is the stereotypical constructive algorithm for Eulerian circuit problems.
It is a good contrast with:
- **Kruskal's algorithm** and **Prim's algorithm**, which solve minimum spanning tree problems
- Hamiltonian problems, which are about visiting every vertex exactly once

## Tiny example
Undirected 4-cycle:
```text
0 -- 1
|    |
3 -- 2
```

One Eulerian circuit is:
```text
0,1,2,3,0
```

## Related terms
- Eulerian path
- Eulerian circuit
- edge-traversal problem
- constructive algorithm
