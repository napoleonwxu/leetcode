func minWindow(s string, t string) string {
    Map := make(map[byte]int)
    for i := range t {
        Map[t[i]]++
    }
    cnt := len(t)
    idx, n := 0, len(s)+1
    l := 0
    for r := range s {
        if Map[s[r]] > 0 {
            cnt--
        }
        Map[s[r]]--
        for cnt == 0 {
            if r-l+1 < n {
                idx, n = l, r-l+1
            }
            if Map[s[l]] == 0 {
                cnt++
            }
            Map[s[l]]++
            l++
        }
    }
    if n == len(s)+1 {
        return ""
    }
    return s[idx:idx+n]
}