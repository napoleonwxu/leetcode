func isValid(S string) bool {
    // O(n)
    stack := []rune{}
    for _, ch := range S {
        if ch == 'c' {
            n := len(stack)
            if n < 2 || stack[n-1] != 'b' || stack[n-2] != 'a' {
                return false
            }
            stack = stack[:n-2]
        } else {
            stack = append(stack, ch)
        }
    }
    return len(stack) == 0
    /* O(n^2)
    for strings.Contains(S, "abc") {
        S = strings.Replace(S, "abc", "", -1)
    }
    return S == ""
    */
}