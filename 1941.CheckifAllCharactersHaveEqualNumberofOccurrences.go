func areOccurrencesEqual(s string) bool {
	cnt := make(map[byte]int, 26)
	tmp := 0
	for _, ch := range []byte(s) {
		cnt[ch]++
		tmp = cnt[ch]
	}
	for _, c := range cnt {
		if c != tmp {
			return false
		}
	}
	return true
}