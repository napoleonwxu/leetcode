/*
dp[n] = dp[n-1] + dp[n-2] + 2*(dp[n-3] + ... + d[0])
      = dp[n-1] + dp[n-3] + (dp[n-2] + dp[n-3] + 2*(dp[n-4] + ... + d[0]))
      = dp[n-1] + dp[n-3] + dp[n-1]
      = 2*dp[n-1] + dp[n-3]
*/
const MOD = 1e9 + 7

func numTilings(N int) int {
    dp := [1001]int{1: 1, 2: 2, 3: 5}
    //dp[1], dp[2], dp[3] = 1, 2, 5
    for i := 4; i <= N; i++ {
        dp[i] = (2*dp[i-1] + dp[i-3])%MOD
    }
    return dp[N]
}