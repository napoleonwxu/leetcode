func findSubstringInWraproundString(p string) int {
    n := len(p)
    cnt := make([]int, 26)
    sub := 0
    for i := 0; i < n; i++ {
        if i > 0 && (p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25) {
            sub++
        } else {
            sub = 1
        }
        cnt[p[i]-'a'] = max(cnt[p[i]-'a'], sub)
    }
    sum := 0
    for _, c := range cnt {
        sum += c
    }
    return sum
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}