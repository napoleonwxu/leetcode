func longestArithSeqLength(A []int) int {
    // O(n^2) + O(n^2)
    n := len(A)
    ans := 2
    dp := make([]map[int]int, n)
    for i := 1; i < n; i++ {
        dp[i] = make(map[int]int)
        for j := 0; j < i; j++ {
            d := A[i] - A[j]
            //dp[i][d] = max(dp[i][d], max(dp[j][d]+1, 2))
            dp[i][d] = max(dp[j][d]+1, 2)
            ans = max(ans, dp[i][d])
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}