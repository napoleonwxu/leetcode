func doesAliceWin(s string) bool {
    for _, ch := range s {
        if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
            return true
        }
    }
    return false
}