func countGood(nums []int, k int) int64 {
	n, ans := len(nums), int64(0)
	cnt := make(map[int]int, n)
	for i, j := 0, 0; j < n; j++ {
		k -= cnt[nums[j]]
		cnt[nums[j]]++
		for k <= 0 {
			cnt[nums[i]]--
			k += cnt[nums[i]]
			i++
		}
		ans += int64(i)
	}
	return ans
}