func reverseParentheses(s string) string {
    stack := make([]byte, len(s))
    i := -1
    for j := 0; j < len(s); j++ {
        if s[j] == ')' {
            queue := []byte{}
            for i >= 0 && stack[i] != '(' {
                queue = append(queue, stack[i])
                i--
            }
            if i >= 0 {
                i--
            }
            for _, char := range queue {
                i++
                stack[i] = char
            }
        } else {
            i++
            stack[i] = s[j]
        }
    }
    return string(stack[:i+1])
}
