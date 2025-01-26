func maximumLengthSubstring(s string) int {
    cnt := make(map[byte]int, 26)
    ans, start := 0, 0
    for i := range s {
        cnt[s[i]]++
        for cnt[s[i]] > 2 {
            cnt[s[start]]--
            start++
        }
        ans = max(ans, i-start+1)
    }
    return ans
}