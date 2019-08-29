func longestWord(words []string) string {
    sort.Strings(words)
    Map := make(map[string]bool)
    ans := ""
    for _, word := range words {
        n := len(word)
        if n == 1 || Map[word[:n-1]] {
            if n > len(ans) {
                ans = word
            }
            Map[word] = true
        }
    }
    return ans
}