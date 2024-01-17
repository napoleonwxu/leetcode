func findNonMinOrMax(nums []int) int {
    if len(nums) < 3 {
        return -1
    }
    mi := min(nums[0], nums[1])
    ma := max(nums[0], nums[1])
    if nums[2] > mi && nums[2] < ma {
        return nums[2]
    }
    if nums[2] < mi {
        return mi
    }
    return ma
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}