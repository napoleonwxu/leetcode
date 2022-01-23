// same with 316. Remove Duplicate Letters
func smallestSubsequence(text string) string {
    cnt, used := make([]int, 26), make([]int, 26)
    for _, ch := range text {
        cnt[ch-'a']++
    }
    ans, n := "", 0
    for _, ch := range text {
        cnt[ch-'a']--
        if used[ch-'a'] > 0 {
            continue
        }
        used[ch-'a']++
        for n > 0 && ans[n-1] > byte(ch) && cnt[ans[n-1]-'a'] > 0 {
            used[ans[n-1]-'a']--
            ans = ans[:n-1]
            n--
        }
        ans += string(ch)
        n++
    }
    return ans
}