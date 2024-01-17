func longestIdealString(s string, k int) int {
    dp := make([]int, 26)
    ans := 0
    for _, ch := range s {
        i := int(ch - 'a')
        for j := max(0, i-k); j <= min(25, i+k); j++ {
            dp[i] = max(dp[i], dp[j])
        }
        dp[i]++
        ans = max(ans, dp[i])
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}