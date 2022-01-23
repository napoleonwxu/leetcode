/*
p: plus
m: minus
sum(p) - sum(m) = S
sum(p) + sum(m) + sum(p) - sum(m) = S + sum(p) + sum(m)
2 * sum(p) = S + sum(nums)
S & sum(nums) must both even or both odd
*/
func findTargetSumWays(nums []int, S int) int {
    // dp, 0ms, O(n*s). s: sum(nums), n: len(sums)
    sum := 0
    for _, n := range nums {
        sum += n
    }
    if sum < S {
        return 0
    }
    if (sum+S)&1 == 1 {
        return 0
    }
    target := (sum+S) >> 1
    dp := make([]int, target+1)
    dp[0] = 1
    for _, num := range nums {
        for i := target; i >= n; i-- {
            dp[i] += dp[i-num]
        }
    }
    return dp[target]
    // dfs, O(2^n), 644ms
    //return dfs(nums, S, 0, 0)
}

func dfs(nums []int, S, i, sum int) int {
    if i == len(nums) {
        if sum == S {
            return 1
        }
        return 0
    }
    return dfs(nums, S, i+1, sum+nums[i]) + dfs(nums, S, i+1, sum-nums[i])
}