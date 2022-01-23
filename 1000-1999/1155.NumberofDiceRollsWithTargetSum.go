const MOD = 1e9+7

func numRollsToTarget(d int, f int, target int) int {
    if d*f < target {
        return 0
    }
    dp := make([][]int, d+1)
    dp[0] = make([]int, target+1)
    dp[0][0] = 1
    for i := 1; i <= d; i++ {
        dp[i] = make([]int, target+1)
        for j := 1; j <= target; j++ {
            if j > i*f {
                break
            }
            for k := 1; k <= f && j-k >= 0; k++ {
                dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % MOD
            }
        }
    }
    return dp[d][target]
}