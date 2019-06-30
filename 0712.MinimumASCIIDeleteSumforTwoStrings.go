func minimumDeleteSum(s1 string, s2 string) int {
    n1, n2 := len(s1), len(s2)
    dp := make([][]int, n1+1)
    for i := range dp {
        dp[i] = make([]int, n2+1)
    }
    for i := n1-1; i >= 0; i-- {
        dp[i][n2] = int(s1[i]) + dp[i+1][n2]
    }
    for j := n2-1; j >= 0; j-- {
        dp[n1][j] = int(s2[j]) + dp[n1][j+1]
    }
    for i := n1-1; i >= 0; i-- {
        for j := n2-1; j >= 0; j-- {
            if s1[i] == s2[j] {
                dp[i][j] = dp[i+1][j+1]
            } else {
                dp[i][j] = min(int(s1[i])+dp[i+1][j], int(s2[j])+dp[i][j+1])
            }
        }
    }
    return dp[0][0]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}