func reportSpam(message []string, bannedWords []string) bool {
    m := make(map[string]bool, len(bannedWords))
    for _, word := range bannedWords {
        m[word] = true
    }
    cnt := 0
    for _, word := range message {
        if m[word] {
            cnt++
        }
        if cnt >= 2 {
            return true
        }
    }
    return false
}