func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    // O(m+n) + O(1)
    i, j := 0, 0
    match, start := 0, -1
    for i < m {
        if j < n && (s[i] == p[j] || p[j] == '?') {
            i, j = i+1, j+1
        } else if j < n && p[j] == '*' {
            match, start = i, j
            j++
        } else if start != -1 {
            i, j = match+1, start+1
            match++
        } else {
            return false
        }
    }
    for j < n && p[j] == '*' {
        j++
    }
    return j == n
    /* DP, O(mn) + O(mn)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true
    for j := 1; j <= n && p[j-1] == '*'; j++ {
        dp[0][j] = true
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '?' {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                dp[i][j] = dp[i][j-1] || dp[i-1][j]
            }
        }
    }
    return dp[m][n]
    */
}