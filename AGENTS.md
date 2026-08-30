# AGENTS.md

## Project conventions

### Hierarchical glossary convention
Maintain the project glossary as a **hierarchy**, not a flat list.

Preferred organization rule:
1. Group by **natural grouping** first.
2. Within each group, sort entries **alphabetically**.
3. If a term belongs to a clear subfamily, place it under a subgroup rather than flattening it.
4. Prefer concept hierarchy over strict global alphabetical order.
5. When no clear definition or subgroup is available yet, place the term in the closest natural group and alphabetize it there.

### What counts as a natural grouping
Examples include:
- foundations
- movement and reachability
- trees and spanning structures
- directed graphs
- Eulerian and Hamiltonian themes
- bipartite, matching, and flow
- coloring and independence
- matrix and algebraic graph theory
- algorithms and optimization
- named graph families
- theorems and structural results

### Glossary maintenance rules
- Update `docs/glossary.md` whenever a lesson introduces a new important term.
- Keep glossary entries concise: term plus one-sentence definition.
- Prefer adding or refining a group rather than appending ad hoc items at the bottom.
- If the glossary starts to outgrow a single level, introduce subgroups.
- Keep notation entries (such as set notation) in the most natural conceptual group unless a notation subgroup becomes large enough to justify its own section.

### Knowledge base alignment
When a glossary term becomes central to instruction or reused across lessons, add or update a corresponding file under:
- `docs/knowledge_base/concepts/`

Then update:
- `docs/knowledge_base/index.md`

### Project-local skill location
Project-local skills belong under:
- `.agents/skills/`

### Learning-skill alignment
The learner-support skill should treat the glossary as hierarchical and should prefer introducing terms in their concept family, not as isolated vocabulary.

### New coding artifact / worked example process
When adding a named algorithm or worked example to the project, apply the following **standard treatment**:

1. **Code**: implement the algorithm in the appropriate package (e.g. `theory/`, `traversal/`, or a new package).
2. **Named example**: add an example under `examples/algorithms/{algorithm_name}.go` using the `ExampleXxx` naming convention.
3. **Referenceable write-up**: add `docs/algorithms/{algorithm_name}.md` with purpose, intuition, key facts, code location, related terms, and a tiny worked example.
4. **KB concept page**: add `docs/knowledge_base/concepts/{algorithm_name}.md` linking to related concepts (depends on / contrasts with / see also).
5. **KB index**: add the new concept to `docs/knowledge_base/index.md` under "Additional concept entries".
6. **Glossary**: add the term to `docs/glossary.md` in the appropriate natural group, alphabetized within that group.
7. **Learning-coach skill**: ensure `.agents/skills/graph-learning-coach/SKILL.md` references `docs/algorithms/` as a learning source.
8. **README**: ensure the top-level `README.md` mentions new directories or packages.

All seven touch points must be updated for a new coding artifact to be considered fully integrated.
