func reverseStr(s string, k int) string {
    n := len(s)
    rs := []rune(s)
    for i := 0; i < n; i += k+k {
        for j, k := i, min(i+k, n)-1; j < k; j, k = j+1, k-1 {
            rs[j], rs[k] = rs[k], rs[j]
        }
    }
    return string(rs)
}

func min(x int, y int) int {
    if x < y {
        return x
    }
    return y
}