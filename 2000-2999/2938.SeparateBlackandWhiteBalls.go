func minimumSteps(s string) int64 {
    var ans, one int64
    for _, ch := range s {
        if ch == '0' {
            ans += one
        } else {
            one++
        }
    }
    return ans
}