func maxTurbulenceSize(A []int) int {
    ans := 1
    inc, dec := 1, 1
    for i := 1; i < len(A); i++ {
        if A[i] > A[i-1] {
            inc, dec = dec+1, 1
        } else if A[i] < A[i-1] {
            inc, dec = 1, inc+1
        } else {
            inc, dec = 1, 1
        }
        ans = max(ans, max(inc, dec))
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}