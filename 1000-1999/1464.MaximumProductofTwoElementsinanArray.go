func maxProduct(nums []int) int {
	max, sec_max := 0, 0
	for _, num := range nums {
		if num > max {
			sec_max = max
			max = num
		} else if num > sec_max {
			sec_max = num
		}
	}
	return (max - 1) * (sec_max - 1)
}