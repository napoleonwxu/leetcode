func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
    ans := make([][]int, 2)
    ans[0] = make([]int, len(colsum))
    ans[1] = make([]int, len(colsum))
    for i, s := range colsum {
        if s == 2 || (s == 1 && upper > lower) {
            ans[0][i] = 1
            upper--
        }
        if s == 2 || (s == 1 && ans[0][i] == 0) {
            ans[1][i] = 1
            lower--
        }
    }
    if upper == 0 && lower == 0 {
        return ans
    }
    return nil
    /*
    idx := []int{}
    sum, sum0 := 0, 0
    for i, s := range colsum {
        sum += s
        if s == 2 {
            ans[0][i], ans[1][i] = 1, 1
            sum0++
        } else if s == 1 {
            idx = append(idx, i)
        }
    }
    if sum != upper+lower || sum0 > upper || sum0 > lower {
        return nil
    }
    for i := 0; i < upper-sum0; i++ {
        ans[0][idx[i]] = 1
    }
    for i := upper-sum0; i < len(idx); i++ {
        ans[1][idx[i]] = 1
    }
    return ans
    */
}