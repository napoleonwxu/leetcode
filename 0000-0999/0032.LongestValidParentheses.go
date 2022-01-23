func longestValidParentheses(s string) int {
    maxLen := 0
    stack := []int{-1}
    for i, ch := range s {
        if ch == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) > 0 {
                maxLen = max(maxLen, i-stack[len(stack)-1])
            } else {
                stack = append(stack, i)
            }
        }
    }
    return maxLen
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}