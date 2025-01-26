func maxScore(a []int, b []int) int64 {
    m, n := len(a), len(b)
    dp := make([][]int64, m+1)
    for i := range dp {
        dp[i] = make([]int64, n+1)
        if i > 0 {
            dp[i][i] = dp[i-1][i-1] + int64(a[i-1]*b[i-1])
        } 
    }
    for i := 1; i <= m; i++ {
        for j := i+1; j <= n; j++ {
            dp[i][j] = max(dp[i][j-1], dp[i-1][j-1] + int64(a[i-1]*b[j-1]))
        }
    }
    return dp[m][n]
}