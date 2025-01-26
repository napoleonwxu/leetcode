func countKeyChanges(s string) int {
    cnt := 0
    for i := 1; i < len(s); i++ {
        if s[i] != s[i-1] && s[i]+'a' != s[i-1]+'A' && s[i]+'A' != s[i-1]+'a' {
            cnt++
        }
    }
    return cnt
}