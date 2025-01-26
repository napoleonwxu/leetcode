func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    dp := make([]int, n)
    for j := 0; j < n; j++ {
        dp[j] = triangle[n-1][j]
    }
    for i := n-2; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
        }
    }
    return dp[0]
}