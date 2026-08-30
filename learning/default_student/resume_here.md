# Resume Here

## Current focus
- Eulerian circuits and Eulerian conditions

## What you have already covered
- graph basics: vertices, edges, adjacency, neighbors, degree
- walks, paths, cycles, connectedness
- trees
- spanning trees, rooted trees, search trees, minimum spanning trees
- arborescences
- digraph basics
- adjacency matrices
- Laplacian basics and sign interpretation
- strong vs weak connectivity in digraphs
- strongly connected components (SCCs)
- spanning trees vs MST with Kruskal's algorithm (3/3 correct)

## PENDING: 3 Eulerian quick-check questions (answer these first next time)
1. Does this graph have an Eulerian circuit?
   ```text
   A -- B
   |    |
   D -- C
   ```
2. Does this graph have an Eulerian path?
   ```text
   A -- B -- C
   ```
3. If a connected graph has exactly 4 odd-degree vertices, can it have an Eulerian path using every edge exactly once?

## What needs repetition
- Laplacian intuition
- why local matrix entries imply global structure
- repeated practice computing $A$, $D$, and $L$
- SCC grouping: remember mutual reachability, not individual listing

## Suggested restart sequence
1. Answer the 3 pending Eulerian quick-check questions above
2. Re-read `docs/exercises/laplacian_practice_sheet.md` and solve Problems 1-3
3. Review `docs/lessons/matrix_tree.md`
4. Later: return to Laplacian practice sheet problems 4-10
5. Then: topological sorting and condensation graphs

## Coaching note
You said you will definitely need to revisit Laplacians multiple times. That is expected and already built into the study flow.
