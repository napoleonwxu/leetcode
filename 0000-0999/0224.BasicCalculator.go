func calculate(s string) int {
    sum := 0
    op, sign := 0, 1
    stack := []int{}
    for _, ch := range s {
        if ch >= '0' && ch <= '9' {
            op = 10*op + int(ch-'0')
        } else if ch == '+' {
            sum += sign * op
            op, sign = 0, 1
        } else if ch == '-' {
            sum += sign * op
            op, sign = 0, -1
        } else if ch == '(' {
            stack = append(stack, sum, sign)
            sum = 0
            sign = 1 // op is 0 now due to pre '(' must be '+/-', ' ', ''
        } else if ch == ')' {
            sum += sign * op
            n := len(stack)
            sum *= stack[n-1]
            sum += stack[n-2]
            stack = stack[:n-2]
            op = 0  // no need to reset sign due to op is 0 and post ')' must be '+/-', ' ', ''
        }
    }
    return sum + sign*op
}