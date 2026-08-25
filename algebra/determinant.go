package algebra

func Minor(a [][]int, row, col int) [][]int {
	out := [][]int{}
	for i := 0; i < len(a); i++ {
		if i == row {
			continue
		}
		line := []int{}
		for j := 0; j < len(a[i]); j++ {
			if j == col {
				continue
			}
			line = append(line, a[i][j])
		}
		out = append(out, line)
	}
	return out
}

func Determinant(a [][]int) int {
	n := len(a)
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a[0][0]
	}
	if n == 2 {
		return a[0][0]*a[1][1] - a[0][1]*a[1][0]
	}
	res := 0
	sign := 1
	for j := 0; j < n; j++ {
		res += sign * a[0][j] * Determinant(Minor(a, 0, j))
		sign = -sign
	}
	return res
}
