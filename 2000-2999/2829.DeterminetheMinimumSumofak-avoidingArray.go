func minimumSum(n int, k int) int {
    sum := 0
    for i := 1; n > 0; i++ {
        if 2*i <= k || i >= k {
            sum += i
            n--
        }
    }
    return sum
}