# Empty Graph

## Definition
The empty graph on #$#n#$# vertices has no edges.

## Go example
```go
g := generators.Empty(5)
order, size := g.Order(), g.Size()
```

## Why it matters
It is the sparsest simple graph family and a baseline object for connectivity, independence, and coloring.

## Expected properties
- order: #$#n#$#
- size: #$#0#$#
- all degrees are #$#0#$#
- connected iff #$#n \le 1#$#
