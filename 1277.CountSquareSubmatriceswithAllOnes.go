func countSquares(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    cnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                if i == 0 || j == 0 {
                    dp[i][j] = matrix[i][j]
                } else {
                    dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
                }
                cnt += dp[i][j]
            }
        }
    }
    return cnt
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}