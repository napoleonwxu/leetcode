func maxSumDivThree(nums []int) int {
    // O(n) + O(1), dp[i]: max sum that divided by 3 is i
    dp := []int{0, math.MinInt32, math.MinInt32}
    for _, num := range nums {
        nxt := make([]int, 3)
        for i := 0; i < 3; i++ {
            nxt[(num+i)%3] = max(dp[(num+i)%3], dp[i]+num)
        }
        dp = nxt
    }
    return dp[0]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}