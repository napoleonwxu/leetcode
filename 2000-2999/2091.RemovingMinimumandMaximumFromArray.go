func minimumDeletions(nums []int) int {
	n := len(nums)
	minIdx, maxIdx := 0, 0
	for i, num := range nums {
		if num < nums[minIdx] {
			minIdx = i
		}
		if num > nums[maxIdx] {
			maxIdx = i
		}
	}
	left, right := min(minIdx, maxIdx), max(minIdx, maxIdx)
	return min(min(right+1, n-left), (left+1)+(n-right))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}