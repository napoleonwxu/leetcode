func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return nil
    }
    cnt := make([]int, 26)
    for _, ch := range p {
        cnt[ch-'a']++
    }
    ans := []int{}
    count := len(p)
    left, right := 0, 0
    for right < len(s) {
        if cnt[s[right]-'a'] > 0 {
            count--
        }
        cnt[s[right]-'a']--
        right++
        if count == 0 {
            ans = append(ans, left)
        }
        if right-left == len(p) {
            if cnt[s[left]-'a'] >= 0 {
                count++
            }
            cnt[s[left]-'a']++
            left++
        }
    }
    return ans
}