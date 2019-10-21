func climbStairs(n int) int {
    pre, cur := 1, 1
    for i := 2; i <= n; i++ {
        pre, cur = cur, pre+cur
    }
    return cur
}