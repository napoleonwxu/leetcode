func findTheLongestBalancedSubstring(s string) int {
    n, ans := len(s), 0
    for i := 0; i < n; {
        zero, one := 0, 0
        for ; i < n && s[i] == '0'; i++ {
            zero++
        }
        for ; i < n && s[i] == '1'; i++ {
            one++
        }
        ans = max(ans, 2*min(zero, one))
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}