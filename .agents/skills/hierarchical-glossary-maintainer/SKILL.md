---
name: hierarchical-glossary-maintainer
description: Maintains the project glossary as a hierarchical concept structure grouped by natural grouping and alphabetized within groups, while aligning glossary entries with the knowledge base.
---

# Hierarchical Glossary Maintainer

Use this skill when the user wants to:
- add terms to the glossary
- reorganize glossary structure
- normalize glossary groupings
- keep glossary and concept pages aligned

## Goal
Treat the glossary as a **hierarchy**:
1. group by **natural grouping**
2. then sort **alphabetically** within each group
3. introduce subgroups if a section becomes too large

## Primary files
- `{project_root}/docs/glossary.md`
- `{project_root}/docs/glossary_conventions.md`
- `{project_root}/docs/knowledge_base/index.md`
- `{project_root}/docs/knowledge_base/concepts/`
- `{project_root}/AGENTS.md`

## Rules
- Do not append terms to the end blindly.
- Place each term in the most natural existing group.
- If no group fits well, create a new natural group.
- If a group becomes crowded or conceptually mixed, split it into subgroups.
- Keep entry definitions concise.
- If a term becomes central to teaching, add or update a corresponding KB concept page.
- Keep KB index references aligned with major glossary concepts.

## Natural grouping examples
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

## Output expectations
When updating the glossary:
- preserve hierarchy
- preserve alphabetical order within group
- mention any new group or subgroup created
- mention any KB concept pages that should be added or updated
