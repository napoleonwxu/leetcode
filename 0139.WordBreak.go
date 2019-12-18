func wordBreak(s string, wordDict []string) bool {
    // O(NML), N: len(s); M: len(wordDict); L: avg len of words in wordDict
    n := len(s)
    dp := make([]bool, n+1)
    dp[0] = true
    for i := 1; i <= n; i++ {
        for _, word := range wordDict {
            m := len(word)
            if i >= m && s[i-m:i] == word && dp[i-m] {
                dp[i] = true
                break
            }
        }
    }
    return dp[n]
}