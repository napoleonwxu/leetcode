func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    } else if n == 1 {
        return nums[0]
    }
    return max(dp(nums, 0, n-1), dp(nums, 1, n))
}

func dp(nums []int, left, right int) int {
    pre, cur := 0, 0
    for i := left; i < right; i++ {
        pre, cur = cur, max(pre+nums[i], cur)
    }
    return  cur
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}