func lengthOfLongestSubstring(s string) int {
    Map := make(map[rune]int)
    ans, idx := 0, 0
    for i, ch := range s {
        if _, ok := Map[ch]; ok {
            idx = max(idx, Map[ch]+1)
        }
        ans = max(ans, i-idx+1)
        Map[ch] = i
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}