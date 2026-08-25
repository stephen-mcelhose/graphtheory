# Laplacian Practice Sheet

This sheet is designed to build Laplacian intuition step by step.

Conventions for this sheet:
- all graphs are undirected simple graphs
- vertex order is always the numerical order shown
- adjacency matrix $A$ uses the same vertex order for rows and columns
- degree matrix $D$ has vertex degrees on the diagonal
- Laplacian is $L = D - A$

Work slowly. For each graph, first identify the neighbors of each vertex, then the degree of each vertex, then the matrix entries.

---

## Problem 1: Path on 2 vertices

Graph:

```text
0 -- 1
```

### Given

$A =
\begin{bmatrix}
0 & 1 \\
1 & 0
\end{bmatrix}
$

$D =
\begin{bmatrix}
1 & 0 \\
0 & 1
\end{bmatrix}
$

### Questions
1. Compute $L = D - A$.
2. What does the $-1$ entry mean here?
3. Why do the diagonal entries equal $1$?

---

## Problem 2: Path on 3 vertices

Graph:

```text
0 -- 1 -- 2
```

### Given

$A =
\begin{bmatrix}
0 & 1 & 0 \\
1 & 0 & 1 \\
0 & 1 & 0
\end{bmatrix}
$

$D =
\begin{bmatrix}
1 & 0 & 0 \\
0 & 2 & 0 \\
0 & 0 & 1
\end{bmatrix}
$

### Questions
1. Compute the Laplacian.
2. Why is the middle diagonal entry larger than the other two?
3. Which entries show that vertices $0$ and $2$ are not neighbors?

---

## Problem 3: Triangle graph

Graph:

```text
0 -- 1
|  / |
2
```

More explicitly:

```text
Edges: (0,1), (1,2), (2,0)
```

### Given

$A =
\begin{bmatrix}
0 & 1 & 1 \\
1 & 0 & 1 \\
1 & 1 & 0
\end{bmatrix}
$

$D =
\begin{bmatrix}
2 & 0 & 0 \\
0 & 2 & 0 \\
0 & 0 & 2
\end{bmatrix}
$

### Questions
1. Compute the Laplacian.
2. Why are all diagonal entries equal?
3. Why are all off-diagonal entries negative here?

---

## Problem 4: Star graph on 4 vertices

Graph:

```text
  1
  |
0-2
  |
  3
```

Use vertex order $0,1,2,3$, where the actual edges are:

```text
(2,0), (2,1), (2,3)
```

### Given

$A =
\begin{bmatrix}
0 & 0 & 1 & 0 \\
0 & 0 & 1 & 0 \\
1 & 1 & 0 & 1 \\
0 & 0 & 1 & 0
\end{bmatrix}
$

$D =
\begin{bmatrix}
1 & 0 & 0 & 0 \\
0 & 1 & 0 & 0 \\
0 & 0 & 3 & 0 \\
0 & 0 & 0 & 1
\end{bmatrix}
$

### Questions
1. Compute the Laplacian.
2. Which row corresponds to the center vertex?
3. How does that row show stronger neighbor coupling than the leaf rows?

---

## Problem 5: Path on 4 vertices

Graph:

```text
0 -- 1 -- 2 -- 3
```

### Tasks
1. Write the adjacency matrix $A$.
2. Write the degree matrix $D$.
3. Compute the Laplacian $L$.
4. Which two rows correspond to endpoints?
5. Which two rows correspond to interior vertices?

---

## Problem 6: Cycle on 4 vertices

Graph:

```text
0 -- 1
|    |
3 -- 2
```

Edges:

```text
(0,1), (1,2), (2,3), (3,0)
```

### Tasks
1. Write $A$.
2. Write $D$.
3. Compute $L$.
4. Compare this Laplacian to the one for the path on 4 vertices. What structural difference shows up in the corner entries?

---

## Problem 7: One isolated vertex

Graph:

```text
0 -- 1    2
```

Vertex $2$ is isolated.

### Tasks
1. Write $A$.
2. Write $D$.
3. Compute $L$.
4. What row and column pattern shows that vertex $2$ has no neighbors?
5. Why does this graph fail to be connected?

---

## Problem 8: Two disconnected edges

Graph:

```text
0 -- 1    2 -- 3
```

### Tasks
1. Write $A$.
2. Write $D$.
3. Compute $L$.
4. Which block structure of the matrix reflects the two connected components?
5. Compare this with Problem 7: how does “disconnected” look different from “isolated vertex”? 

---

## Problem 9: Complete graph on 4 vertices

Graph:

```text
Every vertex is connected to every other vertex.
```

Edges:

```text
(0,1), (0,2), (0,3), (1,2), (1,3), (2,3)
```

### Tasks
1. Write $A$.
2. Write $D$.
3. Compute $L$.
4. Why are all off-diagonal entries negative?
5. Why are all diagonal entries equal to $3$?

---

## Problem 10: Mixed 5-vertex graph

Graph:

```text
0 -- 1 -- 2
|         |
3 ------- 4
```

Use these edges exactly:

```text
(0,1), (1,2), (2,4), (4,3), (3,0)
```

### Tasks
1. Write $A$.
2. Write $D$.
3. Compute $L$.
4. Which vertices have degree $2$?
5. Which entries of $L$ show that $0$ and $2$ are not neighbors?
6. Is this graph connected?

---

# Hints

## General method
For every problem, do these steps in order:

1. List the neighbors of each vertex.
2. Count the degree of each vertex.
3. Build $D$ from the degrees.
4. Build $A$ from adjacency.
5. Compute $L = D - A$ entry by entry.

## How to fill $A$
For undirected simple graphs:
- put $1$ in entry $(i,j)$ if vertices $i$ and $j$ are adjacent
- put $0$ otherwise
- diagonal entries are $0$ because there are no self-loops in these examples
- the matrix will be symmetric

## How to fill $D$
- diagonal entry $D_{ii}$ is the degree of vertex $i$
- all off-diagonal entries are $0$

## How to recognize Laplacian entries directly
Without fully subtracting matrices first, you can remember:
- $L_{ii} = \deg(i)$
- $L_{ij} = -1$ if $i$ and $j$ are neighbors
- $L_{ij} = 0$ if $i$ and $j$ are not neighbors

That rule works for all problems on this sheet.

## What the entries mean
- positive diagonal = how many neighbors this vertex has
- negative off-diagonal = this pair of vertices are adjacent
- zero off-diagonal = this pair are not adjacent

## Structural clues to watch for
- endpoint of a path: diagonal entry $1$
- interior point of a path: larger diagonal entry
- isolated vertex: an entire zero row except possibly its diagonal, which is also $0$ here
- disconnected graph: matrix separates into blocks corresponding to components
- highly connected graph: many negative off-diagonal entries

## Self-check questions
After every Laplacian, ask yourself:
1. Are the diagonal entries equal to the degrees?
2. Are adjacent pairs marked by negative entries?
3. Are non-neighbor pairs marked by zero?
4. Does the matrix visually reflect the graph’s shape?

## If you get stuck
Start from the graph itself, not the matrix.

Write for each vertex:
- neighbors:
- degree:

Then the matrix almost fills itself in.