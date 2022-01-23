func sortString(s string) string {
    cnt := make([]int, 26)
    for _, ch := range s {
        cnt[ch-'a']++
    }
    ans := make([]byte, len(s))
    for i := 0; i < len(ans); {
        for j := 0; j < len(cnt); j++ {
            if cnt[j] > 0 {
                ans[i] = 'a' + byte(j)
                cnt[j]--
                i++
            }
        }
        for j := len(cnt)-1; j >= 0; j-- {
            if cnt[j] > 0 {
                ans[i] = 'a' + byte(j)
                cnt[j]--
                i++
            }
        }
    }
    return string(ans)
}