func minDistance(word1 string, word2 string) int {
    n1, n2 := len(word1), len(word2)
    dp := make([][]int, n1+1)
    for i := range dp {
        dp[i] = make([]int, n2+1)
    }
    // dp[i][j]: w1[:i] & w2[:j]
    for i := 0; i <= n1; i++ {
        for j := 0; j <= n2; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = i + j
            } else if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }
    return dp[n1][n2]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}