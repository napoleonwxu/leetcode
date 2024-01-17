func isValid(s string) bool {
    stack := make([]byte, len(s))
    n := 0
    for _, ch := range []byte(s) {
        if ch == '(' || ch == '{' || ch == '[' {
            stack[n] = ch
            n++
        } else if n > 0 && (stack[n-1] == '(' && (ch == ')') ||
            (stack[n-1] == '{' && ch == '}') ||
            (stack[n-1] == '[' && ch == ']')) {
            n--
        } else {
            return false
        }
    }
    return n == 0
}