func minCost(nums []int, x int) int64 {
    n := len(nums)
    ans := make([]int64, n)
    for i := 0; i < n; i++ {
        ans[i] += int64(i * x)
        cur := nums[i]
        for k := 0; k < n; k++ {
            cur = min(cur, nums[(i-k+n)%n])
            ans[k] += int64(cur)
        }
    }
    mi := ans[0]
    for i := 1; i < n; i++ {
        if ans[i] < mi {
            mi = ans[i]
        }
    }
    return mi
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}