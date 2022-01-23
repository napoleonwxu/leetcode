// Similar with 44. Wildcard Matching 
// https://leetcode.com/problems/regular-expression-matching/discuss/5651/Easy-DP-Java-Solution-with-detailed-Explanation
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true
    for j := 1; j <= n; j++ {
        dp[0][j] = (p[j-1] == '*' && dp[0][j-2])
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '.' {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                if s[i-1] == p[j-2] || p[j-2] == '.' {
                    dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i][j-2]
                } else {
                    dp[i][j] = dp[i][j-2]
                }
            }
        }
    }
    return dp[m][n]
}