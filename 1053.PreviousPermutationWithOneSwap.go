func prevPermOpt1(A []int) []int {
    n := len(A)
    left, right := n-2, n-1
    for ; left >= 0 && A[left] <= A[left+1]; left-- {}
    if left < 0 {
        return A
    }
    for ; A[left] <= A[right]; right-- {}
    for ; A[right-1] == A[right]; right-- {}
    A[left], A[right] = A[right], A[left]
    return A
}