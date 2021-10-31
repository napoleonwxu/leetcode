func singleNumber(nums []int) []int {
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	diff &= -diff // get highest bit
	ans := []int{0, 0}
	for _, num := range nums {
		if num&diff != 0 {
			ans[0] ^= num
		} else {
			ans[1] ^= num
		}
	}
	return ans
}