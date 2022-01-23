func numberOfSubstrings(s string) int {
    cnt := []int{0, 0, 0}
    sum, i := 0, 0
    for _, ch := range s {
        cnt[ch-'a']++
        for ; cnt[0]>0 && cnt[1]>0 && cnt[2]>0; i++ {
            cnt[s[i]-'a']--
        }
        sum += i
    }
    return sum
}