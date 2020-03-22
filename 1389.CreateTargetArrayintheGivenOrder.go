func createTargetArray(nums []int, index []int) []int {
    ans := []int{}
    for i := range nums {
        ans = append(ans, 0)
        for j := len(ans)-1; j > index[i]; j-- {
            ans[j] = ans[j-1]
        }
        ans[index[i]] = nums[i]
    }
    return ans
}