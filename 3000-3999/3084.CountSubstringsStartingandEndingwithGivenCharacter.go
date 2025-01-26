func countSubstrings(s string, c byte) int64 {
    cnt := int64(0)
    for i := range s {
        if s[i] == c {
            cnt++
        }
    }
    return cnt*(cnt+1)/2
}