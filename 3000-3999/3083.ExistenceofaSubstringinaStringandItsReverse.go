func isSubstringPresent(s string) bool {
    m := make(map[string]bool)
    for i := 1; i < len(s); i++ {
        m[s[i-1:i+1]] = true
        if m[string([]byte{s[i], s[i-1]})] {
            return true
        }
    }
    return false
}