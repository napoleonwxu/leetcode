func numDecodings(s string) int {
    if s == "" || s[0] == '0' {
        return 0
    }
    n := len(s)
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        if s[i-1] != '0' {
            dp[i] = dp[i-1]
        }
        //if s[i-2] != '0' && 10*int(s[i-2]-'0')+int(s[i-1]-'0') <= 26 {
        if s[i-2] != '0' && s[i-2:i] <= "26" {
            dp[i] += dp[i-2]
        }
    }
    return dp[n]
}