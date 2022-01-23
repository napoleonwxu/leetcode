// n=5: [-4, -2, 0, 2, 4]
func sumZero(n int) []int {
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i] = 2*i - n + 1
    }
    return ans
}

/* n=5: [-2, -1, 0, 1, 2]
func sumZero(n int) []int {
    ans := make([]int, n)
    i, j := n/2-1, n/2+n&1
    for k := 1; k <= n/2; k++ {
        ans[i], ans[j] = -k, k
        i, j = i-1, j+1
    }
    return ans
}
*/