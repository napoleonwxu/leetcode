func canJump(nums []int) bool {
    reach := 0
    for i, num := range nums {
        if reach < i {
            return false
        }
        reach = max(reach, i+num)
    }
    return true
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}