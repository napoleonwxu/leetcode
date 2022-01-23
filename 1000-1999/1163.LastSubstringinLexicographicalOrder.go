func lastSubstring(s string) string {
    ans := s
    for i := 1; i < len(s); i++ {
        if s[i:] > ans {
            ans = s[i:]
        }
    }
    return ans
}