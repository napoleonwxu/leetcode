func integerBreak(n int) int {
    // O(n) + O(1)
    if n == 2 || n == 3 {
        return n-1
    }
    prod := 1
    for n > 4 {
        prod *= 3
        n -= 3
    }
    prod *= n
    return prod
    /* DP, O(n^2) + O(n)
    dp := make([]int, n+1)
    dp[1] = 1
    for i := 2; i <= n; i++ {
        for j := 1; j <= i/2; j++ {
            dp[i] = max(dp[i], max(j, dp[j])*max(i-j, dp[i-j]))
        }
    }
    return dp[n]
    */
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}