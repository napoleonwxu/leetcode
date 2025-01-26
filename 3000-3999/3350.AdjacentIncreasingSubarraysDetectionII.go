func maxIncreasingSubarrays(nums []int) int {
    // O(n)
    ans := 1
    pre, cur := 0, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            cur++
        } else {
            pre, cur = cur, 1
        }
        ans = max(ans, max(cur/2, min(cur, pre)))
    }
    return ans
    /* binary search O(nlgn)
    l, r := 1, len(nums)/2 + 1
    for l < r {
        m := (l+r) / 2
        if check(nums, m) {
            l = m + 1
        } else {
            r = m
        }
    }
    return l - 1
    */
}

func check(nums []int, k int) bool {
    pre, cur := 0, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            cur++
        } else {
            pre, cur = cur, 1
        }
        if (pre >= k && cur >= k) || (cur >= 2*k) {
            return true
        }
    }
    return false
}