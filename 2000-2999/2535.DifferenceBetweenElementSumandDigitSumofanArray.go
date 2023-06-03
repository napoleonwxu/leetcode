func differenceOfSum(nums []int) int {
	sum1, sum2 := 0, 0
	for _, num := range nums {
		sum1 += num
		for ; num > 0; num /= 10 {
			sum2 += num % 10
		}
	}
	return sum1 - sum2
}
