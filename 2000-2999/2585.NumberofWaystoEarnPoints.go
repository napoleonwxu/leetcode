const MOD = int(1e9 + 7)

func waysToReachTarget(target int, types [][]int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for _, t := range types {
		for i := target; i > 0; i-- {
			for k := 1; k <= t[0] && i-k*t[1] >= 0; k++ {
				dp[i] = (dp[i] + dp[i-k*t[1]]) % MOD
			}
		}
	}
	return dp[target]
}