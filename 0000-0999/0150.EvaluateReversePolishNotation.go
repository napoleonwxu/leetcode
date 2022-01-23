func evalRPN(tokens []string) int {
    stack := make([]int, 0, len(tokens))
    for _, t := range tokens {
        num, ok := strconv.Atoi(t)
        if ok == nil {
            stack = append(stack, num)
        } else {
            n := len(stack)
            if t == "+" {
                stack[n-2] += stack[n-1]
            } else if t == "-" {
                stack[n-2] -= stack[n-1]
            } else if t == "*" {
                stack[n-2] *= stack[n-1]
            } else if t == "/" {
                stack[n-2] /= stack[n-1]
            }
            stack = stack[:n-1]
        }
    }
    return stack[0]
}