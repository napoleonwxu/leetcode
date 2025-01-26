const MOD = int(1e9 + 7)

func sumOfGoodSubsequences(nums []int) int {
    ans := 0
    cnt := make([]int, 100010)
    sum := make([]int, 100010)
    for _, num := range nums {
        cnt[num+1] = (cnt[num] + cnt[num+1] + cnt[num+2] + 1) % MOD
        cur := sum[num] + sum[num+2] + num*(cnt[num] + cnt[num+2] + 1)
        sum[num+1] = (sum[num+1] + cur) % MOD
        ans = (ans + cur) % MOD
    }
    return ans
}