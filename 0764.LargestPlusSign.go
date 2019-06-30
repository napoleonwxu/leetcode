func orderOfLargestPlusSign(N int, mines [][]int) int {
    Map := make(map[int]int)
    for _, mine := range mines {
        Map[mine[0]*N+mine[1]] = 1
    }
    dp := make([][]int, N)
    for i := range dp {
        dp[i] = make([]int, N)
    }
    ans := 0
    for i := 0; i < N; i++ {
        cnt := 0
        for j := 0; j < N; j++ {
            if Map[i*N+j] == 1 {
                cnt = 0
            } else {
                cnt++
            }
            dp[i][j] = cnt
        }
        cnt = 0
        for j := N-1; j >= 0; j-- {
            if Map[i*N+j] == 1 {
                cnt = 0
            } else {
                cnt++
            }
            dp[i][j] = min(dp[i][j], cnt)
        }
    }
    for j := 0; j < N; j++ {
        cnt := 0
        for i := 0; i < N; i++ {
            if Map[i*N+j] == 1 {
                cnt = 0
            } else {
                cnt++
            }
            dp[i][j] = min(dp[i][j], cnt)
        }
        cnt = 0
        for i := N-1; i >= 0; i-- {
            if Map[i*N+j] == 1 {
                cnt = 0
            } else {
                cnt++
            }
            dp[i][j] = min(dp[i][j], cnt)
            if dp[i][j] > ans {
                ans = dp[i][j]
            }
        }
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}