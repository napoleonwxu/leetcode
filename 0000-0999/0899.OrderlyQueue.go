func orderlyQueue(s string, k int) string {
    if k > 1 {
        chs := []byte(s)
        sort.Slice(chs, func(i, j int) bool {
            return chs[i] < chs[j]
        })
        return string(chs)
    }
    ans := s
    for i := 1; i < len(s); i++ {
        t := string(s[i:]) + string(s[:i])
        if strings.Compare(t, ans) < 0 {
            ans = t
        }
    }
    return ans
}