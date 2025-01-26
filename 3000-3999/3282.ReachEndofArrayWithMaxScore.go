func findMaximumScore(nums []int) int64 {
    ans := int64(0)
    ma := 0
    for _, num := range nums {
        ans += int64(ma)
        ma = max(ma, num)
    }
    return ans
}
