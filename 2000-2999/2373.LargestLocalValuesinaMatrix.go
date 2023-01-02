func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		ans[i] = make([]int, n-2)
		for j := 0; j < n-2; j++ {
			ans[i][j] = grid[i][j]
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					ans[i][j] = max(ans[i][j], grid[x][y])
				}
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}