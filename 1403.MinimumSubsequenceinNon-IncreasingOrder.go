func minSubsequence(nums []int) []int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    sort.Ints(nums)
    tmp, ans := 0, []int{}
    //for i := len(nums)-1; i >= 0 && tmp <= sum-tmp; i-- {
	for i := len(nums)-1; tmp <= sum-tmp; i-- {
        ans = append(ans, nums[i])
        tmp += nums[i]
    }
    return ans
}