func getDescentPeriods(prices []int) int64 {
	cnt := int64(0)
	i, n := 0, len(prices)
	for i < n {
		j := i + 1
		for ; j < n && prices[j]+1 == prices[j-1]; j++ {
		}
		cnt += int64((j - i) * (j - i + 1) / 2)
		i = j
	}
	return cnt
}