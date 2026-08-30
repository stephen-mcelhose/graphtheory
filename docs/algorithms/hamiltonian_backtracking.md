# Hamiltonian Backtracking

## Purpose
Backtracking is the standard exact-search baseline for Hamiltonian path and cycle problems.

## Problem family
This is a **vertex-traversal** exact search method.

## In this project
- Code: `theory/hamiltonian.go` (`HamiltonianPath`, `HamiltonianCycle`)
- Example: `examples/algorithms/hamiltonian.go`

## Intuition
Grow a partial path one vertex at a time. If the current choice cannot be extended to a full Hamiltonian solution, backtrack and try another branch.
