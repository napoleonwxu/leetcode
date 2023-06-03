import "sort"

func beautifulSubsets(nums []int, k int) int {
	cnt := make(map[int]int)
	mods := make([][]int, k)
	for i := range mods {
		mods[i] = []int{}
	}
	for _, num := range nums {
		if cnt[num] == 0 {
			mods[num%k] = append(mods[num%k], num)
		}
		cnt[num]++
	}
	ans := 1
	for _, mod := range mods {
		sort.Ints(mod)
		pre, dp0, dp1 := 0, 1, 0
		for _, num := range mod {
			if pre+k == num {
				dp0, dp1 = dp0+dp1, dp0*(1<<cnt[num]-1)
			} else {
				dp0, dp1 = dp0+dp1, (dp0+dp1)*(1<<cnt[num]-1)
			}
			pre = num
		}
		ans *= (dp0 + dp1)
	}
	return ans - 1
}