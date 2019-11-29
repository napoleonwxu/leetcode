func circularPermutation(n int, start int) []int {
    m := 1 << uint(n)
    gray := make([]int, m)
    for i := 0; i < m; i++ {
        gray[i] = start ^ i ^ (i>>1)
    }
    return gray
}