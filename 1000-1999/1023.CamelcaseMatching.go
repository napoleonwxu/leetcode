func camelMatch(queries []string, pattern string) []bool {
    n := len(queries)
    ans := make([]bool, n)
    for i := 0; i < n; i++ {
        ans[i] = isMatch(queries[i], pattern)
        //ans[i] = match(queries[i], pattern)
    }
    return ans
}

// Two pointer, O(m+n)
func isMatch(word string, pattern string) bool {
    j, n := 0, len(pattern)
    for i := 0; i < len(word); i++ {
        if j < n && word[i] == pattern[j] {
            j++
        } else if word[i] < 'a' {
            return false
        }
    }
    return j == n
}

// DP, O(mn)
func match(word string, pattern string) bool {
    m, n := len(word), len(pattern)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true
    for j := 1; j <= n; j++ {
        dp[0][j] = false
    }
    for i := 1; i <= m; i++ {
        if word[i-1] >= 'a' && word[i-1] <= 'z' {
            dp[i][0] = dp[i-1][0]
        } else {
            dp[i][0] = false
        }
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word[i-1] == pattern[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else if word[i-1] >= 'a' && word[i-1] <= 'z' {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = false
            }
        }
    }
    return dp[m][n]
}