func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n < 0 {
        x = 1/x
        n = -n
    }
    if n&1 == 0 {
        return myPow(x*x, n>>1)
    }
    return x*myPow(x*x, n>>1)
}