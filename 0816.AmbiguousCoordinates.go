func ambiguousCoordinates(S string) []string {
    ans := []string{}
    n := len(S)
    for i := 2; i < n-1; i++ {
        xs, ys := build(S[1:i]), build(S[i:n-1])
        for _, x := range xs {
            for _, y := range ys {
                //ans = append(ans, "(" + x + ", " + y + ")")
                ans = append(ans, connect(x, y))
            }
        }
    }
    return ans
}

func build(s string) []string {
    ans := []string{}
    n := len(s)
    for i := 1; i < n+1; i++ {
        a, b := s[:i], s[i:]
        if (a == "0" || !strings.HasPrefix(a, "0")) && !strings.HasSuffix(b, "0") {
            if i < n {
                ans = append(ans, s[:i] + "." + s[i:])
            } else {
                ans = append(ans, s)
            }
        }
    }
    return ans
}

func connect(x, y string) string {
    return fmt.Sprintf("(%s, %s)", x, y)
}