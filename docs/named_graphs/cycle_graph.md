# Cycle Graph

## Definition
The cycle graph #$#C_n#$# is formed by closing a path into a ring.

## Go example
```go
g := generators.Cycle(6)
bip := theory.IsBipartite(g)
```

## Why it matters
It is the standard closed walk example and fundamental for Eulerian and parity-based arguments.

## Expected properties
- size: #$#n#$#
- each degree is #$#2#$#
- Hamiltonian by definition
- bipartite iff #$#n#$# is even
