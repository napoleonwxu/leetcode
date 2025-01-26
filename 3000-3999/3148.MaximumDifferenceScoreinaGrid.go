func maxScore(grid [][]int) int {
    ans := math.MinInt32
    for i := range grid {
        for j := range grid[i] {
            mi := math.MaxInt32
            if i > 0 {
                mi = min(mi, grid[i-1][j])
            }
            if j > 0 {
                mi = min(mi, grid[i][j-1])
            }
            ans = max(ans, grid[i][j]-mi)
            grid[i][j] = min(grid[i][j], mi)
        }
    }
    return ans
}