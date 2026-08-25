# Hypercube Graph

## Definition
The hypercube #$#Q_d#$# has binary strings of length #$#d#$# as vertices, with edges between strings differing in one coordinate.

## Go example
```go
g := generators.Hypercube(3)
count := algebra.SpanningTreeCount(g)
```

## Why it matters
It is a highly symmetric family used in Hamiltonicity, bipartiteness, and recursive constructions.

## Expected properties
- order: #$#2^d#$#
- size: #$#d 2^{d-1}#$#
- each degree is #$#d#$#
- bipartite
