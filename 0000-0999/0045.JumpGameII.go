func jump(nums []int) int {
    curFar, curEnd := 0, 0
    jumps := 0
    for i := 0; i < len(nums)-1; i++ {
        curFar = max(curFar, i+nums[i])
        if i == curEnd {
            jumps++
            curEnd = curFar
        }
    }
    return jumps
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}