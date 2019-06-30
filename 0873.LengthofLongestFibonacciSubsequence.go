func lenLongestFibSubseq(A []int) int {
    n := len(A)
    ans := 0
    index := make(map[int]int)
    /*
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for k := 0; k < n; k++ {
        index[A[k]] = k
        for j := 0; j < k; j++ {
            i, ok := index[A[k]-A[j]]
            if ok && i < j {
                if dp[i][j] > 0 {
                    dp[j][k] = dp[i][j] + 1
                } else {
                    dp[j][k] = 3
                }
            }
            if dp[j][k] > ans {
                ans = dp[j][k]
            }
        }
    }
    return ans
    */
    cnt := make(map[int]int)
    for k := 0; k < n; k++ {
        index[A[k]] = k
        for j := 0; j < k; j++ {
            i, ok := index[A[k]-A[j]]
            if ok && i < j {
                tmp, ok := cnt[i*n+j]
                if ok {
                    tmp += 1
                } else {
                    tmp = 3
                }
                cnt[j*n+k] = tmp
                if tmp > ans {
                    ans = tmp
                }
            }
        }
    }
    return ans
}