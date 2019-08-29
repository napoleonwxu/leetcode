func flipAndInvertImage(A [][]int) [][]int {
    n := len(A[0])
    for i := range A {
        for l, r := 0, n-1; l <= r; l, r = l+1, r-1 {
            A[i][l], A[i][r] = A[i][r]^1, A[i][l]^1
        }
    }
    return A
}