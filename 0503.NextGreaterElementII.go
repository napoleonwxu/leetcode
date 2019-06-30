func nextGreaterElements(nums []int) []int {
    // O(n) + O(n)
    n := len(nums)
    ans := make([]int, n)
    stack := make([]int, n)
    for i := n-1; i >= 0; i-- {
        stack = append(stack, i)
    }
    for i := n-1; i >= 0; i-- {
        s := len(stack)
        for s > 0 && nums[stack[s-1]] <= nums[i] {
            stack = stack[:s-1]
            s--
        }
        ans[i] = -1
        if s > 0 {
            ans[i] = nums[stack[s-1]]
        }
        stack = append(stack, i)
    }
    return ans
}