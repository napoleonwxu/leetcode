const MOD = int(1e9 + 7)

func valueAfterKSeconds(n int, k int) int {
    nums := make([]int, n)
    for i := range nums {
        nums[i] = 1
    }
    for ; k > 0; k-- {
        for i := 1; i < n; i++ {
            nums[i] = (nums[i] + nums[i-1]) % MOD
        }
    }
    return nums[n-1]
}