# Greedy Algorithm

## Definition
A **greedy algorithm** makes the locally optimal choice at each step, without backtracking, in the hope of finding a globally optimal solution.

## Intuition
Instead of planning ahead, a greedy algorithm always takes the best immediate option. For some problems this produces a provably optimal result; for others it produces a good approximation or may fail.

## Key facts
- Greedy algorithms are typically simple and fast.
- They work optimally when the problem has **matroid structure** or satisfies a **greedy-choice property**.
- Examples where greedy is optimal: Kruskal's MST, Prim's MST, Dijkstra's shortest paths (non-negative weights).
- Examples where greedy is not always optimal: general shortest path with negative weights, some scheduling problems.

## Related concepts
- **depends on**: weighted graph, spanning tree
- **used by**: Kruskal's algorithm, Prim's algorithm, Dijkstra's algorithm
- **contrasts with**: dynamic programming, backtracking
- **see also**: minimum spanning tree, constructive algorithm
