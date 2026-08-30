# Hierholzer's Algorithm

## Definition
Hierholzer's algorithm is a constructive procedure for building an **Eulerian circuit** when one exists.

## Intuition
The algorithm walks along unused edges, forms cycles, then splices them together into a single Eulerian circuit.

## Key facts
- Works for undirected graphs where every vertex has even degree
- Works for directed graphs where every vertex has equal in-degree and out-degree
- Runs in linear time relative to the number of edges
- This is an **edge-traversal** algorithm, not a minimum spanning tree algorithm

## In this project
- Code: `theory/eulerian.go`
  - `EulerianCircuit` (undirected)
  - `EulerianCircuitDigraph` (directed)
- Example: `examples/algorithms/hierholzer.go`
- Write-up: `docs/algorithms/hierholzer.md`

## Related concepts
- **depends on**: Eulerian path, Eulerian circuit, edge-traversal problem, constructive algorithm
- **contrasts with**: Kruskal's algorithm, Prim's algorithm (which are MST algorithms, not edge-traversal)
- **see also**: Eulerian path, Eulerian circuit
