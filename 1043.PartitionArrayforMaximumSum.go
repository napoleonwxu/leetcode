func maxSumAfterPartitioning(A []int, K int) int {
    n := len(A)
    dp := make([]int, n+1)
    dp[0] = 0
    for i := 1; i <= n; i++ {
        maxInSub := 0
        for j := 1; j <= K && i-j >= 0; j++ {
            maxInSub = max(maxInSub, A[i-j])
            dp[i] = max(dp[i], dp[i-j] + j*maxInSub)
        }
    }
    return dp[n]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}