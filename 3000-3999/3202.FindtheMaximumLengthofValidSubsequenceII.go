func maximumLength(nums []int, k int) int {
    dp := make([][]int, k)
    for i := range dp {
        dp[i] = make([]int, k)
    }
    ans := 0
    for _, num := range nums {
        for j := 0; j < k; j++ {
            dp[num%k][j] = dp[j][num%k] + 1
            ans = max(ans, dp[num%k][j])
        }
    }
    return ans
}