func isAcronym(words []string, s string) bool {
    chs := make([]byte, len(words))
    for i, word := range words {
        chs[i] = word[0]
    }
    return string(chs) == s
}