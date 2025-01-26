func calculateScore(s string) int64 {
    chs := []byte(s)
    m := make(map[byte][]int, 26)
    ans := int64(0)
    for i, ch := range chs {
        mirror := 25 - (ch-'a') + 'a'
        n := len(m[mirror])
        if n > 0 {
            ans += int64(i - m[mirror][n-1])
            m[mirror] = m[mirror][:n-1]
        } else {
            m[ch] = append(m[ch], i)
        }
    }
    return ans
}