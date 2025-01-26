func clearStars(s string) string {
    n := len(s)
    removed := make([]bool, n)
    buckets := make([][]int, 26)
    for i := range s {
        if s[i] == '*' {
            removed[i] = true
            for j, b := range buckets {
                if len(b) > 0 {
                    m := len(b)
                    removed[b[m-1]] = true
                    buckets[j] = buckets[j][:m-1]
                    break
                }
            }
        } else {
            idx := int(s[i]-'a')
            buckets[idx] = append(buckets[idx], i)
        }
    }
    chs := make([]byte, 0, n)
    for i := range s {
        if !removed[i] {
            chs = append(chs, s[i])
        }
    }
    return string(chs)
}