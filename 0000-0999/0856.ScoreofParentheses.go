func scoreOfParentheses(S string) int {
    stack := []int{0}
    for _, ch := range S {
        if ch == '(' {
            stack = append(stack, 0)
        } else {
            n := len(stack)
            stack[n-2] += max(1, 2*stack[n-1])
            /*
            if stack[n-1] == 0 {
                stack[n-2]++
            } else {
                stack[n-2] += 2 * stack[n-1]
            }
            */
            stack = stack[:n-1]
        }
    }
    return stack[0]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}