# Topological Sort

## Definition
A **topological sort** (or topological ordering) of a directed acyclic graph (DAG) is an ordering of its vertices such that every arc goes from an earlier vertex to a later vertex in the ordering.

## Intuition
If you think of a DAG as encoding dependencies (e.g. task A must happen before task B), a topological sort gives you a valid schedule that respects all dependencies.

## Key facts
- Only defined for DAGs (no directed cycles).
- A DAG may have multiple valid topological orderings.
- If a digraph has a unique topological sort, it must contain a Hamiltonian path.
- Computed via DFS (postorder) or via repeated removal of vertices with in-degree 0 (Kahn's algorithm).

## Example
```text
A -> B
A -> C
B -> D
C -> D
```

Valid topological sorts: A, B, C, D or A, C, B, D.

## Related concepts
- **depends on**: DAG, directed path, in-degree
- **used by**: scheduling, dependency resolution, course prerequisite ordering
- **contrasts with**: Eulerian/Hamiltonian traversals (which are about edges/vertices visited, not dependency order)
- **see also**: digraph, sink component, condensation graph
