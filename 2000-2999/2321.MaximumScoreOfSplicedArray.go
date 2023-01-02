func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	return max(kadane(nums1, nums2), kadane(nums2, nums1))
}

func kadane(nums1, nums2 []int) int {
	sum := 0
	diffSum, diffMax := 0, 0
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i]
		diffSum = max(0, diffSum+nums2[i]-nums1[i])
		diffMax = max(diffMax, diffSum)
	}
	return sum + diffMax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
