func maximumAmount(coins [][]int) int {
    m, n := len(coins), len(coins[0])
    dp := make([][][]int, m)
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, 3)
            for k := range dp[i][j] {
                dp[i][j][k] = math.MinInt32
            }
        }
    }
    dp[0][0][0] = coins[0][0]
    dp[0][0][1] = max(0, coins[0][0])
    dp[0][0][2] = max(0, coins[0][0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k < 3; k++ {
                if i > 0 {
                    //dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k]+coins[i][j])
                    dp[i][j][k] = dp[i-1][j][k] + coins[i][j]
                    if coins[i][j] < 0 && k > 0 {
                        dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k-1])
                    }
                }
                if j > 0 {
                    dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k]+coins[i][j])
                    if coins[i][j] < 0 && k > 0 {
                        dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k-1])
                    }
                }
            }
        }
    }
    return dp[m-1][n-1][2]
}