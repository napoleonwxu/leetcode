import "sort"

func rearrangeArray(nums []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
	return nums
}