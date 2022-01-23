func maxProfit(k int, prices []int) int {
    n := len(prices)
    if k == 0 || n < 2 {
        return 0
    }
    if k >= n/2 {
        profit := 0
        for i := 1; i < n; i++ {
            profit += max(0, prices[i]-prices[i-1])
        }
        return profit
    }
    // O(kn) + O(kn)
    dp := make([][]int, k+1)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i := 1; i <= k; i++ {
        profit := -prices[0]
        for j := 1; j < n; j++ {
            profit = max(profit, dp[i-1][j-1] - prices[j]) // buy
            dp[i][j] = max(dp[i][j-1], profit + prices[j]) // sell
        }
    }
    return dp[k][n-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}