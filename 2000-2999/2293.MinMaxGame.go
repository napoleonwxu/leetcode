func minMaxGame(nums []int) int {
    n := len(nums)
    for n > 1 {
        for i := 0; i < n/2; i++ {
            if i&1 == 0 {
                nums[i] = min(nums[2*i], nums[2*i+1])
            } else {
                nums[i] = max(nums[2*i], nums[2*i+1])
            }
        }
        n /= 2
    }
    return nums[0]
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