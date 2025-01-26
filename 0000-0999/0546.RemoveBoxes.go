func removeBoxes(boxes []int) int {
    n := len(boxes)
    if n == 0 {
        return 0
    }
    dp := make([][][]int, n)
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, n)
        }
    }
    for i := 0; i < n; i++ {
        for k := 0; k <= i; k++ {
            dp[i][i][k] = (k+1) * (k+1)
        }
    }
    for l := 1; l < n; l++ {
        for j := l; j < n; j++ {
            i := j - l
            for k := 0; k <= i; k++ {
                dp[i][j][k] = (k+1) * (k+1) + dp[i+1][j][0]
                for m := i+1; m <= j; m++ {
                    if boxes[m] == boxes[i] {
                        dp[i][j][k] = max(dp[i][j][k], dp[i+1][m-1][0]+dp[m][j][k+1])
                    }
                }
            }
        }
    }
    return dp[0][n-1][0]
}