func minWindow(s string, t string) string {
    Map := make(map[byte]int)
    for i, _ := range t {
        Map[t[i]]++
    }
    cnt := len(t)
    start, n := 0, len(s)+1
    idx := 0
    for i, _ := range s {
        if Map[s[i]] > 0 {
            cnt--
        }
        Map[s[i]]--
        for cnt == 0 {
            if i-idx+1 < n {
                start, n = idx, i-idx+1
            }
            if Map[s[idx]] == 0 {
                cnt++
            }
            Map[s[idx]]++
            idx++
        }
    }
    if n == len(s)+1 {
        return ""
    }
    return s[start:start+n]
}