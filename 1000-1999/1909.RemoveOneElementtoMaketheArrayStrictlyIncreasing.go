func canBeIncreasing(nums []int) bool {
	cnt, idx := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			cnt++
			idx = i
		}
	}
	if cnt == 0 {
		return true
	}
	if cnt == 1 {
		if idx == 1 || idx == len(nums)-1 {
			return true
		}
		if (idx >= 2 && nums[idx-2] < nums[idx]) || nums[idx-1] < nums[idx+1] {
			return true
		}
	}
	return false
}