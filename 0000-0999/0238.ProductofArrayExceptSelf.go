func productExceptSelf(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    // O(n) + O(1)
    ans[0] = 1
    for i := 1; i < n; i++ {
        ans[i] = ans[i-1] * nums[i-1]
    }
    right := 1
    for i := n-1; i >= 0; i-- {
        ans[i] *= right
        right *= nums[i]
    }
    /* O(n) + O(n)
    left, right := make([]int, n), make([]int, n)
    left[0], right[n-1] = 1, 1
    for i := 1; i < n; i++ {
        left[i] = left[i-1] * nums[i-1]
    }
    for i := n-2; i >= 0; i-- {
        right[i] = right[i+1] * nums[i+1]
    }
    for i := range ans {
        ans[i] = left[i] * right[i]
    } */
    return ans
}