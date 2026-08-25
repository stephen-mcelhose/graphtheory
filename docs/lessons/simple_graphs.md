# Lesson: Simple Graphs

## Definition
A simple graph has a finite vertex set and an edge set consisting of unordered pairs of distinct vertices.

## Intuition
Simple graphs encode symmetric, pairwise relationships without parallel edges or self-loops.

## Tiny example
```text
V = {0,1,2,3}
E = {(0,1),(1,2),(2,3)}
```

## In code
- `graph.SimpleGraph`
- `generators.Empty`
- `generators.Complete`
- `generators.Path`
- `generators.Cycle`

## Common confusion
Adjacency is a relation between two vertices. Degree is a count of how many neighbors one vertex has.

## Status
Implemented constructively.

## Exercise prompt
Construct a simple graph on 5 vertices with exactly 4 edges and no cycles.
