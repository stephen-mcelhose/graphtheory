# Weighted Graph

## Definition
A **weighted graph** is a graph in which each edge is assigned a numerical value such as a cost, length, capacity, or similarity score.

## Intuition
Ordinary graphs say whether a connection exists. Weighted graphs also say how expensive, long, strong, or important that connection is.

## Key facts
- In a weighted graph, paths can be compared by total weight.
- Minimum spanning trees are defined for connected weighted graphs.
- Shortest-path problems often become weight-sensitive in weighted graphs.
- Flow networks can be seen as directed weighted structures where weights represent capacities.

## Example
If edge weights represent road distances, then the graph becomes a transportation model where shortest paths and cheapest spanning structures matter.

## Related concepts
- **used by**: minimum spanning tree, shortest path, flow network
- **contrasts with**: unweighted graph
- **see also**: Kruskal's algorithm, Prim's algorithm, greedy algorithm
