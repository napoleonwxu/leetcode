func thousandSeparator(n int) string {
    s := strconv.Itoa(n)
    chs := make([]byte, 0, len(s))
    for i := range s {
        if i > 0 && (len(s)-i)%3 == 0 {
            chs = append(chs, '.')
        }
        chs = append(chs, s[i])
    }
    return string(chs)
    /*
    s := strconv.Itoa(n)
    strs := make([]string, 0, (len(s)+2)/3)
    start, end := 0, len(s)%3
    if end == 0 {
        end = 3
    }
    for end <= len(s) {
        strs = append(strs, s[start:end])
        start = end
        end += 3
    }
    return strings.Join(strs, ".")
    */
}