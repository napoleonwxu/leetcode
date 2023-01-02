func findWinners(matches [][]int) [][]int {
    loss := make(map[int]int)
    for _, match := range matches {
        loss[match[0]] = loss[match[0]]
        loss[match[1]]++
    }
    ans := make([][]int, 2)
    for l, c := range loss {
        if c <= 1 {
            ans[c] = append(ans[c], l)
        }
    }
    sort.Ints(ans[0])
    sort.Ints(ans[1])
    return ans
}