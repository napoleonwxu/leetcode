func maximumLength(nums []int) int {
    cnts, dp := make(map[int]int), make(map[int]int)
    for _, num := range nums {
        cnts[num]++
    }
    ans := 0
    if cnts[1]%2 == 0 {
        ans = cnts[1] - 1
    } else {
        ans = cnts[1]
    }
    sort.Ints(nums)
    for i := len(nums)-1; i >= 0; i-- {
        if cnts[nums[i]] > 0 {
            if dp[nums[i]*nums[i]] > 0 && cnts[nums[i]] >= 2 {
                dp[nums[i]] = dp[nums[i]*nums[i]] + 2
            } else {
                dp[nums[i]] = 1
            }
            ans = max(ans, dp[nums[i]])
            cnts[nums[i]] = 0
        }
    }
    return ans
}