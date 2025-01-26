func lengthOfLongestSubstring(s string) int {
    m := make(map[byte]int)
    ans, start := 0, -1
    for i, ch := range []byte(s) {
        if pre, ok := m[ch]; ok {
            start = max(start, pre)
        }
        ans = max(ans, i-start)
        m[ch] = i
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}