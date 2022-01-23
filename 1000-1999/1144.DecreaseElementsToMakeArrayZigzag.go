func movesToMakeZigzag(nums []int) int {
    n := len(nums)
    d := [2]int{}
    for i := 0; i < n; i++ {
        left, right := 1001, 1001
        if i > 0 {
            left = nums[i-1]
        }
        if i < n-1 {
            right = nums[i+1]
        }
        d[i&1] += max(0, nums[i]-min(left, right)+1)
    }
    return min(d[0], d[1])
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