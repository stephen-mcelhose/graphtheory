# graphtheory

A pedagogical Go library for studying graph theory, organized around Darij Grinberg's notes *[An introduction to graph theory](https://arxiv.org/abs/2308.04512)* (arXiv:2308.04512).

This repo is a course companion, not a finished graph SDK. The code, chapter summaries, lessons, and the graph learning coach are a **starting point**. They will be wrong, incomplete, or one-sided in places. Do not trust them as authority. When something looks thin, missing, or suspicious, plug the gap against the paper and other references, then fix the repo.

Grinberg's notes are dedicated to the public domain under [CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/). That covers the notes. This repository's own code and docs are under [MIT](LICENSE).

## How to use this

`START_HERE.md` is the re-entry map. Current learner progress is in `learning/default_student/`, including `resume_here.md`. Laplacian drill is `docs/exercises/laplacian_practice_sheet.md`.

Work through `docs/chapters/` first. Then lessons, then the matching Go packages. Keep `docs/research_to_implementation_gap.md` open. That table is the honest map of what is implemented, what is a checker only, and what is still just prose.

When you find a hole, fill it. Add a test, a lesson note, or a pointer to an external source. The paper is the spine. Bondy and Murty, Diestel, and the usual algorithm references are fair game when the notes skip an implementation detail.

## Graph learning coach

`.agents/skills/graph-learning-coach/` is a project-local coaching skill. It reads the chapter summaries, lessons, glossary, and algorithm notes, then writes progress under `learning/{student}/` (default: `learning/default_student/`).

It tracks covered topics, fumbles, quizzes, and chapter progress. Use it to quiz, review, or pick a next topic. Treat its explanations the same way you treat the rest of this repo. If it disagrees with Grinberg or another source, believe the source and correct the learner files.

Point the skill at this checkout as `project_root` and pick a `student` id if you are not using `default_student`.

`.agents/skills/hierarchical-glossary-maintainer/` keeps `docs/glossary.md` grouped by concept family and alphabetized within each group.

## Package layout

- `graph`: core graph and digraph data structures
- `traversal`: BFS, DFS, components, SCCs, bridges, distances
- `flow`: max flow and bipartite matching
- `mst`: weighted graphs, Kruskal, and Prim
- `algebra`: adjacency matrices, Laplacians, determinant and spanning-tree counting
- `generators`: complete/path/cycle/bipartite/hypercube/de Bruijn graph generators
- `theory`: Eulerian, Hamiltonian, tree, bipartite, and degree-condition helpers
- `examples`: small examples, including named graph families and algorithm demos
- `examples/algorithms`: Hierholzer, Kruskal, Prim, and Hamiltonian search examples
- `docs`: chapters, knowledge base, glossary, algorithm write-ups, lessons, exercises, sources, and the implementation-gap table
- `docs/algorithms`: referenceable algorithm notes tied to the code
- `learning`: per-student notes and progress
- `.agents/skills/graph-learning-coach`: coaching skill
- `.agents/skills/hierarchical-glossary-maintainer`: keeps `docs/glossary.md` hierarchical and aligned with the knowledge base

## Scope

The notes are theorem-heavy. Not every theorem is an efficient algorithm. Code here is split into:

- polynomial-time constructive algorithms
- exact small-instance search
- checkers and educational helpers

Stdlib only, module `github.com/stephen-mcelhose/graphtheory`.
