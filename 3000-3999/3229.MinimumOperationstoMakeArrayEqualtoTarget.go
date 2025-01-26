func minimumOperations(nums []int, target []int) int64 {
    ans := int64(0)
    pre := 0
    for i := range nums {
        diff := nums[i] - target[i]
        if diff * pre <= 0 {
            ans += int64(abs(pre))
        } else if abs(pre) > abs(diff) {
            ans += int64(abs(pre-diff))
        }
        pre = diff
    }
    ans += int64(abs(pre))
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}