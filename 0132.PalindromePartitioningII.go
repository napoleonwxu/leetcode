func minCut(s string) int {
    // O(n^2) + O(n)
    n := len(s)
    dp := make([]int, n+1)
    for i := range dp {
        dp[i] = i-1
    }
    for i := range dp {
        // odd length palindrome
        for l, r := i, i; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
            dp[r+1] = min(dp[r+1], dp[l]+1)
        }
        // even length palindrome
        for l, r := i-1, i; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
            dp[r+1] = min(dp[r+1], dp[l]+1)
        }
    }
    return dp[n]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}