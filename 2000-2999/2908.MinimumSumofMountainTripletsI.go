func minimumSum(nums []int) int {
    ans, n := -1, len(nums)
    minRights := make([]int, n)
    minRights[n-1] = nums[n-1]
    for i := n - 2; i >= 0; i-- {
        minRights[i] = min(minRights[i+1], nums[i])
    }
    minLeft := nums[0]
    for i := 1; i < n-1; i++ {
        if nums[i] > minLeft && nums[i] > minRights[i] {
            sum := minLeft + nums[i] + minRights[i]
            if ans == -1 || sum < ans {
                ans = sum
            }
        }
        minLeft = min(minLeft, nums[i])
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
