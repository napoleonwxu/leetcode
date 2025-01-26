func maxOperations(s string) int {
    ans, pre := 0, 0
    for i := 1; i < len(s); i++ {
        if s[i-1] == '1' {
            pre++
            if s[i] == '0' {
                ans += pre
            }
        }
    }
    return ans
}