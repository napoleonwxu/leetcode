func numberOfArithmeticSlices(A []int) int {
    n := len(A)
    cnt := 0
    // O(n) + O(n)
    dp := make([]int, n)
    for i := 2; i < n; i++ {
        if A[i]-A[i-1] == A[i-1]-A[i-2] {
            dp[i] = dp[i-1] + 1
            cnt += dp[i]
        }
    }
    /* O(n^2) + O(1)
    for i := 1; i < n-1; i++ {
        for j := i+1; j < n; j++ {
            if A[j]-A[j-1] == A[j-1]-A[j-2] {
                cnt++
            } else {
                break
            }
        }
    }*/
    return cnt
}