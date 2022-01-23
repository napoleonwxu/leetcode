func pushDominoes(dominoes string) string {
    n := len(dominoes)
    d := make([]int, n)
    f := 0
    for i := 0; i < n; i++ {
        if dominoes[i] == 'R' {
            f = n
        } else if dominoes[i] == 'L' {
            f = 0
        } else {
            f = max(f-1, 0)
        }
        d[i] = f
    }
    f = 0
    for i := n-1; i >= 0; i-- {
        if dominoes[i] == 'L' {
            f = n
        } else if dominoes[i] == 'R' {
            f = 0
        } else {
            f = max(f-1, 0)
        }
        d[i] -= f
    }
    s := make([]string, n)
    for i := 0; i < n; i++ {
        if d[i] < 0 {
            s[i] = "L"
        } else if d[i] > 0 {
            s[i] = "R"
        } else {
            s[i] = "."
        }
    }
    return strings.Join(s, "")
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}