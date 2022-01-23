func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    if n < k || k <= 0 {
        return nil
    }
    maxL, maxR := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        if i%k == 0 {
            maxL[i] = nums[i]
        } else {
            maxL[i] = max(maxL[i-1], nums[i])
        }
    }
    maxR[n-1] = nums[n-1]
    for i := n-2; i >= 0; i-- {
        if i%k == 0 {
            maxR[i] = nums[i]
        } else {
            maxR[i] = max(maxR[i+1], nums[i])
        }
    }
    ans := make([]int, n-k+1)
    for i := 0; i <= n-k; i++ {
        ans[i] = max(maxR[i], maxL[i+k-1])
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}