func isValid(s string) bool {
    stack, n := make([]byte, len(s)), 0
    for i := 0; i < len(s); i++ {
        if s[i] == ')' || s[i] == ']' || s[i] == '}' {
            if n > 0 && (stack[n-1] == '(' && s[i] == ')' || 
                         stack[n-1] == '[' && s[i] == ']' || 
                         stack[n-1] == '{' && s[i] == '}') {
                n--
            } else {
                return false
            }
        } else {
            stack[n] = s[i]
            n++
        }
    }
    return n == 0
}