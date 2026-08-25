## Graph Theory Knowledge Graph

Below is a **broad graph theory knowledge graph** in **structured Markdown**. It is organized as concepts plus relationships, so you can read it like a map.

---

# 1. Core object: Graph theory

**Graph theory** studies structures made of:
- **vertices** (nodes)
- **edges** (connections)

It branches into:
- **undirected graphs**
- **directed graphs (digraphs)**
- **weighted graphs**
- **special graph families**
- **structural graph theory**
- **algorithmic graph theory**
- **algebraic/spectral graph theory**

---

# 2. Foundations

## 2.1 Graph
A **graph** consists of:
- a vertex set #$#V#$#
- an edge set #$#E#$#

### Related concepts
- has → **vertices**
- has → **edges**
- may be → **simple graph**
- may be → **multigraph**
- may be → **digraph**
- may be → **weighted graph**

---

## 2.2 Vertex
A **vertex** is a point/node in a graph.

### Related concepts
- connected by → **edges**
- has → **neighbors**
- has → **degree**
- may be part of → **walk**, **path**, **cycle**

---

## 2.3 Edge
An **edge** connects two vertices.

### Related concepts
- defines → **adjacency**
- may have → **weight**
- in digraphs becomes → **arc**

---

## 2.4 Adjacency
Two vertices are **adjacent** if an edge connects them.

### Related concepts
- determines → **neighbors**
- encoded by → **adjacency matrix**
- used in → **colorings**, **paths**, **connectivity**

---

## 2.5 Neighbor
A **neighbor** of a vertex is an adjacent vertex.

### Related concepts
- counted by → **degree**
- influences → **Laplacian**

---

## 2.6 Degree
The **degree** of a vertex is the number of incident edges / neighbors in a simple graph.

### Related concepts
- stored in → **degree matrix**
- used in → **handshaking lemma**
- important for → **trees**, **Eulerian graphs**, **Laplacian**

---

# 3. Movement and reachability

## 3.1 Walk
A **walk** is a sequence of vertices where consecutive vertices are connected.

### Related concepts
- generalizes → **path**
- may repeat → vertices and edges

---

## 3.2 Path
A **path** is a walk with no repeated vertices.

### Related concepts
- special case of → **walk**
- used to define → **connectedness**
- central in → **trees**, **Hamiltonian path**

---

## 3.3 Cycle
A **cycle** is a closed path.

### Related concepts
- absence defines → **acyclic**
- forbidden in → **trees**
- central in → **Eulerian** and **Hamiltonian** ideas

---

## 3.4 Connectedness
A graph is **connected** if every pair of vertices has a path between them.

### Related concepts
- failure produces → **components**
- strengthened in digraphs to → **strong connectivity**
- reflected by → **Laplacian**

---

## 3.5 Component
A **component** is a maximal connected piece of a graph.

### Related concepts
- counted via → **Laplacian nullity** (later/algebraically)
- found by → **BFS**, **DFS**

---

# 4. Tree world

## 4.1 Tree
A **tree** is a connected, acyclic graph.

### Related concepts
- has → unique simple path between any two vertices
- has → leaves
- satisfies → #$#|E| = |V| - 1#$#
- parent concept for → **rooted trees**, **spanning trees**

---

## 4.2 Leaf
A **leaf** is a vertex of degree #$#1#$# in a tree.

### Related concepts
- appears in → inductive proofs on trees
- useful in → pruning arguments

---

## 4.3 Rooted tree
A **rooted tree** is a tree with a chosen root.

### Related concepts
- induces → parent/child relation
- basis for → **search trees**
- directed analogue → **arborescence**

---

## 4.4 Search tree
A **search tree** is a rooted tree produced by graph traversal.

### Related concepts
- produced by → **BFS**
- produced by → **DFS**
- useful for → exploration, shortest paths, proofs

---

## 4.5 Spanning tree
A **spanning tree** is a tree subgraph using all vertices of a connected graph.

### Related concepts
- removes → cycles
- keeps → connectivity
- counted by → **Matrix-Tree theorem**
- optimized by → **minimum spanning tree**

---

## 4.6 Minimum spanning tree (MST)
A **minimum spanning tree** is a spanning tree with minimum total edge weight.

### Related concepts
- belongs to → weighted graph theory
- found by → **Kruskal’s algorithm**, **Prim’s algorithm**
- distinct from → arbitrary **spanning tree**

---

## 4.7 Arborescence
An **arborescence** is a directed rooted tree structure in a digraph.

### Related concepts
- directed analogue of → **rooted tree**
- counted by → directed forms of **Matrix-Tree theorem**

---

# 5. Directed graph world

## 5.1 Digraph
A **digraph** is a graph with directed edges.

### Related concepts
- edges become → **arcs**
- has → **in-degree**, **out-degree**
- supports → directed paths, strong connectivity, tournaments, flows

---

## 5.2 Arc
An **arc** is a directed edge #$#u \to v#$#.

### Related concepts
- contributes to → in-degree/out-degree
- encoded by → adjacency matrix in directed form

---

