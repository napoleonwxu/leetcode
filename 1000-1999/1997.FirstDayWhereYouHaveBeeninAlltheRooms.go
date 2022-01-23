const MOD = 1e9 + 7

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = (2*dp[i-1] - dp[nextVisit[i-1]] + 2 + MOD) % MOD
	}
	return dp[n-1]
}