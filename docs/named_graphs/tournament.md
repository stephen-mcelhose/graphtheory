# Tournament

## Definition
A tournament is an orientation of a complete graph, with exactly one directed edge chosen for each unordered pair.

## Go example
```go
g := generators.TournamentFromUpperTriangle(n, orient)
path := theory.TournamentHamiltonianPath(g)
```

## Why it matters
It is a named directed graph family with strong Hamiltonian path structure and important extremal behavior.

## Expected properties
- exactly one arc for each unordered pair
- every tournament has a Hamiltonian path
- strong tournaments have Hamiltonian cycles
