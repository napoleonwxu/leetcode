func removeStars(s string) string {
    chs := make([]byte, len(s))
    i := 0
    for _, ch := range []byte(s) {
        if ch == '*' {
            if i > 0 {
                i--
            }
        } else {
            chs[i] = ch
            i++
        }
    }
    return string(chs[:i])
}