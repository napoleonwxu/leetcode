func stoneGame(piles []int) bool {
    /* Alex can always take all odd piles or always take all even piles.
       Since sum of all piles is odd then sum of all odd piles won't equals sum of all even piles, Alex could just take the bigger ones. */
    //return true
    // DP: dp[i][j] means the biggest number of stones you can get more than opponent picking piles in piles[i] ~ piles[j].
    n := len(piles)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][i] = piles[i]
    }
    for i := n-2; i >= 0; i-- {
        for j := i+1; j < n; j++ {
            dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
        }
    }
    return dp[0][n-1] > 0
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}