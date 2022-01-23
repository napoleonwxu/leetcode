func numSteps(s string) int {
    chs := make([]byte, len(s)+1)
    for i := 0; i < len(s); i++ {
        chs[i+1] = s[i]
    }
    step := 0
    for i := len(chs)-1; i > 1; {
        if chs[i] == '0' {
            step++
            i--
        } else {
            j := i - 1
            for ; chs[j] == '1'; j-- {}
            chs[j] = '1'
            step += i - j + 1
            i = j
        }
    }
    return step
}