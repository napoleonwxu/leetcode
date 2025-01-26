func countKConstraintSubstrings(s string, k int) int {
    zero, one := 0, 0
    ans := 0
    for l, r := 0, 0; r < len(s); r++ {
        if s[r] == '0' {
            zero++
        } else {
            one++
        }
        for ; zero > k && one > k; l++ {
            if s[l] == '0' {
                zero--
            } else {
                one--
            }
        }
        ans += r - l + 1
    }
    return ans
}