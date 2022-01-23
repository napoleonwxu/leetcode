func maxSubArray(nums []int) int {
    // O(n) + O(1)
    maxEnd, maxSub := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        maxEnd = max(nums[i], maxEnd+nums[i])
        maxSub = max(maxSub, maxEnd)
    }
    return maxSub
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}