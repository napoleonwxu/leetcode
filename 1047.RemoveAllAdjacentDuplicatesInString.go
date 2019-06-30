func removeDuplicates(S string) string {
    n := len(S)
    stack := make([]string, n)
    i := -1
    for _, ch := range S {
        s := string(ch)
        if i >= 0 && s == stack[i] {
            i--
        } else {
            i++
            stack[i] = s
        }
    }
    return strings.Join(stack[:i+1], "")
}