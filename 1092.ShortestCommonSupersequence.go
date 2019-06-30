func shortestCommonSupersequence(str1 string, str2 string) string {
    lcs := LCS(str1, str2)
    ans := []string{}
    i, j := 0, 0
    for k := 0; k < len(lcs); k++ {
        for ; str1[i] != lcs[k]; i++ {
            ans = append(ans, string(str1[i]))
        }
        for ; str2[j] != lcs[k]; j++ {
            ans = append(ans, string(str2[j]))
        }
        ans = append(ans, string(lcs[k]))
        i++
        j++
    }
    return strings.Join(ans, "") + str1[i:] + str2[j:]
}

func LCS(s1, s2 string) string {
    m, n := len(s1), len(s2)
    dp := make([][]string, m+1)
    for i := range dp {
        dp[i] = make([]string, n+1)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if s1[i] == s2[j] {
                dp[i+1][j+1] = dp[i][j] + string(s1[i])
            } else {
                if len(dp[i+1][j]) >= len(dp[i][j+1]) {
                    dp[i+1][j+1] = dp[i+1][j]
                } else {
                    dp[i+1][j+1] = dp[i][j+1]
                }
            }
        }
    }
    return dp[m][n]
}