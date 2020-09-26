func canBeEqual(target []int, arr []int) bool {
	cnt := make(map[int]int)
	for _, num := range arr {
		cnt[num]++
	}
	for _, t := range target {
		if cnt[t] <= 0 {
			return false
		}
		cnt[t]--
	}
	return true
}