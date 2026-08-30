# Hamiltonian Dynamic Programming on Subsets

## Purpose
Dynamic programming on subsets is a classic exact method for Hamiltonian path and cycle problems on small graphs.

## Problem family
This is a **vertex-traversal** exact algorithm based on subset states.

## In this project
- Code: `theory/hamiltonian.go` (`HamiltonianPathDP`)
- Example: `examples/algorithms/hamiltonian.go`

## Intuition
Track whether there is a Hamiltonian path that visits exactly a subset of vertices and ends at a chosen final vertex.
