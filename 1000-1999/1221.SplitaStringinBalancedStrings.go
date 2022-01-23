func balancedStringSplit(s string) int {
    cnt, sum := 0, 0
    for _, ch := range s {
        if ch == 'L' {
            sum++
        } else {
            sum--
        }
        if sum == 0 {
            cnt++
        }
    }
    return cnt
}