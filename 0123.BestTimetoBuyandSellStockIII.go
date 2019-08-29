// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/135704/Detail-explanation-of-DP-solution
func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    // DP: dp[k, i] = max(dp[k, i-1], prices[i] - prices[j] + dp[k-1, j-1]), j=[0..i-1]
    // O(KN) + O(KN), K is 2 here
    K, N := 2, len(prices)
    dp := make([][]int, K+1)
    for k := range dp {
        dp[k] = make([]int, N)
    }
    for k := 1; k <= K; k++ {
        profit := -prices[0]
        for i := 1; i < N; i++ {
            profit = max(profit, dp[k-1][i-1]-prices[i])
            dp[k][i] = max(dp[k][i-1], prices[i]+profit)
        }
    }
    return dp[K][N-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}