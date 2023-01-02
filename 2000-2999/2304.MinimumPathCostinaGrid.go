func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[m-1] = grid[m-1]
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			cur := grid[i][j]
			dp[i][j] = moveCost[cur][0] + dp[i+1][0]
			for k := 1; k < n; k++ {
				dp[i][j] = min(dp[i][j], moveCost[cur][k]+dp[i+1][k])
			}
			dp[i][j] += grid[i][j]
		}
	}
	ans := dp[0][0]
	for j := 1; j < n; j++ {
		ans = min(ans, dp[0][j])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}