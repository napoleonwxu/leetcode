func lengthOfLongestSubstring(s string) int {
    idx := make(map[byte]int)
    ans, start := 0, 0
    for i := 0; i < len(s); i++ {
        if pre, ok := idx[s[i]]; ok {
            start = max(start, pre+1)
        }
        ans = max(ans, i-start+1)
        idx[s[i]] = i
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}