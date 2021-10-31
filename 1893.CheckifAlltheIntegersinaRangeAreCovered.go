const MaxValue = 50

func isCovered(ranges [][]int, left int, right int) bool {
	nums := make([]int, MaxValue+2)
	for _, r := range ranges {
		nums[r[0]]++
		nums[r[1]+1]--
	}
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for i := left; i <= right; i++ {
		if nums[i] == 0 {
			return false
		}
	}
	return true
}