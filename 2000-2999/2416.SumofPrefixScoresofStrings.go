func sumPrefixScores(words []string) []int {
    cnts := make(map[string]int)
    for _, w := range words {
        for i := 1; i <= len(w); i++ {
            cnts[w[:i]]++
        }
    }
    ans := make([]int, len(words))
    for i, w := range words {
        for j := 1; j <= len(w); j++ {
            ans[i] += cnts[w[:j]]
        }
    }
    return ans
}