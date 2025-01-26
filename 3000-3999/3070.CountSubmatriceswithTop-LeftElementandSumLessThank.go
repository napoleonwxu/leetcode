func countSubmatrices(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    cnt := 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + grid[i-1][j-1]
            if dp[i][j] <= k {
                cnt++
            }
        }
    }
    return cnt
}