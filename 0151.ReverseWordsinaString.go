func reverseWords(s string) string {
    words := strings.Split(s, " ")
    ans := make([]string, 0, len(words))
    for i := len(words)-1; i >= 0; i-- {
        if words[i] != "" {
            ans = append(ans, words[i])
        }
    }
    return strings.Join(ans, " ")
}