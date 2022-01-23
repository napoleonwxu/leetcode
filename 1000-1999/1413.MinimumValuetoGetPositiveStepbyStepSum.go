func minStartValue(nums []int) int {
	sum, minSum := 0, 0
	for _, num := range nums {
		sum += num
		minSum = min(minSum, sum)
	}
	return 1 - minSum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}