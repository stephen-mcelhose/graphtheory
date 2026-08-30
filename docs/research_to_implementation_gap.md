# Research-to-Implementation Gap Analysis

| Topic | In the paper | In this SDK | Learning note |
|---|---|---|---|
| Simple graphs | foundational definitions and examples | implemented | best starting point |
| Walks, paths, cycles | foundational | partially implemented via traversal and examples | understand terms before algorithms |
| Connectedness | core theorem layer | implemented | directly constructive |
| Trees and spanning trees | structural + constructive | partially implemented | good bridge from theory to code |
| Eulerian circuits | theorem + proof + applications | Hierholzer construction in `theory` (`EulerianCircuit`, `EulerianCircuitDigraph`) plus degree checks | strong constructive learning topic |
| Hamiltonian paths/cycles | criteria + examples | backtracking path/cycle, subset DP (`HamiltonianPathDP`), and exact cycle search for small instances | useful for discussing complexity |
| Minimum spanning trees | trees and weighted variants | Kruskal and Prim in `mst` with weighted-graph helpers | good algorithm-to-theorem bridge |
| Tournaments | dedicated directed family | partially implemented | excellent for named-family study |
| Matchings | major chapter | implemented at core level | strong algorithmic topic |
| Flows | major chapter | implemented at core level | central optimization module |
| Matrix-Tree theorem | algebraic counting theorem | implemented for counting helpers | strong bridge to linear algebra |
| Chromatic polynomial | enumerative graph invariant | implemented for small graphs | exponential and educational |
| Menger / Gallai–Milgram | theorem-heavy | partial support / indirect support | best taught conceptually first |
| Ramsey / Turán / Caro–Wei | theorem-oriented | helpers/checkers only | mostly conceptual in this SDK |

## Status labels
- **Implemented constructively**: direct algorithm or constructor exists
- **Partially supported**: some parts are implemented, but not the full theory
- **Checker only**: condition-testing support rather than full constructive realization
- **Small-instance only**: exact but exponential routines
- **Documented only**: represented primarily in the learning materials
