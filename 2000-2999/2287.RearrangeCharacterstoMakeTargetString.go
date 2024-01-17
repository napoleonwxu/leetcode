func rearrangeCharacters(s string, target string) int {
    cntS := make([]int, 26)
    cntT := make([]int, 26)
    for _, ch := range s {
        cntS[ch-'a']++
    }
    for _, ch := range target {
        cntT[ch-'a']++
    }
    ans := len(s) / len(target)
    for i := range cntS {
        if cntT[i] > 0 {
            ans = min(ans, cntS[i]/cntT[i])
        }
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}