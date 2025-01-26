func getSmallestString(s string, k int) string {
    t := []byte(s)
    for i := 0; i < len(s) && k > 0; i++ {
        step := min(int(t[i]-'a'), int('a'+26-t[i]))
        if k >= step {
            t[i] = 'a'
            k -= step
        } else {
            t[i] = t[i] - byte(k)
            k = 0
        }
    }
    return string(t)
}