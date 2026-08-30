# Graph Theory Glossary

This glossary tracks important terms from the lesson plan and project knowledge base. It is maintained as a lightweight hierarchy: group by natural grouping first, then alphabetize within each group. See `docs/glossary_conventions.md` for the standing convention.

---

## Core foundations
- **Adjacency** — the relation of two vertices being connected by an edge.
- **Component** — a maximal connected subgraph.
- **Connectedness** — the property that every pair of vertices has a path between them.
- **Degree** — the number of incident edges, or equivalently the number of neighbors in a simple graph.
- **Edge** — a connection between two vertices.
- **Graph** — a structure with a vertex set and an edge set.
- **Graph structure** — the arrangement and relation pattern among vertices, edges, paths, cycles, components, and other substructures of a graph.
- **Loop** — an edge from a vertex to itself; absent in simple graphs.
- **Multigraph** — a graph that may have multiple edges between the same pair of vertices.
- **Neighbor** — a vertex adjacent to another vertex.
- **Path** — a walk with no repeated vertices.
- **Set notation** — the use of curly braces $\{\,\}$ to denote a set of grouped elements, e.g. $\{A, B\}$ means one set containing both $A$ and $B$ together; preferred notation for components, SCCs, matchings, cliques, and edge sets.
- **Simple graph** — a graph with no loops and no multiple edges.
- **Trail** — a walk with no repeated edges.
- **Vertex** — a node in a graph.
- **Walk** — a sequence of vertices where consecutive vertices are connected by edges.
- **Weighted graph** — a graph in which each edge has a numerical weight or cost.

## Movement and reachability
- **Bridge** — an edge whose removal disconnects the graph.
- **Closed walk** — a walk that starts and ends at the same vertex.
- **Cycle** — a closed path.
- **Directed path** — a path that follows edge directions.
- **Reachability** — whether one vertex can be reached from another by following edges or arcs.
- **Shortest path** — a path of minimum total length or weight between two vertices.

## Trees and spanning structures
- **Arborescence** — a directed rooted tree structure in a digraph.
- **Center of a tree** — the vertex or edge midpoint that minimizes the maximum distance to any other vertex.
- **Leaf** — a vertex of degree $1$ in a tree.
- **Minimum spanning tree (MST)** — a spanning tree of minimum total edge weight in a weighted graph.
- **Rooted tree** — a tree with a distinguished root vertex.
- **Search tree** — a rooted tree produced by an exploration procedure such as BFS or DFS.
- **Spanning tree** — a tree subgraph containing all vertices of a connected graph.
- **Tree** — a connected acyclic graph.

## Directed graphs
- **Arc** — a directed edge.
- **Condensation graph** — the digraph obtained by collapsing each SCC into a single node.
- **Digraph** — a graph whose edges have direction.
- **Directed acyclic graph (DAG)** — a digraph with no directed cycles.
- **In-degree** — number of arcs entering a vertex.
- **Out-degree** — number of arcs leaving a vertex.
- **Sink component** — an SCC with no outgoing arcs to other SCCs.
- **Strong connectivity** — every vertex can reach every other by directed paths.
- **Strongly connected component (SCC)** — a maximal set of vertices that are mutually reachable.
- **Topological sort** — an ordering of vertices of a DAG such that every arc goes from an earlier vertex to a later vertex.
- **Tournament** — a digraph where every pair of vertices has exactly one directed edge between them.
- **Weak connectivity** — the underlying undirected graph is connected.

## Eulerian and Hamiltonian themes
- **Camion's theorem** — every strongly connected tournament has a Hamiltonian cycle.
- **Dirac's theorem / Dirac condition** — if every vertex in an $n$-vertex simple graph has degree at least $n/2$, then the graph is Hamiltonian.
- **Edge-traversal problem** — a problem whose objective is defined in terms of using, covering, or ordering edges, rather than visiting each vertex once.
- **Eulerian circuit** — a closed trail using every edge exactly once.
- **Eulerian path** — a trail using every edge exactly once.
- **Hamiltonian backtracking** — an exact search method that grows a vertex-by-vertex path and backtracks when a partial choice fails.
- **Hamiltonian cycle** — a cycle visiting every vertex exactly once.
- **Hamiltonian dynamic programming on subsets** — a subset-state exact algorithm for small Hamiltonian instances.
- **Hamiltonian exact search** — an exponential-time family of methods that explicitly searches for Hamiltonian paths or cycles.
- **Hamiltonian path** — a path visiting every vertex exactly once.
- **Hierholzer's algorithm** — a constructive algorithm for building an Eulerian circuit.
- **Vertex-traversal problem** — a problem whose objective is defined in terms of visiting every vertex, rather than covering every edge.

