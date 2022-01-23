func maxSumTwoNoOverlap(A []int, L int, M int) int {
    n := len(A)
    sum := make([]int, n)
    sum[0] = A[0]
    for i := 1; i < n; i++ {
        sum[i] = sum[i-1] + A[i]
    }
    maxL, maxM := sum[L-1], sum[M-1]
    ans := sum[L+M-1]
    for i := L+M; i < n; i++ {
        maxL = max(maxL, sum[i-M]-sum[i-M-L])
        maxM = max(maxM, sum[i-L]-sum[i-L-M])
        ans = max(ans, max(maxL+sum[i]-sum[i-M], maxM+sum[i]-sum[i-L]))
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}