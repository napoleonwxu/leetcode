func minOperations(nums1 []int, nums2 []int, k int) int64 {
	ops, bal := int64(0), int64(0)
	for i := range nums1 {
		diff := nums1[i] - nums2[i]
		if diff == 0 {
			continue
		}
		if k == 0 || diff%k != 0 {
			return -1
		}
		if diff > 0 {
			ops += int64(diff / k)
		}
		bal += int64(diff)
	}
	if bal == 0 {
		return ops
	}
	return -1
}
