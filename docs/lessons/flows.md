# Lesson: Flows

## Definition
A flow assigns values to directed edges subject to capacity bounds and conservation constraints.

## Intuition
Think of water moving through a network from a source to a sink. The max-flow problem asks how much can be pushed through.

## Tiny example
```text
0 -> 1 (3)
0 -> 2 (2)
1 -> 3 (2)
2 -> 3 (4)
```

## In code
- `graph.Network`
- `flow.EdmondsKarp`

## Common confusion
A path in the graph is not the same thing as a feasible flow. Flow can split across multiple paths.

## Status
Implemented constructively.

## Exercise prompt
Build a 4-node network and compute a max flow by reasoning about bottlenecks.
