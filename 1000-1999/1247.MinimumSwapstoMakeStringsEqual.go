func minimumSwap(s1 string, s2 string) int {
    x1, y1 := 0, 0
    x2 := 0
    for i := 0; i < len(s1); i++ {
        if s1[i] == s2[i] {
            continue
        }
        if s1[i] == 'x' {
            x1++
        } else {
            y1++
        }
        if s2[i] == 'x' {
            x2++
        }
    }
    if (x1+x2)&1 == 1 {
        return -1
    }
    return x1/2 + y1/2 + x1&1 * 2
}