import "strings"

func minimumString(a string, b string, c string) string {
    ans := minString(solve(a, b, c), solve(b, c, a))
    return minString(ans, solve(c, a, b))
}

func solve(a, b, c string) string {
    s1, s2 := addString(a, b), addString(b, a)
    ans1 := minString(addString(s1, c), addString(c, s1))
    ans2 := minString(addString(s2, c), addString(c, s2))
    return minString(ans1, ans2)
}

func addString(x, y string) string {
    if strings.Contains(x, y) {
        return x
    }
    if strings.Contains(y, x) {
        return y
    }
    n := len(x)
    for i := 0; i < n; i++ {
        if strings.HasPrefix(y, x[i:]) {
            return x + y[n-i:]
        }
    }
    return x + y
}

func minString(x, y string) string {
    if len(x) < len(y) {
        return x
    }
    if len(y) < len(x) {
        return y
    }
    for i := range x {
        if x[i] < y[i] {
            return x
        }
        if y[i] < x[i] {
            return y
        }
    }
    return x
}