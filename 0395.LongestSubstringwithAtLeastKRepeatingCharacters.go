func longestSubstring(s string, k int) int {
    n := len(s)
    if k <= 1 {
        return n
    }
    if n < k {
        return 0
    }
    cnt := make(map[byte]int)
    for i := range s {
        cnt[s[i]]++
    }
    var minCh byte
    minCnt := n+1
    for ch, c := range cnt {
        if c < minCnt {
            minCh, minCnt = ch, c
        }
    }
    if minCnt >= k {
        return n
    }
    subs := strings.Split(s, string(minCh))
    ans := 0
    for _, sub := range subs {
        ans = max(ans, longestSubstring(sub, k))
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}