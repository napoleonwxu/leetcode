func numberOfArithmeticSlices(A []int) int {
    n := len(A)
    dp := make([]int, n)
    cnt := 0
    for i := 2; i < n; i++ {
        if A[i]-A[i-1] == A[i-1]-A[i-2] {
            dp[i] = dp[i-1] + 1
            cnt += dp[i]
        }
        //cnt += dp[i]
    }
    return cnt
}