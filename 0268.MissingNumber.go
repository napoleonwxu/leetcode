func missingNumber(nums []int) int {
	sum, n := 0, len(nums)
	for _, num := range nums {
		sum += num
	}
	return (1+n)*n/2 - sum
}