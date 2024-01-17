const MOD = 1e9 + 7

func peopleAwareOfSecret(n int, delay int, forget int) int {
    dp := make([]int, n+1)
    dp[1] = 1
    share := 0
    for i := 2; i <= n; i++ {
        share = (share + dp[max(0, i-delay)] - dp[max(0, i-forget)] + MOD) % MOD
        dp[i] = share
    }
    ans := 0
    for i := n - forget + 1; i <= n; i++ {
        ans = (ans + dp[i]) % MOD
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}