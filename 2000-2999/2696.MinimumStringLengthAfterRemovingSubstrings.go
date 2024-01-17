func minLength(s string) int {
    stack := make([]byte, len(s))
    i := 0
    for _, ch := range []byte(s) {
        if i > 0 && ((stack[i-1] == 'A' && ch == 'B') || (stack[i-1] == 'C' && ch == 'D')) {
            i--
        } else {
            stack[i] = ch
            i++
        }
    }
    return i
    /*
    for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
        s = strings.Replace(s, "AB", "", -1)
        s = strings.Replace(s, "CD", "", -1)
    }
    return len(s)
    */
}