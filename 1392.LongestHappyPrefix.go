func longestPrefix(s string) string {
    n := len(s)
    for i := n-1; i > 0; i-- {
        if s[:i] == s[n-i:] {
            return s[:i]
        }
    }
    return ""
}