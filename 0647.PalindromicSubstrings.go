func countSubstrings(s string) int {
    n := len(s)
    cnt := 0
    // Expand From Center, O(n*n) + O(1)
    for i := 0; i < n<<1-1; i++ {
        l, r := i>>1, i>>1 + i&1
        for ; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
            cnt++
        }
    }
    /* DP, O(n*n) + O(n*n)
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }
    for i := n-1; i >= 0; i-- {
        for j := i; j < n; j++ {
            if s[i] == s[j] && (i+1 >= j-1 || dp[i+1][j-1]) {
                dp[i][j] = true
                cnt++
            }
        }
    }
    */
    return cnt
}

func isPalindromic(s string, i, j int) bool {
    for ; i < j && s[i] == s[j]; i, j = i+1, j-1 {}
    return i >= j
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}