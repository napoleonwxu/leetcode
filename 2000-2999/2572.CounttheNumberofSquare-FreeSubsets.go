const MOD = int(1e9 + 7)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var dp = make([][]int, 1001) // max length of nums

func squareFreeSubsets(nums []int) int {
	for i := range dp {
		dp[i] = make([]int, 1<<len(primes))
	}
	return (dfs(nums, 0, 0) - 1 + MOD) % MOD
}

func dfs(nums []int, i, mask int) int {
	if i >= len(nums) {
		return 1
	}
	if dp[i][mask] != 0 {
		return dp[i][mask]
	}
	add := true
	newMask := mask
	for j, prime := range primes {
		tmp, cnt := nums[i], 0
		for tmp%prime == 0 {
			tmp /= prime
			cnt++
		}
		if cnt >= 2 || (cnt == 1 && (mask>>j)&1 == 1) {
			add = false
			break
		}
		if cnt > 0 {
			newMask |= 1 << j
		}
	}
	if add {
		dp[i][mask] = (dfs(nums, i+1, mask) + dfs(nums, i+1, newMask)) % MOD
	} else {
		dp[i][mask] = dfs(nums, i+1, mask)
	}
	return dp[i][mask]
}