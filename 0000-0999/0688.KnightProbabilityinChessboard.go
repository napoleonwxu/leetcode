func knightProbability(N int, K int, r int, c int) float64 {
    step := [][]int{{-2,-1},{-2,1},{-1,-2},{-1,2},{1,-2},{1,2},{2,-1},{2,1}}
    dp := make([][]float64, N)
    for i := range dp {
        dp[i] = make([]float64, N)
    }
    dp[r][c] = 1.0
    for k := 1; k <= K; k++ {
        nxt := make([][]float64, N)
        for i := range nxt {
            nxt[i] = make([]float64, N)
        }
        for i := 0; i < N; i++ {
            for j := 0; j < N; j++ {
                p := 0.0
                for _, s := range step {
                    ii, jj := i + s[0], j + s[1]
                    if ii >= 0 && ii < N && jj >=0 && jj < N {
                        p += dp[ii][jj]
                    }
                }
                nxt[i][j] = p/8
            }
        }
        dp = nxt
    }
    ans := 0.0
    for i := range dp {
        for j := range dp[i] {
            ans += dp[i][j]
        }
    }
    return ans
}