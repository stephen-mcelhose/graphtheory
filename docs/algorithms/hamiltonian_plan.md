# Hamiltonian Algorithm Plan

## Current coverage
- Backtracking for Hamiltonian path/cycle
- Dynamic programming on subsets for Hamiltonian path
- Exact search for Hamiltonian cycle
- Theorem-based sufficient-condition checks already present:
  - Dirac condition
  - Ore condition
- Tournament Hamiltonian path construction

## Planned expansions
1. Add Hamiltonian cycle via subset DP
2. Add branch-and-bound improvements over plain backtracking
3. Add worked examples contrasting:
   - theorem says yes
   - search finds one
   - theorem inconclusive but search still succeeds
   - theorem inconclusive and no Hamiltonian cycle exists
4. Add more named graph families for Hamiltonian counterexamples and examples
5. Add lesson material on why Hamiltonian problems are harder than Eulerian problems

## Teaching contrast to preserve
- Eulerian: edge-traversal, Hierholzer, simple degree-based characterization
- Hamiltonian: vertex-traversal, exact search / DP / sufficient conditions, much harder in general
