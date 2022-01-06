func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		m = min(m, op[0])
		n = min(n, op[1])
	}
	return m * n
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}