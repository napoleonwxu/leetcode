func minFlips(a int, b int, c int) int {
    flip := 0
    for a+b+c > 0 {
        cnt1 := a&1 + b&1
        if cnt1 <= 1 {
            if cnt1 != c&1 {
                flip++
            }
        } else { //cnt1 == 2
            if c&1 == 0 {
                flip += 2
            }
        }
        a >>= 1
        b >>= 1
        c >>= 1
    }
    return flip
}