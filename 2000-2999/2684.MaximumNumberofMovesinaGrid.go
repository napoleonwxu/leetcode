func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := n - 2; j >= 0; j-- {
		for i := 0; i < m; i++ {
			for ii := i - 1; ii <= i+1; ii++ {
				if ii >= 0 && ii < m && grid[ii][j+1] > grid[i][j] {
					dp[i][j] = max(dp[i][j], dp[ii][j+1]+1)
				}
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans = max(ans, dp[i][0])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}