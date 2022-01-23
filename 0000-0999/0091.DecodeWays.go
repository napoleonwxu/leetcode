func numDecodings(s string) int {
    if s == "" || s[0] == '0' {
        return 0
    }
    n := len(s)
    dp := make([]int, n+1)
    dp[0] = 1
    for i := 1; i <= n; i++ {
        if s[i-1] != '0' {
            dp[i] += dp[i-1]
        }
        //if i >= 2 && s[i-2] != '0' && s[i-2:i] <= "26" {
        if i >= 2 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
            dp[i] += dp[i-2]
        }
    }
    return dp[n]
}