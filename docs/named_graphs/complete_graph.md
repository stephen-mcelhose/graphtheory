# Complete Graph

## Definition
The complete graph $K_n$ joins every pair of distinct vertices.

## Go example
```go
g := generators.Complete(5)
order, size := g.Order(), g.Size()
```

## Why it matters
It is the densest simple graph family and appears in Hamiltonicity, Ramsey theory, and spanning-tree counting.

## Expected properties
- size: $\binom{n}{2}$
- each degree is $n-1$
- chromatic number is $n$
- Dirac and Ore conditions hold for $n \ge 3$
