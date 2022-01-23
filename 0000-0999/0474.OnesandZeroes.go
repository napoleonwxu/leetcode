func findMaxForm(strs []string, m int, n int) int {
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    for _, str := range strs {
        cnt0 := strings.Count(str, "0")
        cnt1 := len(str) - cnt0
        for i := m; i >= cnt0; i-- {
            for j := n; j >= cnt1; j-- {
                dp[i][j] = max(dp[i][j], dp[i-cnt0][j-cnt1]+1)
            }
        }
    }
    return dp[m][n]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}