const MOD = 1e9 + 7

func findPaths(m int, n int, N int, i int, j int) int {
    direction := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
    dp := make([][]int, m)
    for im := range dp {
        dp[im] = make([]int, n)
    }
    dp[i][j] = 1
    cnt := 0
    for k := 1; k <= N; k++ {
        next := make([][]int, m)
        for im := range next {
            next[im] = make([]int, n)
        }
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                if dp[i][j] == 0 {
                    continue
                }
                for _, d := range direction {
                    ii, jj := i + d[0], j + d[1]
                    if ii < 0 || ii >= m || jj < 0 || jj >= n {
                        cnt += dp[i][j]
                        cnt %= MOD
                    } else {
                        next[ii][jj] += dp[i][j]
                        next[ii][jj] %= MOD
                    }
                }
            }
        }
        dp = next
    }
    return cnt
}