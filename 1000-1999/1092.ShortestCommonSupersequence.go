func shortestCommonSupersequence(str1 string, str2 string) string {
    // DP, O(mn)
    lcs := LCS(str1, str2)
    i, j := 0, 0
    ans := []byte{}
    for k := 0; k < len(lcs); k++ {
        for ; str1[i] != lcs[k]; i++ {
            ans = append(ans, str1[i])
        }
        for ; str2[j] != lcs[k]; j++ {
            ans = append(ans, str2[j])
        }
        ans = append(ans, lcs[k])
        i, j = i+1, j+1
    }
    return string(ans) + str1[i:] + str2[j:]
}

func LCS(s, t string) string {
    m, n := len(s), len(t)
    dp := make([][]string, m+1)
    for i := range dp {
        dp[i] = make([]string, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i-1] == t[j-1] {
                dp[i][j] = dp[i-1][j-1] + string(s[i-1])
            } else {
                if len(dp[i-1][j]) >= len(dp[i][j-1]) {
                    dp[i][j] = dp[i-1][j]
                } else {
                    dp[i][j] = dp[i][j-1]
                }
            }
        }
    }
    return dp[m][n]
}