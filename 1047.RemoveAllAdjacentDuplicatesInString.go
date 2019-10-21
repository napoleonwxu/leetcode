func removeDuplicates(S string) string {
    stack := make([]byte, len(S))
    n := 0
    for i := 0; i < len(S); i++ {
        if n > 0 && S[i] == stack[n-1] {
            n--
        } else {
            stack[n] = S[i]
            n++
        }
    }
    return string(stack[:n])
}