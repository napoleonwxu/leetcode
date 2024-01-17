func smallestString(s string) string {
    chs := []byte(s)
    i, n := 0, len(chs)
    for ; i < n && chs[i] == 'a'; i++ {
    }
    if i >= n && n > 0 {
        chs[n-1] = 'z'
        return string(chs)
    }
    for ; i < n && chs[i] != 'a'; i++ {
        chs[i]--
    }
    return string(chs)
}