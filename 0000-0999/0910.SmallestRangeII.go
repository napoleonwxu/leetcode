func smallestRangeII(A []int, K int) int {
    n := len(A)
    sort.Ints(A)
    ans := A[n-1] - A[0]
    for i := 0; i < n-1; i++ {
        high := max(A[n-1]-K, A[i]+K)
        low := min(A[0]+K, A[i+1]-K)
        ans = min(ans, high-low)
    }
    return ans
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}