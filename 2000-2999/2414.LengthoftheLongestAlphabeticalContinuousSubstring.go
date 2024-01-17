func longestContinuousSubstring(s string) int {
    chs := []byte(s)
    cur, ans := 1, 1
    for i := 1; i < len(chs); i++ {
        if chs[i] == chs[i-1]+1 {
            cur++
            ans = max(ans, cur)
        } else {
            cur = 1
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}