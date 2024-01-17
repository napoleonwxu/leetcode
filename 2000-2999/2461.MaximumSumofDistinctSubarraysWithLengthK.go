func maximumSubarraySum(nums []int, k int) int64 {
    n := len(nums)
    cnts := make(map[int]int, n)
    ans := int64(0)
    sum, i := int64(0), 0
    for ; i < k; i++ {
        cnts[nums[i]]++
        sum += int64(nums[i])
    }
    if len(cnts) == k {
        ans = sum
    }
    for ; i < n; i++ {
        cnts[nums[i]]++
        cnts[nums[i-k]]--
        if cnts[nums[i-k]] == 0 {
            delete(cnts, nums[i-k])
        }
        sum += int64(nums[i] - nums[i-k])
        if len(cnts) == k {
            ans = max(ans, sum)
        }
    }
    return ans
}

func max(x, y int64) int64 {
    if x > y {
        return x
    }
    return y
}