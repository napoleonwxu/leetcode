func longestAlternatingSubarray(nums []int, threshold int) int {
    ans, n := 0, len(nums)
    for i := 0; i < n; i++ {
        if nums[i]%2 == 0 && nums[i] <= threshold {
            l, r := i, i
            for ; r+1 < n && nums[r+1] <= threshold && (nums[r]+nums[r+1])%2 == 1; r++ {
            }
            ans = max(ans, r-l+1)
            i = r
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}