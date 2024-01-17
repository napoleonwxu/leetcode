func getGoodIndices(variables [][]int, target int) []int {
    ans := make([]int, 0, len(variables))
    for i, v := range variables {
        p := power(v[0], v[1], 10)
        p = power(p, v[2], v[3])
        if p == target {
            ans = append(ans, i)
        }
    }
    return ans
}

func power(a, b, m int) int {
    if b == 0 {
        return 1
    }
    p := power(a, b/2, m)
    p = p * p % m
    if b&1 == 1 {
        p = p * a % m
    }
    return p
}