## 5.3 In-degree / Out-degree
In a digraph:
- **out-degree** = number of arcs leaving a vertex
- **in-degree** = number of arcs entering a vertex

### Related concepts
- used in → Eulerian digraph conditions
- used in → flow balance ideas

---

## 5.4 Directed path
A **directed path** must follow edge directions.

### Related concepts
- defines → reachability in digraphs
- used in → arborescences, flows, topological ideas

---

## 5.5 Strong connectivity
A digraph is **strongly connected** if every vertex can reach every other by directed paths.

### Related concepts
- stronger than → weak connectivity
- important in → tournaments, arborescences, digraph structure

---

## 5.6 Weak connectivity
A digraph is **weakly connected** if the underlying undirected graph is connected.

### Related concepts
- weaker than → strong connectivity

---

## 5.7 Tournament
A **tournament** is an orientation of a complete graph: every pair of vertices has exactly one directed edge between them.

### Related concepts
- special class of → digraph
- linked to → Hamiltonian paths/cycles
- rich in → extremal and structural results

---

# 6. Special graph families

## 6.1 Complete graph
A **complete graph** #$#K_n#$# has every pair of distinct vertices adjacent.

### Related concepts
- extreme example for → density
- used in → spanning-tree counts, colorings

---

## 6.2 Path graph
A **path graph** is a graph shaped like a line.

### Related concepts
- simple example for → Laplacian
- is a → tree

---

## 6.3 Cycle graph
A **cycle graph** is a single cycle.

### Related concepts
- simplest example of → cyclic connected graph

---

## 6.4 Complete bipartite graph
A **complete bipartite graph** #$#K_{m,n}#$# connects every vertex in one part to every vertex in the other.

### Related concepts
- key in → matchings, Hall’s theorem, flows
- example of → bipartite graph

---

## 6.5 Hypercube
The **hypercube** #$#Q_n#$# has binary strings as vertices, with edges between strings differing in one coordinate.

### Related concepts
- example of → bipartite graph
- important in → combinatorics, coding, parallelism

---

## 6.6 de Bruijn graph
A **de Bruijn graph** encodes overlaps of strings.

### Related concepts
- digraph
- related to → Eulerian cycles, sequence generation

---

# 7. Eulerian and Hamiltonian themes

## 7.1 Eulerian path / circuit
An **Eulerian path** uses every edge exactly once.  
An **Eulerian circuit** is an Eulerian path that starts and ends at the same vertex.

### Related concepts
- depends on → degrees
- constructive algorithm → **Hierholzer**
- contrasted with → Hamiltonian concepts

---

## 7.2 Hamiltonian path / cycle
A **Hamiltonian path** visits every vertex exactly once.  
A **Hamiltonian cycle** is a cycle visiting every vertex exactly once.

### Related concepts
- vertex-based analogue of → Eulerian ideas
- generally harder computationally
- appears in → tournaments, extremal conditions

---

# 8. Bipartite, matching, and flow cluster

## 8.1 Bipartite graph
A graph is **bipartite** if its vertices split into two parts with all edges crossing between the parts.

### Related concepts
- equivalent to → 2-colorable
- central for → matchings, Hall’s theorem, flows

---

## 8.2 Matching
A **matching** is a set of edges with no shared endpoints.

### Related concepts
- special cases → perfect matching
- central theorem → **Hall’s marriage theorem**
- tied to → bipartite graphs

---

## 8.3 Perfect matching
A **perfect matching** matches every vertex.

### Related concepts
- special case of → matching
- often studied in → bipartite graphs

---

## 8.4 Hall’s marriage theorem
Gives a condition for when a bipartite graph has a matching covering one side.

### Related concepts
- foundational theorem in → matching theory
- connected to → flows

---

## 8.5 Network flow
A **flow network** has directed edges with capacities and a source/sink.

### Related concepts
- central theorem → **max-flow min-cut**
- supports proofs of → Hall’s theorem, Menger-type results

---

## 8.6 Max-flow min-cut
The maximum value of a flow equals the minimum capacity of a cut separating source and sink.

### Related concepts
- major bridge between → algorithms and theorems
- used in → matchings, connectivity questions

---

# 9. Coloring and independence cluster

## 9.1 Proper coloring
A **proper coloring** assigns colors to vertices so adjacent vertices get different colors.

### Related concepts
- minimum number of colors = → **chromatic number**
- polynomial form = → **chromatic polynomial**
- special case → bipartite = 2-colorable

---

## 9.2 Chromatic number
The least number of colors needed for a proper coloring.

### Related concepts
- hard optimization problem
- connected to → clique structure, bipartiteness

---

## 9.3 Chromatic polynomial
Counts the number of proper colorings as a function of number of colors.

### Related concepts
- algebraic invariant
- belongs to → enumerative graph theory

---

## 9.4 Independent set
An **independent set** is a set of vertices with no edges between them.

### Related concepts
- dual flavor to → cliques and colorings
- important in → extremal graph theory

---

## 9.5 Turán-type ideas
Study how many edges force certain subgraphs or avoid them.

### Related concepts
- extremal graph theory
- connected to → cliques, independent sets, density

---

# 10. Matrix and algebraic graph theory

