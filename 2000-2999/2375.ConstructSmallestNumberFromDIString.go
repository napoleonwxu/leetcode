func smallestNumber(pattern string) string {
    n := len(pattern)
    ans := make([]byte, 0, n+1)
    stack := make([]byte, 0, n+1)
    for i := 0; i <= n; i++ {
        stack = append(stack, '1'+byte(i))
        if i == n || pattern[i] == 'I' {
            for j := len(stack) - 1; j >= 0; j-- {
                ans = append(ans, stack[j])
            }
            stack = make([]byte, 0, n-i)
        }
    }
    return string(ans)
}