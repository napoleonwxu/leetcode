func countSubarrays(nums []int, k int) int {
	idx, n := 0, len(nums)
	for i, num := range nums {
		if num == k {
			idx = i
			break
		}
	}
	cnt := make(map[int]int)
	balance := 0
	for i := idx; i >= 0; i-- {
		balance += increase(nums[i], k)
		cnt[balance]++
	}
	balance = 0
	ans := 0
	for i := idx; i < n; i++ {
		balance += increase(nums[i], k)
		ans += cnt[-balance] + cnt[1-balance]
	}
	return ans
}

func increase(num, k int) int {
	if num < k {
		return -1
	} else if num > k {
		return 1
	}
	return 0
}
