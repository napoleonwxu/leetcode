import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] < 0 && k > 0; i++ {
		nums[i] = -nums[i]
		k--
	}
	sum, mi := 0, 100
	for _, num := range nums {
		sum += num
		mi = min(mi, num)
	}
	return sum - (k%2)*2*mi
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
