func findTheArrayConcVal(nums []int) int64 {
	ans := int64(0)
	i, j := 0, len(nums)-1
	for ; i < j; i, j = i+1, j-1 {
		l := 1
		for tmp := nums[j]; tmp > 0; tmp /= 10 {
			l *= 10
		}
		ans += int64(nums[i]*l + nums[j])
	}
	if i == j {
		ans += int64(nums[i])
	}
	return ans
}