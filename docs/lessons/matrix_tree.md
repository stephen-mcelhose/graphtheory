# Lesson: Matrix-Tree Theorem

## Definition
The Matrix-Tree theorem counts spanning trees using a determinant of a Laplacian cofactor.

## Intuition
A counting problem on graphs becomes a linear algebra problem on matrices.

## Tiny example
For $K_4$, the spanning-tree count is $16$.

## In code
- `algebra.LaplacianMatrix`
- `algebra.SpanningTreeCount`

## Common confusion
The theorem counts spanning trees; it does not itself construct one.

## Status
Implemented constructively for counting, not proof generation.

## Exercise prompt
Use the library to compare spanning-tree counts of $P_4$, $C_4$, and $K_4$.
