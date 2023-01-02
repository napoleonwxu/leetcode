func maximumTop(nums []int, k int) int {
    n := len(nums)
    if n == 1 && k%2 == 1 {
        return -1
    }
    ma := -1
    for i := 0; i < k-1 && i < n; i++ {
        ma = max(ma, nums[i])
    }
    if k < n {
        ma = max(ma, nums[k])
    }
    return ma
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}