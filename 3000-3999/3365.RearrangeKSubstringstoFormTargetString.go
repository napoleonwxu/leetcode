func isPossibleToRearrange(s string, t string, k int) bool {
    n := len(s)
    l := n / k
    m := make(map[string]int)
    for i := 0; i < n; i += l {
        m[s[i:i+l]]++
        m[t[i:i+l]]--
    }
    for _, cnt := range m {
        if cnt != 0 {
            return false
        }
    }
    return true
}