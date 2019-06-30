func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    m, n := len(matrix)+1, len(matrix[0])+1
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    Len := 0
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
                if dp[i][j] > Len {
                    Len = dp[i][j]
                }
            }
        }
    }
    return Len*Len
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}