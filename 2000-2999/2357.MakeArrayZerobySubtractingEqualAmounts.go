func minimumOperations(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if num != 0 {
			m[num] = true
		}
	}
	return len(m)
}