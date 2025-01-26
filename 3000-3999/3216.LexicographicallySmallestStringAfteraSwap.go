func getSmallestString(s string) string {
    chs := []byte(s)
    for i := 1; i < len(chs); i++ {
        if (chs[i]-'0')%2 == (chs[i-1]-'0')%2 && chs[i] < chs[i-1] {
            chs[i], chs[i-1] = chs[i-1], chs[i]
            return string(chs)
        }
    }
    return s
}