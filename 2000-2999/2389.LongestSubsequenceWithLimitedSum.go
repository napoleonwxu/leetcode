import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = binSearch(nums, query)
	}
	return ans
}

func binSearch(nums []int, query int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= query {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}