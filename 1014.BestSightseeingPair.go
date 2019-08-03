func maxScoreSightseeingPair(A []int) int {
    a_i := A[0] // a_i: A[i] + i
    ans := 0
    for j := 1; j < len(A); j++ {
        ans = max(ans, a_i+A[j]-j)
        a_i = max(a_i, A[j]+j)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}