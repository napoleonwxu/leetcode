func minSteps(s string, t string) int {
    cnt := make([]int, 26)
    for _, ch := range s {
        cnt[ch-'a']++
    }
    for _, ch := range t {
        cnt[ch-'a']--
    }
    ans := 0
    for _, c := range cnt {
        if c > 0 {
            ans += c
        }
    }
    return ans
}
