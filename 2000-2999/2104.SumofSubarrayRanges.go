func subArrayRanges(nums []int) int64 {
	ans := int64(0)
	for i := 0; i < len(nums)-1; i++ {
		mi, ma := nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			mi = min(mi, nums[j])
			ma = max(ma, nums[j])
			ans += int64(ma - mi)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}