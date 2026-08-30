# Path Graph

## Definition
The path graph $P_n$ is a chain of $n$ vertices with consecutive edges.

## Go example
```go
g := generators.Path(6)
centers := traversal.Centers(g)
```

## Why it matters
It is the simplest nontrivial tree and a core example for traversal, distance, and center computations.

## Expected properties
- size: $n-1$
- it is a tree
- it is bipartite
- it has one center when $n$ is odd, two centers when $n$ is even
