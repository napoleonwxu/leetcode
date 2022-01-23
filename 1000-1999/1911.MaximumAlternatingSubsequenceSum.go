func maxAlternatingSum(nums []int) int64 {
	ans := int64(nums[0])
	for i := 1; i < len(nums); i++ {
		ans += int64(max(nums[i]-nums[i-1], 0))
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}