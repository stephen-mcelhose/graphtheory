# graphtheory

`graphtheory` is a Go library blueprint and reference implementation inspired by Darij Grinberg's *An introduction to graph theory* (arXiv:2308.04512).

It is designed to cover the paper's material in a practical way, with a focus on:

- graph data models
- traversals and connectivity
- trees and arborescences
- Eulerian and Hamiltonian utilities
- bipartite graphs, matchings, and flows
- adjacency matrices, Laplacians, and Matrix-Tree style counting helpers
- classical graph generators
- theorem-checking helpers for educational use

## Package layout

- `graph`: core graph and digraph data structures
- `traversal`: BFS, DFS, components, SCCs, bridges, distances
- `flow`: max flow and bipartite matching
- `algebra`: adjacency matrices, Laplacians, determinant and spanning-tree counting
- `generators`: complete/path/cycle/bipartite/hypercube/de Bruijn graph generators
- `theory`: Eulerian, Hamiltonian, tree, bipartite, and degree-condition helpers
- `examples`: small examples showing how modules fit together
- `examples/named_graphs`: named graph family examples
- `docs`: concept map, library plan, lesson plan, sources, and named-graph writeups
- `learning`: recommended per-student learning state layout

## Scope note

The source paper is course-scale and theorem-heavy. Not every theorem maps directly to an efficient general-purpose algorithm. This library therefore distinguishes between:

- polynomial-time constructive algorithms
- exact small-instance algorithms
- theorem-checking utilities and educational helpers

## Named graph coverage

The library now includes dedicated examples and write-ups for these named graph families:
- empty graph
- complete graph
- path graph
- cycle graph
- complete bipartite graph
- hypercube graph
- de Bruijn digraph
- tournament

See:
- `examples/named_graphs/`
- `docs/named_graphs/`

## Learning aid additions
- `docs/learning_path.md`
- `docs/chapters/`
- `docs/lesson_plan.md`
- `docs/research_to_implementation_gap.md`
- `docs/lessons/`
- `docs/exercises/`
- `docs/projects/`
- `docs/sources.md`
- `learning/README.md`

## Status

This implementation is intentionally self-contained and standard-library-only so it is easy to inspect and extend in constrained environments.
