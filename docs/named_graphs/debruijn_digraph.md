# de Bruijn Digraph

## Definition
The #$#k#$#-ary order-#$#n#$# de Bruijn digraph encodes overlap transitions between length-#$#n-1#$# words.

## Go example
```go
g := generators.DeBruijnDigraph(2, 3)
circuit := theory.EulerianCircuitDigraph(g, 0)
```

## Why it matters
It links Eulerian circuits with cyclic sequence generation and appears explicitly in the paper.

## Expected properties
- order: #$#k^{n-1}#$#
- size: #$#k^n#$#
- each indegree and outdegree is #$#k#$#
- Eulerian
