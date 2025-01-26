func hasIncreasingSubarrays(nums []int, k int) bool {
    preLen, curLen := 0, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            curLen++
        } else {
            preLen, curLen = curLen, 1
        }
        if (curLen >= k && preLen >= k) || (curLen >= 2*k) {
            return true
        }
    }
    return false
}