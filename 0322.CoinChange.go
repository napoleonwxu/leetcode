func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for money := 1; money <= amount; money++ {
        dp[money] = amount + 1
    }
    for money := 1; money <= amount; money++ {
        for _, coin := range coins {
            if coin <= money {
                dp[money] = min(dp[money], dp[money-coin]+1)
            }
        }
    }
    if dp[amount] <= amount {
        return dp[amount]
    }
    return -1
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}