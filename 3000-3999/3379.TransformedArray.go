func constructTransformedArray(nums []int) []int {
    n := len(nums)
    ans := make([]int, len(nums))
    for i, num := range nums {
        ans[i] = nums[((i+num)%n+n)%n]
    }
    return ans
}