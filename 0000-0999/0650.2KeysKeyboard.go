func minSteps(n int) int {
    // math
    cnt := 0
    i := 2
    for n > 1 {
        for n%i == 0 {
            cnt += i
            n /= i
        }
        i++
    }
    return cnt
    /* dp2, simpler
    dp := make([]int, n+1)
    for i := 2; i <= n; i++ {
        dp[i] = i
        for j := i/2; j > 1; j-- {
            if i%j == 0 {
                dp[i] = dp[j] + i/j
                break
            }
        }
    }
    return dp[n]
    */
    /* dp1
    dp := make([]int, n+1)
    for i := 2; i <= n; i++ {
        dp[i] = i
        for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
            if i%j == 0 {
                k := i/j
                dp[i] = min(dp[i], min(dp[j]+k, dp[k]+j))
                // actually dp[j]+k == dp[k]+j
            }
        }
    }
    return dp[n]
    */
}

func min(x int, y int) int {
    if x < y {
        return x
    }
    return y
}