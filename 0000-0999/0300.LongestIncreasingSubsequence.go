func lengthOfLIS(nums []int) int {
    // binSearch, O(nlgn) + O(n)
    inc := make([]int, 0, len(nums))
    for _, num := range nums {
        idx := binInsertLeft(inc, num)
        if idx == len(inc) {
            inc = append(inc, num)
        } else {
            inc[idx] = num
        }
    }
    return len(inc)
    /* DP, O(n^2) + O(n)
    n := len(nums)
    dp := make([]int, n)
    ans := 0
    for i := 0; i < n; i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        ans = max(ans, dp[i])
    }
    return ans
    */
}

func binInsertLeft(nums[]int, target int) int {
    l, r := 0, len(nums)
    for l < r {
        mid := (l+r) / 2
        if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}