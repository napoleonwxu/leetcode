func maximumTotalCost(nums []int) int64 {
    add, sub := int64(nums[0]), int64(nums[0])
    for i := 1; i < len(nums); i++ {
        tmpAdd := max(add, sub) + int64(nums[i])
        sub = add - int64(nums[i])
        add = tmpAdd
    }
    return max(add, sub)
}