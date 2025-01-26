func getPermutation(n int, k int) string {
    nums := make([]int, n)
    fact := make([]int, n+1)
    fact[0] = 1
    for i := 1; i <= n; i++ {
        nums[i-1] = i
        fact[i] = i * fact[i-1]
    }
    ans := ""
    k--
    for i := 1; i <= n; i++ {
        idx := k / fact[n-i]
        ans += strconv.Itoa(nums[idx])
        nums = append(nums[:idx], nums[idx+1:]...)
        k %= fact[n-i]
    }
    return ans
}