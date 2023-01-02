func arithmeticTriplets(nums []int, diff int) int {
	m := make(map[int]bool, len(nums))
	cnt := 0
	for _, num := range nums {
		if m[num-diff] && m[num-2*diff] {
			cnt++
		}
		m[num] = true
	}
	return cnt
}