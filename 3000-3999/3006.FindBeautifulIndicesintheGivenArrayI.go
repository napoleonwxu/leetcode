func beautifulIndices(s string, a string, b string, k int) []int {
    ls, la, lb := len(s), len(a), len(b)
    ai := make([]int, 0)
    bi := make([]int, 0)
    for i := range s {
        if i+la <= ls && s[i:i+la] == a {
            ai = append(ai, i)
        }
        if i+lb <= ls && s[i:i+lb] == b {
            bi = append(bi, i)
        }
    }
    ans := make([]int, 0, len(ai))
    for _, i := range ai {
        for _, j := range bi {
            if abs(i-j) <= k {
                ans = append(ans, i)
                break
            }
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}