func longestPalindrome(s string) string {
    // O(n^2)
    ans := ""
    i, n := 0, len(s)
    for i < n {
        l, r := i-1, i+1
        for ; r < n && s[r] == s[i]; r++ {}
        i = r
        for ; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {}
        if r-l-1 > len(ans) {
            ans = s[l+1:r]
        }
    }
    return ans
    /* Expand, O(n^2) + O(1)
    n := len(s)
    idx, maxLen := 0, 0
    for i := 0; i < n && n-i > maxLen>>1; i++ {
        expand(s, i, i, &idx, &maxLen)
        expand(s, i, i+1, &idx, &maxLen)
    }
    return s[idx:idx+maxLen]
    */
    /* DP, O(n^2) + O(n^2)
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }
    ans := ""
    for i := n-1; i >= 0; i-- {
        for j := i; j < n; j++ {
            dp[i][j] = s[i]==s[j] && (j-i<2 || dp[i+1][j-1])
            if dp[i][j] && j-i+1>len(ans) {
                ans = s[i:j+1]
            }
        }
    }
    return ans
    */
}

func expand(s string, i, j int, idx, maxLen *int) {
    for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {}
    if j-i-1 > *maxLen {
        *idx, *maxLen = i+1, j-i-1
    }
}
