func grayCode(n int) []int {
    gray := make([]int, 1<<uint(n))
    for i := range gray {
        gray[i] = i ^ (i>>1)
    }
    return gray
}