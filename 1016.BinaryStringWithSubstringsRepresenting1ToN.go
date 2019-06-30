func queryString(S string, N int) bool {
    //for i := N; i > 0; i-- {
    for i := N; i > N>>1; i-- {
        n, t := i, ""
        for n > 0 {
            t = strconv.Itoa(n&1) + t
            n >>= 1
        }
        if !strings.Contains(S, t) {
            return false
        }
    }
    return true
}