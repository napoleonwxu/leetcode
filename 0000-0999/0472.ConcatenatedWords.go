func findAllConcatenatedWordsInADict(words []string) []string {
    // O(N*L^3), N: len(words); L: avg len of words
    ans := []string{}
    Map := make(map[string]bool)
    for _, word := range words {
        Map[word] = true
    }
    for _, word := range words {
        Map[word] = false
        if canForm(word, Map) {
            ans = append(ans, word)
        }
        Map[word] = true
    }
    return ans
}

func canForm(word string, Map map[string]bool) bool {
    if len(word) == 0 {
        return false
    }
    n := len(word)
    dp := make([]bool, n+1)
    dp[0] = true
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            if dp[j] && Map[word[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[n]
}