## 10.1 Adjacency matrix
The **adjacency matrix** records which vertices are adjacent.

### Related concepts
- for digraphs, standard convention: #$#A_{ij}#$# means edge/arc from #$#i#$# to #$#j#$#
- combined with degree matrix to make → **Laplacian**

---

## 10.2 Degree matrix
The **degree matrix** is diagonal, with degrees on the diagonal.

### Related concepts
- used in → Laplacian

---

## 10.3 Laplacian
The **graph Laplacian** is:

#$#L = D - A#$#

### Related concepts
- diagonal encodes → degrees
- off-diagonal encodes → adjacency / neighbor coupling
- reveals → connectedness, components, spanning-tree counts
- central in → spectral graph theory

---

## 10.4 Matrix-Tree theorem
Counts spanning trees using a determinant of a Laplacian-derived matrix.

### Related concepts
- connects → combinatorics + linear algebra
- used for → spanning-tree counting
- extends to → directed arborescence counting

---

## 10.5 Spectral graph theory
Studies graphs via eigenvalues/eigenvectors of matrices like the Laplacian or adjacency matrix.

### Related concepts
- reveals → connectivity, bottlenecks, expansion, clustering

---

# 11. Algorithmic cluster

## 11.1 BFS
**Breadth-first search** explores layer by layer.

### Related concepts
- builds → BFS tree
- finds → shortest paths in unweighted graphs
- finds → components

---

## 11.2 DFS
**Depth-first search** explores by going deep first.

### Related concepts
- builds → DFS tree
- useful for → connectivity, cycles, structural proofs

---

## 11.3 Kruskal / Prim
Algorithms for **minimum spanning trees**.

### Related concepts
- belong to → greedy algorithms
- solve → MST problem

---

## 11.4 Max-flow algorithms
Algorithms such as Ford–Fulkerson / Edmonds–Karp.

### Related concepts
- solve → network flow problems
- support → matching and cut results

---

# 12. Knowledge-graph style relation list

## Core dependencies
- **graph theory** → studies → **graphs**
- **graph** → consists of → **vertices**, **edges**
- **edges** → define → **adjacency**
- **adjacency** → defines → **neighbors**
- **neighbors** → counted by → **degree**

## Movement
- **walk** → generalizes → **path**
- **path** → can close into → **cycle**
- **path** → defines → **connectedness**
- **lack of connectedness** → splits into → **components**

## Tree branch
- **tree** → means → connected + acyclic
- **tree** → implies → unique simple path
- **rooted tree** → specializes → **tree**
- **search tree** → is produced by → **BFS/DFS**
- **spanning tree** → is a tree subgraph of → **connected graph**
- **minimum spanning tree** → optimizes → **spanning tree**
- **arborescence** → directed analogue of → **rooted tree**
- **Matrix-Tree theorem** → counts → **spanning trees**

## Directed branch
- **digraph** → replaces edges with → **arcs**
- **arcs** → induce → **in-degree/out-degree**
- **digraph** → supports → **directed paths**
- **directed paths** → define → **strong connectivity**
- **tournament** → is a special kind of → **digraph**

## Algebraic branch
- **adjacency matrix** → encodes → **adjacency**
- **degree matrix** → encodes → **degrees**
- **Laplacian** → equals → #$#D - A#$#
- **Laplacian** → encodes → local neighbor coupling
- **Laplacian** → reveals → global structure
- **Matrix-Tree theorem** → uses → **Laplacian**

## Algorithmic branch
- **BFS/DFS** → compute → search trees, connectivity structure
- **Kruskal/Prim** → compute → MST
- **flow algorithms** → compute → max flow
- **max flow** → supports → matchings, min-cuts, connectivity theorems

---

# 13. Best way to use this map
A good learning order through this knowledge graph is:

1. **Foundations**  
   graph, vertex, edge, adjacency, degree

2. **Movement**  
   walk, path, cycle, connectedness, components

3. **Trees**  
   tree, rooted tree, spanning tree, MST

4. **Directed graphs**  
   digraph, in/out-degree, directed path, strong connectivity, tournament

5. **Bipartite/matching/flow**  
   bipartite, matching, Hall, flow, max-flow min-cut

6. **Coloring/independence**  
   coloring, chromatic number, independent set, Turán

7. **Matrices/algebra**  
   adjacency matrix, degree matrix, Laplacian, Matrix-Tree theorem, spectral ideas

---

# 14. Your current location in the map
Based on our lesson so far, you’ve already covered:
- graph
- vertex
- edge
- adjacency
- neighbor
- degree
- walk
- path
- cycle
- connectedness
- components
- tree
- spanning tree
- rooted tree
- search tree
- MST
- digraph
- adjacency matrix
- Laplacian
- Matrix-Tree intuition

So your current frontier is:

- **strong vs weak connectivity**
- **deeper Laplacian intuition**
- **spanning tree vs minimum spanning tree examples**
- later: **Matrix-Tree theorem**

## One-sentence summary

**Graph theory is a network of ideas where local connections (adjacency, degree, neighbors) grow into global structure (paths, trees, connectivity, flows, colorings, and matrix methods).**