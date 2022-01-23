func subarraysMostKDistinct(s string, K int) int {
    cnt := make(map[byte]int)
    ans := 0
    for i, j := 0, 0; j < len(s); j++ {
        if cnt[s[j]] == 0 {
            K--
        }
        cnt[s[j]]++
        for ; K < 0; i++ {
            cnt[s[i]]--
            if cnt[s[i]] == 0 {
                K++
            }
        }
        ans += j - i + 1
    }
    return ans
}