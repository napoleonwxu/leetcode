func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i], nums[i+1] = 2*nums[i], 0
		}
	}
	i := 0
	for j := 0; j < n; j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for ; i < n; i++ {
		nums[i] = 0
	}
	return nums
}