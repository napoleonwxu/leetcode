import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := nums[k-1] - nums[0]
	for i := 0; i+k <= len(nums); i++ {
		ans = min(ans, nums[i+k-1]-nums[i])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
