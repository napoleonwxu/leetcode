func maxDepthAfterSplit(seq string) []int {
    ans := make([]int, len(seq))
    level := 0
    for i, s := range seq {
        if s == '(' {
            ans[i] = level & 1
            level++
        } else {
            level--
            ans[i] = level & 1
        }
    }
    return ans
}