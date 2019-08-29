func calculate(s string) int {
    n := len(s)
    stack := []int{}
    num, sign := 0, '+'
    for i, ch := range s {
        if ch >= '0' && ch <= '9' {
            num = 10*num + int(ch-'0')
        }
        if ch=='+' || ch=='-' || ch=='*' || ch=='/' || i==n-1 {
            switch sign {
                case '+':
                stack = append(stack, num)
                case '-':
                stack = append(stack, -num)
                case '*':
                stack[len(stack)-1] *= num
                case '/':
                stack[len(stack)-1] /= num
            }
            num, sign = 0, ch
        }
    }
    ans := 0
    for _, num := range stack {
        ans += num
    }
    return ans
}