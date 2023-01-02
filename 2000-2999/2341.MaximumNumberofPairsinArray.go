func numberOfPairs(nums []int) []int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}
	pair, left := 0, 0
	for _, cnt := range m {
		pair += cnt / 2
		left += cnt % 2
	}
	return []int{pair, left}
}