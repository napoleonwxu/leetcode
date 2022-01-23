func numberOfArithmeticSlices(A []int) int {
    // O(n^2) + O(n^2)
    n := len(A)
    dp := make([]map[int]int, n)
    cnt := 0
    for i := 1; i < n; i++ {
        dp[i] = make(map[int]int)
        for j := 0; j < i; j++ {
            d := A[i] - A[j]
            if d <= math.MinInt32 || d >= math.MaxInt32 {
                continue
            }
            cnt += dp[j][d]
            dp[i][d] += dp[j][d] + 1
        }
    }
    return cnt
}