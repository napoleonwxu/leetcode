func minimumChairs(s string) int {
    ans, cur := 0, 0
    for i := range s {
        if s[i] == 'E' {
            cur++
        } else if s[i] == 'L' {
            cur--
        }
        ans = max(ans, cur)
    }
    return ans
}