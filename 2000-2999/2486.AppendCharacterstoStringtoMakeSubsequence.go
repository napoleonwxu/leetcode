func appendCharacters(s string, t string) int {
    chs := []byte(t)
    idx, n := 0, len(chs)
    for _, ch := range []byte(s) {
        if idx < n && chs[idx] == ch {
            idx++
        }
    }
    return n - idx
}