## Bipartite, matching, and flow
- **Bipartite graph** — a graph whose vertices split into two parts with all edges crossing between the parts.
- **Cut** — a partition of vertices into two sets; in flow networks, separates source from sink.
- **Flow network** — a digraph with edge capacities, a source, and a sink.
- **Hall's marriage theorem** — gives a condition for when a bipartite graph has a matching covering one side.
- **Matching** — a set of edges with no shared endpoints.
- **Max-flow min-cut theorem** — the maximum value of a flow equals the minimum capacity of a cut separating source and sink.
- **Perfect matching** — a matching that covers every vertex.
- **Residual network** — the network of remaining capacity after a partial flow.

## Coloring and independence
- **Chromatic number** — the least number of colors needed for a proper coloring.
- **Chromatic polynomial** — counts the number of proper colorings as a function of the number of colors.
- **Clique** — a set of vertices where every pair is adjacent.
- **Independent set** — a set of vertices with no edges between them.
- **Proper coloring** — an assignment of colors to vertices so adjacent vertices get different colors.
- **Turán-type problem** — extremal questions about how many edges force or avoid certain subgraphs.

## Matrix and algebraic terms
- **Adjacency matrix** — a matrix encoding which vertices are adjacent; in the standard digraph convention, $A_{ij}$ indicates an edge or arc from $i$ to $j$.
- **Degree matrix** — a diagonal matrix with vertex degrees on the diagonal.
- **Determinant** — a scalar value computed from a square matrix, used in Matrix-Tree theorem counting.
- **Laplacian** — the matrix $L = D - A$.
- **Matrix-Tree theorem** — a theorem counting spanning trees via determinants of Laplacian-derived matrices.
- **Neighbor coupling** — the way the Laplacian compares a vertex value to the values on adjacent vertices.
- **Spectral graph theory** — the study of graphs via eigenvalues and eigenvectors of matrices like the Laplacian or adjacency matrix.

## Algorithms and optimization
- **BFS (Breadth-first search)** — an algorithm exploring vertices layer by layer from a start vertex.
- **Constructive algorithm** — an algorithm that explicitly builds the required object, not just proves existence.
- **DFS (Depth-first search)** — an algorithm exploring by going deep along one branch before backtracking.
- **Edmonds-Karp algorithm** — an implementation of Ford-Fulkerson using BFS to find augmenting paths.
- **Ford-Fulkerson algorithm** — a method for computing max flow by repeatedly augmenting along residual paths.
- **Greedy algorithm** — an algorithm that makes the locally optimal choice at each step, e.g. Kruskal's or Prim's.
- **Hierholzer's algorithm** — a constructive algorithm for building an Eulerian circuit.
- **Kruskal's algorithm** — a greedy algorithm for minimum spanning trees based on adding the cheapest non-cycle edge.
- **Ore's theorem** — if every pair of non-adjacent vertices in a graph has degree-sum at least $n$, the graph is Hamiltonian.
- **Prim's algorithm** — a greedy algorithm for minimum spanning trees based on growing a tree outward from a start vertex.

## Named graph families
- **Complete bipartite graph ($K_{m,n}$)** — connects every vertex in one part to every vertex in the other.
- **Complete graph ($K_n$)** — every pair of distinct vertices is adjacent.
- **Cycle graph ($C_n$)** — a single cycle on $n$ vertices.
- **de Bruijn graph** — a digraph encoding overlaps of strings; related to Eulerian cycles and sequence generation.
- **Empty graph** — a graph with no edges.
- **Hypercube ($Q_n$)** — vertices are binary strings of length $n$, with edges between strings differing in one coordinate.
- **Path graph ($P_n$)** — a graph shaped like a line on $n$ vertices.

## Theorem and structural terms
- **Acyclic** — containing no cycles.
- **BEST theorem** — counts Eulerian circuits in directed Eulerian graphs using de Bruijn sequences and matrix-tree counting.
- **Caro-Wei bound** — a lower bound on the independence number of a graph based on vertex degrees.
- **Handshaking lemma** — the sum of all vertex degrees equals twice the number of edges.
- **Menger's theorem** — relates the maximum number of internally disjoint paths to the minimum separating set.
- **Vizing's theorem** — the edge chromatic number of a simple graph is either $\Delta$ or $\Delta+1$.

## Notes
- Prefer curly-brace set notation ($\{\,\}$) when grouping objects like SCCs, matchings, cliques, or edge sets.
- This glossary should be extended whenever a lesson introduces a new important term.
