func tribonacci(n int) int {
    // O(n) + O(1)
    if n < 2 {
        return n
    }
    a, b, c := 0, 1, 1
    for ; n > 2; n-- {
        a, b, c = b, c, a+b+c
    }
    return c
    /* O(n) + O(n)
    if n == 0 {
        return 0
    }
    if n == 1 || n == 2 {
        return 1
    }
    tri := make([]int, n+1)
    tri[1], tri[2] = 1, 1
    for i := 3; i <= n; i++ {
        tri[i] = tri[i-3] + tri[i-2] + tri[i-1]
    }
    return tri[n]
    */
}