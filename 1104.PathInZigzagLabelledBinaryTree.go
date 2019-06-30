func pathInZigZagTree(label int) []int {
    ans := []int{}
    for ; label > 0; label >>= 1 {
        ans = append(ans, label)
    }
    n := len(ans)
    for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
        ans[l], ans[r] = ans[r], ans[l]
    }
    sum := make([]int, n)
    sum[0] = 2
    for i := 1; i < n; i++ {
        sum[i] = sum[i-1]<<1 + 1
    }
    for i := n&1; i < n; i += 2 {
        ans[i] = sum[i] - ans[i]
    }
    return ans
}