func subarraysMostTwoDistinct(s string) int {
	k := 2
    cnt := make(map[byte]int)
    ans := 0
    for i, j := 0, 0; j < len(s); j++ {
        if cnt[s[j]] == 0 {
            k--
        }
        cnt[s[j]]++
        for ; k < 0; i++ {
            cnt[s[i]]--
            if cnt[s[i]] == 0 {
                k++
            }
        }
        ans += j - i + 1
    }
    return ans
}