# Go Library Plan for arXiv:2308.04512 Coverage

## Revised conceptual partition

1. Foundations and abstract interfaces
2. Structural families and generators
3. Traversal, metric, and connectivity engine
4. Traversal sequences and paths
5. Optimization, flows, and min-max duality
6. Algebraic, enumerative, and invariant analysis

## Package plan

### graph
Core data types and interfaces.
- `Graph`
- `Digraph`
- `SimpleGraph`
- `SimpleDigraph`
- subgraph/export helpers

### traversal
Core constructive algorithms.
- BFS/DFS
- connected components
- weak/strong components
- bridge detection
- shortest unweighted distances
- tree centers

### theory
Recognition and theorem-condition helpers.
- bipartite test and 2-coloring
- tree recognition
- Eulerian conditions and constructions
- Hamiltonian exact search for small graphs
- tournament Hamiltonian path
- Ore/Dirac condition checks

### flow
Optimization packages.
- Edmonds-Karp max flow
- residual network utilities
- bipartite maximum matching via augmenting paths
- Menger-style edge-connectivity by flow reduction helpers

### algebra
Algebraic and counting layer.
- adjacency matrix
- Laplacian matrix
- determinant
- Matrix-Tree spanning-tree count
- chromatic polynomial by deletion-contraction (small graphs)

### generators
Classical constructors.
- complete/path/cycle/empty graphs
- complete bipartite graphs
- hypercubes
- de Bruijn digraphs
- tournaments from orientation tables

## Coverage philosophy
- Cover central objects and constructions from the paper directly.
- Provide exact but exponential routines only when clearly labeled.
- Prefer interfaces so algorithms are decoupled from storage.
- Include examples demonstrating conceptual bridges across packages.
