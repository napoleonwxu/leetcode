func nextGreaterElements(nums []int) []int {
    // O(n) + O(n)
    n := len(nums)
    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    stack, m := make([]int, 2*n), 0
    for i := 0; i < 2*n; i++ {
        for m > 0 && nums[stack[m-1]] < nums[i%n] {
            ans[stack[m-1]] = nums[i%n]
            m--
        }
        stack[m] = i%n
        m++
    }
    return ans
}