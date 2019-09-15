func maxProduct(nums []int) int {
    n := len(nums)
    ans := nums[0]
    Max, Min := nums[0], nums[0]
    for i := 1; i < n; i++ {
        if nums[i] < 0 {
            Max, Min = Min, Max
        }
        Max = max(Max*nums[i], nums[i])
        Min = min(Min*nums[i], nums[i])
        ans = max(ans, Max)
    }
    /*
    l, r := 0, 0
    for i := 0; i < n; i++ {
        if l != 0 {
            l *= nums[i]
        } else {
            l = nums[i]
        }
        if r != 0 {
            r *= nums[n-1-i]
        } else {
            r = nums[n-1-i]
        }
        ans = max(ans, max(l, r))
    }*/
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}