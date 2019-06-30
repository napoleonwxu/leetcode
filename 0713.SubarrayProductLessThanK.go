func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1 {
        return 0
    }
    cnt, pro := 0, 1
    left, right := 0, 0
    for right < len(nums) {
        pro *= nums[right]
        for pro >= k {
            pro /= nums[left]
            left++
        }
        cnt += right - left + 1
        right++
    }
    return cnt
}