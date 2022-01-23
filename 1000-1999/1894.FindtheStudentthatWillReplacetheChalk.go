func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, ch := range chalk {
		sum += ch
	}
	k %= sum
	for i, ch := range chalk {
		if k < ch {
			return i
		}
		k -= ch
	}
	return -1
}